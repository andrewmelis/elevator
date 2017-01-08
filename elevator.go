package elevator

import (
	_ "fmt"
	"math"
)

type direction string

const up direction = "up"
const down direction = "down"

type action string

const open action = "open" // turn these into types that implement command interface
const close action = "close"
const travelUp action = "travelUp"

type Elevator struct {
	Queue         []action
	Floor         int
	PassengerBays []map[*Passenger]bool // one waiting area per floor -- extract me
	Riders        map[*Passenger]bool   // extract me. should be callback for riding passengers
}

func NewElevator(floors int) *Elevator {
	var queue = make([]action, 0, 2)
	var bays = make([]map[*Passenger]bool, floors)
	for i := range bays {
		bays[i] = make(map[*Passenger]bool)
	}
	var riders = make(map[*Passenger]bool)

	return &Elevator{
		Queue:         queue,
		Floor:         0,
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
		case travelUp:
			e.Floor++

		}
	}
}

func (e *Elevator) Call(p *Passenger, d direction) {
	// TODO figure out how to add all necessary actions
	e.Queue = append(e.Queue, open)
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
	// add necessary "travel" actions to go from currentFloor to input floor
	// TODO start with one tick per floor
	diff := math.Abs(float64(destination - e.Floor))

	if diff > 0 {
		e.Queue = append(e.Queue, close) // eww -- how handle open, then no button press?

		for i := 0; i < int(diff); i++ {
			e.Queue = append(e.Queue, travelUp)
		}

		// close? open?
		// just hard code open for now
		e.Queue = append(e.Queue, open)
	}
}

type Passenger struct {
	CurrentFloor int
	Destination  int
	InTransit    bool
}

func NewPassenger(destination int) *Passenger {
	return &Passenger{Destination: destination}
}

func (p *Passenger) Call(e *Elevator, d direction) {
	e.Call(p, d)
}

func (p *Passenger) enter(e *Elevator) {
	p.InTransit = true
	e.removeWaitingPassenger(p) // elevator + bays should be separate
	e.addRidingPassenger(p)
	p.inputDestination(e)
}

func (p *Passenger) inputDestination(e *Elevator) {
	e.enqueueDestination(p.Destination)
}

func (p *Passenger) exit(e *Elevator) {
	if p.shouldExit(e) {
		p.InTransit = false
		p.CurrentFloor = e.Floor
		e.removeRidingPassenger(p) // again, don't have elevator do bookkeeping
	}
}

func (p *Passenger) shouldExit(e *Elevator) bool {
	return p.Destination == e.Floor
}
