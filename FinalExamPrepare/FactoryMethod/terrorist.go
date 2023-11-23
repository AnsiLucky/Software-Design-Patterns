package main

type Terrorist struct {
	Player
}

func newTerrorist(name string) IPlayer {
	return &Terrorist{Player{name, "Terrorist"}}
}
