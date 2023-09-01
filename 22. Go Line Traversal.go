/*
Реализуйте функцию shiftASCII(s string, step int) string, которая принимает на вход состоящую из ASCII символов строку s и возвращает новую строку, где каждый символ из входящей строки сдвинут вперед на число step. Например:

shiftASCII("abc", 0) // "abc"
shiftASCII("abc1", 1) // "bcd2"
shiftASCII("bcd2", -1) // "abc1"
shiftASCII("hi", 10) // "rs"
Если после сдвига код символа выходит за рамки ASCII, то число должно быть взято по модулю 256:

shiftASCII("abc", 256) // "abc"
shiftASCII("abc", -511) // "bcd"
*/

package solution

// BEGIN (write your solution here)

// shiftASCII shifts each byte in the string s and if there is an overflow of ASCII it takes a new code by module of 256.
func shiftASCII(s string, step int) string {
	if step == 0 {
		return s
	}

	shifted := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		shifted[i] = byte(int(s[i]) + step)
	}

	return string(shifted)
}

// END
