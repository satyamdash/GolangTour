package main

func main() {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i := 0; i < len(arr); i++ {
		if arr[i]%2 == 0 {
			println(arr[i], "is even")
		} else {
			println(arr[i], "is odd")
		}
	}
}
