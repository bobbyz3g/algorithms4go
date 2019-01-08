package base

import "testing"

func TestIntCompareFunc(t *testing.T) {
	testdata := [][]interface{}{
		{1, 1, 0},
		{1, 2, -1},
		{2, 1, 1},
		{19, 22, -1},
		{0, 0, 0},
		{1, 0, 1},
		{0, 1, -1},
		{1, 1, 0},
	}

	for _, data := range testdata {
		actual := IntCompareFunc(data[0], data[1])
		expected := data[2]
		if actual != expected {
			t.Errorf("Got %v expected %v", actual, expected)
		}
	}
}

func TestFloat64CompareFunc(t *testing.T) {
	testdata := [][]interface{}{
		{1.5, 1.5, 0},
		{1.9, 2.1, -1},
		{2.0, 1.2, 1},
		{19.11, 22.43, -1},
		{0.00, 0.0, 0},
		{1.1, 0.2, 1},
		{-0.23, 1.1, -1},
		{1.4, 1.4, 0},
	}
	for _, data := range testdata {
		actual := Float64CompareFunc(data[0], data[1])
		expected := data[2]
		if actual != expected {
			t.Errorf("Got %v expected %v", actual, expected)
		}
	}
}

func TestStringCompareFunc(t *testing.T) {
	testdata := [][]interface{}{
		{"aaa", "aaa", 0},
		{"a", "b", -1},
		{"b", "a", 1},
		{"aa", "aab", -1},
		{"", "", 0},
		{"a", "", 1},
		{"", "a", -1},
		{"", "aaaaaaa", -1},
		{"aaa", "ccc", -1},
	}
	for _, data := range testdata {
		actual := StringCompareFunc(data[0], data[1])
		expected := data[2]
		if actual != expected {
			t.Errorf("Got %v expected %v", actual, expected)
		}
	}
}
