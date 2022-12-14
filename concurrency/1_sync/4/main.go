package main

import (
	"fmt"
	"math/rand"
	"sync"
)

// починить
func main() {
	wg := sync.WaitGroup{}
	ch := make(chan string)

	for i := 0; i < rand.Intn(100000); i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ch <- fmt.Sprintf("Goroutine %d", i)
		}()
	}

	for v := range ch {
		fmt.Println(v)
	}

}
