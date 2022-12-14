package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	ch := make(chan string)

	for i := 0; i < rand.Intn(100000); i++ {
		wg.Add(1)
		go func(iIn int) {
			defer wg.Done()
			ch <- fmt.Sprintf("Goroutine %d", iIn)
		}(i)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for v := range ch {
		fmt.Println(v)
	}

}
