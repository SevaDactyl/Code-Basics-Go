/*
Для работы со строками часто используется стандартная библиотека strings. В данном задании вам понадобятся следующие функции:

// обрезает символы, переданные вторым аргументом, с обеих сторон строки
Trim(s, cutset string) string
// пример
strings.Trim(" hello ", " ") // "hello"

// преобразует все буквы в строке в нижний регистр
strings.ToLower(s string) string
// пример
strings.ToLower("пРиВеТ") // "привет"

// озаглавливает первую букву в каждом слове в строке
strings.Title(s string) string
// пример
strings.Title("привет, джон") // "Привет, Джон"
Реализуйте функцию Greetings(name string) string, которая вернет строку-приветствие. Например, при передаче имени Иван, возвращается Привет, Иван!. Учтите, что имя может быть написано в разном регистре и содержать пробелы.
*/

package solution

import (
	"fmt"
	"strings"
)

// BEGIN (write your solution here)
func Greetings(name string) string {
	name = strings.Trim(name, " ")
	name = strings.ToLower(name)
	name = strings.Title(name) //nolint

	return fmt.Sprintf("Привет, %s!", name)
} 

// END
