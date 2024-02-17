package main

import (
	"errors"
	"fmt"
)

func main() {
	err01 := errors.New("something wrong")
	fmt.Printf("%[1]p %[1]T %[1]v\n", err01)
}
