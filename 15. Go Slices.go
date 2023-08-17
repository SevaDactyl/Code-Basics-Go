/*
В Go нет встроенной функции удаления элемента из слайса. Реализуйте функцию Remove(nums []int, i int) []int, которая удаляет элемент по индексу i из слайса nums. Если приходит несуществующий индекс, то из функции возвращается исходный слайс. Порядок элементов может быть нарушен после удаления элемента.
*/

package solution

// BEGIN (write your solution here)
func Remove(nums []int, i int) []int {
	if i < 0 || i > len(nums)-1 {
		return nums
	}

	nums[i] = nums[len(nums)-1]

	return nums[:len(nums)-1]
}

// END
