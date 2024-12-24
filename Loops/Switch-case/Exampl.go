package main

import (
	"fmt"
)

func main() {
	fmt.Print("Go runs on ")
	switch os := "olko"; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}
