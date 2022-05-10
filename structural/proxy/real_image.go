package proxy

import "fmt"

type ProxyImage struct {
	Name      string
	realImage *RealImage
}

func (i *ProxyImage) Display() {
	if i.realImage == nil {
		fmt.Printf("loading image from disk, name:%+v\n", i.Name)
		i.realImage = &RealImage{Name: i.Name}
	}
	i.realImage.Display()
}
