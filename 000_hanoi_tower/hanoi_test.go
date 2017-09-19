package main

import (
	"reflect"
	"testing"
)

func Test_Pop(t *testing.T) {
	l := Stack{3, 2, 1}

	el := l.Pop()

	if !reflect.DeepEqual(l, Stack{3, 2}) {
		t.Errorf("Not equals: %#v", l)
	}

	if el != 1 {
		t.Error("Removed element must be 1")
	}
}

func Test_Push(t *testing.T) {
	l := Stack{3, 2}

	l.Push(1)

	if !reflect.DeepEqual(l, Stack{3, 2, 1}) {
		t.Errorf("Not equals: %#v", l)
	}
}

func Test_move_lt_1(t *testing.T) {
	A := &Stack{1}
	B := &Stack{10}
	C := &Stack{100}

	Move(A, B, C, 0)

	if !reflect.DeepEqual(*A, Stack{1}) {
		t.Error("Not equals")
	}

	if !reflect.DeepEqual(*B, Stack{10}) {
		t.Error("Not equals")
	}

	if !reflect.DeepEqual(*C, Stack{100}) {
		t.Error("Not equals")
	}
}

func Test_move_eq_1(t *testing.T) {
	A := &Stack{1}
	B := &Stack{}
	C := &Stack{}

	Move(A, B, C, 1)

	if !reflect.DeepEqual(*A, Stack{}) {
		t.Error("Not equals")
	}

	if !reflect.DeepEqual(*B, Stack{}) {
		t.Error("Not equals")
	}

	if !reflect.DeepEqual(*C, Stack{1}) {
		t.Error("Not equals")
	}
}

func Test_move_gt_1(t *testing.T) {
	A := &Stack{2, 1}
	B := &Stack{}
	C := &Stack{}

	Move(A, B, C, 2)

	if !reflect.DeepEqual(*A, Stack{}) {
		t.Error("Not equals")
	}

	if !reflect.DeepEqual(*B, Stack{}) {
		t.Error("Not equals")
	}

	if !reflect.DeepEqual(*C, Stack{2, 1}) {
		t.Error("Not equals")
	}
}
