package main

// Написать программу, которая конкурентно рассчитает значение квадратов чисел
// взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.
import (
	"fmt"
	"sync"
)

// Square outputs squares.. and decrements WaitGroup counter.
func square(num int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("%d*%d = %d\n", num, num, num*num)
}

func main() {
	numbers := []int{2, 4, 6, 8, 10}

	var wg sync.WaitGroup
	// Setting wg counter to number of elements in slice.
	wg.Add(len(numbers))

	for _, v := range numbers {
		go square(v, &wg)
	}
	//Blocks main until wg != 0. So main will be executed until other goroutines not done.
	wg.Wait()
}
