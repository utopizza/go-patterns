package template

import "fmt"

type Game interface {
	Init()
	StartPlay()
	EndPlay()
}

type GameA struct {
	Name string
}

func (g *GameA) Init() {
	fmt.Printf("init GameA:%s\n", g.Name)
}

func (g *GameA) StartPlay() {
	fmt.Printf("start GameA:%s\n", g.Name)
}

func (g *GameA) EndPlay() {
	fmt.Printf("end GameA:%s\n", g.Name)
}

type GameB struct {
	Name string
}

func (g *GameB) Init() {
	fmt.Printf("init GameB:%s\n", g.Name)
}

func (g *GameB) StartPlay() {
	fmt.Printf("start GameB:%s\n", g.Name)
}

func (g *GameB) EndPlay() {
	fmt.Printf("end GameB:%s\n", g.Name)
}

type GamePlayer struct {
	Game Game
}

func (p *GamePlayer) Play() {
	p.Game.Init()
	p.Game.StartPlay()
	p.Game.EndPlay()
}
