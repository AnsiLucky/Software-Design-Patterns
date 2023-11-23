package main

type CinemaPark struct {
	screen *Screen
	tuner  *Tuner
	lights *Lights
}

func NewCinemaPark(screen *Screen, tuner *Tuner, lights *Lights) *CinemaPark {
	return &CinemaPark{screen, tuner, lights}
}

func (c *CinemaPark) watchMovie() {
	c.screen.on()
	c.tuner.on()
	c.lights.off()
}

func (c *CinemaPark) endMovie() {
	c.screen.off()
	c.tuner.off()
	c.lights.on()
}
