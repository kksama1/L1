package main

// Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из
// массива, во второй — результат операции x*2, после чего данные из второго
// канала должны выводиться в stdout.
import (
	"fmt"
	"sync"
)

// https://golangify.com/goroutines#multi

// Reader sends array data to c channel.
func reader(arr []int, downstream chan int) {
	for _, v := range arr {
		downstream <- v
	}
}

// Squarer reads data from c channel, squares it, and sends it to squares channel.
func squarer(upstream, downstream chan int) {
	for {
		val, opened := <-upstream
		if opened == true {
			//fmt.Printf("IN squarer %d", val)
			downstream <- val * val
		}
	}
}

// PrintSquares reads data from squares channel and sends it to stdout.
func printSquares(upstream chan int, wg *sync.WaitGroup) {
	for {

		val, opened := <-upstream
		if opened == true {
			fmt.Printf("%d\n", val)
			wg.Done()
		}
	}

}

func main() {
	arr := []int{2, 4, 6, 8, 10}
	var c = make(chan int)
	var squares = make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(arr))
	go reader(arr, c)
	go squarer(c, squares)
	go printSquares(squares, &wg)

	wg.Wait()
	close(c)
	close(squares)
}
