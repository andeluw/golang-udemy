package conversion

import (
	"errors"
	"fmt"
	"strconv"
)

func StringsToFloats(strings []string) ([]float64, error) {
	var floats []float64
	for _, stringVal := range strings {
		floatVal, err := strconv.ParseFloat(stringVal, 64)

		if err != nil {
			fmt.Println("Failed to convert to float!")
			fmt.Println(err)
			return nil, errors.New("failed to convert to float")
		}

		floats = append(floats, floatVal)
	}

	return floats, nil
}