package main

import "log"

func main() {
	// Basic of loop
	for i := 0; i <= 5; i++ {
		log.Println(i)
	}

	// Loop/range over slice
	animals := []string{"dog", "fish", "horse", "cow", "cat"}

	for i, animal := range animals {
		log.Println(i, animal)
	}

	// Loop without index
	for _, animal := range animals {
		log.Println(animal)
	}

	// Range over a map
	animals2 := make(map[string]string)

	animals2["cheetah"] = "Citah"
	animals2["giraffe"] = "Jerapah"

	for animalType, animal := range animals2 {
		log.Println(animalType, animal)
	}

	// Range over a string
	var firstLine = "Once upon a midnight dreary"

	for i, l := range firstLine {
		// it will print the byte of the letter
		log.Println(i, ":", l)
	}

	// Range over a custom type
	type User struct {
		FirstName string
		LastName  string
		Email     string
		Age       int
	}

	var users []User
	users = append(users, User{"John", "Smith", "john@smith.com", 30})
	users = append(users, User{"Mary", "Jones", "mary@jones.com", 20})
	users = append(users, User{"Sally", "Brown", "sally@brown.com", 45})
	users = append(users, User{"Alex", "Anderson", "alex@anderson.com", 17})

	for _, user := range users {
		log.Println(user.FirstName, user.LastName, user.Email, user.Age)
	}
}
