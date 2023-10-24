package main

import (
	"fmt"
	"sync"
)

var mu = &sync.Mutex{}

type bomb struct {
	timer    int
	password int
}

var singleBomb *bomb

func getBomb() *bomb {
	if singleBomb == nil {
		mu.Lock()
		defer mu.Unlock()
		if singleBomb == nil {
			singleBomb = &bomb{
				timer:    40,
				password: 1234567,
			}
			fmt.Println("Bomb created.")
		} else {
			fmt.Println("Bomb already exists.")
		}
	} else {
		fmt.Println("Bomb already exists.")
	}

	return singleBomb
}
