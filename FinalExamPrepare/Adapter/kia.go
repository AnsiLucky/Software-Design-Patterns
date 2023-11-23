package main

import "fmt"

type Kia struct{}

func (k *Kia) turnOnEngine() {
	fmt.Println("Kia engine start")
}

func (k *Kia) turnOffEngine() {
	fmt.Println("Kia engine stopped")
}

func (k *Kia) turnOnLights() {
	fmt.Println("Kia lights on")
}

func (k *Kia) turnOffLights() {
	fmt.Println("Kia lights off")
}
