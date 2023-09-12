package main

// Разработать программу, которая проверяет, что все символы в строке уникальные
// (true — если уникальные, false etc). Функция проверки должна быть
// регистронезависимой.

import (
	"fmt"
	"strings"
)

// Если элементы не стоят на одной позиции и не равны выводит false. Если Символыв совпали то возвразает true.
func UniqueString(mass []string) bool {

	for i, _ := range mass {
		for j, _ := range mass {
			if (mass[i] == mass[j]) && (i != j) {
				return false
			}
		}
	}
	return true
}

func main() {
	var str string
	fmt.Println("Введите строку для проверки уникальности:")
	fmt.Scan(&str)
	mass := strings.Split(strings.ToLower(str), "")
	fmt.Println(UniqueString(mass))
}
