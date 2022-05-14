package strategy

type Strategy interface {
	DoOperation(x int, y int) int
}

type OperationAdd struct{}

func (o *OperationAdd) DoOperation(x int, y int) int {
	return x + y
}

type OperationSubtract struct{}

func (o *OperationSubtract) DoOperation(x int, y int) int {
	return x - y
}

type OperationMultiply struct{}

func (o *OperationMultiply) DoOperation(x int, y int) int {
	return x * y
}
