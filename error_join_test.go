package notes

import (
	"errors"
	"fmt"
	"testing"
)

type JoinedError struct {
	number int
}

func (e *JoinedError) Error() string {
	return fmt.Sprintf("Join %d", e.number)
}

func fun_w_error(i int) error {
	return &JoinedError{number: i}
}

func TestJoinErrors(t *testing.T) {
	var err error
	for i := 0; i < 5; i++ {
		err = errors.Join(err, fun_w_error(i))
	}

	fmt.Println(err)

}
