package mul

import "log"

func (c *Calculator) Mul() {
	log.Println("Multiplication of two numbers:", c.a*c.b)
}
