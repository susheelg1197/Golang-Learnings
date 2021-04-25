/* TO calculate area and perimeter based on shapes */

package main

import "fmt"

type shape interface {
	area() float64
	perimeter() float64
}
type circle struct {
	radius float64
}

func (c *circle) area() float64 {
	return c.radius * 3.14 * c.radius
}
func (c *circle) perimeter() float64 {
	return c.radius * 3.14 * 2
}

type rectangle struct {
	length  float64
	breadth float64
}

func (r *rectangle) area() float64 {
	return r.length * r.breadth
}
func (r *rectangle) perimeter() float64 {
	return 2 * (r.length + r.breadth)
}

func main() {
	c := &circle{
		radius: 3.14,
	}
	r := &rectangle{
		length:  10,
		breadth: 20,
	}
	var s shape

	s = c
	fmt.Println("Area of circle:: ", s.area(), "\nPerimeter of circle:: ", s.perimeter())

	s = r
	fmt.Println("Perimeter of rectangle:: ", s.perimeter(), "\nArea of Rectangle", s.area())
}
