package add

import (
	"log"
)

func (c *Calculator) Add() {
	log.Println("Addition of two numbers:", c.a+c.b)
}
