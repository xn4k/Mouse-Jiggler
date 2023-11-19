package main

import (
	"fmt"
	"github.com/go-vgo/robotgo"
	"time"
)

const textForEditor = "Ugh i am working"

func main() {
	//windowsEditorKeyboard()
	windowsMouse()

}

func windowsEditorKeyboard() {

	err := robotgo.KeyTap("cmd")
	if err != nil {
		return
	}
	robotgo.Sleep(5)
	robotgo.TypeStr("editor")
	robotgo.Sleep(5)

	fmt.Println("Pressing enter")
	robotgo.KeySleep = 100

	err = robotgo.KeyTap("enter")
	if err != nil {
		return
	}

	robotgo.Sleep(5)
	fmt.Println("Keyboard Jiggler started. Press Ctrl+C to exit.")

	for {
		robotgo.TypeStr(textForEditor)
	}

}

func windowsMouse() {
	//get screen size
	screenX, screenY := robotgo.GetScreenSize()
	fmt.Println("Screen size is", screenX, screenY)
	locateX, locateY := robotgo.Location()
	fmt.Println("Current mouse location is", locateX, locateY)

	for {
		robotgo.MoveSmooth(300, 300)
		time.Sleep(5 * time.Second)
		robotgo.MoveSmooth(600, 600)
		time.Sleep(5 * time.Minute)
	}
}
