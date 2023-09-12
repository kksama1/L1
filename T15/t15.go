package main

// К каким негативным последствиям может привести данный фрагмент кода, и как это
// исправить? Приведите корректный пример реализации.
import (
	"fmt"
)

// Выделяется памяти больше чем используется, v[:100] может превысить длинну слайса.
//panic: runtime error: slice bounds out of range [:100] with length 1
//
//var justString string
//func someFunc() {
//	v := createHugeString(1 << 10)
//	justString = v[:100]
//}
//
//func main() {
//	someFunc()
//}
//

var justString string

func createHugeString(size int) string {
	res := []byte{}
	for i := 0; i < size; i++ {
		res = append(res, 0)
	}
	return string(res)
}

// Не знаю зачем именно нужна функция createHugeString, если можно через make.
func someFunc(size int) (string, string) {
	v := createHugeString(size)
	s := string(make([]byte, size))
	return v, s
}

func main() {
	justStringOne, justStringTwo := someFunc(100)
	fmt.Printf("justStringOne:%s\n", justStringOne)
	fmt.Println(len(justStringOne))
	fmt.Printf("justStringTwo:%s\n", justStringTwo)
	fmt.Println(len(justStringTwo))
}
