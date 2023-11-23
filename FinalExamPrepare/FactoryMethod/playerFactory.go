package main

import "fmt"

func GetPlayer(team, name string) (IPlayer, error) {
	if team == "terrorist" {
		return newTerrorist(name), nil
	} else if team == "counter-terrorist" {
		return nil, nil
	}

	return nil, fmt.Errorf("Wrong Team")
}
