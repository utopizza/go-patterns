package state

import "fmt"

type State interface {
	DoAction(m *Machine)
}

type StartState struct{}

func (s *StartState) DoAction(m *Machine) {
	fmt.Println("start state")
	m.SetState(s)
}

type StopState struct{}

func (s *StopState) DoAction(m *Machine) {
	fmt.Println("stop state")
	m.SetState(s)
}
