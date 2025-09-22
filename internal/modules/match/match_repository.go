package match

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

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{BaseRepository: repository.BaseRepository{DB: db}}
}

func (r *Repository) FindAll() ([]models.Match, error) {
	var match []models.Match
	if err := r.DB.Find(&match).Error; err != nil {
		return nil, fmt.Errorf("match not found")
	}
	return match, nil
}

func (r *Repository) FindByID(id uint) (*models.Match, error) {
	var match models.Match
	if err := r.DB.First(&match, id).Error; err != nil {
		return nil, fmt.Errorf("match not found")
	}
	return &match, nil
}

func (r *Repository) Create(match *models.Match) error {
	return r.WithTransaction(func(tx *gorm.DB) error {
		// Valida se o torneio existe
		var tournament models.Tournament
		if err := tx.First(&tournament, match.TournamentID).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return fmt.Errorf("tournament with ID %d not found", match.TournamentID)
			}
			return fmt.Errorf("failed to validate tournament: %w", err)
		}

		// Valida se o time A existe
		var teamA models.Team
		if err := tx.First(&teamA, match.TeamAID).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return fmt.Errorf("team A with ID %d not found", match.TeamAID)
			}
			return fmt.Errorf("failed to validate team A: %w", err)
		}

		// Valida se o time B existe
		var teamB models.Team
		if err := tx.First(&teamB, match.TeamBID).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return fmt.Errorf("team B with ID %d not found", match.TeamBID)
			}
			return fmt.Errorf("failed to validate team B: %w", err)
		}

		// Valida se os times são diferentes
		if match.TeamAID == match.TeamBID {
			return fmt.Errorf("team A and team B cannot be the same")
		}

		// Valida se o sport existe (se fornecido)
		if match.SportID != nil {
			var sport models.Sport
			if err := tx.First(&sport, *match.SportID).Error; err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					return fmt.Errorf("sport with ID %d not found", *match.SportID)
				}
				return fmt.Errorf("failed to validate sport: %w", err)
			}
		}

		// Verifica se já existe uma partida entre os mesmos times no mesmo torneio
		var existingMatch models.Match
		err := tx.Where("tournament_id = ? AND ((team_a_id = ? AND team_b_id = ?) OR (team_a_id = ? AND team_b_id = ?))",
			match.TournamentID, match.TeamAID, match.TeamBID, match.TeamBID, match.TeamAID).
			First(&existingMatch).Error

		if err == nil {
			return fmt.Errorf("match between these teams already exists in this tournament")
		} else if !errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("failed to check existing match: %w", err)
		}

		// Cria a partida
		if err := tx.Create(match).Error; err != nil {
			return fmt.Errorf("failed to create match: %w", err)
		}

		return nil
	})
}

func (r *Repository) Update(id uint, match *models.Match) (*models.Match, error) {
	return nil, nil
}

func (r *Repository) Delete(id uint) error {
	return nil
}
