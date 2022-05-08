package singleton

import (
	"fmt"
	"testing"
)

func TestSinglton(t *testing.T) {
	for i := 0; i < 10; i++ {
		go GetSingleInstance()
	}
	fmt.Scanln()
}
