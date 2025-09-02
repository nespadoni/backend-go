package utils

// ErrorResponse representa uma resposta de erro padronizada
type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message,omitempty"`
}

// NewErrorResponse cria uma nova resposta de erro
func NewErrorResponse(error, message string) ErrorResponse {
	return ErrorResponse{
		Error:   error,
		Message: message,
	}
}

// NewSimpleErrorResponse cria uma resposta de erro apenas com o campo error
func NewSimpleErrorResponse(error string) ErrorResponse {
	return ErrorResponse{
		Error: error,
	}
}
