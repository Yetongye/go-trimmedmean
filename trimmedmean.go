package trimmedmean

import (
	"errors"
	"math"
	"sort"
)

// TrimmedMean computes the trimmed mean of a slice.
// if only one trim proportion is provided, symmetric trimming is applied
// if two values are provided, asymmetric trimming is applied
func TrimmedMean(data []float64, trimLow float64, trimHighOptional ...float64) (float64, error) {
	n := len(data)
	if n == 0 {
		return 0, errors.New("data slice is empty")
	}

	// Ensure proportions are valid
	if trimLow < 0 || trimLow >= 0.5 {
		return 0, errors.New("trimLow must be between 0 and 0.5")
	}

	trimHigh := trimLow
	if len(trimHighOptional) > 0 {
		trimHigh = trimHighOptional[0]
		if trimHigh < 0 || trimHigh >= 0.5 {
			return 0, errors.New("trimHigh must be between 0 and 0.5")
		}
	}

	// Sort the slice
	sort.Float64s(data)

	// Calculate number of elements to trim
	low := int(math.Floor(trimLow * float64(n)))
	high := int(math.Floor(trimHigh * float64(n)))

	if low+high >= n {
		return 0, errors.New("trimming too much data")
	}

	trimmed := data[low : n-high]
	sum := 0.0
	for _, val := range trimmed {
		sum += val
	}

	return sum / float64(len(trimmed)), nil
}
