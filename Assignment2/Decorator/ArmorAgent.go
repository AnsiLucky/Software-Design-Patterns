package main

import "strings"

type ArmorAgent struct {
	agent ILook
}

func (b *ArmorAgent) getLook() string {
	return strings.TrimSpace("Armor | " + b.agent.getLook())
}
