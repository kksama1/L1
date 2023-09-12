package main

// Реализовать конкурентную запись данных в map.
import (
	"fmt"
	"math/rand"
	"sync"
)

var i int

// writeToMap writes rand value to map. Using mutex to escape race condition.
func writeToMap(wg *sync.WaitGroup, mutex *sync.Mutex, mapp map[int]int) {
	mutex.Lock()
	i += 1
	mapp[i] = rand.Intn(500)
	mutex.Unlock()
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	var mutex sync.Mutex
	mapp := map[int]int{}

	for i := 0; i < 15; i++ {
		wg.Add(1)
		go writeToMap(&wg, &mutex, mapp)
	}

	wg.Wait()

	fmt.Println(mapp)

}
