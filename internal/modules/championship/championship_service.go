package championship

import (
	"backend-go/internal/models"
	"fmt"
	"time"

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

func (s *Service) FindAll() ([]ListResponse, error) {

	championships, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	var listResponses []ListResponse
	if err := copier.Copy(&listResponses, &championships); err != nil {
		return nil, fmt.Errorf("erro ao converter dados: %w", err)
	}

	return listResponses, nil
}

func (s *Service) FindById(championshipId int) (Response, error) {
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

func (s *Service) Create(championship CreateRequest) (Response, error) {
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

func (s *Service) Update(id int, req UpdateRequest) (Response, error) {
	if err := s.validate.Struct(&req); err != nil {
		return Response{}, fmt.Errorf("dados invalidos: %w", err)
	}

	// Validar datas
	if err := s.validateDates(req.StartDate, req.EndDate); err != nil {
		return Response{}, fmt.Errorf("datas inválidas: %w", err)
	}

	// Converter para Modelo
	var championship models.Championship
	if err := copier.Copy(&championship, &req); err != nil {
		return Response{}, fmt.Errorf("erro ao processar dados: %w", err)
	}

	updatedChampionship, err := s.repo.Update(id, &championship)
	if err != nil {
		return Response{}, fmt.Errorf("erro ao atualiar campeonato: %w", err)
	}

	var championshipResponse Response
	if err := copier.Copy(&championshipResponse, &updatedChampionship); err != nil {
		return Response{}, fmt.Errorf("erro ao converter resposta: %w", err)
	}

	return championshipResponse, nil
}

func (s *Service) Delete(id int) error {

	if err := s.repo.Delete(id); err != nil {
		return fmt.Errorf("erro no serviço de deletar campeonato com ID %d: %w", id, err)
	}

	return nil
}

func (s *Service) validateDates(startDate, endDate time.Time) error {
	if endDate.Before(startDate) {
		return fmt.Errorf("data de término deve ser posterior à data de início")
	}

	if startDate.Before(time.Now().Truncate(24 * time.Hour)) {
		return fmt.Errorf("data de início não pode ser no passado")
	}

	return nil
}
