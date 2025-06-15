package utils

import (
	"errors"
	"fmt"
	"strconv"
)

func StingsToFloats(strings []string) ([]float64, error) {

	var floats []float64

	for _, strVal := range strings {
		floatVal, err := strconv.ParseFloat(strVal, 64)
		if err != nil {
			fmt.Println("Error converting prices to float")
			return nil, errors.New(fmt.Sprintf("Error converting prices to float: %s", strVal))
		}
		floats = append(floats, floatVal)
	}

	return floats, nil
}
