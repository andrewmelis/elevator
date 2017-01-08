package elevator

import (
	"fmt"
	"testing"
)

const maxTicks = 10

// CAVEAT: either tests must know how many 'ticks' a given set of actions takes,
//         or the elevator needs to return a list of actions so we can check
// TODO figure this out. consider testing DSL like Clean Code(ch 9)

func TestSingleRider(t *testing.T) {
	var singleRiderTests = []struct {
		nFloors int
		p       *Passenger
	}{
		{2, NewPassenger(0, 1)},
	}

	for _, tt := range singleRiderTests {
		if pass, output := SingleRiderFixture(tt.nFloors, tt.p); !pass {
			t.Errorf(output)
		}
	}
}

// SingleRiderFixture accepts scenarios and returns error string
func SingleRiderFixture(nFloors int, p *Passenger) (bool, string) {
	topFloor := nFloors - 1
	e := NewElevator(nFloors)

	dest := topFloor
	p.Call(e)

	var foundPassengerInTransit bool
	for i := 0; i < maxTicks; i++ {
		// fmt.Printf("bays: %v | queue: %v\n", e.PassengerBays, e.Queue)
		e.Tick() // advance time

		if p.InTransit {
			foundPassengerInTransit = true
		}
	}

	if e.PassengerBays[p.CurrentFloor][p] {
		return false, fmt.Sprintf("passenger %v never left elevator bay %v", p, e)
	}

	if !foundPassengerInTransit {
		return false, fmt.Sprintf("passenger %v never entered elevator %v", p, e)
	}

	if p.InTransit {
		return false, fmt.Sprintf("passenger %v never left elevator %v", p, e)
	}

	if p.CurrentFloor != dest {
		return false, fmt.Sprintf("passenger %v did not arrive at destination %d", p, dest)
	}
	return true, ""
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

// Situation: Rider calls elevator then walks away

// Situation: multiple riders -- serial, same bay, same dest

// Situation: multiple riders -- same bay, same direction

// Situation: multiple riders -- same bay, different directions

// Situation: multiple riders -- different bays, different directions

// all those ^^^, but don't start at level 0
