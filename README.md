# stringslice

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