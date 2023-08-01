package sort

import (
	sorting "sort"
	"testing"
)

func testIsSortedInt(t *testing.T, slice []int) {
	t.Helper()

	expected := sorting.IsSorted(sorting.IntSlice(slice))
	received := IsSorted(slice, func(a, b int) bool { return a < b })

	if expected != received {
		t.Errorf("expected %v", expected)
		t.Errorf("received %v", received)
	}
}

func FuzzIsSortedInt4(f *testing.F) {
	f.Add(0, 0, 0, 0)
	f.Add(0, 1, 2, 3)
	f.Add(3, 2, 1, 0)
	f.Fuzz(func(t *testing.T, a, b, c, d int) {
		slice := []int{a, b, c, d}
		testIsSortedInt(t, slice)
	})
}

func TestIsSortedInt(t *testing.T) {
	tests := []struct {
		name  string
		slice []int
	}{
		{
			name:  "[]",
			slice: []int{},
		},
		{
			name:  "[0]",
			slice: []int{0},
		},
		{
			name:  "[0,0]",
			slice: []int{0, 0},
		},
		{
			name:  "[0,1]",
			slice: []int{0, 1},
		},
		{
			name:  "[1,0]",
			slice: []int{1, 0},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			testIsSortedInt(t, test.slice)
		})
	}
}

func testSortInt(t *testing.T, sort func(slice []int, less func(a, b int) bool), slice []int) {
	t.Helper()

	less := func(a, b int) bool { return a < b }

	received := make([]int, len(slice))
	copy(received, slice)
	sort(received, less)

	if !IsSorted(received, less) {
		expected := make([]int, len(slice))
		copy(expected, slice)
		sorting.Sort(sorting.IntSlice(expected))

		t.Errorf("expected %v", expected)
		t.Errorf("received %v", received)
	}

}

func FuzzInsertionSortInt(f *testing.F) {
	f.Add(0, 0, 0, 0)
	f.Add(0, 1, 2, 3)
	f.Add(3, 2, 1, 0)
	f.Fuzz(func(t *testing.T, a, b, c, d int) {
		slice := []int{a, b, c, d}
		testSortInt(t, InsertionSort[int], slice)
	})
}

func TestSortInt(t *testing.T) {
	implementations := []struct {
		name string
		sort func(slice []int, less func(a, b int) bool)
	}{
		{
			name: "InsertionSort",
			sort: InsertionSort[int],
		},
		{
			name: "SelectionSort",
			sort: SelectionSort[int],
		},
	}

	tests := []struct {
		name  string
		slice []int
	}{
		{
			name:  "[]",
			slice: []int{},
		},
		{
			name:  "[0]",
			slice: []int{0},
		},
		{
			name:  "[0,0]",
			slice: []int{0, 0},
		},
		{
			name:  "[0,1]",
			slice: []int{0, 1},
		},
		{
			name:  "[1,0]",
			slice: []int{1, 0},
		},
		{
			name:  "[0,0,0]",
			slice: []int{0, 0, 0},
		},
		{
			name:  "[0,1,2]",
			slice: []int{0, 1, 2},
		},
		{
			name:  "[0,2,1]",
			slice: []int{0, 2, 1},
		},
		{
			name:  "[1,0,2]",
			slice: []int{1, 0, 2},
		},
		{
			name:  "[1,2,0]",
			slice: []int{1, 2, 0},
		},
		{
			name:  "[2,0,1]",
			slice: []int{2, 0, 1},
		},
		{
			name:  "[2,1,0]",
			slice: []int{2, 1, 0},
		},
		{
			name:  "[74,59,238,-784,9845,959,905,0,0,42,7586,-5467984,7586]",
			slice: []int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586},
		},
	}

	for _, implementation := range implementations {
		t.Run(implementation.name, func(t *testing.T) {
			for _, test := range tests {
				t.Run(test.name, func(t *testing.T) {
					testSortInt(t, implementation.sort, test.slice)
				})
			}
		})
	}
}
