package puppy


import (
	"github.com/shan793/dog"
)

func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "Woof! Woof! Woof!"
}

/* we are going to use this function to illustrate how to use
a function that is in another package in this code as well
we want to use the function in the dog package 
*/
func BigBark() string {

	return dog.WhenGrownUp(Bark())

}
func BigBarks() string {

	return dog.WhenGrownUp(Barks())

}


func FromV11() string {
	return "I'm from version 1.1.0"
}

func FromV12() string {
	return "I'm from version 1.2.0"
}

func FromV13() string {
	return "I'm from version 1.3.X (since I kinda messed up haha)"
}