package main

import (
	"github.com/eugenewyj/go-sample/goinaction/ch05/ch0504/counters"
	"fmt"
)

func main() {
	counter62 := counters.New(10)
	fmt.Printf("Counter6_2: %d\n", counter62)
}
