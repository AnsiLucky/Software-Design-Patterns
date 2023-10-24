package main

type UnlockDoorsCommand struct {
	car VehicleWithFourWheels
}

func (o *UnlockDoorsCommand) execute() {
	o.car.unlockDoors()
}
