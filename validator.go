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
