package main

type GameInstance struct {
	name          string
	playersAmount int
	observers     []IObserver
}

func (g *GameInstance) changeName(name string) {
	g.name = name
	g.notifyAll()
}

func (g *GameInstance) addPlayer() {
	g.playersAmount++
	g.notifyAll()
}

func (g *GameInstance) removePlayer() {
	g.playersAmount--
	g.notifyAll()
}

func (g *GameInstance) regObserver(o IObserver) {
	g.observers = append(g.observers, o)
}

func (g *GameInstance) deregObserver(o IObserver) {
	g.observers = removeFromObserverList(g.observers, o)

}

func (g *GameInstance) notifyAll() {
	for _, observer := range g.observers {
		observer.update(g.name, g.playersAmount)
	}
}

func removeFromObserverList(observerList []IObserver, observerToRemove IObserver) []IObserver {
	observerListLength := len(observerList)
	observerToRemoveId := observerToRemove.getId()
	for i, observer := range observerList {
		if observerToRemoveId == observer.getId() {
			observerList[observerListLength-1], observerList[i] =
				observerList[i], observerList[observerListLength-1]

			return observerList[:observerListLength-1]
		}
	}
	return observerList
}
