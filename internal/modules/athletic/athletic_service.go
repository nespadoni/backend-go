package athletic

import (
	"backend-go/internal/models"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/copier"
)

type Service struct {
	repo     *Repository
	validate *validator.Validate
}

func NewAthleticService(repo *Repository, validate *validator.Validate) *Service {
	return &Service{
		repo:     repo,
		validate: validate,
	}
}

func (s *Service) FindAll() ([]ListResponse, error) {
	athletics, err := s.repo.FindAll()
	if err != nil {
		return nil, fmt.Errorf("erro no serviço de buscar todos campeonatos: %w", err)
	}

	var listResponse []ListResponse
	if err := copier.Copy(&listResponse, &athletics); err != nil {
		return nil, fmt.Errorf("erro ao converter dados: %w", err)
	}

	return listResponse, nil
}

func (s *Service) FindById(athleticId uint) (Response, error) {
	athletic, err := s.repo.FindById(athleticId)
	if err != nil {
		return Response{}, fmt.Errorf("erro no serviço de buscar atlética: %w", err)
	}

	var athleticResponse Response
	if err := copier.Copy(&athleticResponse, &athletic); err != nil {
		return Response{}, fmt.Errorf("erro ao converter dados: %w", err)
	}

	return athleticResponse, nil
}

func (s *Service) Create(athletic CreateRequest) (Response, error) {
	if err := s.validate.Struct(athletic); err != nil {
		return Response{}, fmt.Errorf("dados inválidos: %w", err)
	}

	var newAthletic models.Athletic
	if err := copier.Copy(&newAthletic, &athletic); err != nil {
		return Response{}, fmt.Errorf("erro ao processar dados: %w", err)
	}

	if err := s.repo.Create(&newAthletic); err != nil {
		return Response{}, fmt.Errorf("erro ao criar atletica: %w", err)
	}

	createdAthletic, err := s.repo.FindById(newAthletic.ID)
	if err != nil {
		return Response{}, fmt.Errorf("erro ao buscar atletica criada: %w", err)
	}

	var athleticResponse Response
	if err := copier.Copy(&athleticResponse, &createdAthletic); err != nil {
		return Response{}, fmt.Errorf("erro ao converter resposta: %w", err)
	}

	return athleticResponse, nil
}

func (s *Service) Update(id uint, req UpdateRequest) (Response, error) {
	if err := s.validate.Struct(&req); err != nil {
		return Response{}, fmt.Errorf("dados inválidos: %w", err)
	}

	var athletic models.Athletic
	if err := copier.Copy(&athletic, &req); err != nil {
		return Response{}, fmt.Errorf("erro ao procesar dados: %w", err)
	}

	updatedAthletic, err := s.repo.Update(id, &athletic)
	if err != nil {
		return Response{}, fmt.Errorf("erro ao atualizar atlética: %w", err)
	}

	var athleticResponse Response
	if err := copier.Copy(&athleticResponse, &updatedAthletic); err != nil {
		return Response{}, fmt.Errorf("erro ao converter resposta: %w", err)
	}

	return athleticResponse, nil
}

func (s *Service) UpdateStatus(id uint, req UpdateStatusRequest) (Response, error) {
	if err := s.validate.Struct(&req); err != nil {
		return Response{}, fmt.Errorf("dados inválidos: %w", err)
	}

	var athletic models.Athletic
	if err := copier.Copy(&athletic, &req); err != nil {
		return Response{}, fmt.Errorf("erro ao processar dados: %w", err)
	}

	updatedAthletic, err := s.repo.Update(id, &athletic)
	if err != nil {
		return Response{}, fmt.Errorf("erro ao atualizar status da atlética: %w", err)
	}

	var athleticResponse Response
	if err := copier.Copy(&athleticResponse, &updatedAthletic); err != nil {
		return Response{}, fmt.Errorf("erro ao converter resposta: %w", err)
	}

	return athleticResponse, nil
}

func (s *Service) Delete(id uint) error {
	if err := s.repo.Delete(id); err != nil {
		return fmt.Errorf("erro no serviço ao deletar atlética com ID %d: %w", id, err)
	}
	return nil
}
