/*
В больших проектах часто используется gRPC — высокопроизводительный RPC (Remote Procedure Call)-фреймворк для коммуникации между микросервисами. Ошибки в gRPC стандартизированы и представлены в виде строк для пользователя, либо в виде чисел для компьютера.

Представим, что нам нужно написать конвертер ошибок в числовой формат для gRPC. Реализуйте функцию ErrorMessageToCode(msg string) int, которая возвращает числовой код для заданного значения. Список сообщений и соответствующих кодов:

OK = 0
CANCELLED = 1
UNKNOWN = 2
В реальности список намного шире. Мы для простоты ограничимся тремя ошибками. Учтите, что если в функцию передать неизвестную ошибку, она должна вернуть код ошибки для сообщения UNKNOWN.
*/

package solution

// BEGIN (write your solution here)
// error messages
const (
  OkMsg        = "OK"
  CancelledMsg = "CANCELLED"
)

// error codes
const (
  OkCode = iota
  CancelledCode
  UnknownCode
)

// ErrorMessageToCode returns a gRPC error code depending on error message.
func ErrorMessageToCode(msg string) int {
  switch msg {
  case OkMsg:
    return OkCode
  case CancelledMsg:
    return CancelledCode
  }

  return UnknownCode
}


// END
