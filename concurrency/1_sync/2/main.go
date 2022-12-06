// Находим максимальное четное число

package main

import "fmt"

func main() {
	var max int

	for i := 1000; i > 0; i-- {
		go func() {
			if i%2 == 0 && i > max {
				max = i
			}
		}()
	}

	fmt.Printf("Maximum is %d", max)
}

// что будет выведено?
// как минусы у этого кода?
