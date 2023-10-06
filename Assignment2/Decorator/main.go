package main

import "fmt"

func main() {
	terrorist := &Agent{}

	terroristWithArmor := &ArmorAgent{terrorist}
	terroristWithArmorAndCasque := &CasqueAgent{terroristWithArmor}

	counterTerrorist := &Agent{}
	counterTerroristWithArmor := &ArmorAgent{counterTerrorist}
	counterTerroristWithArmorAndCasque := &CasqueAgent{counterTerroristWithArmor}
	counterTerroristWithArmorAndCasqueAndNightVisionGoggles := &NightVisionGoggleAgent{counterTerroristWithArmorAndCasque}

	fmt.Printf("Terrorist look: %s\n", terroristWithArmorAndCasque.getLook())
	fmt.Printf("Counter Terrorist look: %s\n", counterTerroristWithArmorAndCasqueAndNightVisionGoggles.getLook())
}
