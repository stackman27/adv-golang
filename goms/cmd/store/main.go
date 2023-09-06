package main

import (
	"flag"
	"fmt"
	"learn-go/adv-go/goms/types"

	"github.com/anthdm/hollywood/actor"
	"github.com/anthdm/hollywood/remote"
)

type store struct {

}

func newStore() actor.Receiver {
	return &store{} 
}


func (s *store) Receive(c *actor.Context) {
	switch msg:= c.Message().(type) {

		case *types.CatFact:
			fmt.Println("stored fact into the db: ", msg.Fact) 
		
		case actor.Started:
			fmt.Println("Store is started")
		case actor.Stopped:
	}	
}

func main() {
	listenAddr := flag.String("listenAddr", "127.0.0.1:4000", "todo")
	flag.Parse()

	e := actor.NewEngine()
	r := remote.New(e, remote.Config{ListenAddr: *listenAddr})
	e.WithRemote(r) 

	e.Spawn(newStore, "store")
	
	select{}

}