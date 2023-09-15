/*
Для выполнения этого задания потребуется функция json.Unmarshal, которая декодирует JSON байты в структуру:

package main

import (
    "encoding/json"
    "fmt"
)

type HelloWorld struct {
    Hello string
}

func main() {
    hw := HelloWorld{}

  // первым аргументом передается JSON-строка в виде слайса байт. Вторым аргументом указатель на структуру, в которую нужно декодировать результат.
    err := json.Unmarshal([]byte("{\"hello\":\"world\"}"), &hw)

    fmt.Printf("error: %s, struct: %+v\n", err, hw) // error: %!s(<nil>), struct: {Hello:world}
}
В API методах часто используются запросы с телом в виде JSON. Такие тела нужно декодировать в структуры и валидировать. Хоть это и не лучшая практика делать функции, в которых происходит несколько действий, но для простоты примера реализуйте функцию DecodeAndValidateRequest(requestBody []byte) (CreateUserRequest, error), которая декодирует тело запроса из JSON в структуру CreateUserRequest и валидирует ее. Если приходит невалидный JSON или структура заполнена неверно, функция должна вернуть ошибку.
Структура запроса:

type CreateUserRequest struct {
    Email                string `json:"email"`
    Password             string `json:"password"`
    PasswordConfirmation string `json:"password_confirmation"`
}
Список ошибок, которые нужно возвращать из функции:

// validation errors
var (
    errEmailRequired                = errors.New("email is required") // когда поле email не заполнено
    errPasswordRequired             = errors.New("password is required") // когда поле password не заполнено
    errPasswordConfirmationRequired = errors.New("password confirmation is required") // когда поле password_confirmation не заполнено
    errPasswordDoesNotMatch         = errors.New("password does not match with the confirmation") // когда поля password и password_confirmation не совпадают
)
Примеры работы функции DecodeAndValidateRequest:

DecodeAndValidateRequest([]byte("{\"email\":\"\",\"password\":\"test\",\"password_confirmation\":\"test\"}")) // CreateUserRequest{}, "email is required"
DecodeAndValidateRequest([]byte("{\"email\":\"test\",\"password\":\"\",\"password_confirmation\":\"test\"}")) // CreateUserRequest{}, "password is required"
DecodeAndValidateRequest([]byte("{\"email\":\"test\",\"password\":\"test\",\"password_confirmation\":\"\"}")) // CreateUserRequest{}, "password confirmation is required"
DecodeAndValidateRequest([]byte("{\"email\":\"test\",\"password\":\"test\",\"password_confirmation\":\"test2\"}")) // CreateUserRequest{}, "password does not match with the confirmation"
DecodeAndValidateRequest([]byte("{\"email\":\"email@test.com\",\"password\":\"passwordtest\",\"password_confirmation\":\"passwordtest\"}")) // CreateUserRequest{Email:"email@test.com", Password:"passwordtest", PasswordConfirmation:"passwordtest"}, nil
*/

package solution

import (
	"encoding/json"
	"errors"
)

// CreateUserRequest is a request to create a new user.
type CreateUserRequest struct {
	Email                string `json:"email"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"password_confirmation"`
}

// validation errors
var (
	errEmailRequired                = errors.New("email is required")
	errPasswordRequired             = errors.New("password is required")
	errPasswordConfirmationRequired = errors.New("password confirmation is required")
	errPasswordDoesNotMatch         = errors.New("password does not match with the confirmation")
)

// BEGIN (write your solution here)

// DecodeAndValidateRequest decodes the JSON body and validates the user creation request.
func DecodeAndValidateRequest(requestBody []byte) (CreateUserRequest, error) {
	req := CreateUserRequest{}

	err := json.Unmarshal(requestBody, &req)
	if err != nil {
		return CreateUserRequest{}, err
	}

	err = validateCreateUserRequest(req)
	if err != nil {
		return CreateUserRequest{}, err
	}

	return req, nil
}

func validateCreateUserRequest(req CreateUserRequest) error {
	if req.Email == "" {
		return errEmailRequired
	}

	if req.Password == "" {
		return errPasswordRequired
	}

	if req.PasswordConfirmation == "" {
		return errPasswordConfirmationRequired
	}

	if req.Password != req.PasswordConfirmation {
		return errPasswordDoesNotMatch
	}

	return nil
}

// END
