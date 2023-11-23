package main

type MercedesAdapter struct {
	mercedes *Mercedes
}

func (m *MercedesAdapter) start() {
	m.mercedes.startEngine()
}

func (m *MercedesAdapter) stop() {
	m.mercedes.stopEngine()
}

func (m *MercedesAdapter) on() {
	m.mercedes.onLights()
}

func (m *MercedesAdapter) off() {
	m.mercedes.offLights()
}
