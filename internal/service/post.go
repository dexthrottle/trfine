package service

import (
	"context"

	"github.com/dexthrottle/trfine/internal/dto"
	"github.com/dexthrottle/trfine/internal/model"
	"github.com/dexthrottle/trfine/internal/repository"
	"github.com/dexthrottle/trfine/pkg/logging"
	"github.com/mashingan/smapping"
)

type PostService interface {
	Insert(ctx context.Context, b dto.PostCreateDTO) (*model.Post, error)
	Delete(ctx context.Context, b model.Post, userId int) error
	All(ctx context.Context, userTgId int) ([]*model.Post, error)
}

type postService struct {
	ctx            context.Context
	postRepository repository.PostRepository
	log            logging.Logger
}

func NewPostService(
	ctx context.Context,
	postRepository repository.PostRepository,
	log logging.Logger,
) PostService {
	return &postService{
		ctx:            ctx,
		postRepository: postRepository,
		log:            log,
	}
}

func (s *postService) Insert(ctx context.Context, p dto.PostCreateDTO) (*model.Post, error) {
	post := model.Post{}
	err := smapping.FillStruct(&post, smapping.MapFields(&p))
	if err != nil {
		s.log.Errorf("Failed map %v: ", err)
		return nil, err
	}
	postM, errI := s.postRepository.InsertPost(ctx, post)
	if errI != nil {
		s.log.Errorf("post insert error: %v", errI)
		return nil, err
	}
	return postM, nil
}

func (s *postService) Delete(ctx context.Context, p model.Post, userId int) error {
	err := s.postRepository.DeletePost(ctx, p, userId)
	if err != nil {
		s.log.Errorf("post delete error: %v", err)
		return err
	}
	return nil
}

func (s *postService) All(ctx context.Context, userId int) ([]*model.Post, error) {
	posts, err := s.postRepository.AllPost(ctx, userId)
	if err != nil {
		s.log.Errorf("get all posts error: %v", err)
		return nil, err
	}
	return posts, nil
}
