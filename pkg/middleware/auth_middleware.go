package middleware

import (
	"backend-go/pkg/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(jwtSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, utils.ErrorResponse{Error: "unauthorized", Message: "Cabeçalho de autorização não encontrado"})
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, utils.ErrorResponse{Error: "unauthorized", Message: "Formato do cabeçalho de autorização inválido"})
			return
		}

		tokenString := parts[1]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, http.ErrAbortHandler
			}
			return []byte(jwtSecret), nil
		})

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, utils.ErrorResponse{Error: "unauthorized", Message: "Token inválido ou expirado"})
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if ok && token.Valid {
			// Adiciona o Id do usuário ao contexto para uso posterior nos controladores
			c.Set("user_id", claims["user_id"])
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, utils.ErrorResponse{Error: "unauthorized", Message: "Claims do token inválidas"})
			return
		}

		c.Next()
	}
}
