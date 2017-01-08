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
	Riders        map[*Passenger]bool   // extract me. should be callback for riding passengers -- something like ElevatorController or ElevatorSimulator?
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
		default:
			// noop
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
