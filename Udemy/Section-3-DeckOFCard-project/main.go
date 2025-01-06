package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func (p Person) print() {
	fmt.Printf("%+v", p)

}

func (p *Person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func hello(p *Person) {
	fmt.Println(p)
}

type slice []int

func (s slice) modify() {
	for i := range s {
		s[i] = 0
	}

}

func main() {
	//cards := createDeck()
	// hand, remainingdec := deal(cards, 5)
	// hand.print()
	// remainingdec.print()
	// cards.toString()
	// // cards.saveToFile("my_cards")
	// cards := newDeckFromFile("my_cards")
	// cards.print()
	// cards.shuffle()
	// cards.print()

	arr := slice{1, 2, 3, 4, 5}
	fmt.Println(arr)
	arr.modify()
	fmt.Println(arr)

	// jim := Person{
	// 	firstName: "Jim",
	// 	lastName:  "Party",
	// 	contact: contactInfo{
	// 		email:   "smdnak@gmail.com",
	// 		zipCode: 94000,
	// 	},
	// }
	// hello(&jim)
	// jim.print()
	// jim.updateName("Jimmy")
	// jim.print()
	// fmt.Println(alex)

}

// var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

// func main() {
// 	var i int
// 	for i = 0; i < len(pow); i++ {
// 		fmt.Printf("2**%d = %d\n", i, pow[i])
// 	}
// 	for i, v := range pow {
// 		fmt.Printf("2**%d = %d\n", i, v)
// 	}
// }
