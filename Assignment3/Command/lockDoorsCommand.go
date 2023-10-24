package main

type LockDoorsCommand struct {
	car VehicleWithFourWheels
}

func (c *LockDoorsCommand) execute() {
	c.car.lockDoors()
}
