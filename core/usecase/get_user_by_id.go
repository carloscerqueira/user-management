package usecase

import (
	"context"

	"github.com/carloscerqueira/user-management/core/domain"
	"github.com/carloscerqueira/user-management/core/repository"
	"github.com/carloscerqueira/user-management/core/usecase/input"
)

type (
	GetUserById interface {
		Execute(ctx *context.Context, i *input.GetUserByIdInput) (*domain.User, error)
	}
	getUserById struct {
		repo repository.UserRepository
	}
)

func NewGetUserById(repo repository.UserRepository) GetUserById {
	return &getUserById{repo: repo}
}

func (g *getUserById) Execute(ctx *context.Context, i *input.GetUserByIdInput) (*domain.User, error) {
	out, err := g.repo.GetUserById(ctx, i.Id)
	if err != nil {
		return &domain.User{}, err
	}
	return &out, nil
}
