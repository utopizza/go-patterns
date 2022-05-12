package bridge

import "fmt"

type Printer interface {
	PrintFile()
}

type PrinterA struct {
}

func (p *PrinterA) PrintFile() {
	fmt.Println("Printing by PrinterA")
}

type PrinterB struct {
}

func (p *PrinterB) PrintFile() {
	fmt.Println("Printing by PrinterB")
}
