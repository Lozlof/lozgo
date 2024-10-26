package lozgo

import (
	"errors"
	"log"
)

// In Go, "error" is a predefined type.
func SanitizationIntInRange(number, minNum, maxNum int) (int, error) { 

	log.SetPrefix("Function: SanitizationIntInRange: ")
	log.SetFlags(2)

	// Declares the variables verifiedInt and err.
	// The variables are initialized when sanitizationIntVerify returns the values.
	verifiedInt, err := sanitizationIntVerify(number)
}

func sanitizationIntVerify(number int) (int, error) {

	log.SetPrefix("Function: sanitizationIntVerify: " )
	log.SetFlags(2)

}