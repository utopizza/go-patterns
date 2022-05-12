package chain

import "testing"

func TestChain(t *testing.T) {
	cashier := &Cashier{}
	medical := &Medical{}
	doctor := &Doctor{}
	reception := &Reception{}
	reception.SetNext(doctor)
	doctor.SetNext(medical)
	medical.SetNext(cashier)

	patient := &Patient{Name: "abc"}
	reception.Execute(patient)
}
