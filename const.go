package constant

import "fmt"

// ConstSet is a typed set of named constants that supports string parsing.
type ConstSet[V fmt.Stringer] struct {
	members    []V
	defaultVal V
}

// NewConstSet creates a ConstSet with the given default value and members.
func NewConstSet[V fmt.Stringer](defaultVal V, members ...V) ConstSet[V] {
	return ConstSet[V]{members: members, defaultVal: defaultVal}
}

// Parse converts a string to a member of the set.
// Returns the default value and an error if no member matches.
func (cs ConstSet[V]) Parse(value string) (V, error) {
	for _, m := range cs.members {
		if m.String() == value {
			return m, nil
		}
	}
	return cs.defaultVal, fmt.Errorf("%q is not a valid constant", value)
}

// Members returns all members of the set.
func (cs ConstSet[V]) Members() []V {
	return cs.members
}

// Default returns the default value of the set.
func (cs ConstSet[V]) Default() V {
	return cs.defaultVal
}
