package leettasks

// SearchInSortedArray Suppose a sorted array is rotated at some pivot unknown to you beforehand.
// You are given a target value to search. If found in the array return its index, otherwise return -1.
// You may assume no duplicate exists in the array.
func SearchInSortedArray(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	rotationIndex := findRotationIndex(&nums)
	foundIndex := findValueIndex(&nums, target, rotationIndex)

	if foundIndex != -1 {
		foundIndex = rotateIndex(&nums, foundIndex, rotationIndex)
	}

	return foundIndex
}

func findValueIndex(nums *[]int, target int, rotationIndex int) int {
	start := 0
	end := len(*nums) - 1

	var middle int
	var startValue, endValue, middleValue int

	for {
		startValue = getValueInRotatedArray(nums, start, rotationIndex)

		if startValue == target {
			return start
		}

		if startValue > target {
			return -1
		}

		endValue = getValueInRotatedArray(nums, end, rotationIndex)

		if endValue == target {
			return end
		}

		if endValue < target {
			return -1
		}

		if (end - start) <= 1 {
			return -1
		}

		middle = (start + end) / 2
		middleValue = getValueInRotatedArray(nums, middle, rotationIndex)

		if middleValue >= target {
			end = middle
		} else {
			start = middle
		}
	}
}

func findRotationIndex(nums *[]int) int {
	start := 0
	end := len(*nums) - 1
	var middle int

	if (*nums)[start] < (*nums)[end] {
		return 0
	}

	for (end - start) > 1 {
		middle = (start + end) / 2

		if (*nums)[middle] > (*nums)[end] {
			start = middle
		} else {
			end = middle
		}
	}

	return end
}

func getValueInRotatedArray(nums *[]int, index int, rotationIndex int) int {
	index = rotateIndex(nums, index, rotationIndex)
	return (*nums)[index]
}

func rotateIndex(nums *[]int, index int, rotationIndex int) int {
	index += rotationIndex
	if index >= len(*nums) {
		index -= len(*nums)
	}
	return index
}
