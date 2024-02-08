package main

import (
	"fmt"
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

	
}
