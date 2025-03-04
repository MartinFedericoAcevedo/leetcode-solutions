package _8_merged_sorted_array

import (
	"testing"
)

func TestMerge(t *testing.T) {
	testCases := []struct {
		name     string
		nums1    []int
		m        int
		nums2    []int
		n        int
		expected []int
	}{
		{
			name:     "Caso 1: Arreglos de igual longitud con ordenación ascendente",
			nums1:    []int{1, 3, 5, 7, 9, 0, 0, 0, 0, 0},
			m:        5,
			nums2:    []int{2, 4, 6, 8, 10},
			n:        5,
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			name:     "Caso 3: Arreglo vacío `nums1`",
			nums1:    []int{0, 0, 0, 0, 0},
			m:        0,
			nums2:    []int{1, 2, 3, 4, 5},
			n:        5,
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Caso 4: Arreglo vacío `nums2`",
			nums1:    []int{1, 2, 3, 4, 5},
			m:        5,
			nums2:    []int{},
			n:        0,
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Caso 5: Elementos duplicados en ambos arreglos",
			nums1:    []int{2, 5, 5, 7, 9, 0, 0, 0, 0, 0},
			m:        5,
			nums2:    []int{3, 4, 5, 6, 8},
			n:        5,
			expected: []int{2, 3, 4, 5, 5, 5, 6, 7, 8, 9},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			merge(testCase.nums1, testCase.m, testCase.nums2, testCase.n)

			for i, v := range testCase.nums1 {
				if v != testCase.expected[i] {
					t.Errorf("Resultado incorrecto: %v, esperado: %v", testCase.nums1, testCase.expected)
					return
				}
			}
		})
	}
}
