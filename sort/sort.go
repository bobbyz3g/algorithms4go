package sort

import "github.com/Kaiser925/algorithms4go/constraints"

// BubbleSort sorts slice by using bubble sort.
func BubbleSort[T constraints.Ordered](a []T) {
	n := len(a)
	if n <= 1 {
		return
	}

	var flag bool

	for i := 0; i < n; i++ {
		flag = false

		for j := 0; j < n-i-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
				flag = true
			}
		}

		if !flag {
			break
		}
	}
}

// InsertionSort sorts slice by using insertion sort.
func InsertionSort[T constraints.Ordered](a []T) {
	n := len(a)

	if n <= 1 {
		return
	}

	for i := 1; i < n; i++ {
		elem := a[i]
		j := i - 1
		for ; j >= 0; j-- {
			if a[j] > elem {
				a[j+1] = a[j]
			} else {
				break
			}
		}
		a[j+1] = elem
	}
}

// SelectionSort sorts slice by using selection sort.
func SelectionSort[T constraints.Ordered](a []T) {
	n := len(a)
	if n <= 1 {
		return
	}

	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if a[j] < a[minIndex] {
				minIndex = j
			}
		}

		a[i], a[minIndex] = a[minIndex], a[i]
	}
}

// merge array
func merge[T constraints.Ordered](a []T, start, mid, end int) {
	temp := make([]T, len(a))

	i := start
	j := mid + 1

	copy(temp, a)
	for k := start; k <= end; k++ {
		if i > mid {
			a[k] = temp[j]
			j++
		} else if j > end {
			a[k] = temp[i]
			i++
		} else if temp[i] < temp[j] {
			a[k] = temp[i]
			i++
		} else {
			a[k] = temp[j]
			j++
		}
	}
}

func mergeSort[T constraints.Ordered](a []T, start int, end int) {
	if start >= end {
		return
	}

	mid := (start + end) / 2

	mergeSort(a, start, mid)
	mergeSort(a, mid+1, end)

	merge(a, start, mid, end)
}

func MergeSort[T constraints.Ordered](a []T) {
	mergeSort(a, 0, len(a)-1)
}

func partition[T constraints.Ordered](a []T, lo, hi int) int {
	pivot := a[hi]

	i := lo
	for j := lo; j < hi; j++ {
		if a[j] < pivot {
			a[i], a[j] = a[j], a[i]
			i++
		}
	}
	a[i], a[hi] = a[hi], a[i]
	return i
}

func quickSort[T constraints.Ordered](a []T, start, end int) {
	if start >= end {
		return
	}
	pivot := partition(a, start, end)
	quickSort(a, start, pivot-1)
	quickSort(a, pivot+1, end)
}

// QuickSort sorts the slice by using quick sort.
func QuickSort[T constraints.Ordered](a []T) {
	quickSort(a, 0, len(a)-1)
}

// CountingSort sorts the slice by using counting sort.
// Element in slice should be nonnegative integer.
func CountingSort[T constraints.Unsigned](a []T) {
	n := len(a)
	if n <= 1 {
		return
	}

	var max T = a[0]

	for _, i := range a {
		if i > max {
			max = i
		}
	}

	c := make([]T, max+1)

	for _, i := range a {
		c[i]++
	}

	var i T = 1
	for i = 1; i <= max; i++ {
		c[i] = c[i-1] + c[i]
	}

	r := make([]T, n)

	for i := n - 1; i >= 0; i-- {
		index := c[a[i]] - 1
		r[index] = a[i]
		c[a[i]]--
	}

	copy(a, r)
}
