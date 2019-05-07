package strings

// BM returns index of the pattern string, if origin string contains pattern string,
// or -1 if there is no match.
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

// KMP returns index of the pattern string, if origin string contains pattern string,
// or -1 if there is no match.
// KMP is an implemention of the Knuth-Morris-Pratt algorithm.
func KMP(origin string, pattern string) int {
	m := len(pattern)
	R := 256

	dfa := make([][]int, R)
	for i := range dfa {
		dfa[i] = make([]int, m)
	}

	dfa[int(pattern[0])][0] = 1

	for x, j := 0, 1; j < m; j++ {
		for c := 0; c < R; c++ {
			dfa[c][j] = dfa[c][x]
		}

		dfa[int(pattern[j])][j] = j + 1
		x = dfa[int(pattern[j])][x]

	}

	n := len(origin)
	var i, j int
	for i, j = 0, 0; i < n && j < m; i++ {
		j = dfa[int(origin[i])][j]
	}
	if j == m {
		return i - m
	}
	return -1
}
