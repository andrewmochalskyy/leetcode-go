package leettasks

import (
	"bytes"
)

// FullJustify Given an array of words and a length L, format the text
// such that each line has exactly L characters and is fully (left and right) justified.
// You should pack your words in a greedy approach; that is, pack as many words as you can in each line.
// Pad extra spaces ' ' when necessary so that each line has exactly L characters.
// Extra spaces between words should be distributed as evenly as possible.
// If the number of spaces on a line do not divide evenly between words,
// the empty slots on the left will be assigned more spaces than the slots on the right.
// For the last line of text, it should be left justified and no extra space is inserted between words.
func FullJustify(words []string, maxWidth int) []string {
	var results []string

	currentWidth := 0
	firstWordIndex := 0
	currentWordIndex := 0

	for currentWordIndex < len(words) {
		currentWordWidth := len(words[currentWordIndex])

		if currentWidth+1+currentWordWidth > maxWidth && currentWordIndex > firstWordIndex {
			nextLine := createJustifiedString(
				&words, firstWordIndex, currentWordIndex-1, maxWidth, currentWidth, false)

			results = append(results, nextLine)
			firstWordIndex = currentWordIndex
			currentWidth = 0
		}

		if currentWordIndex > firstWordIndex {
			currentWidth++
		}

		currentWidth += currentWordWidth
		currentWordIndex++
	}

	nextLine := createJustifiedString(
		&words, firstWordIndex, currentWordIndex-1, maxWidth, currentWidth, true)

	results = append(results, nextLine)

	return results
}

func createJustifiedString(words *[]string, firstWordIndex int, lastWordIndex int, maxWidth int, currentWidth int, isLastLine bool) string {

	extraSpaces := maxWidth - currentWidth
	extraSpacePlaceholders := lastWordIndex - firstWordIndex

	extraSpaceCount := 0
	overflowSpacePlaceholderCount := 0
	trailingSpaceCount := 0

	if !isLastLine && extraSpacePlaceholders > 0 {
		extraSpaceCount = extraSpaces / extraSpacePlaceholders
		overflowSpacePlaceholderCount = extraSpaces % extraSpacePlaceholders
	}

	if isLastLine || extraSpacePlaceholders == 0 {
		trailingSpaceCount = extraSpaces
	}

	var buffer bytes.Buffer

	for i := firstWordIndex; i <= lastWordIndex; i++ {
		if i > firstWordIndex {
			buffer.WriteString(" ")
			if !isLastLine {
				addSpaces(&buffer, extraSpaceCount)
				if i-firstWordIndex <= overflowSpacePlaceholderCount {
					buffer.WriteString(" ")
				}
			}
		}
		buffer.WriteString((*words)[i])
	}

	if trailingSpaceCount > 0 {
		addSpaces(&buffer, trailingSpaceCount)
	}

	return buffer.String()
}

func addSpaces(buffer *bytes.Buffer, spaceCount int) {
	for i := 0; i < spaceCount; i++ {
		(*buffer).WriteString(" ")
	}
}
