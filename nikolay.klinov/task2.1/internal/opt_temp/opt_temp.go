package opt_temp

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	"task2.1/internal/input"
)

const (
	MinOfficeTemp = 15
	MaxOfficeTemp = 30
)

func calculateComfortRange(currentMin, currentMax int, temperatureConstraint string) (int, int, error) {
	var operator string
	var preferredTemp int

	_, scanErr := fmt.Sscanf(temperatureConstraint, "%s %d", &operator, &preferredTemp)
	if scanErr != nil {
		return currentMin, currentMax, errors.New("invalid constraint format")
	}

	if preferredTemp > MaxOfficeTemp || preferredTemp < MinOfficeTemp {
		return currentMin, currentMax, errors.New("temperature out of range")
	}

	switch operator {
	case "<=":
		if preferredTemp < currentMax {
			currentMax = preferredTemp
		}
	case ">=":
		if preferredTemp > currentMin {
			currentMin = preferredTemp
		}
	default:
		return currentMin, currentMax, errors.New("unsupported operation")
	}

	return currentMin, currentMax, nil
}

func processDepartmentEmployees(employeeCount int) error {
	inputReader := bufio.NewReader(os.Stdin)
	departmentMinTemp := MinOfficeTemp
	departmentMaxTemp := MaxOfficeTemp

	for employeeIndex := 0; employeeIndex < employeeCount; employeeIndex++ {
		fmt.Printf("Enter constraint for employee %d (for example, '>= 20'): ", employeeIndex+1)
		employeeConstraint, readErr := inputReader.ReadString('\n')
		if readErr != nil {
			return errors.New("failed to read input")
		}

		var calcErr error
		departmentMinTemp, departmentMaxTemp, calcErr = calculateComfortRange(
			departmentMinTemp,
			departmentMaxTemp,
			employeeConstraint,
		)
		if calcErr != nil {
			return calcErr
		}

		if departmentMinTemp > departmentMaxTemp {
			fmt.Println("Optimal temperature: -1")
		} else {
			fmt.Printf("Optimal temperature: %d\n", departmentMinTemp)
		}
	}

	return nil
}

func Run() error {
	fmt.Print("Enter the number of departments: ")
	departmentCount, deptCountErr := input.InputNumber()
	if deptCountErr != nil {
		return deptCountErr
	}

	for deptIndex := 0; deptIndex < departmentCount; deptIndex++ {
		fmt.Print("Enter the number of employees: ")
		employeeCount, empCountErr := input.InputNumber()
		if empCountErr != nil {
			return empCountErr
		}

		processErr := processDepartmentEmployees(employeeCount)
		if processErr != nil {
			return processErr
		}
	}

	return nil
}
