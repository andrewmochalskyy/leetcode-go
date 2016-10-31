package leettasks

// IsInterleave Given s1, s2, s3, find whether s3 is formed by the interleaving of s1 and s2.
func IsInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}

	return isInterleave(&s1, &s2, &s3, 0, 0, 0)
}

func isInterleave(s1 *string, s2 *string, s3 *string, offset1 int, offset2 int, offset3 int) bool {
	if len(*s1) == offset1 && len(*s2) == offset2 && len(*s3) == offset3 {
		return true
	}

	if offset1 < len(*s1) && (*s1)[offset1] == (*s3)[offset3] && isInterleave(s1, s2, s3, offset1+1, offset2, offset3+1) {
		return true
	}

	if offset2 < len(*s2) && (*s2)[offset2] == (*s3)[offset3] && isInterleave(s1, s2, s3, offset1, offset2+1, offset3+1) {
		return true
	}

	return false
}
