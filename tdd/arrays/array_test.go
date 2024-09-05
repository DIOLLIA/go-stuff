package arrays

import (
	"reflect"
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {

		numbers := []int{1, 2, 3, 4, 5}

		result := Sum(numbers)
		expected := 15

		if result != expected {
			t.Errorf("Expected %d but got %d, given %v", expected, result, numbers)
		}
	})
}
func TestSumAll(t *testing.T) {
	t.Run("make the sums of some slices", func(t *testing.T) {
		expected := []int{3, 9}
		result := SumAll([]int{1, 2}, []int{0, 9})

		if !slices.Equal(expected, result) {
			t.Errorf("expected %v result %v", expected, result)
		}
	})
}

func TestAppendAllTails(t *testing.T) {
	checkSums := func(t testing.TB, expected, result []int) {
		t.Helper()
		if !reflect.DeepEqual(expected, result) {
			t.Errorf("expected %v result %v", expected, result)
		}
	}
	t.Run("Sum all elements from many to one slice", func(t *testing.T) {
		result := AppendAllTails([]int{1, 2}, []int{0, 9})
		expected := []int{2, 9}

		checkSums(t, expected, result)
	})
	t.Run("Sum elements with zero length slice", func(t *testing.T) {
		result := AppendAllTails([]int{}, []int{1, 4, 5})
		expected := []int{0, 9}

		checkSums(t, expected, result)
	})
}
