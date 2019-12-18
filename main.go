package main

import(
	e "github.com/PlatformBrigade/go_training/exportpackage"
	"fmt"
)
func main() {
	counter := e.AlertCount(10)
	fmt.Printf("Counter: %d\n", counter)

}
