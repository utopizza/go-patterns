package factory

import "fmt"

type Rectangle struct{}

func (r *Rectangle) Draw() {
	fmt.Println("Inside Rectangle::draw() method.")
}
