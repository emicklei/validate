package validate

import "fmt"

type intVar struct {
	validator
}

// Int returns a new named variable of type int.
func Int(name string, v int) *intVar {
	return &intVar{validator: newValidator(name, v)}
}

// IsPositive validates a number for positive.
// Error format: name,actual
func (i *intVar) IsPositive(format string) *intVar {
	actual := i.value.(int)
	if actual < 0 {
		i.errors = append(i.errors, fmt.Errorf(format, i.name, actual))
	}
	return i
}

// IsBetween validates a number to an interval.
// Error format: name,min,max,actual
func (i *intVar) IsBetween(min, max int, format string) *intVar {
	actual := i.value.(int)
	if actual > max || actual < min {
		i.errors = append(i.errors, fmt.Errorf(format, i.name, min, max, actual))
	}
	return i
}

type stringVar struct {
	validator
}

// String returns a new named variable of type string.
func String(name string, s string) *stringVar {
	return &stringVar{validator: newValidator(name, s)}
}
