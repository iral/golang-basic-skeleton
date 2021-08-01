package utils

import "fmt"

func CapitalizeFirstLetter(str string) string {
	// logic goes here
	var upperStr string
	vv := []rune(str) // Introduced later
	for i := 0; i < len(vv); i++ {
		if i == 0 {
			if vv[i] >= 97 && vv[i] <= 122 { // There will be introduction later
				vv[i] -= 32 // string code table differs by 32 bits
				upperStr += string(vv[i])
			} else {
				fmt.Println("Not begins with lowercase letter")
				return str
			}
		} else {
			upperStr += string(vv[i])
		}
	}
	return upperStr
}
