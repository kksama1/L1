package main

// Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.
import "fmt"

type emptyStract struct{}

// SetToMany deletes duplicates by using set data as map keys and then uses them as slice values.
func setToMany(set []string) []string {
	many := map[string]emptyStract{}
	var res []string
	for _, v := range set {
		many[v] = emptyStract{}
	}
	for v, _ := range many {
		res = append(res, v)
	}

	return res

}

func main() {
	set := []string{"cat", "cat", "dog", "cat", "tree"}

	fmt.Println(setToMany(set))
}
