package main

import (
	"fmt"

	humanize "github.com/dustin/go-humanize"
)

func main() {
	fmt.Printf("Hello, world.\n")
	fmt.Printf("That file is %s. \n", humanize.Bytes(82854982))

	dog := Animal{name: "spot", age: 21, species: "dog"}
	dog.name = "wally"
	fmt.Printf("me %s , me can %v \n me am %v\n", dog.name, dog.speak(), dog.isType())
}
