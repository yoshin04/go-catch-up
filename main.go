package main

import "fmt"

type controller interface {
	speedUp() int
	speedDown() int
}

type vehicle struct {
	speed       int
	enginePower int
}

type bycycle struct {
	speed      int
	humanPower int
}

func (v *vehicle) speedUp() int {
	v.speed += 10 * v.enginePower
	return v.speed
}

func (v *vehicle) speedDown() int {
	v.speed -= 5 * v.enginePower
	return v.speed
}

func (v *bycycle) speedUp() int {
	v.speed += 3 * v.humanPower
	return v.speed
}
func (v *bycycle) speedDown() int {
	v.speed -= 1 * v.humanPower
	return v.speed
}
func speedUpAndDown(c controller) {
	fmt.Printf("current speed: %v\n", c.speedUp())
	fmt.Printf("current speed: %v\n", c.speedDown())
}

func main() {
	v := &vehicle{0, 5}
	speedUpAndDown(v)
	b := &bycycle{0, 5}
	speedUpAndDown(b)
}
