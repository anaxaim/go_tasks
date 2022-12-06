package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

const numRequests = 10000

var count int32

func networkRequest(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Millisecond) // Эмуляция сетевого запроса.
	atomic.AddInt32(&count, 1)
}

func main() {

	var wg sync.WaitGroup
	for i := 0; i < numRequests; i++ {
		wg.Add(1)
		go networkRequest(&wg)
	}
	wg.Wait()

	fmt.Println(atomic.LoadInt32(&count))
}
