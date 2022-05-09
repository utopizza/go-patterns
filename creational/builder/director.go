package builder

type Director struct {
	Builder HouseBuilder
}

func NewDirector(b HouseBuilder) *Director {
	return &Director{
		Builder: b,
	}
}

func (d *Director) SetBuilder(b HouseBuilder) {
	d.Builder = b
}

func (d *Director) BuildHouse() House {
	d.Builder.SetWindow()
	d.Builder.SetDoor()
	d.Builder.SetWindow()
	return d.Builder.GetHouse()
}
