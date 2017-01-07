package elevator

import (
	_ "fmt"
)

type direction string

const up direction = "up"
const down direction = "down"

type action string

const open action = "open"
const close action = "close"

type Elevator struct {
	Queue         []action
	Floor         int
	PassengerBays []map[*Passenger]bool
}

func NewElevator(floors int) *Elevator {
	var queue = make([]action, 0, 2)
	var bays = make([]map[*Passenger]bool, floors) // accept # floors param?
	for i := range bays {
		bays[i] = make(map[*Passenger]bool)
	}
	return &Elevator{
		Queue:         queue,
		Floor:         0,
		PassengerBays: bays,
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

		switch current {
		case open:
			for p, _ := range e.PassengerBays[e.Floor] {
				p.enter(e)
			}
		}
	}
}

func (e *Elevator) Call(p *Passenger, d direction) {
	// TODO figure out how to add all necessary actions
	e.Queue = append(e.Queue, open)
	e.Queue = append(e.Queue, close)
	//

	e.addWaitingPassenger(p)
}

func (e *Elevator) addWaitingPassenger(p *Passenger) {
	e.PassengerBays[p.CurrentFloor][p] = true
}

func (e *Elevator) removeWaitingPassenger(p *Passenger) {
	// handle error case?
	bay := e.PassengerBays[p.CurrentFloor]
	delete(bay, p)
}

type Passenger struct {
	CurrentFloor int
	InTransit    bool
}

func NewPassenger() *Passenger {
	return &Passenger{}
}

func (p *Passenger) Call(e *Elevator, d direction) {
	e.Call(p, d)
}

func (p *Passenger) enter(e *Elevator) {
	p.InTransit = true
	e.removeWaitingPassenger(p)
}
