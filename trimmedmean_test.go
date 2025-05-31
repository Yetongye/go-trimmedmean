package trimmedmean

import (
	"math"
	"testing"
)

func TestSymmetricTrimmedMean(t *testing.T) {
	data := []float64{1, 2, 3, 4, 5, 100}
	result, err := TrimmedMean(data, 0.17) // trim 1 from each end
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
	expected := (2 + 3 + 4 + 5) / 4.0
	if math.Abs(result-expected) > 0.01 {
		t.Errorf("Expected %.2f, got %.2f", expected, result)
	}
}

func TestAsymmetricTrimmedMean(t *testing.T) {
	data := []float64{1, 2, 3, 4, 5, 6, 100}
	result, err := TrimmedMean(data, 0.1428, 0.2857)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
	expected := (2 + 3 + 4 + 5) / 4.0
	if result != expected {
		t.Errorf("Expected %.2f, got %.2f", expected, result)
	}
}
