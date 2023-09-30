package main

func main() {

	scholarship := newItem("Scholarship")

	observerFirst := &Student{"Tima Neprostoy"}
	observerSecond := &Master{"Bob Zolotoy"}
	observerThird := &Master{"Bob fasdf"}
	observerFourth := &Master{"Bob John"}

	scholarship.register(observerFirst)
	scholarship.register(observerSecond)
	scholarship.register(observerThird)
	scholarship.register(observerFourth)
	scholarship.deregister(observerThird)

	scholarship.update()
}
