package championship

import (
	"backend-go/internal/models"
	"backend-go/pkg/utils"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/copier"
)

type Service struct {
	repo     *Repository
	validate *validator.Validate
}

func NewChampionshipService(repo *Repository, validate *validator.Validate) *Service {
	return &Service{
		repo:     repo,
		validate: validate,
	}

}

func (s *Service) FindAll() ([]ListResponse, error) {
	championships, err := s.repo.FindAll()
	if err != nil {
		return nil, fmt.Errorf("erro ao chamar metodo find all no repositório: %w", err)
	}

	var listResponses []ListResponse
	if err := copier.Copy(&listResponses, &championships); err != nil {
		return nil, fmt.Errorf("erro ao converter dados: %w", err)
	}

	return listResponses, nil
}

func (s *Service) FindById(championshipId uint) (Response, error) {
	championship, err := s.repo.FindById(championshipId)
	if err != nil {
		return Response{}, fmt.Errorf("erro no serviço de buscar campeonato: %w", err)
	}

	var champResponse Response
	if err := copier.Copy(&champResponse, championship); err != nil {
		return Response{}, fmt.Errorf("erro ao converter dados: %w", err)
	}

	return champResponse, nil
}

func (s *Service) Create(championship CreateRequest) (Response, error) {
	if err := s.validate.Struct(championship); err != nil {
		return Response{}, fmt.Errorf("dados inválidos: %w", err)
	}

	if err := utils.ValidateEventDates(championship.StartDate, championship.EndDate); err != nil {
		return Response{}, fmt.Errorf("datas inválidas: %w", err)
	}

	var newChampionship models.Championship
	if err := copier.Copy(&newChampionship, &championship); err != nil {
		return Response{}, fmt.Errorf("erro ao processar dados: %w", err)
	}

	if err := s.repo.Create(&newChampionship); err != nil {
		return Response{}, fmt.Errorf("erro ao criar campeonato: %w", err)
	}

	createdChampionship, err := s.repo.FindById(newChampionship.ID)
	if err != nil {
		return Response{}, fmt.Errorf("erro ao buscar campeonato criado: %w", err)
	}

	var championshipResponse Response
	if err := copier.Copy(&championshipResponse, &createdChampionship); err != nil {
		return Response{}, fmt.Errorf("erro ao converter resposta: %w", err)
	}

	return championshipResponse, nil
}

func (s *Service) Update(id uint, req UpdateRequest) (Response, error) {
	if err := s.validate.Struct(&req); err != nil {
		return Response{}, fmt.Errorf("dados invalidos: %w", err)
	}

	// Validar datas
	if err := utils.ValidateEventDates(req.StartDate, req.EndDate); err != nil {
		return Response{}, fmt.Errorf("datas inválidas: %w", err)
	}

	// Converter para Modelo
	var championship models.Championship
	if err := copier.Copy(&championship, &req); err != nil {
		return Response{}, fmt.Errorf("erro ao processar dados: %w", err)
	}

	updatedChampionship, err := s.repo.Update(id, &championship)
	if err != nil {
		return Response{}, fmt.Errorf("erro ao atualizar campeonato: %w", err)
	}

	var championshipResponse Response
	if err := copier.Copy(&championshipResponse, &updatedChampionship); err != nil {
		return Response{}, fmt.Errorf("erro ao converter resposta: %w", err)
	}

	return championshipResponse, nil
}

func (s *Service) UpdateStatus(id uint, req UpdateStatusRequest) (Response, error) {
	if err := s.validate.Struct(&req); err != nil {
		return Response{}, fmt.Errorf("dados inválidos: %w", err)
	}

	var championship models.Championship
	if err := copier.Copy(&championship, &req); err != nil {
		return Response{}, fmt.Errorf("erro ao processar dados: %w", err)
	}

	updatedChampionship, err := s.repo.Update(id, &championship)
	if err != nil {
		return Response{}, fmt.Errorf("erro ao atualizar status do campeonato: %w", err)
	}

	var championshipResponse Response
	if err := copier.Copy(&championshipResponse, &updatedChampionship); err != nil {
		return Response{}, fmt.Errorf("erro ao converter resposta: %w", err)
	}

	return championshipResponse, nil
}

func (s *Service) Delete(id uint) error {

	if err := s.repo.Delete(id); err != nil {
		return fmt.Errorf("erro no serviço de deletar campeonato com ID %d: %w", id, err)
	}

	return nil
}
