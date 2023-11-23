package main

import "fmt"

type Screen struct {
}

func (t *Screen) on() {
	fmt.Println("Screen turned on")
}

func (t *Screen) off() {
	fmt.Println("Screen turned off")
}
