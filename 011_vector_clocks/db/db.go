package db

import (
	"errors"
	"fmt"
	"sync"

	"github.com/wwgberlin/timelord/vc"
)

type Record struct {
	Value     string
	Timestamp vc.Timestamp
}

type DB struct {
	mu     sync.RWMutex
	values []Record
}

func (db *DB) String() string {
	return fmt.Sprintf("DB: %v", db.values)
}

func (db *DB) Get() []Record {
	db.mu.RLock()
	defer db.mu.RUnlock()
	return db.values
}

func (db *DB) Set(value string, ts vc.Timestamp) error {
	db.mu.Lock()
	defer db.mu.Unlock()
	r, err := db.findRecord(ts)
	if err != nil {
		return err
	}
	if r == nil {
		db.values = append(db.values, Record{
			Value:     value,
			Timestamp: ts,
		})
	} else {
		r.Value = value
		r.Timestamp = ts
	}
	return nil
}

var TimePassed = errors.New("too late")

func (db *DB) findRecord(ts vc.Timestamp) (*Record, error) {
	for i := range db.values {
		r := ts.HappensBefore(db.values[i].Timestamp)
		if r == vc.NotComparable {
			continue
		}
		if r == vc.HappensBefore {
			return nil, TimePassed
		}
		return &db.values[i], nil
	}
	return nil, nil
}
