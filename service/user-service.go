package service

import (
	"log"

	"github.com/LaouiniSofiene/golang_api/dto"
	"github.com/LaouiniSofiene/golang_api/entity"
	"github.com/LaouiniSofiene/golang_api/repository"
	"github.com/mashingan/smapping"
)

type UserService interface {
	Insert(u dto.UserCreateDTO) entity.User
	Update(u dto.UserUpdateDTO) entity.User
	Delete(u entity.User)
	All() []entity.User
	FindByID(userID uint64) entity.User
	FindByEmail(email string) entity.User
	IsDuplicateEmail(email string) bool
	GetAllUsers(u *entity.User, p *entity.Pagination) (*[]entity.User, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepo,
	}
}

func (service *userService) Insert(u dto.UserCreateDTO) entity.User {
	user := entity.User{}
	err := smapping.FillStruct(&user, smapping.MapFields(&u))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.userRepository.InsertUser(user)
	return res
}

func (service *userService) Update(u dto.UserUpdateDTO) entity.User {
	user := entity.User{}
	err := smapping.FillStruct(&user, smapping.MapFields(&u))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.userRepository.UpdateUser(user)
	return res
}

func (service *userService) Delete(u entity.User) {
	service.userRepository.DeleteUser(u)
}

func (service *userService) All() []entity.User {
	return service.userRepository.AllUser()
}
func (service *userService) GetAllUsers(u *entity.User, p *entity.Pagination) (*[]entity.User, error) {
	return service.userRepository.GetAllUsers(u, p)
}

func (service *userService) FindByID(userID uint64) entity.User {
	return service.userRepository.FindUserById(userID)
}

func (service *userService) FindByEmail(email string) entity.User {
	return service.userRepository.FindByEmail(email)
}

func (service *userService) IsDuplicateEmail(email string) bool {
	res := service.userRepository.IsDuplicateEmail(email)
	return !(res.Error == nil)
}
