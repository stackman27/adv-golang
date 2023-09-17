package main

import (
	"encoding/json"
	"flag"
	"learn-go/adv-go/goms/types"
	"log"
	"net/http"
	"time"

	"github.com/anthdm/hollywood/actor"
	"github.com/anthdm/hollywood/remote"
)

// 2 micro services
// - scarper
// - storage

// we dont want to scrape API too fast because we can get locked out
const scrapeInterval = time.Second
const url = "https://catfact.ninja/fact"

type Scraper struct {
	url string
	// ProcessId = id to get the engine that is used to locate where the process is and execute them
	// Normally replaced this by cluster
	storePID *actor.PID
	engine   *actor.Engine
}

func newScraper(url string, storePID *actor.PID) actor.Producer {
	return func() actor.Receiver {
		return &Scraper{
			url:      url,
			storePID: storePID,
		}
	}
}

func (s *Scraper) Receive(c *actor.Context) {
	switch msg := c.Message().(type) {
	case actor.Started:
		s.engine = c.Engine()
		go s.scrapeLoop() // goRoutine: this is so that we dont block the message scraping when we recall
	case actor.Stopped:

	default:
		_ = msg
	}
}

func (s *Scraper) scrapeLoop() {
	for {
		resp, err := http.Get(s.url)
		if err != nil {
			panic(err)
		}
		var res CatFact
		if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
			log.Print("failed to decode the response bytes: ", err)
			continue
		}

		s.engine.Send(s.storePID, &types.CatFact{
			Fact: res.Fact,
		})

		time.Sleep(scrapeInterval)
	}
}

type CatFact struct {
	Fact string `json:"fact"`
}

func main() {

	listenAddr := flag.String("listenAddr", "127.0.0.1:3000", "todo")
	flag.Parse()

	e := actor.NewEngine()
	r := remote.New(e, remote.Config{ListenAddr: *listenAddr})
	e.WithRemote(r)

	// pid 127.0.0.1/store
	storePID := actor.NewPID("127.0.0.1:4000", "store")

	e.Spawn(newScraper(url, storePID), "store")

	select {}
}
