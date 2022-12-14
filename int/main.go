package main

import (
	"fmt"
	"time"
)

// что будет выведено
func main() {
	var i int8
	for {
		time.Sleep(50 * time.Millisecond)
		fmt.Println(i)
		i++
	}
}
