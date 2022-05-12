package bridge

import "fmt"

type Computer interface {
	Print()
	SetPrinter(Printer)
}

type Mac struct {
	Printer Printer
}

func (m *Mac) Print() {
	fmt.Println("Print request for mac")
	m.Printer.PrintFile()
}

func (m *Mac) SetPrinter(p Printer) {
	m.Printer = p
}

type Windows struct {
	Printer Printer
}

func (w *Windows) Print() {
	fmt.Println("Print request for windows")
	w.Printer.PrintFile()
}

func (w *Windows) SetPrinter(p Printer) {
	w.Printer = p
}
