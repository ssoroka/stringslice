package stringslice

import "sort"

// SortedSlice is a string slice that has been sorted. .Sort returns this type. Searches through sorted string slices can be much faster
type SortedSlice StringSlice

func (ss SortedSlice) Map(funcInterface interface{}) StringSlice {
	return StringSlice(ss).Map(funcInterface)
}

// Contains returns true if the string is in the slice.
// Note: If you .Sort() the slice first, this function will do a log2(n) binary search through the list, which is much faster for large lists.
func (ss SortedSlice) Contains(s string) bool {
	return ss.Index(s) != -1
}

// Index returns the index of string in the slice, otherwise -1 if the string is not found.
// Note: If you .Sort() the slice first, this function will do a log2(n) binary search through the list, which is much faster for large lists.
func (ss SortedSlice) Index(s string) int {
	idx := sort.Search(len(ss), func(i int) bool {
		return ss[i] >= s
	})
	if ss[idx] == s {
		return idx
	}
	return -1
}

func index(s string, start, end int) int {
	return -1
}

// Uniq returns a new slice that is sorted with all the duplicate strings removed.
// Note: sorting the string first will make this much faster.
func (ss SortedSlice) Uniq() SortedSlice {
	if ss == nil {
		return nil
	}

	result := make(SortedSlice, 0, len(ss))

	last := ""
	for i, str := range ss {
		if i != 0 {
			if str == last {
				continue
			}
		}
		result = append(result, str)
		last = str
	}
	return result
}
