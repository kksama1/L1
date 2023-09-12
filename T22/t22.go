package main

// Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a,b, значение которых > 2^20.
import (
	"fmt"
	"log"
	"math/big"
)

func main() {
	a := new(big.Int)
	// Sscan scans the argument string, storing successive space-separated values into successive arguments.
	_, err := fmt.Sscan("78446744073709551617", a)
	if err != nil {
		log.Println("error scanning value:", err)
	}
	b := new(big.Int)
	_, err = fmt.Sscan("18446231273709551617", b)
	if err != nil {
		log.Println("error scanning value:", err)
	}

	res := new(big.Int)
	fmt.Printf("a*b = %v\n", res.Mul(a, b))
	fmt.Printf("a+b = %v\n", res.Add(a, b))
	fmt.Printf("a/b = %v\n", res.Div(a, b))
	fmt.Printf("a-b = %v\n", res.Sub(a, b))

}
