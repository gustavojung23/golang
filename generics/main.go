package main

import "fmt"

func main() {
	fmt.Println("Generics")

	sliceInt := []int{2, 3, 5, 1}
	sliceString := []string{"A", "C", "D", "B"}

	newSliceInt := reverse(sliceInt)
	newSliceString := reverse(sliceString)

	fmt.Println(newSliceInt)
	fmt.Println(newSliceString)

}

type constraintCustom interface {
	int | string
}

func reverse[T constraintCustom](slice []T) []T {

	newInts := make([]T, len(slice))
	newIntsLen := len(slice) - 1

	for i := 0; i < len(slice); i++ {
		newInts[newIntsLen] = slice[i]
		newIntsLen--
	}
	return newInts
}
