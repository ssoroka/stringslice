package stringslice

import (
	"strings"
)

// StringSlice is an alias for []string that adds some functions
type StringSlice struct {
	sl     []string
	sorted bool
}

// New is a convenience wrapper for treating a slice of strings as a "stringslice"
func New(slice []string) StringSlice {
	return StringSlice{sl: slice}
}

// Sort returns a new slice that is the sorted copy of the slice it was called on. Unlike sort.Strings, it does not mutate the original slice
func (ss StringSlice) Sort() StringSlice {
	if ss.sl == nil {
		return ss
	}
	return StringSlice{sl: Sort(ss.sl), sorted: true}
}

// Uniq returns a new slice that is sorted with all the duplicate strings removed.
// Note: sorting the string first will make this much faster.
func (ss StringSlice) Uniq() StringSlice {
	if ss.sorted {
		return StringSlice{sl: SortedUniq(ss.sl), sorted: true}
	}
	return StringSlice{sl: Uniq(ss.sl)}
}

// Subtract the passed slice from the stringslice, returning a new slice of the result.
func (ss StringSlice) Subtract(str ...string) StringSlice {
	return StringSlice{sl: Subtract(ss.sl, str...)}
}

// Add is a convenience alias for append. it returns a nice slice with the passed slice appended
func (ss StringSlice) Add(slice ...string) StringSlice {
	return StringSlice{sl: Add(ss.sl, slice...)}
}

// Map over each element in the slice and perform an operation on it. the result of the operation will replace the element value.
// Normal func structure is func(i int, s string) string.
// Also accepts func structure func(s string) string
func (ss StringSlice) Map(funcInterface interface{}) StringSlice {
	return StringSlice{sl: Map(ss.sl, funcInterface)}
}

// Reduce (aka inject) iterates over the slice of items and calls the accumulator function for each pass, storing the state in the acc variable through each pass.
func (ss StringSlice) Reduce(initialAccumulator string, f AccumulatorFunc) string {
	return Reduce(ss.sl, initialAccumulator, f)
}

// ReduceInt (aka inject) iterates over the slice of items and calls the accumulator function for each pass, storing the state in the acc variable through each pass.
func (ss StringSlice) ReduceInt(initialAccumulator int64, f AccumulatorIntFunc) int64 {
	return ReduceInt(ss.sl, initialAccumulator, f)
}

// Slice returns the stringslice typecast to a []string slice
func (ss StringSlice) Slice() []string {
	return ss.sl
}

// InterfaceSlice returns the stringslice typecast to a []interface slice, as some libraries expect this.
func (ss StringSlice) InterfaceSlice() []interface{} {
	return StringSliceToInterfaceSlice(ss.sl)
}

// Contains returns true if the string is in the slice.
// Note: If you .Sort() the slice first, this function will do a log2(n) binary search through the list, which is much faster for large lists.
func (ss StringSlice) Contains(s string) bool {
	if ss.sorted {
		return SortedContains(ss.sl, s)
	}
	return Contains(ss.sl, s)
}

// Index returns the index of string in the slice, otherwise -1 if the string is not found.
// Note: If you .Sort() the slice first, this function will do a log2(n) binary search through the list, which is much faster for large lists.
func (ss StringSlice) Index(s string) int {
	if ss.sorted {
		return SortedIndex(ss.sl, s)
	}
	return Index(ss.sl, s)
}

// First returns the First element, or "" if there are no elements in the slice.
func (ss StringSlice) First() string {
	return First(ss.sl)
}

// Last returns the Last element, or "" if there are no elements in the slice.
func (ss StringSlice) Last() string {
	return Last(ss.sl)
}

// Any returns true if the length is greater than zero
func (ss StringSlice) Any() bool {
	return Any(ss.sl)
}

// Len returns the slice length for convenience
func (ss StringSlice) Len() int {
	return len(ss.sl)
}

// Join is a convenience wrapper for strings.Join
func (ss StringSlice) Join(sep string) string {
	return strings.Join(ss.sl, sep)
}
