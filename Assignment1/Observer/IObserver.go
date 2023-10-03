package main

type Observer interface {
	update(string)
	getFullName() string
}
