package animal

// Animal represents a class?
type Animal struct {
	age     int
	name    string
	species string
}

func (d *Animal) talk() (sound string) {
	if d.species == "dog" {
		sound = "woof"
	} else {
		sound = "yellp"
	}
	return
}
