package main

import "fmt"

func main() {

	var s Shapes = Circle{12.0}
	fmt.Printf("Shapes Type = %T, Shapes Value = %#v\n", s, s)
	fmt.Printf("Area = %f, Perimeter = %f\n\n", s.Area(), s.Perimeter())

	s = Rectangle{20.0, 30.0}
	fmt.Printf("Shapes Type = %T, Shapes Value = %#v\n", s, s)
	fmt.Printf("Area = %f, Perimeter = %f\n", s.Area(), s.Perimeter())

	PrintMeasurement(Circle{12.0})
	PrintMeasurement(Rectangle{20.0, 30.0})

	shapes := []Shapes{
		Circle{8},
		Rectangle{8, 12},
		Rectangle{3, 9},
		Circle{7},
	}
	totalArea := CalculateTotalArea(shapes...)
	fmt.Printf("Total Area: %f\n\n", totalArea)
}

type Shapes interface {
	Area() float64
	Perimeter() float64
}

func PrintMeasurement(s Shapes) {

	fmt.Printf("Shapes Type = %T, Shapes Value = %#v\n", s, s)
	if c, ok := s.(Circle); ok {
		fmt.Printf("Area = %f, Perimeter = %f, Diameter = %f\n\n", s.Area(), s.Perimeter(), c.Diameter())
	} else {
		fmt.Printf("Area = %f, Perimeter = %f\n\n", s.Area(), s.Perimeter())
	}
}

func CalculateTotalArea(shapes ...Shapes) float64 {

	totalArea := 0.0
	for _, s := range shapes {
		totalArea += s.Area()
	}

	return totalArea
}
