package handler

import (
	"fmt"
)

type AppError struct {
	Code    int
	Message string
	Err     error
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("[%s] %s: %v", e.Code, e.Message, e.Err)
	}
	return fmt.Sprintf("[%s] %s", e.Code, e.Message)
}

func NewError(code int, message string, err error) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}

func NewDBQueryError(err error) *AppError {
	return NewError(500, "Erro ao consultar o banco de dados", err)
}

func NewDBConnectionError(err error) *AppError {
	return NewError(500, "Erro ao conectar ao banco de dados", err)
}

func NewBadRequestError(message string, err error) *AppError {
	return NewError(400, message, err)
}

func NewRequestError(err error) *AppError {
	return NewError(400, "Erro na requisição", err)
}

func NewNotFoundError(entity string, err error) *AppError {
	return NewError(404, fmt.Sprintf("%s não encontrado(a)", entity), err)
}

func NewCreateError(entity string, err error) *AppError {
	return NewError(500, fmt.Sprintf("Erro ao criar %s", entity), err)
}

func NewUpdateError(entity string, err error) *AppError {
	return NewError(500, fmt.Sprintf("Erro ao atualizar %s", entity), err)
}

func NewDeleteError(entity string, err error) *AppError {
	return NewError(500, fmt.Sprintf("Erro ao deletar %s", entity), err)
}

func NewReadError(entity string, err error) *AppError {
	return NewError(500, fmt.Sprintf("Erro ao ler %s", entity), err)
}
