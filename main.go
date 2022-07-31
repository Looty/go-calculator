package main

import (
	"github.com/Looty/go-calculator/calculator"
	"log"
)

func main() {
	a, _ := calculator.Add(1, 2)
	log.Println(a)

	b, err := calculator.Div(1, 4)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(b)
}
