package main

// Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной
// среде. По завершению программа должна выводить итоговое значение счетчика.
import (
	"fmt"
	"sync"
)

type counter struct {
	mutex sync.Mutex
	val   int
}

// mutex used to escape race condition.
func (c *counter) increment() {
	c.mutex.Lock()
	c.val += 1
	c.mutex.Unlock()
}
func counterInit() *counter {
	c := counter{
		val: 0,
	}
	return &c
}

func main() {
	var wg sync.WaitGroup
	cnt := counterInit()
	//cnt := counter{val: 0}

	//incAmm := rand.Intn(15) + 5
	incAmm := 10000
	for i := 0; i < incAmm; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			cnt.increment()
		}()
	}

	wg.Wait()
	fmt.Println(cnt.val)
}
