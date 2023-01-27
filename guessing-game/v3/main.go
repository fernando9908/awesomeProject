package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	maxNum := 100
	rand.Seed(time.Now().UnixNano())
	secretNumber := rand.Intn(maxNum)
	fmt.Println("The secret number is ", secretNumber)

	// 获取输入后续有更简洁的写法
	fmt.Println("Please input your guess")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		return
	}
	// Windows系统上换行需要\r\n，即回到行首同时换行
	input = strings.TrimSuffix(input, "\r\n")
	guess, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Invalid input. Please enter an integer value")
		return
	}
	fmt.Println(guess)

}
