/*
Реализуйте функцию IntsCopy(src []int, maxLen int) []int, которая создает копию слайса src с длиной maxLen. Если maxLen равен нулю или отрицательный, то функция возвращает пустой слайс []int{}. Если maxLen больше длины src, то возвращается полная копия src.
*/

package solution

// BEGIN (write your solution here)

// IntsCopy copies a slice of ints src with max length len.
// If maxLen is greater than src it returns a full copy of src.
// If maxLen is zero or negative it returns empty slice.
func IntsCopy(src []int, maxLen int) []int {
	if maxLen <= 0 {
		return []int{}
	}

	if maxLen > len(src) {
		maxLen = len(src)
	}

	cp := make([]int, maxLen)
	copy(cp, src)

	return cp
}

// END
