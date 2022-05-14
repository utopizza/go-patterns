package strategy

import (
	"fmt"
	"testing"
)

func TestCalculator(t *testing.T) {
	cal := &Calculator{}
	x, y := 10, 5

	cal.SetStrategy(&OperationAdd{})
	fmt.Printf("%d+%d=%d\n", x, y, cal.Execute(x, y))

	cal.SetStrategy(&OperationSubtract{})
	fmt.Printf("%d-%d=%d\n", x, y, cal.Execute(x, y))

	cal.SetStrategy(&OperationMultiply{})
	fmt.Printf("%d*%d=%d\n", x, y, cal.Execute(x, y))
}
