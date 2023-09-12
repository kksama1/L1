package main

// Разработать программу, которая переворачивает подаваемую на ход строку
// (например: «главрыба — абырвалг»). Символы могут быть unicode.
import "fmt"

// Reverse converts string to rune then appends it in opposite order to rev byte slice.
func reverse(str string) string {
	orig := []rune(str)
	rev := make([]rune, 0)
	fmt.Println(orig)

	for i := len(orig) - 1; i >= 0; i-- {
		rev = append(rev, orig[i])
	}

	fmt.Println(rev)
	return string(rev)
}

func main() {
	str := "Кошка уронила чашку"
	revStr := reverse(str)
	fmt.Printf("%s - %s", str, revStr)
}
