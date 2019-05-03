# stringslice

Adds functions to string slices: 

- Sort (does not mutate)
- Uniq (currently implies also sorted)
- Map
- Reduce (aka inject)
- Subtract (subtract one slice's elements from another slice and return the result)
- Add (add two slices together and return the result)
- Contains 
- Index

Future support; perhaps (let me know if you'd like to see one of these – @ssoroka on twitter):

- compact (remove nils)
- reject (aka delete_if)
- select (aka filter)
- each
- join (alias from strings.Join)
- pop
- push
- shift
- unshift
- reverse
- sort_by (like sort, but you can give the sorting function)

## usage

```go
  s := []string{"echo", "alpha", "bravo", "delta", "charlie", "Charlie"}

  s2 := stringslice.New(s).Sort().Map(func(i int, s string) string {
    return strings.ToUpper(s)
  }).Subtract([]string{"ALPHA"}).Uniq().Slice()

  fmt.Println(s2)
```

prints out `[BRAVO CHARLIE DELTA ECHO]`

## install

`go get github.com/ssoroka/stringslice`