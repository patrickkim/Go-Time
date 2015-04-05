package main

// Animal represents a class?
type Animal struct {
	age     int
	name    string
	species string
}

func (animal *Animal) speak() (sound string) {
	if animal.species == "dog" {
		sound = "woof"
	} else {
		sound = "yellp"
	}
	return
}

func (animal *Animal) isType() string {
	return animal.species
}
