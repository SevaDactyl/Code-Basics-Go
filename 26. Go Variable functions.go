/*
Реализуйте функцию MergeNumberLists(numberLists ...[]int) []int, которая принимает вариативный список слайсов чисел и объединяет их в 1, сохраняя последовательность:

MergeNumberLists([]int{1, 2}, []int{3}, []int{4}) // [1, 2, 3, 4]
*/

package solution

// BEGIN (write your solution here)

// MergeNumberLists merges all the number lists into the single one and returns it.
// It preserves the order of the number list items.
func MergeNumberLists(numberLists ...[]int) []int {
	mergedCap := 0
	for i := 0; i < len(numberLists); i++ {
		mergedCap += len(numberLists[i])
	}

	merged := make([]int, 0, mergedCap)
	for _, nl := range numberLists {
		merged = append(merged, nl...)
	}

	return merged
}

// END
