package strategy

type Calculator struct {
	s Strategy
}

func (c *Calculator) SetStrategy(s Strategy) {
	c.s = s
}

func (c *Calculator) Execute(x int, y int) int {
	return c.s.DoOperation(x, y)
}
