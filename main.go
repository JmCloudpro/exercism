package main

import (
	"exercism/luhn"
	"fmt"
)

func main() {
	if luhn.Valid("555646458546848654 6 4 84 654 6") {
		fmt.Println("Valid")
	} else {
		fmt.Println("Invalid Valid")

	}

}
