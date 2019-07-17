# stringslice

See tests for example usage.

Adds functions to string slices: 

- Sort (does not mutate)
- SortBy (Sort, but supply your own sorting function)
- Uniq (currently implies also sorted)
- SortedUniq (only Uniq on an already-sorted string slice)
- Map (iterate elements and call a function on each element, replacing the element with the return value)
- Each (iterate elements and call a function on each element)
- Reduce (aka inject)
- Subtract (subtract one slice's elements from another slice and return the result)
- Add (add two slices together and return the result)
- Contains 
- Index (assumes not sorted, `O(n)` time)
- SortedIndex (index operation on a sorted slice, `O(log(n))` operating time)
- First (first element in slice)
- Last (last element in slice)
- Any (true if len > 0)
- ToStringSlice (converts anything like a []string slice to a []string slice, eg `type MyString string` and `[]MyString`)
- Filter (select elements from the slice based on a selection function)
- DeleteIf (reject elements from the slice if they match a selection function)

Functions for operations on string slices that represent sets:

- Difference
- Union
- Intersect
- Compare (allows custom set operations)

Functions that operate on maps and return string slices:

- GetKeys
- GetValues

Future support; perhaps (let me know if you'd like to see one of these – @ssoroka on twitter):

- Pop
- Push
- Shift
- Unshift
- Reverse

## usage

Functions are available as one-off functions, or as a chainable StringSlice object

```go
  s := []string{"echo", "alpha", "bravo", "delta", "charlie", "Charlie"}

  s2 := stringslice.New(s).Sort().Map(func(i int, s string) string {
    return strings.ToUpper(s)
  }).Subtract("ALPHA").Uniq().Slice()

  fmt.Println(s2)
```

prints out `[BRAVO CHARLIE DELTA ECHO]`

See tests for more usage.

## install

`go get github.com/ssoroka/stringslice`