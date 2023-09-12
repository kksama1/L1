package main

// Дана структура Human (с произвольным набором полей и методов). Реализовать
// встраивание методов в структуре Action от родительской структуры Human (аналог
// наследования).
// https://golangify.com/composition-and-forwarding
import "fmt"

type Human struct {
	age     int
	name    string
	surname string
}

type Action struct {
	human Human
}

func (h Human) Greet() {
	fmt.Printf("Hi, I'm %s %s! I'm %d age old!", h.name, h.surname, h.age)
}

func main() {
	humanOne := Human{42, "John", "White"}
	act := Action{humanOne}
	act.human.Greet()
}
