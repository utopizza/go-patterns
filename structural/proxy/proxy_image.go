package proxy

import "fmt"

type RealImage struct {
	Name string
}

func (i *RealImage) Display() {
	fmt.Printf("displaying %+v\n", i.Name)
}
