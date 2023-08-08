/*
На сервер приходит HTTP-запрос. Тело запроса парсится и мапится в модель. Сразу работать с моделью нельзя, потому что данные могут быть неверными.
Реализуйте функцию Validate(req UserCreateRequest) string, которая валидирует запрос и возвращает строку с ошибкой "invalid request", если модель невалидна. Если запрос корректный, то функция возвращает пустую строку. Правила валидации и структура модели описаны ниже. Не используйте готовые библиотеки и опишите правила самостоятельно.

type UserCreateRequest struct {
  FirstName string // не может быть пустым; не может содержать пробелы
  Age int // не может быть 0 или отрицательным; не может быть больше 150
}
Наличие пробелов можно проверить с помощью функции strings.Contains(firstName, " ").
*/

package solution

import (
	"strings"
)

// UserCreateRequest is a request to create a new user.
type UserCreateRequest struct {
	FirstName string
	Age       int
}

// BEGIN (write your solution here)

var (
	invalidRequest = "invalid request"
)

// Validate validates the UserCreateRequest and returns an error if the model is invalid.
func Validate(req UserCreateRequest) string {
	if req.FirstName == "" || strings.Contains(req.FirstName, " ") {
		return invalidRequest
	}

	if req.Age <= 0 || req.Age > 150 {
		return invalidRequest
	}

	return ""
}

// END
