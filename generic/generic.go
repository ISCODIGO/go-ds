/*
Autor: https://github.com/zyedidia/generic
*/

package generic

import (
	"golang.org/x/exp/constraints"
)

// EqualsFn defines a function type that checks if 'a' and 'b' are equal.
type EqualsFn[T any] func(a, b T) bool

// LessFn defines a function type that checks if 'a' is less than 'b'.
type LessFn[T any] func(a, b T) bool

// HashFn defines a function type that computes a hash for 't'.
type HashFn[T any] func(t T) uint64

// Equals is a wrapper for the '==' operator for comparable types.
func Equals[T comparable](a, b T) bool {
	return a == b
}

// Less is a wrapper for the '<' operator for ordered types.
func Less[T constraints.Ordered](a, b T) bool {
	return a < b
}

// Compare returns an integer indicating the ordering of 'a' and 'b' using the provided 'less' function.
// It returns:
//   - -1 if a < b
//   -  1 if a > b
//   -  0 if a == b
func Compare[T any](a, b T, less LessFn[T]) int {
	switch {
	case less(a, b):
		return -1
	case less(b, a):
		return 1
	default:
		return 0
	}
}
