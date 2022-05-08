package factory

import "errors"

type Shape interface {
	Draw()
}

type ShapeFactory struct{}

func (f *ShapeFactory) GetShape(shapeType string) (Shape, error) {
	if shapeType == "circle" {
		return &Circle{}, nil
	} else if shapeType == "square" {
		return &Square{}, nil
	} else if shapeType == "rectangle" {
		return &Rectangle{}, nil
	} else {
		return nil, errors.New("unknown shape type")
	}
}
