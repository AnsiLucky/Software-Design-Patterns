package main

import "fmt"

type Process struct {
	id int
}

func (p *Process) getId() int {
	return p.id
}

func (p *Process) update(name string, playersAmount int) {
	fmt.Printf("Name of game: %s and Players Amount: %v\n", name, playersAmount)
}
