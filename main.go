package main

import (
	"fmt"

	humanize "github.com/dustin/go-humanize"
	"github.com/patrickkim/gosample/animal"
)

func main() {
	fmt.Printf("Hello, world.\n")
	fmt.Printf("That file is %s. \n", humanize.Bytes(82854982))

	dog := animal.Animal{Name: "spot", Age: 21, Species: "dog"}
	dog.Name = "wally"
	fmt.Printf("me %s , me can %v \n me am %v\n", dog.Name, dog.Speak(), dog.IsType())
}
