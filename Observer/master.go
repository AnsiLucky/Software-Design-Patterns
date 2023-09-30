package main

import "fmt"

type Master struct {
	fullName string
}

func (c *Master) getFullName() string {
	return c.fullName
}

func (c *Master) update(itemName string) {
	fmt.Printf("%s in the amount of 97 024tg deposited into %s account\n", itemName, c.fullName)
}
