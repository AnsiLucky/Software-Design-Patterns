package main

import "fmt"

type Awp struct{}

func (a *Awp) shoot(p *Player) {
	fmt.Printf("shooting from awp to %s\n", p.nickName)
}
