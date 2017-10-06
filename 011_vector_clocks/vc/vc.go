package vc

// Timestamp is a represenation of timestamap based on vector clock
type Timestamp struct {
	vc       []int
	line     int
	resolver Resolver
}

type Resolver interface {
	Resolve(t Timestamp, o Timestamp) int
}

type MajorityResolver struct{}

// New creates a new timestamp in a group timelines. For examle each server or
// actor in distributed system may have its own timeline.
// size is a group size - e.g. number of servers
// line - index of a particular timeline in the group which will be considered as a local
// time for this vector clock (like time in particular server)
func New(size, line int, resolver Resolver) Timestamp {
	return Timestamp{
		vc:       make([]int, size),
		line:     line,
		resolver: resolver,
	}
}

// Tick is clock move in local timeline. Whenever a process does work, increment the logical clock value of the node in the vector
func (t Timestamp) Tick() {
	t.vc[t.line]++
}

const (
	// Equals means that two timestamps are equal
	Equals = 0
	// HappensBefore means that the reciver timestamp happened before
	HappensBefore = -1
	// HappensAfter means that the reciver timestamp happened after
	HappensAfter = 1
	// NotComparable means that "happens before" relatioinship does not exists for 2 given timestamps, they are independent
	NotComparable = -100
)
