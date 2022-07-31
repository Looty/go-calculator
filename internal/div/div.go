package div

import "log"

func (c *Calculator) Div() {
	log.Println("Division of two numbers:", c.a/c.b)
}
