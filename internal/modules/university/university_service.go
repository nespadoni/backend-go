package university

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

func NewUniversityService(repo *Repository, validate *validator.Validate) *Service {
	return &Service{
		repo:     repo,
		validate: validate,
	}
}

func (s *Service) FindAll() ([]ListResponse, error) {
	universities, err := s.repo.FindAll()
	if err != nil {
		return nil, fmt.Errorf("erro no serviço de buscar universidades: %w", err)
	}

	var listResponses []ListResponse
	if err := copier.Copy(&listResponses, &universities); err != nil {
		return nil, fmt.Errorf("erro ao converter dados: %w", err)
	}

	return listResponses, nil
}

func (s *Service) FindById(universityId uint) (Response, error) {
	university, err := s.repo.FindById(universityId)
	if err != nil {
		return Response{}, fmt.Errorf("erro no serviço de buscar universidade: %w", err)
	}

	var universityResponse Response
	if err := copier.Copy(&universityResponse, &university); err != nil {
		return Response{}, fmt.Errorf("erro ao converter dados: %w", err)
	}

	return universityResponse, nil
}

func (s *Service) Create(university CreateRequest) (Response, error) {
	if err := s.validate.Struct(university); err != nil {
		return Response{}, fmt.Errorf("dados inválidos: %w", err)
	}

	var newUniversity models.University
	if err := copier.Copy(&newUniversity, &university); err != nil {
		return Response{}, fmt.Errorf("erro ao processar dados: %w", err)
	}

	if err := s.repo.Create(&newUniversity); err != nil {
		return Response{}, fmt.Errorf("erro ao criar universidade: %w", err)
	}

	createdUniversity, err := s.repo.FindById(newUniversity.ID)
	if err != nil {
		return Response{}, fmt.Errorf("erro ao buscar universidade criada: %w", err)
	}

	var universityResponse Response
	if err := copier.Copy(&universityResponse, &createdUniversity); err != nil {
		return Response{}, fmt.Errorf("erro ao converter resposta: %w", err)
	}

	return universityResponse, nil
}

func (s *Service) Update(id uint, req UpdateRequest) (Response, error) {
	if err := s.validate.Struct(&req); err != nil {
		return Response{}, fmt.Errorf("dados inválidos: %w", err)
	}

	var university models.University
	if err := copier.Copy(&university, &req); err != nil {
		return Response{}, fmt.Errorf("erro ao processar dados: %w", err)
	}

	updatedUniversity, err := s.repo.Update(id, &university)
	if err != nil {
		return Response{}, fmt.Errorf("erro ao atualizar universidade: %w", err)
	}

	var universityResponse Response
	if err := copier.Copy(&universityResponse, updatedUniversity); err != nil {
		return Response{}, fmt.Errorf("erro ao converter resposta: %w", err)
	}

	return universityResponse, nil
}

func (s *Service) Delete(id uint) error {
	if err := s.repo.Delete(id); err != nil {
		return fmt.Errorf("erro no serviço ao deletar universidade com ID %d: %w", id, err)
	}
	return nil
}
