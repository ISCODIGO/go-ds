/*

Autor: https://github.com/zyedidia/generic
*/

package generic

import (
	"golang.org/x/exp/constraints"
)

// EqualsFn is a function that returns whether 'a' and 'b' are equal.
type EqualsFn[T any] func(a, b T) bool

// LessFn is a function that returns whether 'a' is less than 'b'.
type LessFn[T any] func(a, b T) bool

// HashFn is a function that returns the hash of 't'.
type HashFn[T any] func(t T) uint64

// Equals wraps the '==' operator for comparable types.
func Equals[T comparable](a, b T) bool {
	return a == b
}

// Less wraps the '<' operator for ordered types.
func Less[T constraints.Ordered](a, b T) bool {
	return a < b
}

// Compare uses a less function to determine the ordering of 'a' and 'b'. It returns:
//
// * -1 if a < b
//
// * 1 if a > b
//
// * 0 if a == b
func Compare[T any](a, b T, less LessFn[T]) int {
	if less(a, b) {
		return -1
	} else if less(b, a) {
		return 1
	}
	return 0
}