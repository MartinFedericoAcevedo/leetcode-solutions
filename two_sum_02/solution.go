package main

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)
	for i, num := range nums {
		complement := target - num
		if _, ok := numMap[complement]; !ok {
			numMap[num] = i
		}
		if j, ok := numMap[complement]; ok {
			if i != j {
				return []int{j, i}
			}
		}
	}
	return []int{}
}

func twoSum2(nums []int, target int) []int {
	numMap := make(map[int]int)
	for i, num := range nums {
		numMap[num] = i
	}
	for i, num := range nums {
		complement := target - num
		if j, ok := numMap[complement]; ok && j != i {
			return []int{i, j}
		}
	}
	return []int{}
}
