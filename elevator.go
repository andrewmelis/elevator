package elevator

import (
	"fmt"
	"math"
)

type direction string

const up direction = "up"
const down direction = "down"

type action string

const open action = "open" // turn these into types that implement command interface
const close action = "close"
const travelUp action = "travelUp"
const travelDown action = "travelDown"

type Elevator struct {
	Floor         int
	Queue         []action
	PassengerBays []map[*Passenger]bool // one waiting area per floor -- extract me
	Riders        map[*Passenger]bool   // extract me. should be callback for riding passengers -- something like ElevatorController or ElevatorSimulator?
}

func NewElevator(currentFloor, floors int) *Elevator {
	var queue = make([]action, 0, 2)
	var bays = make([]map[*Passenger]bool, floors)
	for i := range bays {
		bays[i] = make(map[*Passenger]bool)
	}
	var riders = make(map[*Passenger]bool)

	return &Elevator{
		Floor:         currentFloor,
		Queue:         queue,
		PassengerBays: bays,
		Riders:        riders,
	}
}

func (e *Elevator) Tick() {
	if len(e.Queue) > 0 {
		// basically, shift + handle empty case
		current := e.Queue[0]
		if len(e.Queue) > 1 {
			e.Queue = e.Queue[1:]
		} else {
			e.Queue = e.Queue[:0] // better way to do this?
		}
		// end extract shift()

		// make this ElevatorAction.perform()
		switch current {
		case open:
			e.allowRidersToExit()
			e.waitingPassengersEnter()
		case close:
			// noop
		case travelUp:
			e.Floor++
		case travelDown:
			e.Floor--
		default:
			// noop
		}
	}
}

func (e *Elevator) Call(p *Passenger, d direction) {
	e.enqueueDestination(p.CurrentFloor)

	e.addWaitingPassenger(p)
}

func (e *Elevator) addRidingPassenger(p *Passenger) {
	e.Riders[p] = true
}

func (e *Elevator) removeRidingPassenger(p *Passenger) {
	delete(e.Riders, p)
}

func (e *Elevator) allowRidersToExit() {
	for p := range e.Riders {
		p.exit(e)
	}
}

func (e *Elevator) addWaitingPassenger(p *Passenger) {
	e.PassengerBays[p.CurrentFloor][p] = true
}

func (e *Elevator) removeWaitingPassenger(p *Passenger) {
	// handle error case?
	delete(e.PassengerBays[p.CurrentFloor], p)
}

func (e *Elevator) waitingPassengersEnter() {
	for p, _ := range e.PassengerBays[e.Floor] {
		p.enter(e)
	}
}

func (e *Elevator) enqueueDestination(destination int) {
	diff := math.Abs(float64(destination - e.Floor))

	if diff > 0 {
		e.Queue = append(e.Queue, close) // eww -- no need to close if already closed?...

		var a action
		dir := tripDirection(e.Floor, destination)
		switch dir {
		case up:
			a = travelUp
		case down:
			a = travelDown
		}

		for i := 0; i < int(diff); i++ {
			e.Queue = append(e.Queue, a)
		}

	}
	e.Queue = append(e.Queue, open) // always need to open
}

// direction is a helper that returns a direction from trip parameters
// extract me to a helpers file
func tripDirection(current, destination int) direction {
	var d direction
	if current < destination {
		d = up
	} else if current > destination {
		d = down
	} else {
		// noop
		panic(fmt.Sprintf("invalid trip params! current: %d, dest: %d", current, destination))
	}
	return d
}
