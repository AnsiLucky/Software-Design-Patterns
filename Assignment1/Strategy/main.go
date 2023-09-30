package main

func main() {
	player1 := initPlayer("Tima")
	player2 := initPlayer("John")

	ak := &Ak47{}
	player1.setGun(ak)
	player1.piupiu(player2)

	deagle := &DesertEagle{}
	player1.setGun(deagle)
	player1.piupiu(player2)

	player2.piupiu(player1)
	awp := &Awp{}
	player2.setGun(awp)
	player2.piupiu(player1)
}
