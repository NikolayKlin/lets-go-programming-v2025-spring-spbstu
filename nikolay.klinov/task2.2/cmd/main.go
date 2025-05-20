package main

import (
	"fmt"

	"task2.2/internal/chooseDishes"
)

func main() {
	runErr := chooseDishes.Run()

	if runErr != nil {
		fmt.Printf("Error: %s\n", runErr)
		return
	}
}
