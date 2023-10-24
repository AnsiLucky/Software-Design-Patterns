package main

type StartEngineCommand struct {
	car VehicleWithFourWheels
}

func (o *StartEngineCommand) execute() {
	o.car.startEngine()
}
