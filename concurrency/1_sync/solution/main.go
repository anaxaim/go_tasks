package main

import (
	"fmt"
	"sync"
)

// print square of range 0...20 in random order
func main() {
	wg := sync.WaitGroup{}
	counter := 20
	for i := 0; i <= counter; i++ {
		wg.Add(1)
		i := i
		go func() {
			defer wg.Done()
			fmt.Println(i * i)
		}()
	}
	wg.Wait()
}
