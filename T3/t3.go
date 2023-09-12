package main

// Дана последовательность чисел: 2,4,6,8,10. Найти сумму их
// квадратов(22+32+42….) с использованием конкурентных вычислений.
import (
	"fmt"
	"sync"
)

// SquareSum sums squares of given slice. sync.Mutex allow to block sum for other
// goroutines, so there is no race condition.
func squareSum(number int, sum *int, mutex *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	mutex.Lock()
	*sum += number * number
	mutex.Unlock()
}

func main() {

	numbers := []int{2, 4, 6, 8, 10}
	var (
		mutex sync.Mutex
		sum   int
		wg    sync.WaitGroup
	)
	// Setting wg counter to number of elements in slice.
	wg.Add(len(numbers))

	for _, v := range numbers {
		go squareSum(v, &sum, &mutex, &wg)
	}
	// Blocks main until wg != 0. So main will be executed until other goroutines not done.
	wg.Wait()
	fmt.Println(sum)
}
