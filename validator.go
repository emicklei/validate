package validate

import "bytes"

type validator struct {
	errors []error
	name   string
	value  interface{}
}

func newValidator(name string, value interface{}) validator {
	return validator{errors: []error{}, name: name, value: value}
}

// IsError returns whether this validator collected at least one error.
func (v *validator) IsError() bool {
	return len(v.errors) > 0
}

// Message returns a newline separated string of all error messsages.
func (v *validator) Message() string {
	var buf bytes.Buffer
	for i, e := range v.errors {
		if i > 0 {
			buf.WriteString("\n")
		}
		buf.WriteString(e.Error())
	}
	return buf.String()
}

// Validation function checks the actual value and should return a readable error message if it fails.
type Validation func(name string, actual interface{}) error

// And calls the validation function passing the name and the actual value.
// As it returns an internal validator instance, this call must be last in the chain.
func (v *validator) And(f Validation) *validator {
	if err := f(v.name, v.value); err != nil {
		v.errors = append(v.errors, err)
	}
	return v
}
