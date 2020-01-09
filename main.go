package main

import (
	"fmt"
	"os"
	"sync"
	"time"

	"bygui86/go-atomic-counters/atomics"
)

const (
	jobInterval     = 5
	jobIntervalUnit = "s"

	goroutinesNum       = 10
	goroutingIterations = 100
)

func main() {
	jobLogger, jobErr := atomics.InitJobLogger(jobInterval, jobIntervalUnit)
	if jobErr != nil {
		fmt.Println("ERROR - creating JobLogger failed:", jobErr.Error())
		os.Exit(1)
	}

	fmt.Println("JobLogger started")
	jobLogger.Start()

	var wg sync.WaitGroup
	wg.Add(goroutinesNum)
	for r := 0; r < goroutinesNum; r++ {
		fmt.Println("Starting Goroutine", r)
		go func(id int) {
			// fmt.Println("[DEBUG] Goroutine", id, "started")
			for i := 0; i < goroutingIterations; i++ {
				// fmt.Println("[DEBUG] TotalCounts increased!")
				atomics.TotalCounts.Increment()
				time.Sleep(1 * time.Second)
			}
			// fmt.Println("[DEBUG] Goroutine", id, "completed")
			wg.Done()
		}(r)
	}
	wg.Wait()

	jobLogger.Stop()
	fmt.Println("JobLogger completed")
}
