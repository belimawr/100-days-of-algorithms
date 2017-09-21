package main

import (
	"reflect"
	"testing"
)

func Test_Stack_Push(t *testing.T) {
	stack := Stack{}
	expected := Stack{1, 2, 3}

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	if !reflect.DeepEqual(expected, stack) {
		t.Errorf("Expecting %#v, got: %#v", expected, stack)
	}
}

func Test_Stack_Pop(t *testing.T) {
	stack := Stack{1, 2, 3}
	expectedStack := Stack{1, 2}
	expectedNumeber := 3.0

	n := stack.Pop()

	if !reflect.DeepEqual(expectedStack, stack) {
		t.Errorf("Expecting %#v, got: %#v", expectedStack, stack)
	}

	if n != expectedNumeber {
		t.Errorf("Expecting %f, got %f", expectedNumeber, n)
	}
}

func Test_multiplication(t *testing.T) {
	expected := float64(42)

	fn := operations["*"]

	result := fn(7, 6)

	if result != expected {
		t.Errorf("Expecting %f, got %f", expected, result)
	}
}

func Test_division(t *testing.T) {
	expected := float64(42)

	fn := operations["/"]

	result := fn(84, 2)

	if result != expected {
		t.Errorf("Expecting %f, got %f", expected, result)
	}
}

func Test_subtraction(t *testing.T) {
	expected := float64(42)

	fn := operations["-"]

	result := fn(44, 2)

	if result != expected {
		t.Errorf("Expecting %f, got %f", expected, result)
	}
}

func Test_addition(t *testing.T) {
	expected := float64(42)

	fn := operations["+"]

	result := fn(40, 2)

	if result != expected {
		t.Errorf("Expecting %f, got %f", expected, result)
	}
}

func Test_postfix_notation(t *testing.T) {
	expression := "40 2 +"
	expected := float64(42)

	result, err := postfix_notation(expression)

	if err != nil {
		t.Error("Expecting an error")
	}

	if result != expected {
		t.Errorf("Expecting %f, got %f", expected, result)
	}
}

func Test_postfix_notation_invalid_token(t *testing.T) {
	expression := "A"
	expected := float64(0)

	result, err := postfix_notation(expression)

	if err == nil {
		t.Error("Expecting an error")
	}

	if result != expected {
		t.Errorf("Expecting %f, got %f", expected, result)
	}
}

func Test_postfix_notation_too_few_elements(t *testing.T) {
	expression := "1 +"
	expected := float64(0)

	result, err := postfix_notation(expression)

	if err == nil {
		t.Error("Expecting an error")
	}

	if result != expected {
		t.Errorf("Expecting %f, got %f", expected, result)
	}
}
