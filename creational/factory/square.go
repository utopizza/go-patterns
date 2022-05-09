package factory

import "fmt"

type Square struct{}

func (s *Square) Draw() {
	fmt.Println("Inside Square::draw() method.")
}
