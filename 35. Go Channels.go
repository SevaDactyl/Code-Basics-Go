/*
Реализуйте функцию-воркера SumWorker(numsCh chan []int, sumCh chan int), которая суммирует переданные числа из канала numsCh и передает результат в канал sumCh:

numsCh := make(chan []int)
sumCh := make(chan int)

go SumWorker(numsCh, sumCh)
numsCh <- []int{10, 10, 10}

res := <- sumCh // 30
*/

package solution

// BEGIN (write your solution here)

// SumWorker runs a worker that calculates a sum of nums from the numsCh and returns a result in the sumCh.
func SumWorker(numsCh chan []int, sumCh chan int) {
	for nums := range numsCh {
		sumCh <- sum(nums)
	}
}

func sum(nums []int) int {
	s := 0
	for _, n := range nums {
		s += n
	}

	return s
}

// END
