package stringslice_test

import (
	"strconv"
	"testing"

	"github.com/ssoroka/stringslice"
	"gotest.tools/assert"
)

func TestSortedSliceIndex(t *testing.T) {
	for i := 0; i < 100; i++ {
		slice := make([]string, i)
		for j := 0; j < i; j++ {
			slice[j] = strconv.Itoa(j)
		}
		unsortedSlice := stringslice.New(slice)
		sortedSlice := unsortedSlice.Sort()
		for j := -1; j < i; j++ {
			str := strconv.Itoa(j)
			assert.Equal(t, unsortedSlice.Index(str), sortedSlice.Index(str))
		}
	}
}
