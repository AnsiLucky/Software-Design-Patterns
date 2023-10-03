package main

type IEquipment interface {
	setName(name string)
	getName() string
	setBehaviour(behaviour string)
	getBehaviour() string
}
type Equipment struct {
	name      string
	behaviour string
}

func (e *Equipment) setName(name string) {
	e.name = name
}

func (e *Equipment) getName() string {
	return e.name
}

func (e *Equipment) setBehaviour(behaviour string) {
	e.behaviour = behaviour
}

func (e *Equipment) getBehaviour() string {
	return e.behaviour
}
