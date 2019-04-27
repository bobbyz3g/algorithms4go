package strings

import "testing"

func TestBM(t *testing.T) {
	a := []string{"a", "ab", "aaa", "aaabad", "aaabaa"}
	b := []string{"b", "ab", "bbbb", "baa", "baa"}
	rs := []int{-1, 0, -1, -1, 3}

	for i := range a {
		if index := BM(a[i], b[i]); index != rs[i] {
			t.Errorf("Excepted %d got %d \n", rs[i], index)
		}
	}
}
