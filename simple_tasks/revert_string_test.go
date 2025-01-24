package notes

import (
	"fmt"
	"testing"
)

func TestRevertString(t *testing.T) {
	var s string = "abcdefg"

	rev_s := []rune(s)

	for i := 0; i < len(s)/2; i++ {
		if i == len(s)-i-1 {
			continue
		}
		rev_s[i], rev_s[len(s)-i-1] = rev_s[len(s)-i-1], rev_s[i]
	}

	fmt.Println(string(rev_s))

}
