/*
В уроке упоминалась функция math.Max, которая сравнивает два числа и возвращает наибольшее. В этом задании следует использовать противоположную функцию math.Min, определяющую наименьшее число.
Напишите функцию MinInt(x, y int) int, которая возвращает наименьшее целое число
*/

package solution

import "math"

// BEGIN (write your solution here)
func MinInt(x, y int) int {
	return int(math.Min(float64(x), float64(y)))
}

// END
