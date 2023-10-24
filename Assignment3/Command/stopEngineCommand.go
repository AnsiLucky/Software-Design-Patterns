package main

type StopEngineCommand struct {
	car VehicleWithFourWheels
}

func (o *StopEngineCommand) execute() {
	o.car.stopEngine()
}
