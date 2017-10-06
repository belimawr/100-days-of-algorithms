package db

import (
	"fmt"
	"testing"

	"github.com/wwgberlin/timelord/vc"
)

const (
	alice = iota
	bob
	cathy
	dave
)

func TestScenario(t *testing.T) {
	db := &DB{}

	ta := vc.New(4, alice, nil)
	tb := vc.New(4, bob, nil)
	tc := vc.New(4, cathy, nil)
	td := vc.New(4, dave, nil)

	ta.Tick()
	if err := db.Set("Wednesday", ta); err != nil {
		t.Errorf("failed to set: %s", err)
	}
	r := db.Get()
	if len(r) != 1 || r[0].Value != "Wednesday" {
		t.Errorf("got smth strange: %v", r)
	}
	fmt.Printf("Response: %v\n", r)

	tb.Merge(ta)
	tb.Tick()
	if err := db.Set("Tuesday", tb); err != nil {
		t.Errorf("failed to set: %s", err)
	}
	r = db.Get()
	if len(r) != 1 || r[0].Value != "Tuesday" {
		t.Errorf("got smth strange: %v", r)
	}
	fmt.Printf("Response: %v\n", r)

	td.Merge(tb)
	td.Tick()
	if err := db.Set("Tuesday", td); err != nil {
		t.Errorf("failed to set: %s", err)
	}
	r = db.Get()
	if len(r) != 1 || r[0].Value != "Tuesday" {
		t.Errorf("got smth strange: %v", r)
	}
	fmt.Printf("Response: %v\n", r)

	tc.Merge(ta)
	tc.Tick()
	if err := db.Set("Thursday", tc); err != nil {
		t.Errorf("failed to set: %s", err)
	}
	r = db.Get()
	if len(r) != 2 {
		t.Errorf("got smth strange: %v", r)
	}
	fmt.Printf("Response: %v\n", r)

	if err := db.Set("Thursday", ta); err == nil {
		t.Error("expected to get an error")
	} else {
		fmt.Printf("As expected got an error: %s\n", err)
	}

}
