package main

import (
	"lecture/demo/pkg/display"
	"lecture/demo/pkg/msg"
)

func main() {
	msg.Hi()
	display.Display("hello from display")
	msg.Exciting("an exciting message")
}
