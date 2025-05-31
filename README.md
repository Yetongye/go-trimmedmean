# go-trimmedmean

The trimmed mean package provides a function to read a slice of integers and/or floats and compute a trimmed mean, where the degree-of-trimming arguments represent the proportions of slice elements to be taken from the lower and upper ends of the slice after ordering from the lowest value to the highest value. 

## Package Overview 

The TrimmedMean function computes a mean after removing a proportion of the lowest and/or highest values from a sorted data slice.

- Symmetric trimming: TrimmedMean(data, 0.05)

- Asymmetric trimming: TrimmedMean(data, 0.05, 0.1)

Supports input of []float64, and returns the trimmed average.

## How to Get and Use the Package

To install and use the trimmed mean package:

```bash 
go get github.com/Yetongye/go-trimmedmean
```


Import the package in your Go code:

```go 
import "github.com/Yetongye/go-trimmedmean"
```


Example usage:

```go 
result, err := trimmedmean.TrimmedMean(data, 0.05)        // symmetric
result, err := trimmedmean.TrimmedMean(data, 0.05, 0.1)   // asymmetric
```

### Unit test:

The`go-trimmedmean`package includes full test coverage in`trimmedmean_test.go`, covering:

- Symmetric and asymmetric trimming
- Empty slice errors
- Invalid trimming parameters

Run tests:

```bash 
go test
```

Sample output:

```go 
PASS
ok      github.com/Yetongye/go-trimmedmean      0.114s
```
