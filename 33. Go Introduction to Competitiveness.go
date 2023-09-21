/*
Написать конкурентный код самостоятельно в данном уроке не получится, поэтому давайте подготовимся к следующему уроку реализацией синхронного кода.

Реализуйте функцию MaxSum(nums1, nums2 []int) []int, которая суммирует числа в каждом слайсе и возвращает слайс с наибольшей суммой. Если сумма одинаковая, то возвращается первый слайс.

Пример работы:

MaxSum([]int{1, 2, 3}, []int{10, 20, 50}) // [10 20 50]

MaxSum([]int{3, 2, 1}, []int{1, 2, 3}) // [3 2 1]
*/

package solution

// BEGIN (write your solution here)

// MaxSum gets sum of each nums slice and returns a slice with the max sum.
// If the slices are the same then the first one will be returned.
func MaxSum(nums1, nums2 []int) []int {
    if sum(nums1) >= sum(nums2) {
        return nums1
    }

    return nums2
}

func sum(nums []int) int {
    s := 0
    for _, n := range nums {
        s += n
    }

    return s
}

// END
