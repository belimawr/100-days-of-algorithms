package vc

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	ts := New(5, 1, nil)
	if len(ts.vc) != 5 {
		t.Errorf("unexpected ts size: %d", len(ts.vc))
	}
}

func TestTick(t *testing.T) {
	ts := New(5, 1, nil)
	ts.Tick()
	if ts.vc[1] != 1 {
		t.Errorf("expected to see a tick: %v", ts)
	}
}

func TestMerge(t *testing.T) {
	ts := New(5, 1, nil)
	ts.Tick()
	r := New(5, 0, nil)
	r.Tick()
	ts.Merge(r)
	if ts.vc[0] != 1 || ts.vc[1] != 2 {
		t.Errorf("expected different merge result: %v", ts)
	}
}

func TestScenario(t *testing.T) {
	const (
		a = iota
		b
		c
	)
	ta := New(3, a, nil)
	tb := New(3, b, nil)
	tc := New(3, c, nil)
	tc.Tick()
	tb.Merge(tc)
	tb.Tick()

	ta.Merge(tb)
	ta.Tick()

	tb.Tick()
	tc.Merge(tb)

	tb.Merge(ta)
	tb.Tick()

	tc.Tick()
	ta.Merge(tc)

	tc.Merge(tb)
	tc.Tick()
	ta.Merge(tc)

	assert := func(ts Timestamp, ref []int) {
		for i := range ts.vc {
			if ts.vc[i] != ref[i] {
				t.Errorf("i: %d, expected %d got %d", i, ref[i], ts.vc[i])
			}
		}
	}

	fmt.Printf("a:%v  b:%v c:%v\n", ta, tb, tc)
	assert(ta, []int{4, 5, 5})
	assert(tb, []int{2, 5, 1})
	assert(tc, []int{2, 5, 5})
}

func TestHappensBefore(t *testing.T) {
	cases := []struct {
		t   Timestamp
		r   Timestamp
		res int
	}{
		{
			t:   Timestamp{vc: []int{0, 0, 0}},
			r:   Timestamp{vc: []int{0, 0, 0}},
			res: Equals,
		},
		{
			t:   Timestamp{vc: []int{0, 1, 0}},
			r:   Timestamp{vc: []int{0, 0, 0}},
			res: HappensAfter,
		},
		{
			t:   Timestamp{vc: []int{0, 0, 0}},
			r:   Timestamp{vc: []int{0, 0, 1}},
			res: HappensBefore,
		},
		{
			t:   Timestamp{vc: []int{0, 1, 0}},
			r:   Timestamp{vc: []int{0, 0, 1}},
			res: NotComparable,
		},
	}
	for i, tc := range cases {
		got := tc.t.HappensBefore(tc.r)
		if got != tc.res {
			t.Errorf("%d: expected %d, got %d", i, tc.res, got)
		}
	}
}

type mockResolver struct{}

func (m mockResolver) Resolve(t, o Timestamp) int {
	return 42
}

func TestHappensBeforeResolver(t *testing.T) {
	expected := 42
	ts := Timestamp{
		vc:       []int{0, 1, 0},
		resolver: mockResolver{},
	}
	rs := Timestamp{vc: []int{0, 0, 1}}

	resp := ts.HappensBefore(rs)

	if resp != expected {
		t.Errorf("Got %d, expecting %d", resp, expected)
	}
}

func TestMajorityResolverResolve(t *testing.T) {
	resolver := MajorityResolver{}

	cases := []struct {
		t   Timestamp
		r   Timestamp
		res int
	}{
		{
			t:   Timestamp{vc: []int{0, 2, 0}},
			r:   Timestamp{vc: []int{0, 0, 1}},
			res: HappensAfter,
		},
		{
			t:   Timestamp{vc: []int{0, 1, 0}},
			r:   Timestamp{vc: []int{2, 0, 1}},
			res: HappensBefore,
		},
		{
			t:   Timestamp{vc: []int{0, 1, 0}},
			r:   Timestamp{vc: []int{0, 1, 0}},
			res: Equals,
		},
	}

	for i, tc := range cases {
		got := resolver.Resolve(tc.t, tc.r)
		if got != tc.res {
			t.Errorf("%d: expected %d, got %d", i, tc.res, got)
		}
	}
}
