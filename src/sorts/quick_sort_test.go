package sorts_test

import (
	"testing"

	"github.com/KevinBoyette/GoAlgorithms/src/sorts"
)

func TestQuickSort(t *testing.T) {
	cases := map[string]struct {
		testParam []int
		expected  []int
	}{
		"QuickSort([]int{})":          {[]int{}, []int{}},
		"QuickSort([]int{1})":         {[]int{1}, []int{1}},
		"QuickSort([]int{5,4,3,2,1})": {[]int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		"QuickSort([]int{2,1,3,5,4})": {[]int{2, 1, 3, 5, 4}, []int{1, 2, 3, 4, 5}},
		"QuickSort([]int{1,2,3,4,5})": {[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			actual := sorts.QuickSortRecur(tc.testParam)
			expected := tc.expected
			testCase := true
			for index := range actual {
				// This if can result in panic: index out of range
				if actual[index] != expected[index] {
					testCase = false
				}
			}
			if testCase != true {
				t.Errorf("During %s; expected %v and got %v",
					name,
					tc.expected,
					actual,
				)
			}
		})
	}

}
