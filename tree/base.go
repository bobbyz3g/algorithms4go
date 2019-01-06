package tree

type Comparable interface {
	// CompareTo compares two Interface.
	// a > p return 1.
	// a = p return 0.
	// a < p return -1.
	CompareTo(p interface{}) int
}

type Key Comparable
type Value interface{}
