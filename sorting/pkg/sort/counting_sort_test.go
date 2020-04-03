package sort_test

import (
	"testing"

	sort "github.com/go-ds-algo/sorting/pkg/sort"
	"github.com/stretchr/testify/assert"
)

func TestCountingSort(t *testing.T) {
	array := []int{2, 4, 1, 9, 6}
	expected := []int{1, 2, 4, 6, 9}

	sort.CountingSort(array)

	assert.Equal(t, expected, array, "Arrays do not match")
}
