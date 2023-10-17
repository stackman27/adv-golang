package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	userName := fetchUser()

	// Buffered channel is usually secure
	respChannel := make(chan any, 2)
	wg := &sync.WaitGroup{}

	wg.Add(2) // we have 2 goRoutine running

	go fetchUserLikes(userName, respChannel, wg)
	go fetchUserMatch(userName, respChannel, wg)

	wg.Wait() // block until we have 2 wg.Done() calls (this is called in goRoutine functions)
	close(respChannel)

	for resp := range respChannel {
		fmt.Println("resp: ", resp)
	}

	fmt.Println("took: ", time.Since(start))
}

func fetchUser() string {
	time.Sleep(time.Millisecond * 100)

	return "BOB"
}

func fetchUserLikes(userName string, respch chan any, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 150)

	respch <- 11
	wg.Done()
}

func fetchUserMatch(userName string, respch chan any, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 100)

	respch <- "ANNA"
	wg.Done()
}
