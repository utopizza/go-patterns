package factory

import "fmt"

type Circle struct{}

func (c *Circle) Draw() {
	fmt.Println("Inside Circle::draw() method.")
}
