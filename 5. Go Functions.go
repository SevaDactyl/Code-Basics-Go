/*
В Go есть стандартная библиотека strconv для конвертации чисел в строки и наоборот. Пример использования:

s := strconv.Itoa(-42) // "-42"
Напишите функцию IntToString, которая преобразует и возвращает входящее число в строку
*/

// Первое решение

package solution

import "strconv"
import "fmt"

// BEGIN (write your solution here)
func IntToString(x int) string {
    var y = strconv.Itoa(x)
    fmt.Println(y)
	return y
}
// END

// Второе решение

package solution

import "strconv"

// BEGIN (write your solution here)

// IntToString converts i number to a string representation.
func IntToString(i int) string {
	return strconv.Itoa(i)
}

// END
