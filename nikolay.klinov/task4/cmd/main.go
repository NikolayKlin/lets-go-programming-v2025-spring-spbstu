package main

import (
	"flag"
	"fmt"
	"os"
	"task4/internal/polynomial"
	"task4/internal/solver"
)

func main() {
	degree := flag.Int("degree", 5, "Degree of polynomial to find")
	useSync := flag.Bool("sync", false, "Use synchronous version")
	flag.Parse()

	s := solver.NewSolver()
	var result *polynomial.Polynomial

	if *useSync {
		fmt.Println("Using synchronous version")
		result = s.FindIrreducibleSync(*degree)
	} else {
		fmt.Println("Using asynchronous version")
		result = s.FindIrreducibleAsync(*degree)
	}

	if result != nil {
		fmt.Printf("Found irreducible polynomial: %s\n", result)
	} else {
		fmt.Println("No irreducible polynomial found")
		os.Exit(1)
	}
}
