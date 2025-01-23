package notes

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var once sync.Once

type singleton struct {
	title string
}

var instance *singleton = nil

type Singleton interface {
	SetTitle(t string)
	GetTitle() string
}

// Setter for singleton variable
func (s *singleton) SetTitle(t string) {
	s.title = t
}

// Getter singleton variable
func (s *singleton) GetTitle() string {
	return s.title
}

func GetInstance() Singleton {
	once.Do(func() {
		instance = new(singleton)
	})
	return instance
}

func TestGetInstance(t *testing.T) {
	var s Singleton
	s = GetInstance()
	if s == nil {
		t.Fatalf("First sigletone is nil")
	}

	s.SetTitle("First value")
	checkTitle := s.GetTitle()

	if checkTitle != "First value" {
		t.Errorf("First value is not setted")
	}

	var s2 Singleton
	s2 = GetInstance()
	if s2 != s {
		t.Error("New instance different")
	}

	s2.SetTitle("New title")
	newTitle := s.GetTitle()
	if newTitle != "New title" {
		t.Errorf("Title different after change")
	}
}

func TestSecondGetInstance(t *testing.T) {
	var w sync.WaitGroup

	var num_instances int = 100
	var num_cycles int = 3000
	w.Add(num_instances)

	t1 := time.Now()
	for inst := 0; inst < num_instances; inst++ {
		var instance = GetInstance()
		var j = inst
		go func() {
			defer w.Done()
			for i := 0; i < num_cycles; i++ {
				instance.SetTitle("Instance #" + strconv.Itoa(j) + ". Cycle #" + strconv.Itoa(i))
			}
		}()
	}
	fmt.Println(GetInstance().GetTitle())
	fmt.Println(-time.Until(t1))
	t2 := time.Now()

	stop_word := strings.Clone(GetInstance().GetTitle())
	for stop_word != GetInstance().GetTitle() {
		time.Sleep(10 * time.Millisecond)
		stop_word = GetInstance().GetTitle()
		fmt.Println(GetInstance().GetTitle())
	}

	fmt.Println(strings.Repeat("-", 100))
	fmt.Println(-time.Until(t2))
	fmt.Println(GetInstance().GetTitle())

	assert.Equal(t, GetInstance().GetTitle(), GetInstance().GetTitle())

}
