/*
Реализуйте функции nextASCII(b byte) byte и prevASCII(b byte) byte, которые возвращают следующий или предыдущий символ ASCII таблицы соответственно. Например:

nextASCII(byte('a')) // 'b'
prevASCII(byte('b')) // 'a'
Допускаем, что в функцию prevASCII не может прийти нулевой символ, а в функцию nextASCII — последний символ ASCII таблицы.
*/

package solution

// BEGIN (write your solution here)

// nextASCII returns the next byte after b according to the ASCII code table.
func nextASCII(b byte) byte {
	return b + 1
}

// prevASCII returns the previous byte after b according to the ASCII code table.
func prevASCII(b byte) byte {
	return b - 1
}

// END
