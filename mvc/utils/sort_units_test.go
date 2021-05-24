package utils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubleSortWorthCase(t *testing.T) {
	els := []int{9, 8, 7, 6, 5}
	Sort(els)
	fmt.Println(els)
	assert.NotNil(t, els)
	assert.EqualValues(t, 5, len(els))
	assert.EqualValues(t, []int{5, 6, 7, 8, 9}, els)
}

func getElements(n int) []int {
	result := make([]int, n)
	for i := 0; i < n; i++ {
		result[i] = i
	}
	return result
}

func TestBubleSortBestCase(t *testing.T) {
	els := getElements(5)
	Sort(els)
	assert.NotNil(t, els)
	assert.EqualValues(t, 5, len(els))
	assert.EqualValues(t, getElements(5), els)
}

func TestBubleSortNilSlice(t *testing.T) {
	var els []int
	Sort(els)
	assert.Nil(t, els, "hello worls")
}

func BenchmarkBableSort10(b *testing.B) {
	els := getElements(10)
	for i := 0; i < b.N; i++ {
		Sort(els)
	}
}
