package usecase

import (
	"context"

	"github.com/carloscerqueira/user-management/core/domain"
	"github.com/carloscerqueira/user-management/core/repository"
)

type (
	ListAllUsers interface {
		Execute(ctx *context.Context) (*[]domain.User, error)
	}
	listAllUsers struct {
		repo repository.UserRepository
	}
)

func NewListAllUsers(repo repository.UserRepository) ListAllUsers {
	return &listAllUsers{repo: repo}
}

func (l *listAllUsers) Execute(ctx *context.Context) (*[]domain.User, error) {
	out, err := l.repo.ListAllUsers(ctx)
	if err != nil {
		return &[]domain.User{}, err
	}
	return &out, nil
}
