package main

import (
	"fmt"
)

type IWeaponFactory interface {
	makeRifle() IGun
	makePistol() IGun
	makeEquipment() IEquipment
}

func GetWeaponFactory(team string) (IWeaponFactory, error) {
	if team == "terrorist" {
		return &Terrorist{}, nil
	}

	if team == "counter-terrorist" {
		return &CounterTerrorist{}, nil
	}

	return nil, fmt.Errorf("Wrong brand type passed")
}
