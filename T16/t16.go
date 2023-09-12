package main

// Реализовать быструю сортировку массива (quicksort) встроенными методами языка.
import (
	"fmt"
	"math/rand"
)

// Если я правильно понял смысл в том чтобы разбить элементы массива на 2 подмассива больше или меньше разделителя.
// Повторять этот этап рекурсивно, с каждым из подмассивов. Если в массиве 1 элемент, то он отсортированный.
//
// Quicksort takes first elem as pivot and divide rest into 2 groups <= OR > then repeat in recursion with received
// arrays until there is less than 2 elems. After that it appends lesser -> pivot -> greater in recursion so at the end
// we get sorted array.
func quicksort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	pivot := arr[0]
	var less, greater []int

	for _, num := range arr[1:] {
		if num <= pivot {
			less = append(less, num)
		} else {
			greater = append(greater, num)
		}
	}

	sortedLess := quicksort(less)
	sortedGreater := quicksort(greater)
	// sortedGreater... Т.к. аппенд поддерживает переменное количество аргументов
	//https://gobyexample.com/variadic-functions
	return append(append(sortedLess, pivot), sortedGreater...)
}

func main() {
	var arr []int
	//arr := []int{2, 1}

	for i := 0; i < rand.Intn(2)+20; i++ {
		arr = append(arr, rand.Intn(30))
	}

	fmt.Println(arr)
	newArr := quicksort(arr)
	fmt.Printf("sortedArr: %v", newArr)

}
