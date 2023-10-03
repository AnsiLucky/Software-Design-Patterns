package main

type CounterTerrorist struct{}

func (*CounterTerrorist) makeRifle() IGun {
	return &CounterTerroristGun{
		Gun{
			category:        "primary",
			name:            "M4A1-S",
			magazineBullets: 20,
			stockBullets:    80,
		},
	}

}

func (*CounterTerrorist) makePistol() IGun {
	return &CounterTerroristGun{
		Gun{
			category:        "secondary",
			name:            "USP-S",
			magazineBullets: 12,
			stockBullets:    24,
		},
	}
}

func (*CounterTerrorist) makeEquipment() IEquipment {
	return &CounterTerroristEquipment{
		Equipment{
			name:      "defusing kits",
			behaviour: "defusing",
		},
	}
}
