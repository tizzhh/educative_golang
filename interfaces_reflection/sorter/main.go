package main

import (
	"fmt"
	"sorter/mysort"
)

func ints() {
	data := []int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}
	a := sorter.IntSlice(data)
	sorter.Sort(a)
	if !sorter.IsSorted(a) {
		panic("fail")
	}
	fmt.Println("sorted:", a)
}

func strings() {
	data := []string{"Monday", "Friday", "Tuesday", "Wednesday", "Sunday", "Thursday", "", "Saturday"}
	a := sorter.StringSlice(data)
	sorter.Sort(a)
	if !sorter.IsSorted(a) {
		panic("fail")
	}
	fmt.Printf("The sorted array is: %v\n", a)
}

type day struct {
	num                 int
	shortName, longName string
}

type dayArray struct {
	data []*day
}

func (p *dayArray) Len() int           { return len(p.data) }
func (p *dayArray) Less(i, j int) bool { return p.data[i].num < p.data[j].num }
func (p *dayArray) Swap(i, j int)      { p.data[i], p.data[j] = p.data[j], p.data[i] }

func days() {
	Sunday := day{0, "SUN", "Sunday"}
	Monday := day{1, "MON", "Monday"}
	Tuesday := day{2, "TUE", "Tuesday"}
	Wednesday := day{3, "WED", "Wednesday"}
	Thursday := day{4, "THU", "Thursday"}
	Friday := day{5, "FRI", "Friday"}
	Saturday := day{6, "SAT", "Saturday"}
	data := dayArray{[]*day{&Tuesday, &Thursday, &Wednesday, &Sunday, &Monday, &Friday, &Saturday}}
	sorter.Sort(&data)
	if !sorter.IsSorted(&data) {
		panic("fail")
	}
	for _, d := range data.data {
		fmt.Printf("%s ", d.longName)
	}
	fmt.Printf("\n")
}

func main() {
	ints()
	strings()
	days()

	data := []int{10, 9, 8}
	fmt.Println(data)
	sorter.SortInts(data)
	fmt.Println(data)

}
