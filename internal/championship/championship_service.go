package championship

import (
	"backend-go/internal/models"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/copier"
)

type Service struct {
	repo     Repository
	validate *validator.Validate
}

func NewChampionshipService(repo *Repository) *Service {
	return &Service{repo: *repo}
}

func (s Service) FindAll() ([]models.Championship, error) {

	championship, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	return championship, nil
}

func (s Service) FindById(championshipId *models.Championship) (Response, error) {
	championship, err := s.repo.FindById(championshipId)
	if err != nil {
		return Response{}, fmt.Errorf("erro no serviço de buscar usuário: %w", err)
	}

	var champResponse Response
	if err := copier.Copy(&champResponse, championship); err != nil {
		return Response{}, fmt.Errorf("erro ao buscar converter dados: %w", err)
	}

	return champResponse, nil
}

func (s Service) Create(championship *models.Championship) (Response, error) {
	if err := s.validate.Struct(championship); err != nil {
		return Response{}, fmt.Errorf("dados inválidos: %w", err)
	}

	var newChampionship models.Championship
	if err := copier.Copy(&newChampionship, &championship); err != nil {
		return Response{}, fmt.Errorf("erro ao processar dados: %w", err)
	}

	if err := s.repo.Create(&newChampionship); err != nil {
		return Response{}, fmt.Errorf("erro ao criar usuario: %w", err)
	}

	var championshipResponse Response
	if err := copier.Copy(&championshipResponse, &newChampionship); err != nil {
		return Response{}, fmt.Errorf("erro ao converter resposta: %w", err)
	}

	return championshipResponse, nil
}

func (s Service) Update(id int, championship *models.Championship) (Response, error) {
	if err := s.validate.Struct(&championship); err != nil {
		return Response{}, fmt.Errorf("dados invalidos: %w", err)
	}

	if err := s.repo.Update(id, championship); err != nil {
		return Response{}, fmt.Errorf("erro ao converter resposta: %w", err)
	}

	var championshipResponse Response
	if err := copier.Copy(&championshipResponse, &championship); err != nil {
		return Response{}, fmt.Errorf("erro ao converter resposta: %w", err)
	}

	return championshipResponse, nil
}

func (s Service) Delete(id int) error {

	if err := s.repo.Delete(id); err != nil {
		return fmt.Errorf("erro no serviço de deletar campeonato com ID %s: %w", id, err)
	}

	return nil
}
