package main

import "fmt"

type Car struct {
	isDoorsOpen bool
	isEngineOn  bool
}

func (c *Car) startEngine() {
	c.isEngineOn = true
	fmt.Println("Engine started.")
}

func (c *Car) stopEngine() {
	c.isEngineOn = false
	fmt.Println("Engine stopped.")
}

func (c *Car) unlockDoors() {
	c.isDoorsOpen = true
	fmt.Println("Doors opened.")
}

func (c *Car) lockDoors() {
	c.isDoorsOpen = false
	fmt.Println("Doors closed.")
}
