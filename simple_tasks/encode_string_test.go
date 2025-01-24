package notes

import (
	"fmt"
	"testing"
)

func TestEncodeString(t *testing.T) {
	var s string = "aaabbccccddaaaaaaaaaaaaaaaaaaaaaaaaa"

	encoded := []rune("")

	var num int = 1
	var letter rune = rune(s[0])
	for i := 1; i < len(s); i++ {
		if rune(s[i]) == letter && i != len(s)-1 {
			num += 1
		} else {
			encoded = append(encoded, rune(letter))
			for j := 0; j < len(fmt.Sprintf("%d", num)); j++ {
				encoded = append(encoded, []rune(fmt.Sprintf("%d", num))[j])
			}

			num = 1
			letter = rune(s[i])
		}
	}

	fmt.Println(string(encoded))

}
