/*
Объявите две переменные firstName и lastName. Переменная firstName должна содержать строку «John», переменная lastName — «Smith».
Выведите значения переменных firstName и lastName через пробел, чтобы получилась строка «John Smith». Используйте функцию Println из пакета fmt
*/

package main

import "fmt"

func main() {
  // BEGIN (write your solution here)
  var firstName string = "John"
  var lastName  string = "Smith"
  fmt.Println(firstName, lastName)
  // END
}
