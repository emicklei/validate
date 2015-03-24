Validate
========

Go package for writing validations

	if result := validate.Int("years", i).IsPositive("%v must be positive, got %v"); result.IsError() {
		log.Println(result.Message())
	}	

	result := validate.Int("years", i).
		IsPositive("%v must be positive, got %v").
		IsBetween(2, 10, "%v must be between %v and %v, got %v")
	if result.IsError() {
		log.Println(result.Message())
	}


- [Documentation on godoc.org](http://godoc.org/github.com/emicklei/validate)

(c) 2015, http://ernestmicklei.com. MIT License