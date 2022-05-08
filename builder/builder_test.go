package builder

import (
	"fmt"
	"testing"
)

func TestBuilder(t *testing.T) {
	woodHouseBuilder := GetHouseBuilder("wood")
	iceHouseBuilder := GetHouseBuilder("ice")

	director := NewDirector(woodHouseBuilder)
	house := director.BuildHouse()
	fmt.Printf("house:%+v\n", house)

	director.SetBuilder(iceHouseBuilder)
	house = director.BuildHouse()
	fmt.Printf("house:%+v\n", house)
}
