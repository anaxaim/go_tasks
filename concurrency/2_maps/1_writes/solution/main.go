package main

import (
	"fmt"
	"sync"
)

func main() {
	writes := 1000
	storage := make(map[int]int, 1000)

	wg := sync.WaitGroup{}
	mu := sync.Mutex{}

	wg.Add(writes)
	for i := 0; i < writes; i++ {
		i := i
		go func() {
			defer wg.Done()
			mu.Lock()
			storage[i] = i
			mu.Unlock()
		}()
	}

	wg.Wait()
	fmt.Println(storage)
}
