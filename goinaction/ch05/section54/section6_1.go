package main

import (
	"github.com/eugenewyj/go-sample/goinaction/ch05/section54/counters"
	"fmt"
)

func main() {
	counter := counters.AlertCounter6_1(10)
	fmt.Printf("Counter6_1: %d\n", counter)
}
