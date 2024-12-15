package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

// we are creating a mutex to not allow for writing to the same memory
// address in the same time
var m = sync.Mutex{}

var dbData = []string{"id1", "id2", "id3"}
var results = []string{}

func main() {
	// Concurrency is not parallel execution
	// concurrency is multiple task running at the same time
	// go before the function that we want to use goroutines

	for i := 0; i < len(dbData); i++ {
		// without wait group and `go` no result will be print
		// it will just go and doesn't wait for result
		// wait group is a counter
		// we can add 1 to a counter and than .Done() to decrement
		wg.Add(1)
		go dbCall(i)
	}
	// with .Wait() we wait for the counter to go down to 0
	wg.Wait()
	fmt.Println("Results are:", results)
}

func dbCall(i int) {
	var delay float32 = 2000
	// cannot use m.Lock() here because it will lock for 2000 seconds
	// it will destroy the concurrency
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Println("The result from the database is:", dbData[i])
	// it checks if some other routine Locked something and if so
	// wait until unlock
	m.Lock()
	results = append(results, dbData[i])
	m.Unlock()
	// decrement the wait group
	wg.Done()
	// with RWMutex() we have the same functionality
	// additionally we have .RLock() and .RUnlock
	// for Read lock and read unlock, for reading data
}
