package validate

import "fmt"

type IntVar struct {
	Validator
}

// Int returns a new named variable of type int.
func Int(name string, v int) *IntVar {
	return &IntVar{Validator: newValidator(name, v)}
}

// IsPositive validates a number for positive.
// Error format: name,actual
func (i *IntVar) IsPositive(format string) *IntVar {
	actual := i.value.(int)
	if actual < 0 {
		i.errors = append(i.errors, fmt.Errorf(format, i.name, actual))
	}
	return i
}

// IsBetween validates the value inside an interval.
// Error format: name,min,max,actual
func (i *IntVar) IsBetween(min, max int, format string) *IntVar {
	actual := i.value.(int)
	if actual > max || actual < min {
		i.errors = append(i.errors, fmt.Errorf(format, i.name, min, max, actual))
	}
	return i
}

type StringVar struct {
	Validator
}

// String returns a new named variable of type string.
func String(name string, s string) *StringVar {
	return &StringVar{Validator: newValidator(name, s)}
}

// HasLengthBetween validates the length of the string.
// Error format: name,min,max,actual
func (s *StringVar) HasLengthBetween(min, max int, format string) *StringVar {
	actual := len(s.value.(string))
	if actual > max || actual < min {
		s.errors = append(s.errors, fmt.Errorf(format, s.name, min, max, actual))
	}
	return s
}

// Condition return a Validator with an error if the condition is false.
// Error format: left,right
func Condition(left, right interface{}, condition bool, format string) *Validator {
	validator := newValidator("", condition)
	if !condition {
		validator.errors = append(validator.errors, fmt.Errorf(format, left, right))
	}
	return &validator
}
