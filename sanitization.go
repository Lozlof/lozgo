package lozgo

import (
	"errors"
	"log"
	"strconv"
)

// In Go, "error" is a predefined type.
// Takes a string input and two int's.
// The string is sent off to sanitizationIntVerify to be verified and converted to an int.
func SanitizationIntInRange(input string, minNum, maxNum int) (int, error) { 

	// Set logging values.
	log.SetPrefix("lozgo/SanitizationIntInRange: ")
	log.SetFlags(2)

	// Attempts to convert "input" to an integer.
	// If it is converted, err will be nil.
	// convertedInt will be initialized to the integer returned by strconv.Atoi.
	convertedInt, err := strconv.Atoi(input)

	if err != nil {
		log.Println('"',input,'"'," is not an integer" )
		return 0, errors.New("input cannot be converted to an integer")
	}



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
