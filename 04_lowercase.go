package main

import (
	"fmt"
	"unicode"
)

const xiaomiId = "9hf0P6MBHM7OY6UOWm"
const expected = "9hf0_p6_m_b_h_m7_o_y6_u_o_wm"

func main() {
	fmt.Println(IdToLowercase(xiaomiId))
}

func IdToLowercase(id string) string {
	newId := ""
	for _, c := range id {
		if unicode.IsUpper(c) {
			newId += "_"
			newId += string(unicode.ToLower(c))
		} else {
			newId+=string(c)
		}
	}
	return newId
}
