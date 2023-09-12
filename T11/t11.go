package main

// Реализовать пересечение двух неупорядоченных множеств.
import (
	"fmt"
	"math/rand"
)

type emptyStract struct{}

// FindIntersection takes first many(?) and puts it in map, then it checks if
// keys which are array elements already exists if so puts it in intersection many.
func findIntersection(first, second []int) []int {
	buffer := map[int]emptyStract{}
	intersection := []int{}
	for _, v := range first {
		buffer[v] = emptyStract{}
	}

	for _, v := range second {
		_, exist := buffer[v]
		if exist {
			intersection = append(intersection, v)

		}
	}
	return intersection
}

func main() {
	first := []int{}
	second := []int{}

	// По идее в множестве не могут повторятся элементы, при подобной генерации
	// множеств элементы могут повторяться, но т.к. это тестовые данные, то, надеюсь,
	// такой вариант тоже пойдёт.
	// Если брать набор данных с повторениями, а не множества, то можно в функцию добавить
	// вторую мапу и писать пересечение в неё и потом уже в слайс, потому что если во
	// втором наборе число повторяется дважды и оно есть в первом, то оно выводится дважды.
	for i := 0; i < 7; i++ {
		first = append(first, rand.Intn(15))
	}
	for i := 0; i < 5; i++ {
		second = append(second, rand.Intn(15))
	}

	fmt.Println(first)
	fmt.Println(second)

	fmt.Printf("Пересечение двух множеств %v", findIntersection(first, second))
}
