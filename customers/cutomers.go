package customers

import (
	"interview/utility"
	"log"
)

func Cutomers() {
	add := utility.Add(4, 8)
	log.Printf("Addition result %f", add)
	subtract := utility.Subtract(8, 4)
	log.Printf("Subtraction result %f", subtract)
}
