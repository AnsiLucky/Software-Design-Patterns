package main

type LaserSight struct {
	weapon IWeapon
}

func (s *LaserSight) getDescription() string {
	return s.weapon.getDescription() + " | Laser Signt"
}

func (s *LaserSight) getPrice() float32 {
	return 2 + s.weapon.getPrice()
}
