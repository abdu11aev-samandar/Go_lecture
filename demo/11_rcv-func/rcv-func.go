package main

import "fmt"

type Space struct {
	occupied bool
}

type ParkingLot struct {
	space []Space
}

func occupySpace(lot *ParkingLot, spaceNum int) {
	lot.space[spaceNum-1].occupied = true
}

func (lot *ParkingLot) occupySpace(spaceNum int) {
	lot.space[spaceNum-1].occupied = true
}

func (lot *ParkingLot) vacateSpace(spaceNum int) {
	lot.space[spaceNum-1].occupied = true
}

func main() {
	lot := ParkingLot{space: make([]Space, 10)}
	fmt.Println("Initial:", lot)

	lot.occupySpace(1)
	occupySpace(&lot, 2)
	fmt.Println("After occupied:", lot)

	lot.vacateSpace(1)
	fmt.Println("After vacate:", lot)
}
