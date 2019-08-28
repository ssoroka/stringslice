package stringslice

import (
	"fmt"
	"reflect"
	"sort"
)

// Sort returns a new slice that is the sorted copy of the slice it was called on. Unlike sort.Strings, it does not mutate the original slice
func Sort(ss []string) []string {
	if ss == nil {
		return nil
	}
	ss2 := make([]string, len(ss))
	copy(ss2, ss)
	sort.Strings(ss2)
	return ss2
}

// SortBy returns a new, slice that is the sorted copy of the slice it was called on, using sortFunc to interpret the string as a sortable integer value. It does not mutate the original slice
func SortBy(ss []string, sortFunc func(slice []string, i, j int) bool) []string {
	if ss == nil {
		return nil
	}
	ss2 := make([]string, len(ss))
	copy(ss2, ss)
	sort.Slice(ss2, func(i, j int) bool {
		return sortFunc(ss2, i, j)
	})
	return ss2
}

// Uniq returns a new slice that is sorted with all the duplicate strings removed.
func Uniq(ss []string) []string {
	return SortedUniq(Sort(ss)) // TODO: uniq without sorting.
}

func SortedUniq(ss []string) []string {
	if ss == nil {
		return nil
	}
	result := []string{}
	last := ""
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
func Subtract(ss []string, str ...string) []string {
	otherElems := map[string]struct{}{}

	for _, e := range str {
		otherElems[e] = struct{}{}
	}

	res := []string{}
	for _, e := range ss {
		if _, contains := otherElems[e]; !contains {
			res = append(res, e)
		}
	}
	return res
}

// Add is a convenience alias for append. it returns a nice slice with the passed slice appended
func Add(ss []string, slice ...string) []string {
	return append(ss, slice...)
}

// Map over each element in the slice and perform an operation on it. the result of the operation will replace the element value.
// Normal func structure is func(i int, s string) string.
// Also accepts func structure func(s string) string
func Map(ss []string, funcInterface interface{}) []string {
	if ss == nil {
		return nil
	}
	if funcInterface == nil {
		return ss
	}
	f := func(i int, s string) string {
		switch tf := funcInterface.(type) {
		case func(int, string) string:
			return tf(i, s)
		case func(string) string:
			return tf(s)
		}
		panic(fmt.Sprintf("Map cannot understand function type %T", funcInterface))
	}
	result := make([]string, len(ss))
	for i, s := range ss {
		result[i] = f(i, s)
	}
	return result
}

// Each iterates over each element in the slice and perform an operation on it.
// Normal func structure is func(i int, s string).
// Also accepts func structure func(s string)
func Each(ss []string, funcInterface interface{}) {
	if ss == nil {
		return
	}
	if funcInterface == nil {
		return
	}
	f := func(i int, s string) {
		switch tf := funcInterface.(type) {
		case func(int, string):
			tf(i, s)
		case func(string):
			tf(s)
		default:
			panic(fmt.Sprintf("Each cannot understand function type %T", funcInterface))
		}
	}
	for i, s := range ss {
		f(i, s)
	}
}

type AccumulatorFunc func(acc string, i int, s string) string
type AccumulatorIntFunc func(acc int64, i int, s string) int64

// Reduce (aka inject) iterates over the slice of items and calls the accumulator function for each pass, storing the state in the acc variable through each pass.
func Reduce(ss []string, initialAccumulator string, f AccumulatorFunc) string {
	if ss == nil {
		return initialAccumulator
	}
	acc := initialAccumulator
	for i, s := range ss {
		acc = f(acc, i, s)
	}
	return acc
}

// ReduceInt (aka inject) iterates over the slice of items and calls the accumulator function for each pass, storing the state in the acc variable through each pass.
func ReduceInt(ss []string, initialAccumulator int64, f AccumulatorIntFunc) int64 {
	if ss == nil {
		return initialAccumulator
	}
	acc := initialAccumulator
	for i, s := range ss {
		acc = f(acc, i, s)
	}
	return acc
}

// Contains returns true if the string is in the slice.
// Note: If you .Sort() the slice first, this function will do a log2(n) binary search through the list, which is much faster for large lists.
func Contains(ss []string, s string) bool {
	return Index(ss, s) != -1
}

func SortedContains(ss []string, s string) bool {
	return SortedIndex(ss, s) != -1
}

// Index returns the index of string in the slice, otherwise -1 if the string is not found.
// Note: If you .Sort() the slice first, this function will do a log2(n) binary search through the list, which is much faster for large lists.
func Index(ss []string, s string) int {
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
func SortedIndex(ss []string, s string) int {
	idx := sort.Search(len(ss), func(i int) bool {
		return ss[i] >= s
	})
	if idx >= 0 && idx < len(ss) && ss[idx] == s {
		return idx
	}
	return -1
}

// First returns the First element, or "" if there are no elements in the slice.
func First(ss []string) string {
	if len(ss) > 0 {
		return ss[0]
	}
	return ""
}

// Last returns the Last element, or "" if there are no elements in the slice.
func Last(ss []string) string {
	if len(ss) > 0 {
		return ss[len(ss)-1]
	}
	return ""
}

// Any returns true if the length is greater than zero
func Any(ss []string) bool {
	return len(ss) != 0
}

// ToStringSlice converts any slice of string-like types to a []string, and panics if the type cannot be converted.
func ToStringSlice(o interface{}) []string {
	v := reflect.ValueOf(o)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Slice && v.Kind() != reflect.Array {
		panic("ToStringSlice cannot convert type " + v.Type().String() + " (kind " + v.Kind().String() + ") to slice.")
	}
	result := make([]string, v.Len())
	for i := 0; i < v.Len(); i++ {
		el := v.Index(i)
		if el.Kind() != reflect.String {
			panic(fmt.Sprintf("ToStringSlice failed to convert Element %d of slice to a string, it's a kind %s", i, el.Kind().String()))
		}
		result[i] = el.String()
	}
	return result
}

func StringSliceToInterfaceSlice(ss []string) []interface{} {
	result := make([]interface{}, len(ss))
	for i := range ss {
		result[i] = ss[i]
	}
	return result
}

func Filter(ss []string, funcInterface interface{}) []string {
	f := func(i int, s string) bool {
		switch tf := funcInterface.(type) {
		case func(int, string) bool:
			return tf(i, s)
		case func(string) bool:
			return tf(s)
		default:
			panic(fmt.Sprintf("Filter cannot understand function type %T", funcInterface))
		}
	}

	result := []string{}

	for i, s := range ss {
		if f(i, s) {
			result = append(result, s)
		}
	}
	return result
}

func DeleteIf(ss []string, funcInterface interface{}) []string {
	f := func(i int, s string) bool {
		switch tf := funcInterface.(type) {
		case func(int, string) bool:
			return tf(i, s)
		case func(string) bool:
			return tf(s)
		default:
			panic(fmt.Sprintf("Filter cannot understand function type %T", funcInterface))
		}
	}

	result := []string{}

	for i, s := range ss {
		if !f(i, s) {
			result = append(result, s)
		}
	}
	return result
}
