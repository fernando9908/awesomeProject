package main

import "fmt"

func main() {

	s := make([]string, 3)
	s[1] = "a"
	fmt.Println(s)
	s = append(s, "b")
	fmt.Println(s)
	fmt.Println(len(s))

	good := []string{"g", "o", "o", "d"}
	good = append(good, "!")
	fmt.Println(good)
}
