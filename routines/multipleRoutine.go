package main

import (
	"fmt"
	"sync"
	"time"
)

func createClientsMultipleRoutines() {
	fmt.Println("Starting multiple (5) routines flow")
	var routines = 5
	var multiMap = make(map[string]string)
	var mutex sync.Mutex
	var waiting sync.WaitGroup
	start := time.Now()

	for i := 0; i < routines; i++ {
		waiting.Add(1)
		go func() {
			defer waiting.Done()
			for i := 0; i < 100/routines; i++ {
				log, pwd := createPair()
				mutex.Lock()
				multiMap[log] = pwd
				mutex.Unlock()
			}
		}()
	}
	waiting.Wait()
	fmt.Println("Multiple routines time elapsed: ", time.Since(start))
	fmt.Println("multiMap capacity: ", len(multiMap))

}
