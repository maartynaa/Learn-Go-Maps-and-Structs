package main

import "fmt"

type Pet struct {
	name    string
	petType string
	age     int
}

func main() {
	nuggets := Pet{name: "Nuggets", petType: "dog", age: 4}
	mittens := Pet{"Mittens", "cat", 7}
	robin := Pet{"Robin", "bird", 2}

	pets := []Pet{nuggets, mittens, robin}

	pets[1].age = 5

	fmt.Println(nuggets)
	fmt.Println(mittens)
	fmt.Println(pets)
	fmt.Println(pets[2].petType)

	pointerToRobin := &robin
	fmt.Println(pointerToRobin.name)
}
