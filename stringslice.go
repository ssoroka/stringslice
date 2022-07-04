package stringslice

import (
	"strings"
)

// StringSlice is an alias for []string that adds some functions
type Slice[T Ordered] struct {
	sl     []T
	sorted bool
}

// New is a convenience wrapper for treating a slice of strings as a "stringslice"
func New[T Ordered](slice []T) Slice[T] {
	return Slice[T]{sl: slice}
}

// Sort returns a new slice that is the sorted copy of the slice it was called on. Unlike sort.Strings, it does not mutate the original slice
func (ss Slice[T]) Sort() Slice[T] {
	if ss.sl == nil {
		return ss
	}
	return Slice[T]{sl: Sort(ss.sl), sorted: true}
}

// Uniq returns a new slice that is sorted with all the duplicate strings removed.
// Note: sorting the string first will make this much faster.
func (ss Slice[T]) Uniq() Slice[T] {
	if ss.sorted {
		return Slice[T]{sl: SortedUniq[T](ss.sl), sorted: true}
	}
	return Slice[T]{sl: Uniq[T](ss.sl)}
}

// Subtract the passed slice from the Slice, returning a new slice of the result.
func (ss Slice[T]) Subtract(str ...T) Slice[T] {
	return Slice[T]{sl: Subtract[T](ss.sl, str...)}
}

// Add is a convenience alias for append. it returns a nice slice with the passed slice appended
func (ss Slice[T]) Add(slice ...T) Slice[T] {
	return Slice[T]{sl: Add[T](ss.sl, slice...)}
}

// Map over each element in the slice and perform an operation on it. the result of the operation will replace the element value.
// Normal func structure is func(i int, s string) string.
// Also accepts func structure func(s string) string
func (ss Slice[T]) Map(funcInterface interface{}) Slice[T] {
	return Slice[T]{sl: Map[T](ss.sl, funcInterface)}
}

// Reduce (aka inject) iterates over the slice of items and calls the accumulator function for each pass, storing the state in the acc variable through each pass.
func (ss Slice[T]) Reduce(initialAccumulator T, f AccumulatorFunc[T]) T {
	return Reduce[T](ss.sl, initialAccumulator, f)
}

// Slice returns the Slice typecast to the underlying slice type
func (ss Slice[T]) Slice() []T {
	return ss.sl
}

// InterfaceSlice returns the Slice typecast to a []interface slice, as some libraries expect this.
func (ss Slice[T]) InterfaceSlice() []interface{} {
	return SliceToInterfaceSlice[T](ss.sl)
}

// Contains returns true if the string is in the slice.
// Note: If you .Sort() the slice first, this function will do a log2(n) binary search through the list, which is much faster for large lists.
func (ss Slice[T]) Contains(s T) bool {
	if ss.sorted {
		return SortedContains[T](ss.sl, s)
	}
	return Contains[T](ss.sl, s)
}

// Index returns the index of string in the slice, otherwise -1 if the string is not found.
// Note: If you .Sort() the slice first, this function will do a log2(n) binary search through the list, which is much faster for large lists.
func (ss Slice[T]) Index(s T) int {
	if ss.sorted {
		return SortedIndex[T](ss.sl, s)
	}
	return Index[T](ss.sl, s)
}

// First returns the First element, or "" if there are no elements in the slice.
func (ss Slice[T]) First() T {
	return First[T](ss.sl)
}

// Last returns the Last element, or "" if there are no elements in the slice.
func (ss Slice[T]) Last() T {
	return Last[T](ss.sl)
}

// Any returns true if the length is greater than zero
func (ss Slice[T]) Any() bool {
	return Any[T](ss.sl)
}

// Len returns the slice length for convenience
func (ss Slice[T]) Len() int {
	return len(ss.sl)
}

// Join is a convenience wrapper for strings.Join
func (ss Slice[T]) Join(sep string) string {
	return strings.Join(ToStringSlice[T](ss.sl), sep)
}
