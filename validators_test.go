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

func ExampleIntVar_IsPositive() {
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
		t.Error(result.Message())
	}
}

func ExampleIntVar_IsBetween() {
	result := Int("grade", 12).IsBetween(1, 10, "%v must be between %v and %v, got %v")
	if result.IsError() {
		log.Println(result.Message())
	}
}

func TestChainInt(t *testing.T) {
	i := -1
	result := Int("years", i).
		IsPositive("%v must be positive, got %v").
		IsBetween(2, 10, "%v must be between %v and %v, got %v")
	if !result.IsError() {
		t.Fail()
	}
	if result.Message() != `years must be positive, got -1
years must be between 2 and 10, got -1` {
		t.Error(result.Message())
	}
}

func TestStringHasLength(t *testing.T) {
	s := "hello"
	result := String("greeting", s).HasLengthBetween(6, 12, "%s length must be between %v and %v, got %v")
	if result.Message() != "greeting length must be between 6 and 12, got 5" {
		t.Error(result.Message())
	}
}

func ExampleStringVar_HasLengthBetween() {
	s := "hello"
	result := String("greeting", s).HasLengthBetween(6, 12, "%s length must be between %v and %v, got %v")
	if result.IsError() {
		log.Println(result.Message())
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
		t.Error(result.Message())
	}
}

func TestCompareTwoVars(t *testing.T) {
	begin := 10
	end := 0
	result := Condition(begin, end, begin < end, "begin = %v must be less than end = %v")
	if !result.IsError() {
		t.Fail()
	}
	if result.Message() != "begin = 10 must be less than end = 0" {
		t.Error(result.Message())
	}
}

func ExampleCondition() {
	start := 5
	end := 1
	result := Condition(start, end, start < end, "start day = %v must be before end day = %v")
	if result.IsError() {
		log.Println(result.Message())
	}
}
