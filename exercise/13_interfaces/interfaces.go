package main

import "fmt"

const (
	SmallLift = iota
	StandardLift
	LargeLift
)

type Lift int

type LiftPicker interface {
	PickLift() Lift
}

type (
	Motorcycle string
	Car        string
	Truck      string
)

func (m Motorcycle) String() string {
	return fmt.Sprintf("Motorcycle: %v", string(m))
}

func (c Car) String() string {
	return fmt.Sprintf("Car: %v", string(c))
}

func (t Truck) String() string {
	return fmt.Sprintf("Truck: %v", string(t))
}

func (m Motorcycle) PickLift() Lift {
	return SmallLift
}

func (c Car) PickLift() Lift {
	return StandardLift
}

func (t Truck) PickLift() Lift {
	return LargeLift
}

func sendToLift(p LiftPicker) {
	switch p.PickLift() {
	case SmallLift:
		fmt.Printf("send %v to small lift\n", p)
	case StandardLift:
		fmt.Printf("send %v to standard lift\n", p)
	case LargeLift:
		fmt.Printf("send %v to large lift\n", p)
	}
}

func main() {
	car := Car("Sporty")
	truck := Truck("MountainCrusher")
	motorcycle := Motorcycle("Croozer")

	sendToLift(car)
	sendToLift(truck)
	sendToLift(motorcycle)
}
