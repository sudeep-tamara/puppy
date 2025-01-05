package puppy

import "github.com/sudeep-tamara/dog"

func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "Woof! Woof! Woof!"
}

func BigBarks() string {
	return dog.WhenGrownUp(Barks())
}

func BigBark() string {
	return dog.WhenGrownUp(Bark())
}
