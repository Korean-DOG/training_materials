package notes

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"testing"
)

type Figure struct {
	sides []float32
}

func (f *Figure) getPerimeter() (float32, error) { return 0.0, errors.New("Must be re-defined") }

type Triangle struct {
	Figure
}

func (t *Triangle) getPerimeter() (float32, error) {
	if len(t.sides) != 3 {
		return 0.0, errors.New("Number of sides must be 3 for Triangle, but it is " + strconv.Itoa(len(t.sides)))

	}
	return t.sides[0] + t.sides[1] + t.sides[2], nil
}

type Circle struct {
	Figure
}

func (c *Circle) getPerimeter() (float32, error) {
	if len(c.sides) != 1 {
		return 0, errors.New("Number of sides must be 1 for Circle (radius), but it is " + strconv.Itoa(len(c.sides)))

	}
	return c.sides[0] * c.sides[0] * math.Pi, nil
}

func TestInheritance2(t *testing.T) {
	var tr = Triangle{}
	tr.sides = []float32{3.0, 4.0, 5.0}
	trPerimeter, err := tr.getPerimeter()
	if err != nil {
		fmt.Println("Error with trianlge")
		//fmt.Errorf(err)
	}
	fmt.Println("Perimeter of Triangle = " + fmt.Sprintf("%v", trPerimeter))

	var c = Circle{}
	c.sides = []float32{1.0}
	c_perimeter, err := c.getPerimeter()
	if err != nil {
		fmt.Println("Error with circle")
		//fmt.Errorf(err)
	}
	fmt.Println("Perimeter of Circle = " + fmt.Sprintf("%v", c_perimeter))

}
