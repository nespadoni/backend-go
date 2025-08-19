package auth

import (
	"encoding/json"
	"fmt"
)

// Definir todas as permissões do sistema
const (
	// Permissões globais
	PermissionCreateAthletic = "create_athletic"
	PermissionManageUsers    = "manage_users"
	PermissionSystemAdmin    = "*"

	// Permissões de atlética
	PermissionManageAthletic      = "manage_athletic"
	PermissionManageMembers       = "manage_members"
	PermissionCreateChampionships = "create_championships"
	PermissionManageFinances      = "manage_finances"
	PermissionManageTeams         = "manage_teams"
	PermissionManageMatches       = "manage_matches"

	// Permissões de conteúdo
	PermissionCreateNews        = "create_news"
	PermissionManageSocialMedia = "manage_social_media"
	PermissionUploadMedia       = "upload_media"
	PermissionModerateComments  = "moderate_comments"

	// Permissões básicas
	PermissionViewContent       = "view_content"
	PermissionViewPublicContent = "view_public_content"
	PermissionComment           = "comment"
	PermissionLike              = "like"
	PermissionViewReports       = "view_reports"
)

// PermissionSet representa um conjunto de permissões
type PermissionSet map[string]bool

// Verifica se tem uma permissão específica
func (p PermissionSet) HasPermission(permission string) bool {
	// Se tem permissão de admin total
	if p[PermissionSystemAdmin] {
		return true
	}

	return p[permission]
}

// Converte JSON string para PermissionSet
func ParsePermissions(permissionsJSON string) (PermissionSet, error) {
	var permissions PermissionSet
	if err := json.Unmarshal([]byte(permissionsJSON), &permissions); err != nil {
		return nil, fmt.Errorf("erro ao decodificar permissões: %w", err)
	}
	return permissions, nil
}

// Converte PermissionSet para JSON string
func (p PermissionSet) ToJSON() (string, error) {
	data, err := json.Marshal(p)
	if err != nil {
		return "", fmt.Errorf("erro ao codificar permissões: %w", err)
	}
	return string(data), nil
}
