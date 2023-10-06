package main

import "strings"

type CasqueAgent struct {
	agent ILook
}

func (c *CasqueAgent) getLook() string {
	return strings.TrimSpace("Casque | " + c.agent.getLook())
}
