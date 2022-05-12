package adapter

import "testing"

func TestAdapter(t *testing.T) {
	a := &A{}
	DoSomething(a)

	b := &B{}
	//DoSomething(b)
	DoSomething(&BAdapter{b})
}
