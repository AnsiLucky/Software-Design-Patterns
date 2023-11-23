package main

import "fmt"

type Lights struct {
}

func (t *Lights) on() {
	fmt.Println("Lights turned on")
}

func (t *Lights) off() {
	fmt.Println("Lights turned off")
}
