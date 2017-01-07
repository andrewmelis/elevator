package elevator

import "testing"
import "fmt"

const maxTicks = 10

// Situation:
// - single rider at floor 0 (ground).
// - elevator idle, empty at floor 0
// - rider going up to floor n; n > 0
func TestElevatorOpensOnBottomFloor(t *testing.T) {
	e := NewElevator(1)
	p := NewPassenger()

	dest := 0 // CHANGE TO HIGHER NUMBER LATER
	p.Call(e, up)

	var foundPassengerInTransit bool
	for i := 0; i < maxTicks; i++ {
		fmt.Printf("bays: %v | queue: %v\n", e.PassengerBays, e.Queue)
		e.Tick() // advance time

		if p.InTransit {
			foundPassengerInTransit = true
		}
	}

	if e.PassengerBays[p.CurrentFloor][p] {
		t.Errorf("passenger %v never left elevator bay %v", p, e)
	}

	if !foundPassengerInTransit {
		t.Errorf("passenger %v never entered elevator %v", p, e)
	}

	if p.InTransit {
		t.Errorf("passenger %v never left elevator %v", p, e)
	}

	if p.CurrentFloor != dest {
		t.Errorf("passenger %v did not arrive at destination %d", p, dest)
	}
}

// Situation:
// - single rider at floor 0 (ground).
// - elevator idle, empty at floor n; n > 0
// - rider going up to floor n; n > 0
//
// CAVEAT: either tests must know how many 'ticks' a given set of actions takes,
//         or the elevator needs to return a list of actions so we can check
// func TestElevatorToBottomFloorFromFloorGreaterThan0(t *testing.T) {
// 	e := NewElevator(3)
// 	p := NewPassenger()

// 	dest := 2 // CHANGE TO HIGHER NUMBER LATER
// 	p.Call(e, up)

// 	var foundPassengerInTransit bool
// 	for i := 0; i < maxTicks; i++ {
// 		e.Tick() // advance time

// 		if p.InTransit {
// 			foundPassengerInTransit = true
// 		}
// 	}
// 	if !foundPassengerInTransit {
// 		t.Errorf("passenger %v never entered elevator %v", p, e)
// 	}

// 	if p.InTransit {
// 		t.Errorf("passenger %v never left elevator %v", p, e)
// 	}

// 	if p.CurrentFloor != dest {
// 		t.Errorf("passenger %v did not arrive at destination %d", p, dest)
// 	}
// }