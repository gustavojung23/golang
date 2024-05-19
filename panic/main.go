package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Panic / defer / recover")

	//Panic
	_, err := os.Open("f:/settings.txt")

	if err != nil {
		panic(err)
	}

	//Defer
	file, err := os.Create("./settings.txt")
	defer file.Close()
	defer ShowText()

	if err != nil {
		panic(err)
	}

	file.Write([]byte("teste"))

}

func ShowText() {
	fmt.Println("Finalizando de manipular o arquivo...")
}
