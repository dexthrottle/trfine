package service

import (
	"context"

	"github.com/dexthrottle/trfine/internal/dto"
	"github.com/dexthrottle/trfine/internal/model"
	"github.com/dexthrottle/trfine/internal/repository"
	"github.com/dexthrottle/trfine/pkg/logging"
	"github.com/mashingan/smapping"
)

type CategoryService interface {
	Insert(ctx context.Context, p dto.CreateCategoryDTO) (*model.Category, error)
	Delete(ctx context.Context, b model.Category, userTgId int) error
	All(ctx context.Context, userTgId int) ([]*model.Category, error)
}

type categoryService struct {
	ctx                context.Context
	categoryRepository repository.CategoryRepository
	log                logging.Logger
}

func NewCategoryService(
	ctx context.Context,
	categoryRepository repository.CategoryRepository,
	log logging.Logger,
) CategoryService {
	return &categoryService{
		ctx:                ctx,
		categoryRepository: categoryRepository,
		log:                log,
	}
}

func (s *categoryService) Insert(ctx context.Context, c dto.CreateCategoryDTO) (*model.Category, error) {
	category := model.Category{}
	err := smapping.FillStruct(&category, smapping.MapFields(&c))
	if err != nil {
		s.log.Errorf("Failed map %v: ", err)
		return nil, err
	}
	categoryM, errI := s.categoryRepository.InsertCategory(ctx, category)
	if errI != nil {
		s.log.Errorf("category insert error: %v", errI)
		return nil, err
	}
	return categoryM, nil
}

func (s *categoryService) Delete(ctx context.Context, b model.Category, userId int) error {
	err := s.categoryRepository.DeleteCategory(ctx, b, userId)
	if err != nil {
		s.log.Errorf("category delete error: %v", err)
		return err
	}
	return nil
}

func (s *categoryService) All(ctx context.Context, userTgId int) ([]*model.Category, error) {
	categories, err := s.categoryRepository.AllCategory(ctx, userTgId)
	if err != nil {
		s.log.Errorf("get all categories error: %v", err)
		return nil, err
	}
	return categories, nil
}
