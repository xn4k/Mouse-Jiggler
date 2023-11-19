package main

import (
	"fmt"
	"github.com/go-vgo/robotgo"
)

const x = "Hello World"

func main() {
	printToEditorWindows()

}

func printToEditorWindows() {
	fmt.Println("Mouse Jiggler started. Press Ctrl+C to exit.")

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

	for {
		robotgo.TypeStr(x)
	}

}
