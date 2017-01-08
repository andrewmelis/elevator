package elevator

import (
	"fmt"
	"testing"
)

const maxTicks = 10

// CAVEAT: either tests must know how many 'ticks' a given set of actions takes,
//         or the elevator needs to return a list of actions so we can check
// TODO figure this out. consider testing DSL like Clean Code (ch 9)

func TestSingleRider(t *testing.T) {
	var singleRiderTests = []struct {
		e Elevator
		p Passenger
	}{
		{*NewElevator(0, 2), *NewPassenger(0, 1)},
		{*NewElevator(0, 2), *NewPassenger(1, 0)},
		{*NewElevator(1, 2), *NewPassenger(0, 1)},
	}

	for _, tt := range singleRiderTests {
		t.Logf("=== RUN   %+v\n", tt)
		pass, output := SingleRiderFixture(&tt.e, &tt.p)
		if !pass {
			t.Errorf("--- FAIL: %+v! %s\n", tt, output)
		} else {
			t.Logf("--- PASS: %+v\n", tt)
		}
	}
}

// SingleRiderFixture accepts scenarios and returns error string
func SingleRiderFixture(e *Elevator, p *Passenger) (bool, string) {
	p.Call(e)

	var foundPassengerInTransit bool
	for i := 0; i < maxTicks; i++ {
		fmt.Printf("bays: %v | riders: %v | queue: %v\n", e.PassengerBays, e.Riders, e.Queue)
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

	if p.CurrentFloor != p.Destination {
		return false, fmt.Sprintf("passenger %v did not arrive at destination %d", p, p.Destination)
	}
	return true, ""
}

// Situation: Rider calls elevator then walks away

// Situation: multiple riders -- serial, same bay, same dest

// Situation: multiple riders -- same bay, same direction

// Situation: multiple riders -- same bay, different directions

// Situation: multiple riders -- different bays, different directions

// all those ^^^, but don't start at level 0
