package main

var ch chan struct{}

func main() {
	ch = make(chan struct{})

	go ready()
	go ready()
	go ready()

	Steady()
	//сделать что бы они начали бежать после вызова фунции, одновременно
}

func Steady() {
	// magic
	close(ch)
}

func ready() {
	// ждем вызова стеди
	<-ch
	run()
}
