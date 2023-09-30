package main

import "fmt"

type Ak47 struct{}

func (a *Ak47) shoot(p *Player) {
	fmt.Printf("shooting from ak-47 to %s\n", p.nickName)
}
