package main

import "fmt"

func main() {
	screen := &Screen{}
	tuner := &Tuner{}
	lights := &Lights{}
	cinemaPark := NewCinemaPark(screen, tuner, lights)

	cinemaPark.watchMovie()
	fmt.Println()
	cinemaPark.endMovie()
}
