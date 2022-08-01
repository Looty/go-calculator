## go-calculator
Simple calculator made in go


## Usage
```golang
package main

import (
  "log"
  "github.com/Looty/go-calculator/calculator"
)

func main() {
	a, _ := calculator.Add(4, 16)
	log.Println(a)
	
	b, _ := calculator.Sub(10, 2)
	log.Println(b)
	
	c, _ := calculator.Mul(3, 3)
	log.Println(c)
	
	d, _ := calculator.Div(20, 5)
	log.Println(d)
}
```
