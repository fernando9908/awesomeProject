package main

import (
	"errors"
	"fmt"
)

type user struct {
	name     string
	password string
}

func findUser(users []user, name string) (v *user, err error) {
	for _, u := range users {
		if u.name == name {
			return &u, nil
		}
	}
	return nil, errors.New("not found")
}

func main() {
	u, err := findUser([]user{{"wang", "haha"}}, "wang")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(u.name)

	if u, err := findUser([]user{{"wang", "haha"}}, "wang"); err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(u.name)
	}
	fmt.Println(u.name)
}
