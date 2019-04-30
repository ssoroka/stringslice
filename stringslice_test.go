package stringslice_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/ssoroka/stringslice"

	"gotest.tools/assert"
)

func TestSubtraction(t *testing.T) {
	type testCase struct {
		input    []string
		args     []string
		expected []string
	}
	tests := []testCase{
		{input: []string{}, args: []string{}, expected: []string{}},
		{input: []string{"a", "b", "c", "d", "e"}, args: []string{"b", "d"}, expected: []string{"a", "c", "e"}},
		{input: []string{"a", "a"}, args: []string{"a", "l"}, expected: []string{}},
		{input: []string{}, args: []string{"a", "l"}, expected: []string{}},
	}
	for _, test := range tests {
		ss := stringslice.New(test.input)
		result := ss.Subtract(test.args).Slice()

		assert.DeepEqual(t, test.expected, result)
	}
}

func TestAdd(t *testing.T) {
	assert.DeepEqual(t, []string{"a", "b"}, stringslice.New([]string{"a"}).Add([]string{"b"}).Slice())
	assert.DeepEqual(t, []string{"a"}, stringslice.New([]string{"a"}).Add([]string{}).Slice())
	assert.DeepEqual(t, []string{}, stringslice.New([]string{}).Add([]string{}).Slice())
}

func TestUniq(t *testing.T) {
	type testCase struct {
		input    []string
		expected []string
	}
	tests := []testCase{
		{input: []string{}, expected: []string{}},
		{input: []string{"a", "b", "a", "c", "c"}, expected: []string{"a", "b", "c"}},
		{input: []string{"b", "a", "a", "b", "b"}, expected: []string{"a", "b"}},
		{input: []string{""}, expected: []string{""}},
	}
	for _, test := range tests {
		ss := stringslice.New(test.input)
		result := ss.Uniq().Slice()

		assert.DeepEqual(t, test.expected, result)
	}
}

func TestMap(t *testing.T) {
	s := []string{"a", "b", "c"}
	ss := stringslice.New(s)
	result := ss.Map(func(i int, s string) string {
		return strings.ToUpper(s)
	}).Slice()
	expected := []string{"A", "B", "C"}

	assert.DeepEqual(t, expected, result)
}

func TestReadme(t *testing.T) {
	s := []string{"echo", "alpha", "bravo", "delta", "charlie", "Charlie"}

	s2 := stringslice.New(s).Sort().Map(func(i int, s string) string {
		return strings.ToUpper(s)
	}).Subtract([]string{"ALPHA"}).Uniq().Slice()

	fmt.Println(s2)
}
