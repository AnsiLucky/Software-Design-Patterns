package main

import "fmt"

type DesertEagle struct{}

func (d *DesertEagle) shoot(p *Player) {
	fmt.Printf("shooting from deagle to %s\n", p.nickName)
}
