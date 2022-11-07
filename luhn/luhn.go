/*Instructions

Given a number determine whether or not it is valid per the Luhn formula.

The Luhn algorithm is a simple checksum formula used to validate a variety of identification numbers, such as credit card numbers and Canadian Social Insurance Numbers.

The task is to check if a given string is valid.
Validating a Number

Strings of length 1 or less are not valid. Spaces are allowed in the input, but they should be stripped before checking. All other non-digit characters are disallowed.
Example 1: valid credit card number

4539 3195 0343 6467

The first step of the Luhn algorithm is to double every second digit, starting from the right. We will be doubling

4_3_ 3_9_ 0_4_ 6_6_

If doubling the number results in a number greater than 9 then subtract 9 from the product. The results of our doubling:

8569 6195 0383 3437

Then sum all of the digits:

8+5+6+9+6+1+9+5+0+3+8+3+3+4+3+7 = 80

If the sum is evenly divisible by 10, then the number is valid. This number is valid!
Example 2: invalid credit card number

8273 1232 7352 0569

Double the second digits, starting from the right

7253 2262 5312 0539

Sum the digits

7+2+5+3+2+2+6+2+5+3+1+2+0+5+3+9 = 57

57 is not evenly divisible by 10, so this number is not valid.

*/

package luhn

import (
	"strconv"
	"strings"
	"unicode"
)

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
