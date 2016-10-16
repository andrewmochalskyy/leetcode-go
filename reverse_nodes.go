package leettasks

// ReverseKGroup given a linked list, reverse the nodes of a linked list
// k at a time and return its modified list
func ReverseKGroup(head *ListNode, k int) *ListNode {

	futureHead := head
	currentHead := head
	groupsFound := 0
	var lastTail *ListNode

	for {
		currentTail := findGroupTail(currentHead, k)
		if currentTail == nil {
			break
		}
		if groupsFound == 0 {
			futureHead = currentTail
		}
		if lastTail != nil {
			lastTail.Next = currentTail
		}
		lastTail = currentHead
		nextHead := currentTail.Next
		for currentHead != currentTail {
			tmp := currentHead.Next
			currentHead.Next = currentTail.Next
			currentTail.Next = currentHead
			currentHead = tmp
		}
		currentHead = nextHead
		groupsFound++
	}

	return futureHead
}

func findGroupTail(head *ListNode, k int) *ListNode {
	currentNode := head
	for currentIndex := 1; currentIndex < k; currentIndex++ {
		if currentNode == nil {
			return nil
		}

		currentNode = currentNode.Next
	}
	return currentNode
}
