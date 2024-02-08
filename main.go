package main

import (
	"fmt"
	"unsafe"
)

const secret = "abc"

type Os int

// 定数の一括定義 (変数の場合も同様)
const (
	Mac Os = iota + 1
	Windows
	Linux
)

func main() {
	// var i int
	i := 1
	ui := uint16(2)
	fmt.Println(i)
	fmt.Printf("i: %[1]v %[1]T\n ui: %[2]v %[2]T", i, ui)

	pi, title := 3.14, "Go"
	fmt.Printf("pi: %v title: %v\n", pi, title)

	x := 10
	y := 1.23
	z := float64(x) + y
	fmt.Printf("z: %v", z)

	fmt.Printf("Mac:%v Windows:%v Linux:%v\n", Mac, Windows, Linux)

	var ui1 uint16
	fmt.Printf("memory address of ui1: %p\n", &ui1)
	var ui2 uint16
	fmt.Printf("memory address of ui1: %p\n", &ui2)
	var p1 *uint16
	fmt.Printf("memory address of p1: %p\n", p1)
	p1 = &ui1
	fmt.Printf("memory address of pi: %p\n", &pi)
	fmt.Printf("value of ui1(dereference): %v\n", *p1)

	var pp1 **uint16 = &p1
	fmt.Printf("value of pp1: %v\n", pp1)
	fmt.Printf("value of pp1(dereference): %v\n", **pp1)
	fmt.Printf("size of pp1: %d[bytes]\n", unsafe.Sizeof(pp1))

	ok, result := true, "A"
	if ok {
		// result := "B" // 同じアドレスを参照せず、スコープ内のみ値が変わる
		result = "B" // 同じアドレスを参照する
		println(result)
	} else {
		result := "C"
		println(result)
	}
	print(result)
}
