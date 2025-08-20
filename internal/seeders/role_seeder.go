package seeders

import (
	"backend-go/internal/models"
	"errors"
	"gorm.io/gorm"
	"log"
)

// GetSystemRoles define todas as roles padrão do sistema
func GetSystemRoles() []models.Role {
	return []models.Role{
		{
			Name:        "Super Admin",
			Description: "Administrador com acesso total ao sistema",
			Level:       100,
			IsSystem:    true,
			Permissions: `{"*": true}`,
		},
		{
			Name:        "Global Admin",
			Description: "Administrador global - pode criar atléticas",
			Level:       90,
			IsSystem:    true,
			Permissions: `{
				"create_athletic": true,
				"manage_users": true,
				"view_reports": true
			}`,
		},
		{
			Name:        "Presidente",
			Description: "Presidente da atlética - controle total",
			Level:       80,
			IsSystem:    true,
			Permissions: `{
				"manage_athletic": true,
				"manage_members": true,
				"create_championships": true,
				"manage_finances": true,
				"manage_teams": true,
				"create_news": true,
				"view_reports": true
			}`,
		},
		{
			Name:        "Vice-Presidente",
			Description: "Vice-presidente da atlética",
			Level:       75,
			IsSystem:    true,
			Permissions: `{
				"manage_members": true,
				"create_championships": true,
				"manage_teams": true,
				"create_news": true,
				"manage_matches": true
			}`,
		},
		{
			Name:        "Secretário",
			Description: "Secretário da atlética",
			Level:       60,
			IsSystem:    true,
			Permissions: `{
				"create_news": true,
				"manage_members": true,
				"view_reports": true,
				"moderate_comments": true
			}`,
		},
		{
			Name:        "Gestor Esportivo",
			Description: "Responsável pela área esportiva",
			Level:       50,
			IsSystem:    true,
			Permissions: `{
				"create_championships": true,
				"manage_teams": true,
				"manage_matches": true,
				"view_content": true
			}`,
		},
		{
			Name:        "Gestor Marketing",
			Description: "Responsável pelo marketing e comunicação",
			Level:       40,
			IsSystem:    true,
			Permissions: `{
				"create_news": true,
				"manage_social_media": true,
				"upload_media": true,
				"view_content": true
			}`,
		},
		{
			Name:        "Moderador",
			Description: "Moderador de conteúdo",
			Level:       30,
			IsSystem:    true,
			Permissions: `{
				"create_news": true,
				"moderate_comments": true,
				"view_content": true
			}`,
		},
		{
			Name:        "Membro",
			Description: "Membro ativo da atlética",
			Level:       10,
			IsSystem:    true,
			Permissions: `{
				"view_content": true,
				"comment": true,
				"like": true
			}`,
		},
		{
			Name:        "Seguidor",
			Description: "Usuário que segue a atlética",
			Level:       5,
			IsSystem:    true,
			Permissions: `{
				"view_public_content": true,
				"like": true
			}`,
		},
	}
}

// SeedRoles insere as roles padrão no banco
func SeedRoles(db *gorm.DB) error {
	log.Println("Iniciando seed das roles...")

	systemRoles := GetSystemRoles()

	for _, role := range systemRoles {
		// Verifica se a role já existe
		var existingRole models.Role
		result := db.Where("name = ? AND is_system = ?", role.Name, true).First(&existingRole)

		if result.Error != nil {
			// Se não existe, cria
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				if err := db.Create(&role).Error; err != nil {
					log.Printf("Erro ao criar role %s: %v", role.Name, err)
					return err
				}
				log.Printf("Role '%s' criada com sucesso", role.Name)
			} else {
				log.Printf("Erro ao verificar role %s: %v", role.Name, result.Error)
				return result.Error
			}
		} else {
			// Se existe, atualiza
			if err := db.Model(&existingRole).Updates(role).Error; err != nil {
				log.Printf("Erro ao atualizar role %s: %v", role.Name, err)
				return err
			}
			log.Printf("Role '%s' atualizada", role.Name)
		}
	}

	log.Println("Seed das roles concluído!")
	return nil
}
