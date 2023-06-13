package main

import "fmt"

const (
	Monday    = 0
	Tuesday   = 0
	Wednesday = 0
	Thursday  = 0
	Friday    = 0
	Saturday  = 0
	Sunday    = 0
)

const (
	Admin      = 10
	Manager    = 20
	Contractor = 30
	Guest      = 40
)

func weekday(day int) bool {
	return day <= 4
}

func accessGranted() {
	fmt.Println("Granted")
}

func accessDenied() {
	fmt.Println("Denied")
}

func main() {
	today, role := Tuesday, Guest

	if role == Admin || role == Manager {
		accessGranted()
	} else if role == Contractor && !weekday(today) {
		accessGranted()
	} else if role == Manager && weekday(today) {
		accessGranted()
	} else if role == Guest && (today == Monday || today == Wednesday || today == Friday) {
		accessGranted()
	} else {
		accessDenied()
	}
}
