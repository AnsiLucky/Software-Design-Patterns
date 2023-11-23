package main

import "fmt"

type Tuner struct {
}

func (t *Tuner) on() {
	fmt.Println("Tuner turned on")
}

func (t *Tuner) off() {
	fmt.Println("Tuner turned off")
}
