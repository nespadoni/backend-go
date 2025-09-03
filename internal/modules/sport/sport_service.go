package sport

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

func NewSportService(repo *Repository, validate *validator.Validate) *Service {
	return &Service{
		repo:     repo,
		validate: validate,
	}
}

func (s *Service) FindAll() ([]ListResponse, error) {
	sports, err := s.repo.FindAll()
	if err != nil {
		return nil, fmt.Errorf("erro no serviço ao buscar esportes: %w", err)
	}

	var listResponses []ListResponse
	if err := copier.Copy(&listResponses, &sports); err != nil {
		return nil, fmt.Errorf("erro ao converter dados: %w", err)
	}

	return listResponses, nil
}

func (s *Service) FindById(sportId uint) (Response, error) {
	sport, err := s.repo.FindById(sportId)
	if err != nil {
		return Response{}, fmt.Errorf("erro no serviço ao buscar esporte: %w", err)
	}

	var sportResponse Response
	if err := copier.Copy(&sportResponse, &sport); err != nil {
		return Response{}, fmt.Errorf("erro ao converter dados: %w", err)
	}

	return sportResponse, nil
}

func (s *Service) Create(req CreateRequest) (Response, error) {
	// Validar request
	if err := s.validate.Struct(req); err != nil {
		return Response{}, fmt.Errorf("dados inválidos: %w", err)
	}

	// Validar lógica de negócio
	if req.MaxPlayers < req.MinPlayers {
		return Response{}, fmt.Errorf("número máximo de jogadores deve ser maior ou igual ao mínimo")
	}

	// Converter para modelo
	var newSport models.Sport
	if err := copier.Copy(&newSport, &req); err != nil {
		return Response{}, fmt.Errorf("erro ao processar dados: %w", err)
	}

	if err := s.repo.Create(&newSport); err != nil {
		return Response{}, fmt.Errorf("erro ao criar esporte: %w", err)
	}

	// Buscar o esporte criado
	createdSport, err := s.repo.FindById(newSport.ID)
	if err != nil {
		return Response{}, fmt.Errorf("erro ao buscar esporte criado: %w", err)
	}

	var sportResponse Response
	if err := copier.Copy(&sportResponse, &createdSport); err != nil {
		return Response{}, fmt.Errorf("erro ao converter resposta: %w", err)
	}

	return sportResponse, nil
}

func (s *Service) Update(id uint, req UpdateRequest) (Response, error) {
	// Validar request
	if err := s.validate.Struct(&req); err != nil {
		return Response{}, fmt.Errorf("dados inválidos: %w", err)
	}

	// Validar lógica de negócio
	if req.MaxPlayers < req.MinPlayers {
		return Response{}, fmt.Errorf("número máximo de jogadores deve ser maior ou igual ao mínimo")
	}

	// Converter para modelo
	var sport models.Sport
	if err := copier.Copy(&sport, &req); err != nil {
		return Response{}, fmt.Errorf("erro ao processar dados: %w", err)
	}

	updatedSport, err := s.repo.Update(id, &sport)
	if err != nil {
		return Response{}, fmt.Errorf("erro ao atualizar esporte: %w", err)
	}

	var sportResponse Response
	if err := copier.Copy(&sportResponse, updatedSport); err != nil {
		return Response{}, fmt.Errorf("erro ao converter resposta: %w", err)
	}

	return sportResponse, nil
}

func (s *Service) UpdateStatus(id uint, req UpdateStatusRequest) (Response, error) {
	if err := s.validate.Struct(&req); err != nil {
		return Response{}, fmt.Errorf("dados inválidos: %w", err)
	}

	var sport models.Sport
	if err := copier.Copy(&sport, &req); err != nil {
		return Response{}, fmt.Errorf("erro ao processar dados: %w", err)
	}

	updatedSport, err := s.repo.Update(id, &sport)
	if err != nil {
		return Response{}, fmt.Errorf("erro ao atualizar status do esporte: %w", err)
	}

	var sportResponse Response
	if err := copier.Copy(&sportResponse, updatedSport); err != nil {
		return Response{}, fmt.Errorf("erro ao converter resposta: %w", err)
	}

	return sportResponse, nil
}

func (s *Service) Delete(id uint) error {
	if err := s.repo.Delete(id); err != nil {
		return fmt.Errorf("erro no serviço ao deletar esporte com ID %d: %w", id, err)
	}
	return nil
}

func (s *Service) FindPopular() ([]ListResponse, error) {
	sports, err := s.repo.FindPopular()
	if err != nil {
		return nil, fmt.Errorf("erro no serviço ao buscar esportes populares: %w", err)
	}

	var listResponses []ListResponse
	if err := copier.Copy(&listResponses, &sports); err != nil {
		return nil, fmt.Errorf("erro ao converter dados: %w", err)
	}

	return listResponses, nil
}
