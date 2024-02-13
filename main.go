package main

import (
	"fmt"

	older "github.com/alex210501/dependentpkg/older"
)

func main() {
	person := older.New("Alex", 20)

	fmt.Println(person)
	older.SetOlder(&person, 10)
	fmt.Println(person)
	older.SetYounger(&person, 20)
	fmt.Println(person)
}
