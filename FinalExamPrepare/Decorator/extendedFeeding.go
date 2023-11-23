package main

type ExtendedFeeding struct {
	weapon IWeapon
}

func (s *ExtendedFeeding) getDescription() string {
	return s.weapon.getDescription() + " | Extended Feeding"
}

func (s *ExtendedFeeding) getPrice() float32 {
	return 5 + s.weapon.getPrice()
}
