package main
import (
	"fmt"
	)

type Circle struct {
radius int
}

func main() {
	var c1 Circle
	c1.radius=10
	fmt.Println("Area of circle(c1)=",c1.getArea()) 
}

func (c Circle) getArea() int {
return 3 * c.radius * c.radius
}

