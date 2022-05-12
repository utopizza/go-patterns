package adapter

import "fmt"

type I interface {
	Func()
}

type A struct{}

func (a *A) Func() {
	fmt.Println("a func")
}

type B struct{}

func (b *B) FuncB() {
	fmt.Println("b func")
}

// 实现I类型，底层调用的是b·
type BAdapter struct {
	b *B
}

func (a *BAdapter) Func() {
	a.b.FuncB()
}

func DoSomething(x I) {
	x.Func()
}
