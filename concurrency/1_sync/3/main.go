package main

import (
	"fmt"
	"time"
)

const numRequests = 10000

var count int32

func networkRequest() {
	time.Sleep(time.Millisecond) // Эмуляция сетевого запроса.
	count++
}

func main() {

	for i := 0; i < numRequests; i++ {
		networkRequest()
	}

	fmt.Println(count)
}

// что выведется в итоге?
// сколько будет выполняться и как можно ускорить?
