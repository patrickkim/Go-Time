package main

import (
	"fmt"
	// "strconv"
	humanize "github.com/dustin/go-humanize"
)

func main() {
	fmt.Printf("Hello, world.\n")
	fmt.Printf("That file is %s. \n", humanize.Bytes(82854982))

	spot := Dog{name: "spot", age: 21, species: "cat"}
	spot.name = "wally"
	fmt.Printf("me %s , me can %v \n", spot.name, spot.bark())
}

type Dog struct {
	age     int // how old is this
	name    string
	species string
}

func (d *Dog) bark() (sound string) {
	if d.species == "dog" {
		sound = "woof"
	} else {
		sound = "hiss"
	}
	return
}
