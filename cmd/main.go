package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Raspberry Pi Keyboard/Mouse Intercept Device")
	// TODO: Initialize event capture and USB HID emulation
	if os.Geteuid() != 0 {
		fmt.Println("This application must be run as root.")
		os.Exit(1)
	}
	// TODO: Setup event listeners and USB gadget
}
