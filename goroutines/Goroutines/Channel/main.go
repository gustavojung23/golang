package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Channels")

	channel := make(chan int, 100)

	go setList(channel)

	for v := range channel {
		fmt.Println("Recebendo:", v)
		time.Sleep(time.Second)
	}
}

func setList(channel chan int) {
	for i := 0; i < 100; i++ {
		channel <- i
		fmt.Println("Enviando:", i)
		time.Sleep(time.Second)
	}
	close(channel)
}

// <-chan - leitura
// chan<- escrita
//channel := make(chan int, 100) - channel buffered
