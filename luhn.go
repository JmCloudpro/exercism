package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func main() {

	if Valid("9999999999 99999999d99 9999999999 9999999999") {
		fmt.Println("Valid")
	} else {
		fmt.Println("Not Valid")

	}

}

func Valid(id string) bool {

	//Eliminate all blanks
	id = strings.ReplaceAll(id, " ", "")

	//validate if length  is  bigger than 1
	if len(id) <= 1 {
		return false
	}

	// This if statement checks if each character passed is a number, if not, it returns false
	for _, c := range id {
		if !unicode.IsNumber(c) {
			return false
		}
	}

	// After all validations, lets check the algorithm

	str := strings.Split(id, "")

	// For lenght -2 ( check if duble is greater than 9 and subtract if necessary)
	for i := len(str); i > 1; {
		i = i - 2
		v, _ := strconv.ParseInt((str[i]), 10, 64)
		v = v * 2
		if v > 9 {
			v = v - 9
		}
		str[i] = strconv.FormatInt(v, 10)

	}

	// For lenght sum each digt
	v := 0

	for i := 0; i < len(str)-1; i++ {

		if v == 0 {
			v, _ = strconv.Atoi((str[0]))
		}

		v2, _ := strconv.Atoi((str[i+1]))
		v = v + v2

	}

	if v%10 == 0 {

		return true
	}

	return false
}
