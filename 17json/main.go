package main

import (
	"encoding/json"
	"fmt"
)

type userinInfo struct {
	Name  string
	Age   int `json:"age"`
	Hobby []string
}

func main() {
	a := userinInfo{Name: "wang", Age: 18, Hobby: []string{"Golang", "TypeScript"}}
	buf, err := json.Marshal(a)
	if err != nil {
		panic(err)
	}
	fmt.Println(buf)
	fmt.Println(string(buf))

	buf, err = json.MarshalIndent(a, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(buf))

	var b userinInfo
	err = json.Unmarshal(buf, &b)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", b)
}
