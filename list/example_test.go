package list

import "fmt"

func ExampleNewList() {
	l := NewList[int]()
	l.Push(0)
	l.Push(1)
	l.Push(2)
	fmt.Println(l.ToSlice())
	fmt.Println(l.Pop())
	fmt.Println(l.Pop())
	fmt.Println(l.Pop())
	fmt.Println(l.Pop())
	fmt.Println(l.Len())
	l.Push(0)
	fmt.Println(l.Len())
	// Output:
	// [0 1 2]
	// 2 true
	// 1 true
	// 0 true
	// 0 false
	// 0
	// 1
}
