package main

import (
	"fmt"

	"task2.1/internal/opt_temp"
)

func main() {
	runErr := opt_temp.Run()

	if runErr != nil {
		fmt.Printf("Error: %s\n", runErr)
		return
	}
}
