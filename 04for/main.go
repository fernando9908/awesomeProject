package main

import "fmt"

func main() {
	i := 1
	for {
		fmt.Println("loop")
		if i == 3 {
			break
		}
		i++
	}
}
