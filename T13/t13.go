package main

// Поменять местами два числа без создания временной переменной.
import "fmt"

// https://go.dev/ref/spec#Assignments The assignment proceeds in two phases.
// First, the operands of index expressions and pointer indirections (including
// implicit pointer indirections in selectors) on the left and the expressions on
// the right are all evaluated in the usual order. Second, the assignments are
// carried out in left-to-right order.
func main() {
	var num1, num2 int

	num1 = 1
	num2 = 2
	fmt.Printf("num1 = %d num2 = %d\n", num1, num2)
	fmt.Printf("num1 = %v num2 = %v\n", &num1, &num2)
	num1, num2 = num2, num1
	fmt.Printf("num1 = %d num2 = %d\n", num1, num2)
	fmt.Printf("num1 = %v num2 = %v\n", &num1, &num2)

}
