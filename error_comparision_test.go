package notes

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

type UserError struct {
	msg string
}

func (e *UserError) Error() string {
	return fmt.Sprintf("User Error: %s", e.msg)
}

const divisionErrorMessage = "Division Error"

type DivisionError struct {
	number int
}

func (e *DivisionError) Error() string {
	return fmt.Sprintf(divisionErrorMessage+" by %d", e.number)
}

func func1(i int) (int, error) {
	return i, &DivisionError{number: i}
}

func TestErrors(t *testing.T) {
	var UserError1 *UserError
	var DivisionError1 *DivisionError
	_, err := func1(0)

	if errors.As(err, &UserError1) {
		fmt.Println("We defined UserError")
	} else if errors.As(err, &DivisionError1) {
		fmt.Println("We defined DivisionError")
	}

	if errors.Is(err, &DivisionError{number: 0}) {
		fmt.Println("We catched DivError by 0")
	}

	fmt.Println(err, &DivisionError{number: 0}, err == &DivisionError{number: 0})

	fmt.Println(reflect.DeepEqual(err, &DivisionError{number: 0}))

	/* fmt.Println(err, DivisionError{number: 0})
	fmt.Println(&DivisionError{number: 0}, errors.Is(err, &DivisionError{number: 0}))
	fmt.Println(errors.Unwrap(fmt.Errorf("run error: %w", err))) */

}
