package auth

import (
	"backend-go/internal/models"
	"fmt"
	"gorm.io/gorm"
)

type PermissionService struct {
	db *gorm.DB
}

func NewPermissionService(db *gorm.DB) *PermissionService {
	return &PermissionService{db: db}
}

// Verifica se um usuário tem permissão específica em uma atlética
func (s *PermissionService) UserHasPermission(userID int, athleticID *int, permission string) (bool, error) {
	// Buscar todas as roles do usuário para essa atlética
	var userRoles []models.UserRoleAthletic
	query := s.db.Preload("Role").Where("user_id = ?", userID)

	if athleticID != nil {
		query = query.Where("athletic_id = ? OR athletic_id IS NULL", *athleticID)
	} else {
		query = query.Where("athletic_id IS NULL")
	}

	if err := query.Find(&userRoles).Error; err != nil {
		return false, fmt.Errorf("erro ao buscar roles do usuário: %w", err)
	}

	// Verificar permissão em cada role
	for _, userRole := range userRoles {
		permissions, err := ParsePermissions(userRole.Role.Permissions)
		if err != nil {
			continue // Pula se não conseguir parsear
		}

		if permissions.HasPermission(permission) {
			return true, nil
		}
	}

	return false, nil
}

// Busca o nível de permissão mais alto do usuário - VERSÃO GORM
func (s *PermissionService) GetUserMaxLevel(userID int, athleticID *int) (int, error) {
	var result struct {
		MaxLevel int `json:"max_level"`
	}

	// Usar GORM com joins
	query := s.db.Table("user_role_athletics").
		Select("COALESCE(MAX(roles.level), 0) as max_level").
		Joins("JOIN roles ON user_role_athletics.role_id = roles.id").
		Where("user_role_athletics.user_id = ?", userID)

	if athleticID != nil {
		query = query.Where("(user_role_athletics.athletic_id = ? OR user_role_athletics.athletic_id IS NULL)", *athleticID)
	} else {
		query = query.Where("user_role_athletics.athletic_id IS NULL")
	}

	if err := query.Scan(&result).Error; err != nil {
		return 0, fmt.Errorf("erro ao buscar nível máximo: %w", err)
	}

	return result.MaxLevel, nil
}

// ALTERNATIVA: Versão mais simples usando preload
func (s *PermissionService) GetUserMaxLevelSimple(userID int, athleticID *int) (int, error) {
	var userRoles []models.UserRoleAthletic
	query := s.db.Preload("Role").Where("user_id = ?", userID)

	if athleticID != nil {
		query = query.Where("athletic_id = ? OR athletic_id IS NULL", *athleticID)
	} else {
		query = query.Where("athletic_id IS NULL")
	}

	if err := query.Find(&userRoles).Error; err != nil {
		return 0, fmt.Errorf("erro ao buscar roles do usuário: %w", err)
	}

	maxLevel := 0
	for _, userRole := range userRoles {
		if userRole.Role.Level > maxLevel {
			maxLevel = userRole.Role.Level
		}
	}

	return maxLevel, nil
}

// Lista todas as permissões de um usuário
func (s *PermissionService) GetUserPermissions(userID int, athleticID *int) (PermissionSet, error) {
	var userRoles []models.UserRoleAthletic
	query := s.db.Preload("Role").Where("user_id = ?", userID)

	if athleticID != nil {
		query = query.Where("athletic_id = ? OR athletic_id IS NULL", *athleticID)
	} else {
		query = query.Where("athletic_id IS NULL")
	}

	if err := query.Find(&userRoles).Error; err != nil {
		return nil, fmt.Errorf("erro ao buscar roles: %w", err)
	}

	// Combinar todas as permissões
	allPermissions := make(PermissionSet)

	for _, userRole := range userRoles {
		rolePermissions, err := ParsePermissions(userRole.Role.Permissions)
		if err != nil {
			continue
		}

		// Merge das permissões
		for perm, value := range rolePermissions {
			if value {
				allPermissions[perm] = true
			}
		}
	}

	return allPermissions, nil
}

// Adiciona uma role para um usuário
func (s *PermissionService) AssignRole(userID, roleID int, athleticID *int) error {
	userRole := models.UserRoleAthletic{
		UserID:     userID,
		RoleID:     roleID,
		AthleticID: athleticID,
	}

	// Verifica se já existe usando FirstOrCreate
	result := s.db.Where(models.UserRoleAthletic{
		UserID:     userID,
		RoleID:     roleID,
		AthleticID: athleticID,
	}).FirstOrCreate(&userRole)

	return result.Error
}

// Remove uma role de um usuário
func (s *PermissionService) RemoveRole(userID, roleID int, athleticID *int) error {
	result := s.db.Where("user_id = ? AND role_id = ? AND athletic_id = ?",
		userID, roleID, athleticID).Delete(&models.UserRoleAthletic{})

	if result.Error != nil {
		return fmt.Errorf("erro ao remover role: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("role não encontrada para remoção")
	}

	return nil
}

// Lista todas as roles de um usuário em uma atlética específica
func (s *PermissionService) GetUserRoles(userID int, athleticID *int) ([]models.UserRoleAthletic, error) {
	var userRoles []models.UserRoleAthletic
	query := s.db.Preload("Role").Preload("Athletic").Where("user_id = ?", userID)

	if athleticID != nil {
		query = query.Where("athletic_id = ?", *athleticID)
	}

	if err := query.Find(&userRoles).Error; err != nil {
		return nil, fmt.Errorf("erro ao buscar roles do usuário: %w", err)
	}

	return userRoles, nil
}

// Verifica se usuário é admin de uma atlética específica
func (s *PermissionService) IsUserAthleticAdmin(userID int, athleticID int) (bool, error) {
	var count int64

	err := s.db.Table("user_role_athletics").
		Joins("JOIN roles ON user_role_athletics.role_id = roles.id").
		Where("user_role_athletics.user_id = ? AND user_role_athletics.athletic_id = ? AND roles.level >= ?",
			userID, athleticID, 50). // Nível 50+ = admin
		Count(&count).Error

	if err != nil {
		return false, fmt.Errorf("erro ao verificar admin: %w", err)
	}

	return count > 0, nil
}

// Busca usuários com determinada role em uma atlética
func (s *PermissionService) GetUsersWithRole(roleID int, athleticID *int) ([]models.User, error) {
	var users []models.User

	query := s.db.Joins("JOIN user_role_athletics ON users.id = user_role_athletics.user_id").
		Where("user_role_athletics.role_id = ?", roleID)

	if athleticID != nil {
		query = query.Where("user_role_athletics.athletic_id = ?", *athleticID)
	} else {
		query = query.Where("user_role_athletics.athletic_id IS NULL")
	}

	if err := query.Find(&users).Error; err != nil {
		return nil, fmt.Errorf("erro ao buscar usuários com role: %w", err)
	}

	return users, nil
}
