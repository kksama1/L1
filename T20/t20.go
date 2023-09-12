package main

//Разработать программу, которая переворачивает слова в строке.
//Пример: «snow dog sun — sun dog snow».
import (
	"fmt"
	"strings"
)

// ReverseStr swaps elements from begining with elements at the end.
func reverseStr(str string) string {
	revStr := strings.Split(str, " ")
	j := len(revStr) - 1
	for i := 0; i < len(revStr)/2; i++ {
		revStr[i], revStr[j] = revStr[j], revStr[i]
		j--
	}
	return fmt.Sprint(strings.Join(revStr, " "))
}

func main() {
	str := "Кошка уронила чашку вчера утром в час дня"
	revStr := reverseStr(str)
	fmt.Printf("%s - %s", str, revStr)

}
