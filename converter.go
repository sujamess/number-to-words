package numbertowords

import (
	"fmt"
	"strconv"
)

func convertStringToFloat64(str string) (float64, error) {
	f64, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return f64, fmt.Errorf("Cannot parse %s to type float64", str)
	}

	return f64, nil
}
