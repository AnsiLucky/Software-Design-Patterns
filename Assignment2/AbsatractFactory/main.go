package main

import (
	"fmt"
)

func main() {
	terroristFactory, _ := GetWeaponFactory("terrorist")
	counterTerroristFactory, _ := GetWeaponFactory("counter-terrorist")

	terroristRifle := terroristFactory.makeRifle()
	terroristPistol := terroristFactory.makePistol()
	terroristEquipment := terroristFactory.makeEquipment()

	counterTerroristRifle := counterTerroristFactory.makeRifle()
	counterTerroristPistol := counterTerroristFactory.makePistol()
	counterTerroristEquipment := counterTerroristFactory.makeEquipment()

	printGunDetails(terroristRifle)
	printGunDetails(terroristPistol)
	printEquipmentDetails(terroristEquipment)

	printGunDetails(counterTerroristRifle)
	printGunDetails(counterTerroristPistol)
	printEquipmentDetails(counterTerroristEquipment)
}

func printGunDetails(g IGun) {
	fmt.Printf("Category: %s", g.getCategory())
	fmt.Println()
	fmt.Printf("Name: %s", g.getName())
	fmt.Println()
	fmt.Printf("Bullets in a magazine: %d", g.getMagazineBullets())
	fmt.Println()
	fmt.Printf("Bullets in a stock: %d\n", g.getStockBullets())
	fmt.Println("-------------------")
}

func printEquipmentDetails(e IEquipment) {
	fmt.Printf("Name: %s", e.getName())
	fmt.Println()
	fmt.Printf("Behaviour: %s\n", e.getBehaviour())
	fmt.Println("-------------------")
}
