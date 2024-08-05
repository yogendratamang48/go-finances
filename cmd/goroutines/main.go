package main

import (
	"fmt"
	"sync"
	"time"
)

var dbData = []string{"data1", "data2", "data3", "data4", "data5"}
var wg = sync.WaitGroup{}
var lock = sync.Mutex{}
var results = []string{}
var enableSequential = false

func dbCall(i int) {
	var delay float32 = 2000

	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Println("Result from the database is: ", dbData[i])
}

func dbCallRoutine(i int) {
	var delay float32 = 2000

	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Println("Result from the database is: ", dbData[i])
	lock.Lock()
	results = append(results, dbData[i])
	lock.Unlock()
	wg.Done()
}
func main() {
	fmt.Println("# Goroutines")
	fmt.Println("Goroutines are lightweight threads managed by the Go runtime.")
	fmt.Println("Goroutines are functions or methods that run concurrently with other functions or methods.")
	fmt.Println("Goroutines are started by using the 'go' keyword followed by a function invocation.")
	fmt.Println("Goroutines are multiplexed onto multiple OS threads so if one should block, such as waiting for I/O, others continue to run.")
	fmt.Println("Goroutines are used to make concurrent programming easier.")
	t0 := time.Now()
	if enableSequential {
		fmt.Println("---> Running Sequentially")

		for i := 0; i < len(dbData); i++ {
			dbCall(i)
		}
		fmt.Printf("\n Total Execution time %v\n", time.Since(t0))

		fmt.Println("---> Running With Goroutines")

	}
	t0 = time.Now()
	for i := 0; i < len(dbData); i++ {
		wg.Add(1)
		go dbCallRoutine(i)
	}
	wg.Wait()
	fmt.Printf("\n Total Execution time %v\n", time.Since(t0))
	fmt.Printf("Results: %v\n", results)

	fmt.Println("Using Mutex")

}
