package validate

import "bytes"

type Validator struct {
	errors []error
	name   string
	value  interface{}
}

func newValidator(name string, value interface{}) Validator {
	return Validator{errors: []error{}, name: name, value: value}
}

// IsError returns whether this Validator collected at least one error.
func (v *Validator) IsError() bool {
	return len(v.errors) > 0
}

// Message returns a newline separated string of all error messsages.
func (v *Validator) Message() string {
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
// As it returns an internal Validator instance, this call must be last in the chain.
func (v *Validator) And(f Validation) *Validator {
	if err := f(v.name, v.value); err != nil {
		v.errors = append(v.errors, err)
	}
	return v
}
