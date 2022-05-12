package chain

type Department interface {
	Execute(p *Patient)
	SetNext(d Department)
}
