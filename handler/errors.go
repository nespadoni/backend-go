package handler

import (
	"errors"
	"fmt"
	"net/http"
)

type AppError struct {
	Code    int // Código HTTP ou de negócio
	Message string
	Err     error // Erro raiz (interno)
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("[%d] %s: %v", e.Code, e.Message, e.Err)
	}
	return fmt.Sprintf("[%d] %s", e.Code, e.Message)
}

// Unwrap permite compatibilidade com errors.Is e errors.As
func (e *AppError) Unwrap() error {
	return e.Err
}

// New cria um novo AppError
func New(code int, message string, err error) *AppError {
	return &AppError{Code: code, Message: message, Err: err}
}

// Helpers genéricos para erros HTTP comuns:
func BadRequest(msg string, err error) *AppError {
	return New(http.StatusBadRequest, msg, err)
}

func NotFound(entity string, err error) *AppError {
	return New(http.StatusNotFound, fmt.Sprintf("%s não encontrado(a)", entity), err)
}

func Internal(msg string, err error) *AppError {
	return New(http.StatusInternalServerError, msg, err)
}

func Validation(err error) *AppError {
	return New(http.StatusUnprocessableEntity, "Erro ao validar dados", err)
}

func DB(err error) *AppError {
	return New(http.StatusInternalServerError, "Erro de banco de dados", err)
}

func Converter(entity string, err error) *AppError {
	return New(http.StatusInternalServerError, fmt.Sprintf("Erro ao converter %s", entity), err)
}

// IsAppError ajuda a identificar rapidamente se um erro é AppError
func IsAppError(err error) bool {
	var ae *AppError
	return errors.As(err, &ae)
}
