package usecase

import (
	"context"

	"github.com/carloscerqueira/user-management/core/domain"
	"github.com/carloscerqueira/user-management/core/repository"
)

type (
	CreateUser interface {
		Execute(ctx *context.Context, u *domain.User) (*domain.User, error)
	}
	createUser struct {
		repo repository.UserRepository
	}
)

func NewCreateUser(repo repository.UserRepository) CreateUser {
	return &createUser{repo: repo}
}

func (c *createUser) Execute(ctx *context.Context, u *domain.User) (*domain.User, error) {
	user := domain.NewUser(
		domain.WithName(u.Name),
		domain.WithEmail(u.Email),
		domain.WithAge(u.Age),
	)

	out, err := c.repo.CreateUser(ctx, user)
	if err != nil {
		return &domain.User{}, err
	}
	return &out, nil
}
