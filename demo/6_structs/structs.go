package main

import "fmt"

type Passenger struct {
	Name         string
	TicketNumber int
	Boarder      bool
}

type Bus struct {
	FrontSeat Passenger
}

func main() {
	casey := Passenger{
		"casey",
		1,
		false,
	}

	fmt.Println(casey)

	var (
		bill = Passenger{Name: "Bill", TicketNumber: 2}
		ella = Passenger{Name: "Ella", TicketNumber: 3}
	)

	fmt.Println(bill, ella)

	var heidi Passenger
	heidi.Name = "Heidi"
	heidi.TicketNumber = 4
	fmt.Println(heidi)

	casey.Boarder = true
	bill.Boarder = true

	if bill.Boarder {
		fmt.Println("Bill has boarded the bus")
	}

	if casey.Boarder {
		fmt.Println(casey.Name, "has boarded the bus")
	}

	heidi.Boarder = true
	bus := Bus{heidi}
	fmt.Println(bus)
	fmt.Println(bus.FrontSeat.Name, "is in the front seat")
}
