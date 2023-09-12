package main

// Удалить i-ый элемент из слайса.
import (
	"fmt"
)

// DeleteElement takes slices before and after given one and appends them.
func DeleteElement(mass []int, i int) []int {
	return append(mass[:i], mass[i+1:]...)
}

func main() {
	mass := []int{1, 2, 3, 4, 5, 7, 8, 9, 11, 23, 45, 6}
	mass = DeleteElement(mass, 6)
	fmt.Println(mass)
}
