package main

//Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.
import (
	"fmt"
	"strconv"
)

// SetBit converts int64 to byte slice and then changes given bit as required.
func setBit(val int64, pos, bit int) (result int64) {
	helpVar := []rune(strconv.FormatInt(val, 2))
	bitToSet := len(helpVar) - pos - 1 // Если считать биты слева то нужно взять только позицию
	if bit == 1 {
		helpVar[bitToSet] = '1'
	} else if bit == 0 {
		helpVar[bitToSet] = '0'
	}
	helpString := string(helpVar)
	result, _ = strconv.ParseInt(helpString, 2, 64)
	return result
}

func main() {
	var (
		numb int64
		pos  int
		bit  int
	)
	fmt.Scan(&numb)
	fmt.Printf("Input numb in binary system:\n%v\n", strconv.FormatInt(numb, 2))
	fmt.Scan(&pos, &bit)
	fmt.Printf("Output to compare\n%v\n", strconv.FormatInt(numb, 2))
	numb = setBit(numb, pos, bit)
	fmt.Printf("%v\n", strconv.FormatInt(numb, 2))
	fmt.Printf("result in decimal system %v", numb)
}
