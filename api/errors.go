package api

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(c *fiber.Ctx, err error) error {
	if apiError, ok := err.(Error); ok {
		return c.Status(apiError.Code).JSON(apiError)
	}
	apiError := NewError(http.StatusInternalServerError, err.Error())
	return c.Status(apiError.Code).JSON(apiError)
}

type Error struct {
	Code int    `json:"code"`
	Err  string `json:"error"`
}

func (e Error) Error() string {
	return e.Err
}

// Error implements the Error interface
func NewError(code int, err string) Error {
	return Error{
		Code: code,
		Err:  err,
	}
}

func ErrUnauthorized() Error {
	return Error{
		Code: http.StatusUnauthorized,
		Err:  "unauthorized request",
	}
}

func ErrInvalidID() Error {
	return Error{
		Code: http.StatusBadRequest,
		Err:  "invalid id given",
	}
}

func ErrBadRequest() Error {
	return Error{
		Code: http.StatusBadRequest,
		Err:  "invalid JSON request",
	}
}

func ErrResourceNotFound(res string) Error {
	return Error{
		Code: http.StatusNotFound,
		Err:  res + " resource not found",
	}
}
