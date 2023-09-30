package main

import "fmt"

type noGun struct{}

func (n *noGun) shoot(p *Player) {
	fmt.Printf("we haven't got gun for shoot to %s\n", p.nickName)
}
