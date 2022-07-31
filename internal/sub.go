package internal

import (
	"log"
)

func (c *Calculator) Sub() {
	log.Println("Subtraction of two numbers:", c.a-c.b)
}
