package main

type IObserver interface {
	getId() int
	update(name string, playersAmount int)
}
