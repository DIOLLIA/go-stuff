package main

import (
	"fmt"
	"sync"
	"time"
)

func createClientWithChannels() {
	fmt.Println("Starting channels flow")

	var waitGroup sync.WaitGroup
	channelMap := make(map[string]string)
	var resultMapChan = make(chan LPair)

	start := time.Now()
	for i := 0; i < 100; i++ {
		waitGroup.Add(1)
		go func() {
			defer waitGroup.Done()
			login, pwd := createPair()
			resultMapChan <- LPair{login, pwd}
		}()
	}
	go func() {
		waitGroup.Wait()
		close(resultMapChan)
	}()

	for pairs := range resultMapChan {
		channelMap[pairs.login] = pairs.pwd
	}

	fmt.Println("With channels time elapsed: ", time.Since(start))
	fmt.Println("result channelMap capacity: ", len(channelMap))
}
