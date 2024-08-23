package main

import (
	"fmt"
	"time"
)

func createClientsSingleRoutine() {
	singleMap := make(map[string]string)
	fmt.Println("Starting single routine flow")

	start := time.Now()
	//single routine
	for i := 0; i < 100; i++ {
		log, pwd := createPair()
		singleMap[log] = pwd
	}
	end := time.Now()

	fmt.Println("Single routine time elapsed: ", end.Sub(start))
	fmt.Println("map capacity: ", len(singleMap))

}
