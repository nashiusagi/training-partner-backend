package usecases

import (
	"training-partner/domains"
	"training-partner/repositories"
)

type PostUsecase interface {
	GetAll() ([]*domains.Post, error)
}

type postUseCase struct {
	postRepository repositories.PostRepository
}

func NewPostUsecase(postRepository repositories.PostRepository) PostUsecase {
	return &postUseCase{postRepository}
}

func (u *postUseCase) GetAll() ([]*domains.Post, error) {
	return u.postRepository.GetAll()
}
