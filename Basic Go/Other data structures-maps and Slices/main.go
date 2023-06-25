package main

import (
	"log"
	"sort"
)

type User struct {
	FirstName string
	LastName  string
}

func main() {
	// string map example
	myMap := make(map[string]string)

	myMap["dog"] = "Samson"
	myMap["other-dog"] = "Cassie"

	myMap["dog"] = "Fido"

	log.Println(myMap["dog"])
	log.Println(myMap["other-dog"])

	// int map example
	myMap2 := make(map[string]int)

	myMap2["first"] = 1
	myMap2["second"] = 2
	log.Println(myMap2["first"], myMap2["second"])

	// custom type map example
	myMap3 := make(map[string]User)

	me := User{
		FirstName: "Ade",
		LastName:  "Boernama",
	}

	myMap3["me"] = me
	log.Println(myMap3["me"].FirstName)

	// Slices
	var mySlice []int

	mySlice = append(mySlice, 2)
	mySlice = append(mySlice, 1)
	mySlice = append(mySlice, 3)

	sort.Ints(mySlice)

	log.Println(mySlice)

	// Another way to declare a slice
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	log.Println(numbers)

	log.Println(numbers[6:9])

	// Slice of string example
	names := []string{"one", "seven", "fish", "cat"}
	log.Println(names)
}
