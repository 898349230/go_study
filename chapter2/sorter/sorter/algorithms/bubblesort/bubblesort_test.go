package bubblesort

import (
	"fmt"
	"testing"
)

func TestBubblesort1(t *testing.T) {
	values := []int{5, 4, 3, 2, 1}
	BubbleSort(values)
	for _, value := range values {
		fmt.Print(value, " ")
	}
}
func TestBubblesort2(t *testing.T) {
	values := []int{5, 8, 3, 7, 1}
	BubbleSort(values)
	for _, value := range values {
		fmt.Print(value, " ")
	}
}
