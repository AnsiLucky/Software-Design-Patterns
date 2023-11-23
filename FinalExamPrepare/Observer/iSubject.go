package main

type Subject interface {
	regObserver(observer IObserver)
	deregObserver(observer IObserver)
	notifyAll()
}
