package decorator

type Pizza interface {
	GetPrice() float32
}

type BreadPizza struct{}

func (p *BreadPizza) GetPrice() float32 {
	return 5.5
}

type TomatoTopping struct {
	Pizza Pizza
}

func (p *TomatoTopping) GetPrice() float32 {
	return p.Pizza.GetPrice() + 0.5
}

type CheeseTopping struct {
	Pizza Pizza
}

func (p *CheeseTopping) GetPrice() float32 {
	return p.Pizza.GetPrice() + 1.2
}
