package main

import "fmt"

type Mercedes struct{}

func (m *Mercedes) startEngine() {
	fmt.Println("Mercedes engine start")
}

func (m *Mercedes) stopEngine() {
	fmt.Println("Mercedes engine stopped")
}

func (m *Mercedes) onLights() {
	fmt.Println("Mercedes lights on")
}

func (m *Mercedes) offLights() {
	fmt.Println("Mercedes lights off")
}
