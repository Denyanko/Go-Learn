package main

import (
	"github.com/Denyanko/packages/helpers"
	"log"
)

func main() {
	log.Println("Hello")

	var myVar helpers.SomeType
	myVar.TypeName = "Some name"
	log.Println(myVar.TypeName)
}
