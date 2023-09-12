package main

// Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0,
// 15.5, 24.5, -21.0, 32.5. Объединить данные значения в группы с шагом в 10
// градусов. Последовательность в подмножноствах не важна.
import (
	"fmt"
)

//func sortTemp(temp float32, sortedTemp map[int][]float32, mutex *sync.Mutex) {
//	group := int(temp) / 10 * 10
//	mutex.Lock()
//	sortedTemp[group] = append(sortedTemp[group], temp)
//	mutex.Unlock()
//}

func main() {

	//var mutex sync.Mutex
	temp := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	sortedTemp := map[int][]float32{}

	//for _, t := range temp {
	//	go sortTemp(t, sortedTemp, &mutex)
	//}

	for _, t := range temp {
		group := int(t) / 10 * 10
		sortedTemp[group] = append(sortedTemp[group], t)
	}

	fmt.Println(sortedTemp)
}
