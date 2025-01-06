package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{Y: 1, X: 2}
	//	v.X = 4
	fmt.Println(v)
	fmt.Printf("X: %d, Y: %d\n", v.X, v.Y)
}
