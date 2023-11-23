package main

type Glock struct {
}

func (g *Glock) getDescription() string {
	return "Glock"
}

func (g *Glock) getPrice() float32 {
	return 5
}
