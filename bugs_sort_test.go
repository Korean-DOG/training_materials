package notes

import (
	"fmt"
	"sort"
	"testing"
)

func TestBugsSort(t *testing.T) {
	var defects = map[string]int{
		"defect1":   1,
		"bug2":      3,
		"big":       10,
		"some bug3": 3,
	}

	type sliceElement struct {
		bug   string
		count int
	}

	var bugs []sliceElement

	for k, v := range defects {
		bugs = append(bugs, sliceElement{bug: k, count: v})
	}

	sort.Slice(bugs, func(i, j int) bool {
		return bugs[i].count > bugs[j].count
	})

	fmt.Println(bugs)

	/* 	var bugSlice struct {
	   		bug   []string
	   		count []int
	   	}

	   	for k, v := range defects {
	   		fmt.Println(len(bugSlice.bug), k)
	   		fmt.Println(len(bugSlice.count), v)
	   		bugSlice.bug = append(bugSlice.bug, k)
	   		bugSlice.count = append(bugSlice.count, v)

	   	}

	   	fmt.Println(bugSlice) */

	//sort.Slice(bugSlice, func(i, j int) {}

	/* var myMap = map[string]int{"1": 1, "3": 5}
	var reverseMap = make(map[int]string)

	fmt.Println(myMap["1"], myMap["2"])

	for k, v := range myMap {
		fmt.Printf("%s:%d,", k, v)
		reverseMap[v] = k
	}

	fmt.Println()
	fmt.Println(myMap)
	fmt.Println(reverseMap) */

	/* var part1 = []int{1, 2, 3}
	var part2 = []int{4, 5, 6}

	mergedSlice := append(part1, part2...)

	fmt.Println(mergedSlice)
	fmt.Println(part2)
	part1[1] = 999
	fmt.Println(mergedSlice)
	fmt.Println(part1)

	var array = [...]int{1, 2, 3, 4, 5}
	var a = array[0:2]
	var b = array[3:]
	fmt.Println(array, a, b)
	a[0] = 8
	b[0] = 8
	fmt.Println(array, a, b) */

}
