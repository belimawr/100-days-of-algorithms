package main

import (
	"reflect"
	"testing"
)

func Test_Pop(t *testing.T) {
	l := List{
		lst: []uint{3, 2, 1},
	}

	el := l.Pop()

	if !reflect.DeepEqual(l.lst, []uint{3, 2}) {
		t.Errorf("Not equals: %#v", l)
	}

	if el != 1 {
		t.Error("Removed element must be 1")
	}
}

func Test_Push(t *testing.T) {
	l := List{
		lst: []uint{3, 2},
	}

	l.Push(1)

	if !reflect.DeepEqual(l.lst, []uint{3, 2, 1}) {
		t.Errorf("Not equals: %#v", l)
	}
}

func Test_move_lt_1(t *testing.T) {
	A := &List{[]uint{1}}
	B := &List{[]uint{10}}
	C := &List{[]uint{100}}

	Move(A, B, C, 0)

	if !reflect.DeepEqual(A.lst, []uint{1}) {
		t.Error("Not equals")
	}

	if !reflect.DeepEqual(B.lst, []uint{10}) {
		t.Error("Not equals")
	}

	if !reflect.DeepEqual(C.lst, []uint{100}) {
		t.Error("Not equals")
	}
}

func Test_move_eq_1(t *testing.T) {
	A := &List{[]uint{1}}
	B := &List{[]uint{}}
	C := &List{[]uint{}}

	Move(A, B, C, 1)

	if !reflect.DeepEqual(A.lst, []uint{}) {
		t.Error("Not equals")
	}

	if !reflect.DeepEqual(B.lst, []uint{}) {
		t.Error("Not equals")
	}

	if !reflect.DeepEqual(C.lst, []uint{1}) {
		t.Error("Not equals")
	}
}

func Test_move_gt_1(t *testing.T) {
	A := &List{[]uint{2, 1}}
	B := &List{[]uint{}}
	C := &List{[]uint{}}

	Move(A, B, C, 2)

	if !reflect.DeepEqual(A.lst, []uint{}) {
		t.Error("Not equals")
	}

	if !reflect.DeepEqual(B.lst, []uint{}) {
		t.Error("Not equals")
	}

	if !reflect.DeepEqual(C.lst, []uint{2, 1}) {
		t.Error("Not equals")
	}
}
