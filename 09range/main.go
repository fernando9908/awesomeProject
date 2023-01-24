package main

import "fmt"

func main() {

	nums := []int{2, 3, 4}
	sum := 0
	for i, num := range nums {
		sum += num
		if num == 2 {
			fmt.Println("index:", i)
		}
	}
	fmt.Println(sum)

}
