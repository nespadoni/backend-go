package user

//
//import (
//	"github.com/go-playground/validator/v10"
//	"github.com/jinzhu/copier"
//)
//
//type UserService struct {
//	repo     UserRepository
//	validate *validator.Validate
//}
//
//func NewUserService(repo UserRepository, validate *validator.Validate) *UserService {
//	return &UserService{
//		repo:     repo,
//		validate: validate,
//	}
//}
//
//func (s *UserService) GetAll() ([]UserResponse, error) {
//	users := s.repo.FindAll()
//
//	var userResponses []UserResponse
//	if err := copier.Copy(&userResponses, users); err != nil {
//		return nil, handler.Converter("User -> UserResponse", err)
//	}
//	return userResponses, nil
//}
//
//func (s *UserService) GetById(userId int) (UserResponse, error) {
//	user := s.repo.FindById(userId)
//	if userId == 0 {
//		return UserResponse{}, handler.NotFound("Usuário", nil)
//	}
//
//	var userResponse UserResponse
//	if err := copier.Copy(&userResponse, user); err != nil {
//		return UserResponse{}, handler.Converter("User -> UserResponse", err)
//	}
//	return userResponse, nil
//}
//
//func (s *UserService) CreateUser(req CreateUserRequest) (UserResponse, error) {
//	if err := s.validate.Struct(req); err != nil {
//		return UserResponse{}, handler.Validation(err)
//	}
//
//	var newUser User
//	if err := copier.Copy(&newUser, &req); err != nil {
//		return UserResponse{}, handler.Converter("CreateUserRequest -> User", err)
//	}
//
//	createdUser := s.repo.SaveUser(newUser)
//
//	var userResponse UserResponse
//	if err := copier.Copy(&userResponse, createdUser); err != nil {
//		return UserResponse{}, handler.Converter("User -> UserResponse", err)
//	}
//
//	return userResponse, nil
//}
