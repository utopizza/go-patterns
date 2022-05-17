package observer

import "testing"

func TestObserver(t *testing.T) {
	shoes := &Item{Name: "nike"}

	customer1 := &Customer{ID: "xiaoming@qq.com"}
	customer2 := &Customer{ID: "xiaohong@qq.com"}

	shoes.Register(customer1)
	shoes.Register(customer2)

	shoes.NotifyAll()

	shoes.Deregister(customer1)
	shoes.NotifyAll()
}
