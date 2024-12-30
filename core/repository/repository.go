package repository

import (
	"context"

	"github.com/carloscerqueira/user-management/core/domain"
)

type UserRepository interface {
	CreateUser(ctx *context.Context, user *domain.User) (domain.User, error)
	ListAllUsers(ctx *context.Context) ([]domain.User, error)
	GetUserById(ctx *context.Context, id int) (domain.User, error)
	UpdateUser(ctx *context.Context, user *domain.User) (domain.User, error)
	DeleteUser(ctx *context.Context, id int) error
}
