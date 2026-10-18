package main

import "fmt"

func fibonacci() func() int {
	s := []int{0, 1}
	return func() int {
		s = []int{s[1], s[0] + s[1]}
		return s[1]
	}
}

func main() {
	f := fibonacci()
	for range 10 {
		fmt.Println(f())
	}
}
