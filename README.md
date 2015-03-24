Validate
========

Go package for writing validations.

First you create a named, typed variable.
Then you call one or more validating functions.
You must provide the error message template ; it is not part of the package.
Inspect the result and process its error message if needed.

	if result := validate.Int("years", i).IsPositive("%v must be positive, got %v"); result.IsError() {
		log.Println(result.Message())
	}	

	result := validate.Int("years", i).
		IsPositive("%v must be positive, got %v").
		IsBetween(2, 10, "%v must be between %v and %v, got %v")
		
	if result.IsError() {
		log.Println(result.Message())
	}

You can create your own validator functions

	func IsOdd(name string, value interface{}) error {
		i, ok := value.(int)
		if !ok {
			return fmt.Errorf("%s is not an int, got %v (%T)", name, value, value)
		}
		if i%2 != 0 {
			return fmt.Errorf("%s is not even, got %v", name, value)
		}
		return nil
	}
	
and use it with together with existing functions

	result := validate.Int("marbles", given)
		.IsPositive("%v must be positive, got %v")
		.And(IsOdd)

- [Documentation on godoc.org](http://godoc.org/github.com/emicklei/validate)

(c) 2015, http://ernestmicklei.com. MIT License