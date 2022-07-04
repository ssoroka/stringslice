package stringslice

import (
	"fmt"
	"sort"
)

type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~float32 | ~float64 | ~string
}

type Sortable interface {
	Less(i, j int) bool
}

// Sort returns a new slice that is the sorted copy of the slice it was called on. Unlike sort.Strings, it does not mutate the original slice
func Sort[T Ordered](ss []T) []T {
	if ss == nil {
		return nil
	}
	ss2 := make([]T, len(ss))
	copy(ss2, ss)
	sort.Slice(ss2, func(i int, j int) bool {
		return ss2[i] <= ss2[j]
	})
	return ss2
}

// SortBy returns a new, slice that is the sorted copy of the slice it was called on, using sortFunc to interpret the string as a sortable integer value. It does not mutate the original slice
func SortBy[T Ordered](ss []T, sortFunc func(slice []T, i, j int) bool) []T {
	if ss == nil {
		return nil
	}
	ss2 := make([]T, len(ss))
	copy(ss2, ss)
	sort.Slice(ss2, func(i, j int) bool {
		return sortFunc(ss2, i, j)
	})
	return ss2
}

// Uniq returns a new slice that is sorted with all the duplicate strings removed.
func Uniq[T Ordered](ss []T) []T {
	return SortedUniq[T](Sort(ss)) // TODO: uniq without sorting.
}

func SortedUniq[T Ordered](ss []T) []T {
	if ss == nil {
		return nil
	}
	result := []T{}
	last := *new(T)
	for i, s := range ss {
		if i != 0 && last == s {
			continue
		}
		result = append(result, s)
		last = s
	}
	return result
}

// Subtract the passed slice from the []string, returning a new slice of the result. It's rather memory abusive, and Difference might be a better option.
func Subtract[T Ordered](ss []T, str ...T) []T {
	otherElems := map[T]struct{}{}

	for _, e := range str {
		otherElems[e] = struct{}{}
	}

	res := []T{}
	for _, e := range ss {
		if _, contains := otherElems[e]; !contains {
			res = append(res, e)
		}
	}
	return res
}

// Add is a convenience alias for append. it returns a nice slice with the passed slice appended
func Add[T Ordered](ss []T, slice ...T) []T {
	return append(ss, slice...)
}

// Map over each element in the slice and perform an operation on it. the result of the operation will replace the element value.
// Normal func structure is func(i int, s string) string.
// Also accepts func structure func(s string) string
func Map[T any](ss []T, funcInterface interface{}) []T {
	if ss == nil {
		return nil
	}
	if funcInterface == nil {
		return ss
	}
	f := func(i int, s T) T {
		switch tf := funcInterface.(type) {
		case func(int, T) T:
			return tf(i, s)
		case func(T) T:
			return tf(s)
		}
		panic(fmt.Sprintf("Map cannot understand function type %T", funcInterface))
	}
	result := make([]T, len(ss))
	for i, s := range ss {
		result[i] = f(i, s)
	}
	return result
}

// Each iterates over each element in the slice and perform an operation on it.
// Normal func structure is func(i int, s string).
// Also accepts func structure func(s string)
func Each[T any](ss []T, funcInterface interface{}) {
	if ss == nil {
		return
	}
	if funcInterface == nil {
		return
	}
	f := func(i int, s T) {
		switch tf := funcInterface.(type) {
		case func(int, T):
			tf(i, s)
		case func(T):
			tf(s)
		default:
			panic(fmt.Sprintf("Each cannot understand function type %T", funcInterface))
		}
	}
	for i, s := range ss {
		f(i, s)
	}
}

type AccumulatorFunc[T any] func(acc T, i int, s T) T

// Reduce (aka inject) iterates over the slice of items and calls the accumulator function for each pass, storing the state in the acc variable through each pass.
func Reduce[T any](items []T, initialAccumulator T, f AccumulatorFunc[T]) T {
	if items == nil {
		return initialAccumulator
	}
	acc := initialAccumulator
	for i, s := range items {
		acc = f(acc, i, s)
	}
	return acc
}

// Contains returns true if the string is in the slice.
// Note: If you .Sort() the slice first, this function will do a log2(n) binary search through the list, which is much faster for large lists.
func Contains[T comparable](ss []T, s T) bool {
	return Index(ss, s) != -1
}

func SortedContains[T Ordered](ss []T, s T) bool {
	return SortedIndex(ss, s) != -1
}

// Index returns the index of string in the slice, otherwise -1 if the string is not found.
// Note: .SortedIndex() will do a log2(n) binary search through the list, which is much faster for large lists.
func Index[T comparable](ss []T, s T) int {
	for i, b := range ss {
		if b == s {
			return i
		}
	}
	return -1
}

// SortedIndex returns the index of string in the slice, otherwise -1 if the string is not found.
// this function will do a log2(n) binary search through the list, which is much faster for large lists.
// The slice must be sorted in ascending order.
func SortedIndex[T Ordered](ss []T, s T) int {
	idx := sort.Search(len(ss), func(i int) bool {
		return ss[i] >= s
	})
	if idx >= 0 && idx < len(ss) && ss[idx] == s {
		return idx
	}
	return -1
}

// First returns the First element, or "" if there are no elements in the slice.
func First[T any](ss []T) T {
	if len(ss) > 0 {
		return ss[0]
	}
	return *new(T)
}

// Last returns the Last element, or "" if there are no elements in the slice.
func Last[T any](ss []T) T {
	if len(ss) > 0 {
		return ss[len(ss)-1]
	}
	return *new(T)
}

// Any returns true if the length is greater than zero
func Any[T any](ss []T) bool {
	return len(ss) != 0
}

// ToStringSlice converts any slice of string-like types to a []string, and panics if the type cannot be converted.
func ToStringSlice[T Ordered](o []T) []string {
	result := make([]string, len(o))

	for i := 0; i < len(o); i++ {
		result[i] = fmt.Sprint(o[i])
	}
	return result
}

func SliceToInterfaceSlice[T any](ss []T) []interface{} {
	result := make([]interface{}, len(ss))
	for i := range ss {
		result[i] = ss[i]
	}
	return result
}

func Filter[T any](ss []T, funcInterface interface{}) []T {
	f := func(i int, s T) bool {
		switch tf := funcInterface.(type) {
		case func(int, T) bool:
			return tf(i, s)
		case func(T) bool:
			return tf(s)
		default:
			panic(fmt.Sprintf("Filter cannot understand function type %T", funcInterface))
		}
	}

	result := []T{}

	for i, s := range ss {
		if f(i, s) {
			result = append(result, s)
		}
	}
	return result
}

func DeleteIf[T any](ss []T, funcInterface interface{}) []T {
	f := func(i int, s T) bool {
		switch tf := funcInterface.(type) {
		case func(int, T) bool:
			return tf(i, s)
		case func(T) bool:
			return tf(s)
		default:
			panic(fmt.Sprintf("Filter cannot understand function type %T", funcInterface))
		}
	}

	result := []T{}

	for i, s := range ss {
		if !f(i, s) {
			result = append(result, s)
		}
	}
	return result
}
