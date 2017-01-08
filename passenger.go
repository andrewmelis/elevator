package elevator

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
