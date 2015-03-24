package validate

import "testing"

func TestPositiveInt(t *testing.T) {
	i := -1
	result := Int("years", i).IsPositive("%v must be positive, got %v")
	if !result.IsError() {
		t.Fail()
	}
	if result.Message() != "years must be positive, got -1" {
		t.Fail()
	}
}

func TestBetweenInt(t *testing.T) {
	i := 12
	result := Int("years", i).IsBetween(2, 10, "%v must be between %v and %v, got %v")
	if !result.IsError() {
		t.Error("must be error")
	}
	if result.Message() != "years must be between 2 and 10, got 12" {
		t.Errorf(result.Message())
	}
}

func TestWithInt(t *testing.T) {
	i := -1
	result := Int("years", i).
		IsPositive("%v must be positive, got %v").
		IsBetween(2, 10, "%v must be between %v and %v, got %v")
	if !result.IsError() {
		t.Fail()
	}
	if result.Message() != `years must be positive, got -1
years must be between 2 and 10, got -1` {
		t.Errorf(result.Message())
	}
}

//func TestPositiveString(t *testing.T) {
//	result := String("title", "mr").IsPositive("%v is not positive")
//	if !result.IsError() {
//		t.Fail()
//	}
//	if result.Message() != "title is not a number, got mr (string)" {
//		t.Errorf("expected %s got %v", "title is not a number", result.Message())
//	}
//}
