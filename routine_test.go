package notes

import (
	"fmt"
	"strconv"
	"sync"
	"testing"
)

func sumOdd(w *sync.WaitGroup, sum *int, in1 <-chan int) {
	defer w.Done()

	fmt.Println("[ODD] sum " + strconv.Itoa(*sum))
	for i := range in1 {
		fmt.Println("[ODD] got " + strconv.Itoa(i))
		if i == -1 {
			break
		}
		if i%2 == 1 {
			*sum += i
		}
		fmt.Println("[ODD] sum " + strconv.Itoa(*sum))
	}
	fmt.Println("[ODD] Done")
}

func sumEven(w *sync.WaitGroup, sum *int, in1 <-chan int) {
	defer w.Done()

	fmt.Println("[EVEN] sum " + strconv.Itoa(*sum))
	for i := range in1 {
		fmt.Println("[EVEN] got " + strconv.Itoa(i))
		if i == -1 {
			break
		}
		if i%2 == 0 {
			*sum += i
		}
		fmt.Println("[EVEN] sum " + strconv.Itoa(*sum))
	}
	fmt.Println("[EVEN] Done")
}

func balancer(w *sync.WaitGroup, sum_odd, sum_even *int, in1 <-chan int) {
	defer w.Done()

	odd_chan := make(chan int)
	even_chan := make(chan int)

	fmt.Println("[BALANCE] odd = " + strconv.Itoa(*sum_odd))

	go sumEven(w, sum_even, even_chan)
	go sumOdd(w, sum_odd, odd_chan)

	for i := range in1 {
		odd_chan <- i
		even_chan <- i
		if i == -1 {
			break
		}
	}
	fmt.Println("Balancer is Done")

}

func TestRoutine(t *testing.T) {
	var w sync.WaitGroup

	num_chan := make(chan int)
	var sum_odd int
	var sum_even int

	w.Add(3)

	go balancer(&w, &sum_odd, &sum_even, num_chan)
	for i := 0; i < 10; i++ {
		num_chan <- i
	}

	fmt.Println("Send -1")
	num_chan <- -1
	w.Wait()
	fmt.Println("Odds = " + strconv.Itoa(sum_odd) + ", Evens = " + strconv.Itoa(sum_even))
}
