package main

import "fmt"

func main() {
	var str [2]string
	// basic definition of an array
	str[0] = "hey"
	str[1] = "there"
	fmt.Println(str[0], str[1])
	fmt.Println(str)

	// shorthand way off adding
	str2 := [3]string{"Loading", "an", "array"}
	fmt.Println(str2)

	// this is a slice, slices are a pointer reference to inside of an array. This is a way of editing arrays.
	evens := [10]int{2, 4, 5, 8, 10, 12, 14, 16, 18, 20}
	var sl []int = evens[0:5]
	fmt.Println(sl)
	a := evens[:5]
	b := evens[1:6]
	fmt.Println(a, b)
	b[0] = 42
	fmt.Println(evens, a, b)
	// notice the second slot for evens has changed to 42 as well as the 1st slot for b.

	//selecting the entire array.
	a1 := evens[:]
	b1 := evens[:5]
	fmt.Println(a1, b1)
	fmt.Println(len(b1))
	fmt.Println(cap(b1))
	// capacity is different than length because you are addressing the entire evens.
	b1 = b1[:6]
	fmt.Println(b1)

	var empty []int
	fmt.Println(empty, len(empty), cap(empty), empty == nil)
	// dynamically creating the array we need.
	c := []int{1, 3, 5, 7, 9}
	fmt.Println(c, len(c), cap(c))
	// all sorts of different types can be used to create arrays and slices
	d := []struct {
		x int
		y int
	}{
		{0, 0},
		{0, 1},
		{1, 1},
		{1, 0},
	}
	fmt.Println(d)

	//using make to make a dynamic slices

	m := make([]int, 5)
	fmt.Println(cap(m))
	m = m[1:]
	fmt.Println(cap(m))

	m = append(m, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(m, cap(m), len(m))
	// m has been automatically resized now very much in how python handles arrays.

	// we can use for to itterate over a slice

	for i, v := range m {
		fmt.Println(i, v)
	}
	// _ symbol indicates to go to drop variable from memory.
	for _, v := range m {
		fmt.Println(v)
	}
}
