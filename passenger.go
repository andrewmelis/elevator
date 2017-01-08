package elevator

import (
	_ "fmt"
)

type Passenger struct {
	CurrentFloor int
	Destination  int
	InTransit    bool
}

func NewPassenger(currentFloor, destination int) *Passenger {
	return &Passenger{currentFloor, destination, false}
}

func (p *Passenger) Call(e *Elevator) {
	d := tripDirection(p.CurrentFloor, p.Destination)
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
