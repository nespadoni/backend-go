package tournament

import (
	"backend-go/internal/models"
	"backend-go/internal/repository"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type Repository struct {
	repository.BaseRepository
}

func NewTournamentRepository(db *gorm.DB) *Repository {
	return &Repository{BaseRepository: repository.BaseRepository{DB: db}}
}

func (r *Repository) FindAll() ([]models.Tournament, error) {
	var tournament []models.Tournament

	if err := r.DB.Find(&tournament).Error; err != nil {
		return nil, fmt.Errorf("error finding tournamentes: %w", err)
	}

	return tournament, nil
}

func (r *Repository) FindById(id uint) (models.Tournament, error) {
	var tournament models.Tournament

	if err := r.DB.Preload("Championship").Preload("Sport").First(&tournament, id).Error; err != nil {
		return models.Tournament{}, fmt.Errorf("error finding tournament: %w", err)
	}

	return tournament, nil
}

func (r *Repository) FindByName(name string) (models.Tournament, error) {
	var tournament models.Tournament

	if err := r.DB.Where("name = ?", name).First(&tournament).Error; err != nil {
		return models.Tournament{}, fmt.Errorf("error finding tournament: %w", err)
	}

	return tournament, nil
}

func (r *Repository) Create(tournament *models.Tournament) error {
	return r.WithTransaction(func(tx *gorm.DB) error {
		var championship models.Championship
		var sport models.Sport
		if err := tx.First(&championship, tournament.ChampionshipID).Error; err != nil {
			return fmt.Errorf("error finding championship: %w", err)
		}
		if err := tx.First(&sport, tournament.SportID).Error; err != nil {
			return fmt.Errorf("error finding sport: %w", err)
		}
		return tx.Create(tournament).Error
	})
}

func (r *Repository) Update(id uint, tournament *models.Tournament) (*models.Tournament, error) {
	if err := r.DB.First(&tournament, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("tournament with Id %d not found", id)
		}
		return nil, fmt.Errorf("error finding tournament: %w", err)
	}

	// Atualizar
	if err := r.DB.Model(&tournament).Where("id = ?", id).Updates(tournament).Error; err != nil {
		return nil, fmt.Errorf("error updating tournament: %w", err)
	}

	// Buscar o registro atualizado com relacionamento
	var updatedTournament models.Tournament
	if err := r.DB.Preload("Championship").Preload("Sport").First(&updatedTournament, id).Error; err != nil {
		return nil, fmt.Errorf("error finding tournament: %w", err)
	}

	return &updatedTournament, nil
}

func (r *Repository) Delete(id uint) error {
	result := r.DB.Delete(&models.Tournament{}, id)

	if result.Error != nil {
		return fmt.Errorf("error deleting tournament: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("tournament with Id %d not found", id)
	}
	return nil
}
