package factory

import (
	"fmt"
	"testing"
)

func TestShapeFactory(t *testing.T) {
	shapeFactory := &ShapeFactory{}

	// circle
	circle, err := shapeFactory.GetShape("circle")
	if err != nil {
		fmt.Println(err)
	}
	circle.Draw()

	// square
	square, err := shapeFactory.GetShape("square")
	if err != nil {
		fmt.Println(err)
	}
	square.Draw()
}
