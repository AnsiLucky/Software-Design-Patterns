package main

func main() {
	gameInstance := &GameInstance{name: "Battle Royal"}
	process1 := &Process{1}
	process2 := &Process{2}
	gameInstance.regObserver(process1)
	gameInstance.regObserver(process2)
	gameInstance.changeName("1 vs 1")
	gameInstance.addPlayer()
	gameInstance.addPlayer()
}
