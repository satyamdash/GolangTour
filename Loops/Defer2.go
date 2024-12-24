package main

import "fmt"

// Defer : A defer statement defers the execution of a function until the surrounding function returns.
// The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.
// Defer is commonly used to simplify functions that perform various clean-up actions.
// Defer uses stack to store the deferred function calls.
func Defer() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
func main() {
	Defer()
}
