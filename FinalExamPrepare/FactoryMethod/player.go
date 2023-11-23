package main

type Player struct {
	name string
	team string
}

func (p *Player) getTeam() string {
	return p.team
}

func (p *Player) setName(name string) {
	p.name = name
}

func (p *Player) getName() string {
	return p.name
}
