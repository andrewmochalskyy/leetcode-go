package leettasks

// JumpGame given an array of non-negative integers, you are initially positioned at the first index of the array.
// Each element in the array represents your maximum jump length at that position.
// Your goal is to reach the last index in the minimum number of jumps.
func JumpGame(nums []int) int {

	minDepth := len(nums)
	depths := make([]int, len(nums))

	if success, depth := tryJump(&nums, &depths, 0, 0, minDepth); success == true {
		return depth
	}

	return 0
}

func tryJump(nums *[]int, depths *[]int, currentIndex int, currentDepth int, minDepth int) (bool, int) {

	if currentIndex+1 == len(*nums) {
		return true, currentDepth
	}

	if currentIndex >= len(*nums) || currentDepth+1 >= minDepth || ((*depths)[currentIndex] != 0 && currentDepth >= (*depths)[currentIndex]) {
		return false, minDepth
	}

	(*depths)[currentIndex] = currentDepth

	maxJumpLength := (*nums)[currentIndex]
	anySuccess := false

	for jumpLength := maxJumpLength; jumpLength >= 1; jumpLength-- {
		if currentIndex+jumpLength < len(*nums) {
			success, depth := tryJump(nums, depths, currentIndex+jumpLength, currentDepth+1, minDepth)
			if success {
				anySuccess = true
				minDepth = depth
			}
		}
	}

	return anySuccess, minDepth
}
