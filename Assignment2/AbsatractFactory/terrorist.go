package main

type Terrorist struct{}

func (*Terrorist) makeRifle() IGun {
	return &TerroristGun{
		Gun{
			category:        "primary",
			name:            "AK-47",
			magazineBullets: 30,
			stockBullets:    90,
		},
	}

}

func (*Terrorist) makePistol() IGun {
	return &TerroristGun{
		Gun{
			category:        "secondary",
			name:            "Glock-18",
			magazineBullets: 20,
			stockBullets:    120,
		},
	}
}

func (*Terrorist) makeEquipment() IEquipment {
	return &TerroristEquipment{
		Equipment{
			name:      "bomb",
			behaviour: "exploding",
		},
	}
}
