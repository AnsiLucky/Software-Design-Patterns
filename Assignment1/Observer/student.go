package main

import "fmt"

type Student struct {
	fullName string
}

func (c *Student) getFullName() string {
	return c.fullName
}

func (c *Student) update(itemName string) {
	fmt.Printf("%s in the amount of 41 898tg deposited into %s account\n", itemName, c.fullName)
}
