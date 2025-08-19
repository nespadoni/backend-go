package championship

import (
	"backend-go/internal/models"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type ChampionshipRepository struct {
	DB *gorm.DB
}

func NewChampionshipRepository(db *gorm.DB) *ChampionshipRepository {
	return &ChampionshipRepository{DB: db}
}

func (repo *ChampionshipRepository) FindAll() ([]models.Championship, error) {
	var championships []models.Championship

	result := repo.DB.Find(&championships)
	if result.Error != nil {
		// Retorna o erro quando houver problema na consulta
		return nil, result.Error
	}

	return championships, nil
}

func (r *ChampionshipRepository) Create(championship *models.Championship) error {
	// Transação para garantir sincronismo
	tx := r.DB.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Validar se o Athletic existe
	var athletic models.Athletic
	if err := tx.First(&athletic, championship.AthleticID).Error; err != nil {
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("atlética com ID %d não encontrada", championship.AthleticID)
		}
		return fmt.Errorf("erro ao validar atlética: %w", err)
	}

	// Criar o championship
	if err := tx.Create(championship).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("erro ao criar campeonatos: %w", err)
	}

	// Commit da transação
	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("erro ao confirmar transação: %w", err)
	}

	// Carrega o relacionamento de Athletic
	if err := r.DB.Preload("Athletic").First(&championship, championship.ID).Error; err != nil {
		return fmt.Errorf("campeonato criado, mas erro ao carregar dados da atlética: %w", err)
	}

	return nil
}
