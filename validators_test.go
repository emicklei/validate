package validate

import (
	"fmt"
	"log"
	"testing"
)

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

func ExampleIsPositive() {
	result := Int("length", -1).IsPositive("%v must be positive, got %v")
	if result.IsError() {
		log.Println(result.Message())
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

func ExampleIsBetween() {
	result := Int("grade", 12).IsBetween(1, 10, "%v must be between %v and %v, got %v")
	if result.IsError() {
		log.Println(result.Message())
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

func TestCustomChecker(t *testing.T) {
	i := 41
	result := Int("years", i).
		IsPositive("%v must be positive, got %v").
		And(IsOdd)
	if !result.IsError() {
		t.Fail()
	}
	if result.Message() != "years is not even, got 41" {
		t.Errorf(result.Message())
	}
}
