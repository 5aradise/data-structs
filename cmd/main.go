package main

import (
	"fmt"

	ds "github.com/5aradise/data-structs"
)

func main() {
	l := ds.NewCSLList(0, 1, 2)

	fmt.Println(l)
	fmt.Println("length =", l.Length())

	fmt.Println("Appending 3 and 4")
	l.Append(3)
	l.Append(4)
	fmt.Println(l)

	fmt.Println("Inserting 69 at 3, -42, 42")
	err := l.Insert(69, 3)
	if err != nil {
		fmt.Println("error:", err)
	}
	err = l.Insert(69, -42)
	if err != nil {
		fmt.Println("error:", err)
	}
	err = l.Insert(69, 42)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(l)

	fmt.Println("Deleting 69 at 3, -42, 42")
	v, err := l.Delete(3)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("deleted:", v)
	}
	v, err = l.Delete(-42)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("deleted:", v)
	}
	v, err = l.Delete(42)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("deleted:", v)
	}
	fmt.Println(l)

	fmt.Println("Getting 3, -42, 42")
	v, err = l.Get(3)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("got:", v)
	}
	v, err = l.Get(-42)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("got:", v)
	}
	v, err = l.Get(42)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("got:", v)
	}
	fmt.Println(l)

	fmt.Println("Cloning")
	l2 := l.Clone()
	fmt.Println("l: ", l, "\nl2:", l2)
	fmt.Println("Inserting 69 into l2 at 3")
	l2.Insert(69, 3)
	fmt.Println("l: ", l, "\nl2:", l2)

	fmt.Println("Extending l by l2")
	l.Extend(l2)
	fmt.Println(l)

	fmt.Println("Deleting all 0")
	l.DeleteAll(0)
	fmt.Println(l)

	fmt.Println("first 2:", l.FindFirst(2))
	fmt.Println("first 42:", l.FindFirst(42))

	fmt.Println("last 2:", l.FindLast(2))
	fmt.Println("last 42:", l.FindLast(42))

	fmt.Println("Reversing")
	l.Reverse()
	fmt.Println(l)

	fmt.Println("Clearing")
	l.Clear()
	fmt.Println(l)
}
