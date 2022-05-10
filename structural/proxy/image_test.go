package proxy

import "testing"

func TestProxy(t *testing.T) {
	// load from disk
	image := &ProxyImage{Name: "image1"}
	image.Display()
	// use cache
	image.Display()
}
