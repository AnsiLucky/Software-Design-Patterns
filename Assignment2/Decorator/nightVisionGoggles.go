package main

import "strings"

type NightVisionGoggleAgent struct {
	agent ILook
}

func (n *NightVisionGoggleAgent) getLook() string {
	return strings.TrimSpace("Night Vision Goggle | " + n.agent.getLook())
}
