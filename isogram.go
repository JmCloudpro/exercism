package isogram

import "strings"

func IsIsogram(word string) bool {
	word = strings.ToLower(word)
	sword := strings.Split(word, "")
	for a, i := range sword {
		if i != " " {
			if i != "-" {
				for b, s := range sword {
					if a != b {
						if i == s {
							return false
						}
					}
				}
			}
		}
	}

	return true
}
