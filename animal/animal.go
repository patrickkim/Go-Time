package animal

// Animal represents a class?
type Animal struct {
	Age     int
	Name    string
	Species string
}

// Speak barks
func (animal *Animal) Speak() (sound string) {
	if animal.Species == "dog" {
		sound = "woof"
	} else {
		sound = "yellp"
	}
	return
}

// IsType spits latin name
func (animal *Animal) IsType() string {
	var latinName string
	if animal.Species == "dog" {
		latinName = "Canis Canis"
	} else {
		latinName = "shite"
	}

	return latinName
}
