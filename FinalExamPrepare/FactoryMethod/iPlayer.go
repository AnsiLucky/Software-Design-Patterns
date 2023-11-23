package main

type IPlayer interface {
	setName(name string)
	getName() string
	getTeam() string
}
