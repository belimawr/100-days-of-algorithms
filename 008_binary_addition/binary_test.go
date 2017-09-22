package main

import (
	"reflect"
	"testing"
)

func Test_reversed(t *testing.T) {
	str := []string{"1", "2", "3"}

	if !reflect.DeepEqual(reversed(str), []string{"3", "2", "1"}) {
		t.Error("The list was not reversed")
	}
}

func Test_binaryAdder_invalid_token(t *testing.T) {
	r, err := binaryAdder("A", "0")

	if err == nil {
		t.Error("Expecting an error")
	}

	if r != "" {
		t.Errorf("Expecting an empty string, got %s", r)
	}
}

func Test_binaryAdder_00(t *testing.T) {
	op1 := "0"
	op2 := "0"
	expected := "0"

	r, err := binaryAdder(op1, op2)
	if err != nil {
		t.Errorf("Did not expect an error: %q", err.Error())
	}

	if r != expected {
		t.Errorf("Got %q, expecting %q", r, expected)
	}
}

func Test_binaryAdder_01(t *testing.T) {
	op1 := "0"
	op2 := "1"
	expected := "1"

	r, err := binaryAdder(op1, op2)
	if err != nil {
		t.Errorf("Did not expect an error: %q", err.Error())
	}

	if r != expected {
		t.Errorf("Got %q, expecting %q", r, expected)
	}
}

func Test_binaryAdder_carry(t *testing.T) {
	op1 := "1"
	op2 := "1"
	expected := "10"

	r, err := binaryAdder(op1, op2)
	if err != nil {
		t.Errorf("Did not expect an error: %q", err.Error())
	}

	if r != expected {
		t.Errorf("Got %q, expecting %q", r, expected)
	}
}

func Test_binaryAdder_middle_carry(t *testing.T) {
	op1 := "11"
	op2 := "11"
	expected := "110"

	r, err := binaryAdder(op1, op2)
	if err != nil {
		t.Errorf("Did not expect an error: %q", err.Error())
	}

	if r != expected {
		t.Errorf("Got %q, expecting %q", r, expected)
	}
}

func Test_binaryAdder_n1_smaller(t *testing.T) {
	op1 := "0"
	op2 := "1000"
	expected := "1000"

	r, err := binaryAdder(op1, op2)
	if err != nil {
		t.Errorf("Did not expect an error: %q", err.Error())
	}

	if r != expected {
		t.Errorf("Got %q, expecting %q", r, expected)
	}
}

func Test_binaryAdder_n2_smaller(t *testing.T) {
	op1 := "1000"
	op2 := "1"
	expected := "1001"

	r, err := binaryAdder(op1, op2)
	if err != nil {
		t.Errorf("Did not expect an error: %q", err.Error())
	}

	if r != expected {
		t.Errorf("Got %q, expecting %q", r, expected)
	}
}
