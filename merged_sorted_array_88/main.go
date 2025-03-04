package main

import "fmt"

func main() {
	nums1 := []int{1, 3, 5, 0, 0, 0}
	nums2 := []int{2, 4, 6}
	merge(nums1, 3, nums2, len(nums2))
	fmt.Println("Merged Array:", nums1)
}
