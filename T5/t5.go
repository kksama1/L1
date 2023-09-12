package main

// Разработать программу, которая будет последовательно отправлять значения в
// канал, а с другой стороны канала — читать. По истечению N секунд программа
// должна завершаться
import (
	"fmt"
	"math/rand"
	"time"
)

// Writer sends random number in channel every second.
func writer(c chan int) {
	for {
		c <- rand.Intn(100)
		time.Sleep(1 * time.Second)
	}
}

// Reader appends data from chan to given array.
func reader(c chan int, arr *[]int) {
	for {
		*arr = append(*arr, <-c)
	}
}

func main() {
	c := make(chan int)
	var arr []int
	go writer(c)
	go reader(c, &arr)

	time.Sleep(7 * time.Second)
	close(c)
	//fmt.Println(len(arr))
	fmt.Println(arr)
}
