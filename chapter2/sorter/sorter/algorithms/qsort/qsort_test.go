package qsort

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	values := []int{5, 4, 9, 2, 1}
	QuickSort(values)
	for _, value := range values {
		fmt.Print(value, " ")
	}
}
