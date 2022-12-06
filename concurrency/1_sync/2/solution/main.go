// Находим максимальное четное число

package main

import (
	"fmt"
	"sync"
)

func main() {
	var max int
	var wg sync.WaitGroup
	var lock sync.RWMutex

	for i := 1000; i > 0; i-- {
		wg.Add(1)
		go func(iIn int) {
			defer wg.Done()
			if iIn%2 == 0 && iIn > max {
				lock.RLock()
				max = iIn
				lock.RUnlock()
			}
		}(i)
	}
	wg.Wait()

	fmt.Printf("Maximum is %d", max)
}

// добавил wg, чтобы пройти по всем числам в диапазоне и не потерять нужные
// добавили мьютекс (на чтение), так как переменная макс - это общий ресурс для всех горутин
