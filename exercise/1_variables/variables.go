package main

import "fmt"

func main() {
	var favoriteColor = "red"
	fmt.Println("my favorite color is", favoriteColor)

	birthYear, ageInYear := 1987, 34
	fmt.Println("Born in", birthYear, "aged", ageInYear, "years")

	var (
		firstInitial = "J"
		lastInitial  = "L"
	)
	fmt.Println("Initials=", firstInitial, lastInitial)

	var ageInDays int
	ageInDays = 365 * ageInYear
	fmt.Println("I am", ageInDays, "days old")
}
