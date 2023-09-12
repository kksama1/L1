package main

// Разработать программу, которая в рантайме способна определить тип переменной: int, string, bool, channel из переменной типа interface{}.
import "fmt"

// Т.к. пустой интерфейс может содержать значения любого типа, то таким образом можно унифицировать аргумент функции.
// %T	a Go-syntax representation of the type of the value
func describeType(i interface{}) {
	fmt.Printf("It is %T\n", i)
}

// Ещё такую штуку видел

func describe(i interface{}) {
	switch val := i.(type) {
	case int:
		fmt.Println("int")

	// Default не отличается от первого решения, но такая штука может быть полезна,
	// когда надо решить, что делать в зависимости от типа переменной
	default:
		fmt.Printf("%T\n", val)
	}

}

type st struct {
	a int
	b string
}

func main() {
	var i int
	var f float64
	var s string
	var b bool
	var c = make(chan int)
	var st = st{1, "hi"}
	var stSt = struct{}{}

	describeType(i)
	describeType(f)
	describeType(s)
	describeType(b)
	describeType(c)
	describeType(st)
	describeType(st.a)
	describeType(st.b)
	describeType(stSt)

	fmt.Printf("Вариант 2\n")
	describe(i)
	describe(f)
	describe(s)
	describe(b)
	describe(c)
	describe(st)
	describe(stSt)
	describe(st.a)
	describe(st.b)
}
