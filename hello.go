package main

import (
	"fmt"

	humanize "github.com/dustin/go-humanize"
)

func main() {
	fmt.Printf("Hello, world.\n")
	fmt.Printf("That file is %s. \n", humanize.Bytes(82854982))

	spot := Animal{name: "spot", age: 21, species: "thing"}
	spot.name = "wally"
	fmt.Printf("me %s , me can %v \n", spot.name, spot.talk())
}

// Animal represents a class?
type Animal struct {
	age     int
	name    string
	species string
}

func (animal *Animal) talk() (sound string) {
	if animal.species == "dog" {
		sound = "woof"
	} else {
		sound = "yellp"
	}
	return
}
