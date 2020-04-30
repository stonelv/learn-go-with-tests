package ArraysAndSlices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 4, 4, 6}

		got := Sum(numbers)
		wanted := 17

		if got != wanted {
			t.Errorf("got %d want %d given %v", got, wanted, numbers)
		}
	})

}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2, 3}, []int{0, 8})
	wanted := []int{6, 8}
	CheckSums(t, got, wanted)
}

func TestSumAllUsingAppend(t *testing.T) {
	got := SumAll([]int{1, 2, 3}, []int{0, 8})
	wanted := []int{6, 8}
	CheckSums(t, got, wanted)
}

func TestSumAllTails(t *testing.T) {
	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3}, []int{0, 8})
		wanted := []int{5, 8}
		CheckSums(t, got, wanted)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3}, []int{})
		wanted := []int{5, 0}
		CheckSums(t, got, wanted)
	})
}

func CheckSums(t *testing.T, got, want []int) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
