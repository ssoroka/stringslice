package stringslice

// Union is a set operation that returns the unique set of elements that are in s1 or in s2.
func Union(s1, s2 []string) []string {
	result := []string{}
	appendFunc := func(s string) {
		result = append(result, s)
	}
	Compare(s1, s2, appendFunc, appendFunc, appendFunc)
	return result
}

// Intersect is a set operation that returns the unique set of elements that are in s1 and in s2.
func Intersect(s1, s2 []string) []string {
	result := []string{}
	Compare(s1, s2, nil, func(s string) {
		result = append(result, s)
	}, nil)
	return result
}

var compareNoop = func(s string) {}

// Compare sorts and iterates s1 and s2. calling left() if the element is only in s1, right() if the element is only in s2, and equal() if it's in both.
// this is used as the speedy basis for other set operations.
func Compare(s1, s2 []string, left, equal, right func(s string)) {
	if left == nil {
		left = compareNoop
	}
	if right == nil {
		right = compareNoop
	}
	if equal == nil {
		equal = compareNoop
	}
	s1 = Uniq(Sort(s1))
	s2 = Uniq(Sort(s2))
	s1Counter := 0
	s2Counter := 0
	for s1Counter < len(s1) && s2Counter < len(s2) {
		if s1[s1Counter] < s2[s2Counter] {
			left(s1[s1Counter])
			s1Counter++
			continue
		}
		if s1[s1Counter] > s2[s2Counter] {
			right(s2[s2Counter])
			s2Counter++
			continue
		}
		// must be equal
		equal(s1[s1Counter])
		s1Counter++
		s2Counter++
	}
	// catch any remaining items
	for i := s1Counter; i < len(s1); i++ {
		left(s1[i])
	}
	for i := s2Counter; i < len(s2); i++ {
		right(s2[i])
	}
}

// Difference is a set operation that returns the elements from s1 that are not in s2.
func Difference(s1, s2 []string) []string {
	result := []string{}
	Compare(s1, s2, func(s string) {
		result = append(result, s)
	}, nil, nil)
	return result
}

// Distinct is an alias for Uniq
var Distinct = Uniq
