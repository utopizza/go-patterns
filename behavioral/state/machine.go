package state

type Machine struct {
	s State
}

func (m *Machine) SetState(s State) {
	m.s = s
}

func (m *Machine) GetState() State {
	return m.s
}
