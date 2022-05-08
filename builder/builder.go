package builder

type House struct {
	Window string
	Door   string
	Floor  string
}

type HouseBuilder interface {
	SetWindow()
	SetDoor()
	SetFloor()
	GetHouse() House
}

func GetHouseBuilder(builderType string) HouseBuilder {
	if builderType == "wood" {
		return NewWoodHouseBuilder()
	}
	if builderType == "ice" {
		return NewIceHouseBuilder()
	}
	return nil
}

// build wood house
type WoodHouseBuilder struct {
	Window string
	Door   string
	Floor  string
}

func NewWoodHouseBuilder() *WoodHouseBuilder {
	return &WoodHouseBuilder{}
}

func (b *WoodHouseBuilder) SetWindow() {
	b.Window = "wood window"
}

func (b *WoodHouseBuilder) SetDoor() {
	b.Door = "wood door"
}

func (b *WoodHouseBuilder) SetFloor() {
	b.Floor = "wood floor"
}

func (b *WoodHouseBuilder) GetHouse() House {
	return House{
		Door:   b.Door,
		Window: b.Window,
		Floor:  b.Floor,
	}
}

// build ice house
type IceHouseBuilder struct {
	Window string
	Door   string
	Floor  string
}

func NewIceHouseBuilder() *IceHouseBuilder {
	return &IceHouseBuilder{}
}

func (b *IceHouseBuilder) SetWindow() {
	b.Window = "ice window"
}

func (b *IceHouseBuilder) SetDoor() {
	b.Door = "ice door"
}

func (b *IceHouseBuilder) SetFloor() {
	b.Floor = "ice floor"
}

func (b *IceHouseBuilder) GetHouse() House {
	return House{
		Door:   b.Door,
		Window: b.Window,
		Floor:  b.Floor,
	}
}
