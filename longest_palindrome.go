package leettasks

// LongestPalindrome Finds longest palindrome in a given string
func LongestPalindrome(s string) string {
	maxLength := 0
	left := 0
	right := 0

	firstCursor := 0
	secondCursor := 0
	stringSize := len(s)

	for firstCursor < stringSize && secondCursor < stringSize {
		found, foundStart, foundEnd := probePalindrome(&s, firstCursor, secondCursor)
		if found == true {
			currentLength := 1 + (foundEnd - foundStart)
			if currentLength > maxLength {
				left = foundStart
				right = foundEnd
				maxLength = currentLength
			}
		}
		if firstCursor == secondCursor {
			secondCursor++
		} else {
			firstCursor++
		}
	}

	if maxLength > 0 {
		return s[left : right+1]
	}

	return ""
}

func probePalindrome(s *string, left int, right int) (found bool, foundStart int, foundEnd int) {
	foundStart = -1
	foundEnd = -1
	testString := *s

	for left >= 0 && right < len(testString) && testString[left] == testString[right] {
		foundStart = left
		foundEnd = right

		left--
		right++
	}

	found = foundStart >= 0 && foundEnd < len(testString)
	return
}
