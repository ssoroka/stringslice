package stringslice

import "sort"

type StringSlice []string

// New is a convenience wrapper for treating a slice of strings as a "stringslice"
func New(slice []string) StringSlice {
	return StringSlice(slice)
}

// Sort returns a new slice that is the sorted copy of the slice it was called on. Unlike sort.Strings, it does not mutate the original slice
func (ss StringSlice) Sort() StringSlice {
	ss2 := make(StringSlice, len(ss))
	copy(ss2, ss)
	sort.Strings(ss2)
	return StringSlice(ss2)
}

// Uniq returns a new slice that is sorted with all the duplicate strings removed.
func (ss StringSlice) Uniq() StringSlice {
	ss2 := ss.Sort()

	result := make(StringSlice, 0, len(ss))

	last := ""
	for i, str := range ss2 {
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

// Subtract the passed slice from the stringslice, returning a new slice of the result.
func (ss StringSlice) Subtract(slice []string) StringSlice {
	// todo: implement this without a map using sorted strings.
	otherElems := map[string]struct{}{}

	for _, e := range slice {
		otherElems[e] = struct{}{}
	}

	res := []string{}
	for _, e := range ss {
		if _, contains := otherElems[e]; !contains {
			res = append(res, e)
		}
	}
	return StringSlice(res)
}

// Map over each element in the slice and perform an operation on it. the result of the operation will replace the element value.
func (ss StringSlice) Map(f func(i int, s string) string) StringSlice {
	result := make(StringSlice, len(ss))
	for i, s := range ss {
		result[i] = f(i, s)
	}
	return result
}

// Slice returns the stringslice typecast to a []string slice
func (ss StringSlice) Slice() []string {
	return []string(ss)
}
