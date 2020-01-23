package base

// CompareFunc returns a number:
//    negative , if a < b
//    zero     , if a == b
//    positive , if a > b
type CompareFunc func(a, b interface{}) int

// IntCompareFunc privates a function to compare two int interface.
func IntCompareFunc(a, b interface{}) int {
	intA := a.(int)
	intB := b.(int)

	switch {
	case intA > intB:
		return 1
	case intA < intB:
		return -1
	default:
		return 0
	}
}

// Float64CompareFunc privates a function to compare two float64 interface.
func Float64CompareFunc(a, b interface{}) int {
	floatA := a.(float64)
	floatB := b.(float64)

	switch {
	case floatA > floatB:
		return 1
	case floatA < floatB:
		return -1
	default:
		return 0
	}
}

// StringCompareFunc privates a function to compare two string interface.
func StringCompareFunc(a, b interface{}) int {
	strA := a.(string)
	strB := b.(string)

	minLen := len(strA)

	if len(strA) > len(strB) {
		minLen = len(strB)
	}

	diff := 0
	for i := 0; i < minLen && diff == 0; i++ {
		diff = int(strA[i]) - int(strB[i])
	}
	if diff == 0 {
		diff = len(strA) - len(strB)
	}
	if diff < 0 {
		return -1
	}
	if diff > 0 {
		return 1
	}
	return 0
}
