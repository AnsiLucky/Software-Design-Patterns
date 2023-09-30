package main

type Item struct {
	observerList []Observer
	name         string
}

func newItem(name string) *Item {
	return &Item{
		name: name,
	}
}
func (i *Item) update() {
	i.notifyAll()
}
func (i *Item) register(o Observer) {
	i.observerList = append(i.observerList, o)
}

func (i *Item) deregister(o Observer) {
	i.observerList = removeFromObserverList(i.observerList, o)

}

func (i *Item) notifyAll() {
	for _, observer := range i.observerList {
		observer.update(i.name)
	}
}

func removeFromObserverList(observerList []Observer, observerToRemove Observer) []Observer {
	observerListLength := len(observerList)
	for i, observer := range observerList {
		if observerToRemove.getFullName() == observer.getFullName() {
			observerList[observerListLength-1], observerList[i] =
				observerList[i], observerList[observerListLength-1]

			return observerList[:observerListLength-1]
		}
	}
	return observerList
}
