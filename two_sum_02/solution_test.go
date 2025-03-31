package main

import (
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		nombre   string
		nums     []int
		target   int
		expected []int
	}{
		// Casos básicos (ejemplos de LeetCode)
		{
			nombre:   "Ejemplo 1: números positivos",
			nums:     []int{2, 7, 11, 15},
			target:   9,
			expected: []int{0, 1},
		},
		{
			nombre:   "Ejemplo 2: target en medio",
			nums:     []int{3, 2, 4},
			target:   6,
			expected: []int{1, 2},
		},
		{
			nombre:   "Ejemplo 3: números repetidos",
			nums:     []int{3, 3},
			target:   6,
			expected: []int{0, 1},
		},

		// Casos adicionales
		{
			nombre:   "Array grande",
			nums:     []int{1, 5, 8, 10, 13, 15, 20, 25, 30},
			target:   35,
			expected: []int{5, 6},
		},
		{
			nombre:   "Números negativos",
			nums:     []int{-3, 4, 3, 90},
			target:   0,
			expected: []int{0, 2},
		},
		{
			nombre:   "Target negativo",
			nums:     []int{-1, -2, -3, -4, -5},
			target:   -8,
			expected: []int{2, 4},
		},
		{
			nombre:   "Números extremos",
			nums:     []int{1000000000, -1000000000},
			target:   0,
			expected: []int{0, 1},
		},
		{
			nombre:   "Elementos no adyacentes",
			nums:     []int{2, 5, 8, 11, 15},
			target:   17,
			expected: []int{0, 4},
		},
		{
			nombre:   "Elementos adyacentes",
			nums:     []int{4, 6, 8, 9, 10},
			target:   14,
			expected: []int{1, 2},
		},
		{
			nombre:   "Elementos desordenados",
			nums:     []int{10, 2, 8, 6, 4},
			target:   14,
			expected: []int{2, 3},
		},
	}

	for _, test := range tests {
		t.Run(test.nombre, func(t *testing.T) {
			result := twoSum(test.nums, test.target)
			if !((result[0] == test.expected[0] && result[1] == test.expected[1]) ||
				(result[0] == test.expected[1] && result[1] == test.expected[0])) {
				t.Errorf("Para nums=%v y target=%d, se esperaba %v pero se obtuvo %v",
					test.nums, test.target, test.expected, result)
			}
		})
	}
}
