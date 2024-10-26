package lozgo

import (
	"errors"
	"log"
)

// In Go, "error" is a predefined type.
// Takes a string input and two int's.
// The string is sent off to sanitizationIntVerify to be verified and converted to an int.
func SanitizationIntInRange(input string, minNum, maxNum int) (int, error) { 

	log.SetPrefix("Function: SanitizationIntInRange: ")
	log.SetFlags(2)

	// Declares the variables verifiedInt and err.
	// The variables are initialized when sanitizationIntVerify returns the values.
	verifiedInt, err := sanitizationIntVerify(number)

	if err != nil {
		log.Fatal(err)
	}
}

func sanitizationIntVerify(number int) (int, error) {

	log.SetPrefix("Function: sanitizationIntVerify: " )
	log.SetFlags(2)

}
