package main

import (
	"fmt"
)

// 課題 2
func culc(fn func(float64, float64) float64, x float64, y float64) (res float64) {
	res = fn(x, y)
	return
}

// 課題 1
func plus(x float64, y float64) float64 {

	return x + y
}

func minus(x float64, y float64) float64 {

	return x - y
}

func multiply(x float64, y float64) float64 {

	return x * y

}

func divided(x float64, y float64) float64 {

	return x / y
}

type Turtle struct {
	msg  string
	name string
	x    float64
	y    float64
	a    float64
}

func (t *Turtle) move(msg string, name string, x float64, y float64, a float64) {
	t.x += x
	t.y += y
	// 課題 3.1
	t.a += a
	// 課題 3.2
	t.name = name
	// 課題 3.3
	t.msg = msg

}

func main() {

	fmt.Println(plus(1, 2))
	fmt.Println(minus(2, 2))
	fmt.Println(multiply(3, 3))
	fmt.Println(divided(4, 4))

	fmt.Println(culc(plus, 1, 3))
	fmt.Println(culc(minus, 1, 3))
	fmt.Println(culc(multiply, 1, 3))
	fmt.Println(culc(divided, 1, 3))

	var t1 Turtle = Turtle{"最高", "師匠", 1000, 5, 180.5}
	var t2 Turtle = Turtle{"msg", "弟子", 10, 250, 270.3}

	fmt.Println(t1)
	fmt.Println(t2)

	// 課題 3.3
	fmt.Println("Turtleの", t1.name, ":", t1.msg)
}
