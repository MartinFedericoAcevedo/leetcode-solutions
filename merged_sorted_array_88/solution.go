package main

import (
	"log"
)

func merge(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}

	// Caso especial: si no hay elementos en nums1, copiamos todos los elementos de nums2
	if m == 0 {
		copy(nums1, nums2)
		return
	}

	p1 := m - 1    // Puntero al último elemento de nums1
	p2 := n - 1    // Puntero al último elemento de nums2
	p := m + n - 1 // Puntero al último elemento del resultado en nums1

	// Combinación progresiva mientras ambos arreglos tengan elementos
	for p1 >= 0 && p2 >= 0 {
		if nums1[p1] > nums2[p2] {
			nums1[p] = nums1[p1]
			p1--
		} else {
			nums1[p] = nums2[p2]
			p2--
		}
		p--
		log.Printf("p1 = %d, p2 = %d, p = %d, nums1 = %v, nums2 = %v", p1, p2, p, nums1, nums2)
	}

	// Copiar los elementos restantes de nums2 a nums1
	for p2 >= 0 {
		nums1[p] = nums2[p2]
		p2--
		p--
	}
}
