package main

// Реализовать бинарный поиск встроенными методами языка.
import (
	"fmt"
	"math/rand"
)

var foundOrNot bool

func binarySearch(arr []int, lookingFor int) bool {
	if len(arr) > 2 {
		if arr[len(arr)/2] > lookingFor {
			binarySearch(arr[:len(arr)/2-1], lookingFor)
		} else if arr[len(arr)/2] < lookingFor {
			binarySearch(arr[len(arr)/2+1:], lookingFor)
		} else {
			foundOrNot = true
		}
	} else {
		if arr[len(arr)/2] != lookingFor {
			foundOrNot = false
		} else {
			foundOrNot = true
		}
	}

	return foundOrNot
}

// Вариант 2 доделать
//var foundOrNotInexed bool
//
//func binarySearchWithIndex(arr []int, lookingFor int) bool {
//	var prevMid int
//	var index int
//	mid := len(arr) / 2
//	for mid/2 >0 || prevMid-mid >  {
//		if arr[mid] > lookingFor {
//			prevMid = mid
//			mid = mid / 2
//		} else if arr[mid] < lookingFor {
//			prevMid = mid
//			mid = (len(arr) - mid) / 2
//		} else {
//			foundOrNotInexed = true
//		}
//	}
//
//}

func main() {
	var arr []int

	for i := rand.Intn(5); i < 100; i += rand.Intn(4) + 1 {
		arr = append(arr, i)
	}
	for i := range arr {
		fmt.Printf("%d:%d\t", i, arr[i])
	}

	var lookingFor int

	lookingFor = rand.Intn(100)
	//fmt.Println()
	//fmt.Scan(&lookingFor)

	fmt.Println()
	fmt.Println(lookingFor)
	found := binarySearch(arr, lookingFor)

	fmt.Printf("found?:%t", found)

}
