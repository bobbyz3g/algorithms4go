package search

import "github.com/Kaiser925/algorithms4go/constraints"

// BinarySearch finds the index of first target value element in slice.
// If there is no target value element, return -1.
func BinarySearch[T constraints.Ordered](a []T, target T) int {
	n := len(a)
	lo := 0
	hi := n - 1

	for lo <= hi {
		mid := lo + ((hi - lo) >> 1)
		if a[mid] >= target {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}

	if lo < n && a[lo] == target {
		return lo
	}
	return -1
}

// BinarySearchLast finds the index of last target value element in slice.
// If there is no target value element, return -1.
func BinarySearchLast[T constraints.Integer | constraints.Float](a []T, target T) int {
	n := len(a)
	lo := 0
	hi := n - 1

	for lo <= hi {
		mid := lo + ((hi - lo) >> 1)
		if a[mid] > target {
			hi = mid - 1
		} else if a[mid] < target {
			lo = mid + 1
		} else {
			if mid == n-1 || a[mid]+1 != target {
				return mid
			} else {
				hi = mid - 1
			}
		}
	}
	return -1
}
