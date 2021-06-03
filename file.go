package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	firstName string
	lastName  string
	birthYear int
}

func (p *Person) uniqueId() string {
	return string(p.firstName[0]) + p.lastName + strconv.Itoa(p.birthYear)
}

func main() {
	gary := Person{"Gary", "Lime", 1982}

	id := gary.uniqueId()

	fmt.Println(id)
}
