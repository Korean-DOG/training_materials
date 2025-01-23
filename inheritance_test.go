package notes

import (
	"fmt"
	"testing"
)

type Coordinate struct {
	X, Y int
}

type Resetter interface {
	Reset()
}

type Player struct {
	health     int
	coordinate Coordinate
}

func (p *Player) Reset() {
	p.coordinate.X = 0
	p.coordinate.Y = 0
	p.health = 100
}

type Monster struct {
	health     int
	coordinate Coordinate
	power      bool
}

func (m *Monster) Reset() {
	m.health = 100
	m.coordinate.X = 0
	m.coordinate.Y = 0
	m.power = false
}

func Reset(r Resetter) {
	r.Reset()
}

func ResetWithPenalityToPlayer(r Resetter) {
	if player, ok := r.(Player); ok {
		player.health = 50
	} else {
		r.Reset()
	}
}

func TestInheritance(t *testing.T) {

	var p1 = Player{50, Coordinate{1, 1}}
	fmt.Println(p1)
	Reset(&p1) // same as p1.Reset
	fmt.Println(p1)

	var m1 = Monster{20, Coordinate{5, 5}, true}
	fmt.Println(m1)
	Reset(&m1) // same as m1.Reset
	fmt.Println(m1)

	//fmt.Println(array, a, b)

}
