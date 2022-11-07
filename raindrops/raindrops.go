/*
Your task is to convert a number into a string that contains raindrop sounds corresponding to certain potential factors. A factor is a number that evenly divides into another number, leaving no remainder. The simplest way to test if one number is a factor of another is to use the modulo operation.

The rules of raindrops are that if a given number:

    has 3 as a factor, add 'Pling' to the result.
    has 5 as a factor, add 'Plang' to the result.
    has 7 as a factor, add 'Plong' to the result.
    does not have any of 3, 5, or 7 as a factor, the result should be the digits of the number.

Examples

    28 has 7 as a factor, but not 3 or 5, so the result would be "Plong".
    30 has both 3 and 5 as factors, but not 7, so the result would be "PlingPlang".
    34 is not factored by 3, 5, or 7, so the result would be "34".

*/

package raindrops

import (
	"strconv"
)

// package raindrops
func Convert(n int) string {
	txt := ""
	txt1 := strconv.Itoa(n)

	for i := n; i > 1; i-- {

		if n%2 == 0 {
			n = n / 2
			i = n

		}

		if n%3 == 0 {
			n = n / 3
			i = n
			if txt != "Pling" {
				txt = txt + "Pling"
			}

		}

		if n%5 == 0 {
			n = n / 5
			i = n

			if txt != "Plang" {
				txt = txt + "Plang"
			}

		}

		if n%7 == 0 {
			n = n / 7
			i = n
			if txt != "Plong" {
				txt = txt + "Plong"
			}
		}

	}

	if txt == "" {
		txt = txt1
	}

	return txt
}
