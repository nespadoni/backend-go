package auth

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Middleware struct {
	permissionService *PermissionService
}

func NewMiddleware(permissionService *PermissionService) *Middleware {
	return &Middleware{
		permissionService: permissionService,
	}
}

// RequirePermission verifica se o usuário tem a permissão necessária
func (m *Middleware) RequirePermission(permission string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userIDStr := c.GetHeader("X-User-ID")
		if userIDStr == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "usuário não autenticado"})
			c.Abort()
			return
		}

		userID, err := strconv.Atoi(userIDStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID de usuário inválido"})
			c.Abort()
			return
		}

		var athleticID *int
		if athleticIDStr := c.Param("athletic_id"); athleticIDStr != "" {
			if id, err := strconv.Atoi(athleticIDStr); err == nil {
				athleticID = &id
			}
		}

		hasPermission, err := m.permissionService.UserHasPermission(userID, athleticID, permission)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "erro ao verificar permissões"})
			c.Abort()
			return
		}

		if !hasPermission {
			c.JSON(http.StatusForbidden, gin.H{"error": "permissão negada"})
			c.Abort()
			return
		}

		c.Set("user_id", userID)
		c.Set("athletic_id", athleticID)
		c.Next()
	}
}

// RequireLevel verifica se o usuário tem nível mínimo necessário
func (m *Middleware) RequireLevel(minLevel int) gin.HandlerFunc {
	return func(c *gin.Context) {
		userIDStr := c.GetHeader("X-User-ID")
		if userIDStr == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "usuário não autenticado"})
			c.Abort()
			return
		}

		userID, err := strconv.Atoi(userIDStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID de usuário inválido"})
			c.Abort()
			return
		}

		var athleticID *int
		if athleticIDStr := c.Param("athletic_id"); athleticIDStr != "" {
			if id, err := strconv.Atoi(athleticIDStr); err == nil {
				athleticID = &id
			}
		}

		// Usar versão simples para maior confiabilidade
		userLevel, err := m.permissionService.GetUserMaxLevelSimple(userID, athleticID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "erro ao verificar nível"})
			c.Abort()
			return
		}

		if userLevel < minLevel {
			c.JSON(http.StatusForbidden, gin.H{"error": "nível de acesso insuficiente"})
			c.Abort()
			return
		}

		c.Set("user_id", userID)
		c.Set("user_level", userLevel)
		c.Set("athletic_id", athleticID)
		c.Next()
	}
}
