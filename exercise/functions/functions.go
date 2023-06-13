package main

import "fmt"

func twoTwos() (int, int) {
	return 2, 2
}

func five() int {
	return 5
}

func addThree(a, b, c int) int {
	return a + b + c
}

func greetPerson(name string) {
	fmt.Println("Hello", name)
}

func hiThere() string {
	return "Hi there!"
}

func main() {
	greetPerson("Alisa")
	fmt.Println(hiThere())

	a, b := twoTwos()
	answer := addThree(five(), a, b)

	fmt.Println(answer)
}
