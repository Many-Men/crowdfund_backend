package service

import (
	"context"
	"github.com/Many-Men/crowdfund_backend/config"
	_interface "github.com/Many-Men/crowdfund_backend/internal/domain/interface"
	"github.com/Many-Men/crowdfund_backend/internal/infrastructure/entity"
	infrastructureInterface "github.com/Many-Men/crowdfund_backend/internal/service/interface"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserServiceImpl struct {
	userRepository infrastructureInterface.UserRepository
	config         *config.Config
}

func NewUserServiceImpl(userRepo infrastructureInterface.UserRepository) _interface.UserService {
	return &UserServiceImpl{
		userRepository: userRepo,
		config:         config.Load(),
	}
}

func (s *UserServiceImpl) CreateUser(user entity.User) (primitive.ObjectID, error) {
	return s.userRepository.CreateUser(context.Background(), user)
}

func (s *UserServiceImpl) GetUserByID(id primitive.ObjectID) (*entity.User, error) {
	return s.userRepository.GetUserByID(context.Background(), id)
}

func (s *UserServiceImpl) GetUserByEmail(email string) (*entity.User, error) {
	return s.userRepository.GetUserByEmail(context.Background(), email)
}

func (s *UserServiceImpl) UpdateUserBalance(id primitive.ObjectID, balance float64) error {
	return s.userRepository.UpdateUserBalance(context.Background(), id, balance)
}

func (s *UserServiceImpl) DeleteUser(id primitive.ObjectID) error {
	return s.userRepository.DeleteUser(context.Background(), id)
}

func (s *UserServiceImpl) ListUsers() ([]entity.User, error) {
	return s.userRepository.ListUsers(context.Background())
}
