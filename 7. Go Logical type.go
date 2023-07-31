/*
Реализуйте функцию IsValid(id int, text string) bool, которая проверяет, что переданный идентификатор id является положительным числом и текст text не пустой.
Например:

IsValid(0, "hello world") // false
IsValid(-22, "hello world") // false
IsValid(22, "") // false
IsValid(225, "hello world") // true
*/

package solution

// BEGIN (write your solution here)
func IsValid(id int, text string) bool {
	return id > 0 && text != ""
} 
// END
