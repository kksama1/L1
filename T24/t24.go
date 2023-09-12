package main

// Разработать программу нахождения расстояния между двумя точками, которые
// представлены в виде структуры Point с инкапсулированными параметрами x,y и
// конструктором.
import (
	"fmt"
	"math"
)

type Point struct {
	x int
	y int
}

// конструктор
func CreatePoint(x int, y int) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

func (pnt_1 *Point) CalcDistance(pnt_2 Point) float64 {
	x := pnt_1.x - pnt_2.x
	y := pnt_1.y - pnt_2.y

	return math.Sqrt(float64(x*x + y*y))
}

func main() {
	first_p := CreatePoint(0, 0)
	second_p := CreatePoint(10, 5)
	distance := first_p.CalcDistance(*second_p)
	fmt.Println("Расстояние между точками: ", distance)
}
