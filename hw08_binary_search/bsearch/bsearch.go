package bsearch

import "errors"

func BinarySearch(m []int, target int) (int, error) {
	if m == nil {
		return -1, errors.New("slice is null")
	}
	if len(m) < 3 {
		return -1, errors.New("uncorrected size")
	}

	beg := 0
	end := len(m) - 1
	for end >= beg {
		mid := (end + beg) / 2
		if target == m[mid] {
			return mid, nil
		}
		if m[mid] > target {
			end = mid - 1
		} else {
			beg = mid + 1
		}
	}

	return -1, errors.New("target not found")
}
