package main

import "fmt"

func main() {
	terrorist, err := GetPlayer("terrorist", "Ansi")
	if err != nil {
		panic("Error occurs when terrorist is creating")
	}

	fmt.Println(terrorist.getName())
}
