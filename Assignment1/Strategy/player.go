package main

type Player struct {
	nickName string
	hp       int
	bomb     bool
	gun      GunType
}

func initPlayer(nick string) *Player {
	return &Player{nick, 100, false, &noGun{}}
}

func (p *Player) setGun(g GunType) {
	p.gun = g
}

func (p *Player) getBomb() {
	p.bomb = true
}

func (p *Player) throwBomb() {
	p.bomb = false
}

func (p *Player) piupiu(e *Player) { // e = enemy
	p.gun.shoot(e)
	// logic of game:
	// decrease enemy hp depending on gun type
}
