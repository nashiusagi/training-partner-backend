package usecases

import (
	"training-partner/domains"
	"training-partner/repositories"
)

type PostUsecase interface {
	GetAll() ([]*domains.Post, error)
	FindById(id int) (*domains.Post, error)
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

func (u *postUseCase) FindById(id int) (*domains.Post, error) {
	return u.postRepository.FindById(id)
}
