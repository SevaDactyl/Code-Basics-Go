/*
Реализуйте функцию isASCII(s string) bool, которая возвращает true, если строка s состоит только из ASCII символов.
*/

package solution

import "unicode"

// BEGIN (write your solution here)

// isASCII checks whether the string s contains only ASCII chars.
func isASCII(s string) bool {
	for _, r := range s {
		if r > unicode.MaxASCII {
			return false
		}
	}

	return true
}

// END
