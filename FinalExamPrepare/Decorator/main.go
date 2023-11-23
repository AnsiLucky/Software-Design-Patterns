package main

import "fmt"

func main() {
	glock := &Glock{}
	glockWithSilencer := &Silencer{glock}
	glockWithSilencerWithExtendedFeeding := &ExtendedFeeding{glockWithSilencer}
	glockWithSilencerWithExtendedFeedingWithLaserSight := LaserSight{glockWithSilencerWithExtendedFeeding}

	fmt.Printf("Description: %s\n Price: %v\n\n", glock.getDescription(), glock.getPrice())
	fmt.Printf("Description: %s\n Price: %v\n\n", glockWithSilencer.getDescription(), glockWithSilencer.getPrice())
	fmt.Printf("Description: %s\n Price: %v\n\n", glockWithSilencerWithExtendedFeeding.getDescription(), glockWithSilencerWithExtendedFeeding.getPrice())
	fmt.Printf("Description: %s\n Price: %v\n\n", glockWithSilencerWithExtendedFeedingWithLaserSight.getDescription(), glockWithSilencerWithExtendedFeedingWithLaserSight.getPrice())

}
