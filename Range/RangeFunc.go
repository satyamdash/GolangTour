package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	var i int
	for i = 0; i < len(pow); i++ {
		fmt.Printf("2**%d = %d\n", i, pow[i])
	}
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
