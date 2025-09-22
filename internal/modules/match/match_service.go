package match

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

func NewTournamntService(repo *Repository, validate *validator.Validate) *Service {
	return &Service{repo: repo, validate: validate}
}

func (s *Service) FindAll() ([]ListResponse, error) {
	tournaments, err := s.repo.FindAll()
	if err != nil {
		return nil, fmt.Errorf("failed to find all tournaments: %w", err)
	}

	var listResponse []ListResponse
	if err := copier.Copy(&listResponse, tournaments); err != nil {
		return nil, fmt.Errorf("failed to copy to list: %w", err)
	}
	return listResponse, nil
}

func (s *Service) FindByID(id uint) (*ListResponse, error) {
	tournament, err := s.repo.FindByID(id)
	if err != nil {
		return nil, fmt.Errorf("failed to find tournament: %w", err)
	}

	var listResponse ListResponse
	if err := copier.Copy(&listResponse, tournament); err != nil {
		return nil, fmt.Errorf("failed to copy to list: %w", err)
	}
	return &listResponse, nil
}

func (s *Service) Create(tournament CreateRequest) (Response, error) {
	if err := s.validate.Struct(&tournament); err != nil {
		return Response{}, fmt.Errorf("failed to validate tournament: %w", err)
	}

	var newMatch models.Match
	if err := copier.Copy(&newMatch, &tournament); err != nil {
		return Response{}, fmt.Errorf("failed to copy match: %w", err)
	}

	if err := s.repo.Create(&newMatch); err != nil {
		return Response{}, fmt.Errorf("failed to insert match: %w", err)
	}

	createdMatch, err := s.repo.FindByID(newMatch.Id)
	if err != nil {
		return Response{}, fmt.Errorf("failed to find match: %w", err)
	}

	var response Response
	if err := copier.Copy(&response, createdMatch); err != nil {
		return Response{}, fmt.Errorf("failed to copy match: %w", err)
	}
	return response, nil
}

func (s *Service) Update(id uint, req UpdateRequest) (Response, error) {
	if err := s.validate.Struct(&req); err != nil {
		return Response{}, fmt.Errorf("failed to validate update: %w", err)
	}

	var match models.Match
	if err := copier.Copy(&match, &req); err != nil {
		return Response{}, fmt.Errorf("failed to copy match: %w", err)
	}

	updatedMatch, err := s.repo.Update(id, &match)
	if err != nil {
		return Response{}, fmt.Errorf("failed to find match: %w", err)
	}

	var response Response
	if err := copier.Copy(&response, &updatedMatch); err != nil {
		return Response{}, fmt.Errorf("failed to copy match: %w", err)
	}

	return response, nil
}

func (s *Service) Delete(id uint) error {
	if err := s.repo.Delete(id); err != nil {
		return fmt.Errorf("failed to delete tournament: %w", err)
	}
	return nil
}
