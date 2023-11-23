package main

func main() {
	mercedes := &Mercedes{}
	kia := &Kia{}

	mercedesAdapter := &MercedesAdapter{mercedes}
	kiaAdapter := &KiaAdapter{kia}

	readyRun(mercedesAdapter)
	readyRun(kiaAdapter)

	stopRun(mercedesAdapter)
	stopRun(kiaAdapter)
}

func readyRun(car ITarget) {
	car.start()
	car.on()
}

func stopRun(car ITarget) {
	car.off()
	car.stop()
}
