package tournament

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

func NewTournamentService(repo *Repository, validate *validator.Validate) *Service {
	return &Service{repo: repo, validate: validate}
}

func (s *Service) FindAll() ([]ListResponse, error) {
	tournaments, err := s.repo.FindAll()
	if err != nil {
		return nil, fmt.Errorf("erro ao chamar metodo find all no repositorio: %v", err)
	}
	var listResponse []ListResponse
	if err := copier.Copy(&listResponse, tournaments); err != nil {
		return nil, fmt.Errorf("erro ao copiar metodo copier: %v", err)
	}
	return listResponse, nil
}

func (s *Service) FindByID(tournamentId uint) (ListResponse, error) {
	tournament, err := s.FindByID(tournamentId)
	if err != nil {
		return ListResponse{}, fmt.Errorf("erro ao chamar metodo find by id: %v", err)
	}

	var listResponse ListResponse
	if err := copier.Copy(&listResponse, tournament); err != nil {
		return ListResponse{}, fmt.Errorf("erro ao copiar metodo copier: %v", err)
	}

	return listResponse, nil
}

func (s *Service) Create(tournament CreateRequest) (Response, error) {
	if err := s.validate.Struct(&tournament); err != nil {
		return Response{}, fmt.Errorf("erro ao validar metodo create a tournament: %v", err)
	}

	if err := utils.ValidateEventDates(tournament.StartDate, tournament.EndDate); err != nil {
		return Response{}, fmt.Errorf("datas inválidas: %v", err)
	}

	var newTournament models.Tournament
	if err := copier.Copy(&newTournament, &tournament); err != nil {
		return Response{}, fmt.Errorf("erro ao copiar metodo copier: %v", err)
	}

	if err := s.repo.Create(&newTournament); err != nil {
		return Response{}, fmt.Errorf("erro ao chamar metodo create a tournament: %v", err)
	}

	createdTournament, err := s.repo.FindById(newTournament.Id)
	if err != nil {
		return Response{}, fmt.Errorf("erro ao buscar torneio criado: %v", err)
	}

	var listResponse Response
	if err := copier.Copy(&listResponse, &createdTournament); err != nil {
		return Response{}, fmt.Errorf("erro ao converter resposta: %v", err)
	}

	return listResponse, nil
}

func (s *Service) Update(id uint, req UpdateRequest) (Response, error) {
	if err := s.validate.Struct(&req); err != nil {
		return Response{}, fmt.Errorf("erro ao validar metodo update: %v", err)
	}

	if err := utils.ValidateEventDates(req.StartDate, req.EndDate); err != nil {
		return Response{}, fmt.Errorf("datas inválidas: %v", err)
	}

	var torunament models.Tournament
	if err := copier.Copy(&torunament, req); err != nil {
		return Response{}, fmt.Errorf("erro ao copiar metodo copier: %v", err)
	}

	updatedTournament, err := s.repo.FindById(id)
	if err != nil {
		return Response{}, fmt.Errorf("erro ao buscar torneio criado: %v", err)
	}

	var tournamentResponse Response
	if err := copier.Copy(&tournamentResponse, &updatedTournament); err != nil {
		return Response{}, fmt.Errorf("erro ao converter resposta: %v", err)
	}

	return tournamentResponse, nil
}

func (s *Service) Delete(id uint) error {
	if err := s.repo.Delete(id); err != nil {
		return fmt.Errorf("erro no serviço de deletar campeonato com Id %d: %v", err)
	}
	return nil
}
