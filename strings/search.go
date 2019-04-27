package strings

// BM returns index of the pattern string in the origin string, or -1 if no match.
// BM is an implemention of the Boyer-Moore algorithm.
func BM(origin string, pattern string) (index int) {
	index = -1

	lenP := len(pattern)
	lenO := len(origin)

	R := 256
	right := make([]int, R)

	for c := 0; c < R; c++ {
		right[c] = -1
	}

	// Find the most right index of each character.
	for index, c := range pattern[:lenP] {
		right[int(c)] = index
	}

	skip := 0

	for i := 0; i <= (lenO - lenP); i++ {
		skip = 0
		for j := lenP - 1; j >= 0; j-- {
			if pattern[j] != origin[i+j] {
				skip = j - right[origin[i+j]]
				if skip < 1 {
					skip = 1
				}
				break
			}
		}
		if skip == 0 {
			index = i
			return
		}
	}

	return
}
