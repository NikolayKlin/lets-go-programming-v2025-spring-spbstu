package calculate

import (
	"errors"
	"fmt"
)

func CalcSuitableTemp(n, employesQuantity uint64) ([]int64, error) {

	var minTemp, maxTemp uint64 = 15, 30
	for j := uint64(0); j < employesQuantity; j++ {
		var op string
		var temp uint64
		fmt.Printf("Предпочтения сотрудника %d в отделе %d: ", j, n)
		_, err := fmt.Scanf("%s %d", &op, &temp)
		if err != nil {
			return nil, err
		}

		switch op {
		case ">=":
			if temp > minTemp {
				minTemp = temp
			}
		case "<=":
			if temp < maxTemp {
				maxTemp = temp
			}
		default:
			err = errors.New("bad op arg")
			return nil, err
		}

	}

	err := checkInput(minTemp, maxTemp)
	if err != nil {
		return []int64{-1}, nil
	}

	return fillSlice(int64(minTemp), int64(maxTemp)), nil
}

func fillSlice(minTemp, maxTemp int64) []int64 {
	slice := []int64{}

	for i := minTemp; i <= maxTemp; i++ {
		slice = append(slice, i)
	}

	return slice
}

func checkInput(minTemp, maxTemp uint64) error {
	if minTemp > maxTemp {
		return errors.New("minTemp > maxTemp")
	}

	if maxTemp < 15 {
		return errors.New("bad maxTemp")
	} else if minTemp > 30 {
		return errors.New("bad minTemp")
	}

	return nil
}
