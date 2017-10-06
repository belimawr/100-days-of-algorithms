package vc

import (
	"reflect"
)

func (t Timestamp) HappensBefore(r Timestamp) int {
	if reflect.DeepEqual(t.vc, r.vc) {
		return Equals
	}

	var greatThan, lessThan bool
	var ts, rs int

	for i := 0; i < len(t.vc); i++ {

		if t.vc[i] > r.vc[i] {
			greatThan = true
		} else if t.vc[i] < r.vc[i] {
			lessThan = true
		}

		ts += t.vc[i]
		rs += r.vc[i]
	}

	if greatThan && lessThan {
		if t.resolver != nil {
			return t.resolver.Resolve(t, r)
		}
		return NotComparable
	}

	if ts < rs {
		return HappensBefore
	}

	return HappensAfter
}

// Merge happens when actors in 2 different timelines communicate, for example when a message with timestamp is recieved.
// What is reuired:
// - update each element in the vector to be max(local, received)
// - increment the logical clock value representing the current node in the vector
// r remote timeline, recieved as part of a message
func (t Timestamp) Merge(r Timestamp) {
	for i := 0; i < len(t.vc); i++ {
		if r.vc[i] > t.vc[i] {
			t.vc[i] = r.vc[i]
		}
	}
	t.Tick()
}

func (r MajorityResolver) Resolve(t Timestamp, o Timestamp) int {
	max := func(v []int) int {
		m := v[0]
		for i := 1; i < len(v); i++ {
			if m < v[i] {
				m = v[i]
			}
		}
		return m
	}

	tm := max(t.vc)
	om := max(o.vc)

	if tm > om {
		return HappensAfter
	} else if om > tm {
		return HappensBefore
	}

	return Equals
}
