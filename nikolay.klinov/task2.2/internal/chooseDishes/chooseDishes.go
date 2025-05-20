package chooseDishes

import (
	"errors"
	"fmt"

	"task2.2/internal/heap"
	"task2.2/internal/input"
)

func Run() error {
	fmt.Print("Enter the number of dishes (1-10000): ")
	dishCount, err := input.InputNumber()

	if err != nil {
		return err
	}

	if !input.CheckRange(dishCount, 1, 10000) {
		return errors.New("Number of dishes out of range\n")
	}

	fmt.Print("Enter dish preferences separated by spaces: ")
	preferences, err := input.InputPreferences(dishCount)

	if err != nil {
		return err
	}

	fmt.Print("Enter the value of k (1 to N): ")
	k, err := input.InputNumber()

	if err != nil {
		return err
	}

	if !input.CheckRange(k, 1, dishCount) {
		return errors.New("k is out of range\n")
	}

	kthLargest := heap.FindKthLargest(preferences, k)
	fmt.Printf("The %d-th most preferred dish has a score of: %d\n", k, kthLargest)

	return nil
}
