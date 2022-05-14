package state

import "testing"

func TestState(t *testing.T) {
	m := &Machine{}

	start := &StartState{}
	stop := &StopState{}

	start.DoAction(m)
	stop.DoAction(m)
}
