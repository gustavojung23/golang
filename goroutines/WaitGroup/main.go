package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(3)

	callDataBase(&wg)
	callAPI(&wg)
	processInternal(&wg)

	wg.Wait()
}

func callDataBase(wg *sync.WaitGroup) {
	fmt.Println("Finalizado callDataBase")
	wg.Done()
}

func callAPI(wg *sync.WaitGroup) {
	fmt.Println("Finalizado callAPI")
	wg.Done()
}

func processInternal(wg *sync.WaitGroup) {
	fmt.Println("Finalizado processInternal")
	wg.Done()
}
