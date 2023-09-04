/*
В пакете unicode есть функция unicode.Is(unicode.Latin, rune), которая проверяет, что руна является латинским символом или нет.

Реализуйте функцию latinLetters(s string) string, которая возвращает только латинские символы из строки s. Например:

latinLetters("привет world!") // "world"
*/

package solution

import (
	"strings"
	"unicode"
)

// BEGIN (write your solution here) (write your solution here)
func latinLetters(s string) string {
	sb := &strings.Builder{}

	for _, r := range s {
		if unicode.Is(unicode.Latin, r)  {
			sb.WriteRune(r)
		}
    }

    return sb.String()
}
// END
