package main

type triangle struct {
	base   float64
	height float64
}

type square struct {
	sideLength float64
}

type shape interface {
	getarea() float64
}

func (t triangle) getarea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) getarea() float64 {
	return s.sideLength * s.sideLength
}

func printArea(s shape) {
	println(s.getarea())
}

// type bot interface {
// 	getGreeting() string
// }

// type englishBot struct{}
// type spanishBot struct{}

// func (englishBot) getGreeting() string {
// 	return "Hi There!"
// }

// func (spanishBot) getGreeting() string {
// 	return "Hola!"
// }

//	func printGreeting(b bot) {
//		fmt.Println(b.getGreeting())
//	}
