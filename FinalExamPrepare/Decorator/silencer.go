package main

type Silencer struct {
	weapon IWeapon
}

func (s *Silencer) getDescription() string {
	return s.weapon.getDescription() + " | Silencer"
}

func (s *Silencer) getPrice() float32 {
	return 1.99 + s.weapon.getPrice()
}
