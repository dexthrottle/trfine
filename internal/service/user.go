package service

import (
	"context"

	"github.com/dexthrottle/trfine/internal/dto"
	"github.com/dexthrottle/trfine/internal/model"
	"github.com/dexthrottle/trfine/internal/repository"
	"github.com/dexthrottle/trfine/pkg/logging"
	"github.com/mashingan/smapping"
)

type UserService interface {
	Insert(ctx context.Context, user dto.CreateUserDTO) (*model.User, error)
	Profile(ctx context.Context, userID string) (*model.User, error)
	IsDuplicateUserTGID(ctx context.Context, tgID int) (bool, error)
	FindUserByTgUserId(ctx context.Context, userTgId int) (*model.User, error)
}

type userService struct {
	ctx            context.Context
	userRepository repository.UserRepository
	log            logging.Logger
}

func NewUserService(ctx context.Context, userRepo repository.UserRepository, log logging.Logger) UserService {
	return &userService{
		ctx:            ctx,
		userRepository: userRepo,
		log:            log,
	}
}

func (service *userService) IsDuplicateUserTGID(ctx context.Context, tgID int) (bool, error) {
	res, err := service.userRepository.IsDuplicateUserTGID(ctx, tgID)
	if err != nil {
		return false, err
	}
	return res, nil
}

// Обновить пользователя
func (s *userService) Insert(ctx context.Context, user dto.CreateUserDTO) (*model.User, error) {
	userToCreate := model.User{}
	err := smapping.FillStruct(&userToCreate, smapping.MapFields(&user))
	if err != nil {
		s.log.Errorf("Failed map %v:", err)
		return nil, err
	}
	updatedUser, err := s.userRepository.InsertUser(ctx, userToCreate)
	if err != nil {
		s.log.Errorf("insert user error: %v", err)
		return nil, err
	}
	return updatedUser, nil
}

// Профиль пользователя
func (s *userService) Profile(ctx context.Context, userID string) (*model.User, error) {
	userProfile, err := s.userRepository.ProfileUser(ctx, userID)
	if err != nil {
		s.log.Errorf("profile user error: %v", err)
		return nil, err
	}
	return userProfile, nil
}

func (s *userService) FindUserByTgUserId(ctx context.Context, userTgId int) (*model.User, error) {
	user, err := s.userRepository.FindUserByTgUserId(ctx, userTgId)
	if err != nil {
		s.log.Errorf("is allowed to edit user error: %v", err)
		return nil, err
	}
	return user, nil
}
