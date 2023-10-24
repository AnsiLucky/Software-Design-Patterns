package main

import "time"

func main() {
	for i := 0; i < 10; i++ {
		go getBomb()
	}
	time.Sleep(3 * time.Second)
}
