package main

import "fmt"

type Rectangle struct {
	length float64
	width  float64
}

type Shape struct {
	Rectangle
}

// Method with value receiver
func (r Rectangle) Area() float64 {
	return r.length * r.width
}

// Method with pointer receiver
func (r *Rectangle) Scale(factor float64) {
	r.length *= factor
	r.width *= factor
}

type MyInt int

// Method in user-defined types
func (m MyInt) IsPositive() bool {
	return m > 0
}

func (MyInt) welcomeMessage() string {
	return "Welcome to MyInt Type"
}

func main() {
	rect := Rectangle{
		length: 10,
		width:  3,
	}
	area := rect.Area()
	fmt.Println(area)
	rect.Scale(2)
	area = rect.Area()
	fmt.Println(area)

	num := MyInt(-5)
	num1 := MyInt(9)

	fmt.Println(num.IsPositive())
	fmt.Println(num1.IsPositive())

	fmt.Println(num.welcomeMessage())

	s := Shape{rect}

	fmt.Println(s.Rectangle.Area())
	fmt.Println(s.Area())

}
