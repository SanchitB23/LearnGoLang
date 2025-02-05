package conversion

import (
	"fmt"
	"strconv"
)

func StringsToFloats(strings []string) []float64 {
	var floats []float64
	for _, str := range strings {
		float, err := strconv.ParseFloat(str, 64)
		if err != nil {
			fmt.Println("Error parsing float:", err)
			continue
		}
		floats = append(floats, float)
	}
	return floats
}
