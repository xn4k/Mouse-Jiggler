package main

import (
	"fmt"
	"time"

	"github.com/go-vgo/robotgo"
)

const (
	jiggleInterval = 5 * time.Second // Adjust the interval as needed
)

func main() {
	fmt.Println("Simple Mouse Jiggler started. Press Ctrl+C to exit.")

	for {
		moveMouseToCenter()
		time.Sleep(jiggleInterval)
	}
}

func moveMouseToCenter() {
	screenWidth, screenHeight := robotgo.GetScreenSize()
	mouseX, mouseY := robotgo.GetMousePos()

	// Determine the center of the screen where the mouse is located
	screenX := mouseX % screenWidth
	screenY := mouseY % screenHeight

	centerX := screenX + screenWidth/2
	centerY := screenY + screenHeight/2

	robotgo.MoveMouse(centerX, centerY)
}
