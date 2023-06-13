package main

import "fmt"

func main() {
	var myName = "Jayson"
	fmt.Println("my name is ", myName)

	var name string = "Kata"
	fmt.Println("name =", name)

	userName := "admin"
	fmt.Println("Username =", userName)

	var sum int
	fmt.Println("The sum is", sum)

	part1, other := 1, 5
	fmt.Println("part 1 is", part1, "other is", other)

	part2, other := 2, 0
	fmt.Println("part 2 is", part2, "other is", other)

	sum = part1 + part2
	fmt.Println("sum=", sum)

	var (
		lessonName = "Variables"
		lessonType = "Demo"
	)
	fmt.Println("lessonName", lessonName, "lessonType", lessonType)

	word1, word2, _ := "hello", "world", "!"
	fmt.Println(word1, word2)
}
