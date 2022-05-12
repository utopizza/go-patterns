package bridge

import (
	"testing"
)

func TestComputer(t *testing.T) {
	printerA := &PrinterA{}
	printerB := &PrinterB{}

	macComputer := &Mac{}
	macComputer.SetPrinter(printerA)
	macComputer.Print()
	macComputer.SetPrinter(printerB)
	macComputer.Print()

	winComputer := &Windows{}
	winComputer.SetPrinter(printerA)
	winComputer.Print()
	winComputer.SetPrinter(printerB)
	winComputer.Print()
}
