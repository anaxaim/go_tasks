package main

import (
	"fmt"
)

func main() {

	done := make(chan string)
	go ready(done)
	go ready(done)
	go ready(done)
	Steady()
	go func() {
		close(done)
	}()
	for c := range done {
		fmt.Println(c)
	}
}

func Steady() {
	fmt.Println("steady magic")
}

func run() string {
	return "run ready"
}

func ready(done chan<- string) {
	done <- run()
}
