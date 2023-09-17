/*
Какая-то функция возвращает критичные и некритичные ошибки:

// некритичная ошибка валидации
type nonCriticalError struct{}

func (e nonCriticalError) Error() string {
    return "validation error"
}

// критичные ошибки
var (
    errBadConnection = errors.New("bad connection")
    errBadRequest    = errors.New("bad request")
)
Реализуйте функцию GetErrorMsg(err error) string, которая возвращает текст ошибки, если она критичная. В случае неизвестной ошибки возвращается строка unknown error:

GetErrorMsg(errors.New("bad connection")) // "bad connection"
GetErrorMsg(errors.New("bad request")) // "bad request"
GetErrorMsg(nonCriticalError{}) // ""
GetErrorMsg(errors.New("random error")) // "unknown error"
*/

package solution

import (
	"errors"
)

type nonCriticalError struct{}

func (e nonCriticalError) Error() string {
	return "validation error"
}

var (
	errBadConnection = errors.New("bad connection")
	errBadRequest    = errors.New("bad request")
)

const unknownErrorMsg = "unknown error"

// BEGIN (write your solution here)

var criticalErrs = []error{errBadRequest, errBadConnection}

// GetErrorMsg returns the err message if the error is critical. Otherwise it returns an empty string.
func GetErrorMsg(err error) string {
	for _, crErr := range criticalErrs {
		if errors.Is(err, crErr) {
			return crErr.Error()
		}
	}

	if errors.As(err, &nonCriticalError{}) {
		return ""
	}

	return unknownErrorMsg
}

// END
