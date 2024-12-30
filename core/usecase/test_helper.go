package usecase

import (
	"context"
	"fmt"

	"github.com/carloscerqueira/user-management/core/domain"
	"github.com/stretchr/testify/mock"
)

type databaseMock struct {
	mock.Mock
}

var databaseErr = fmt.Errorf("Database error")

func NewMockDatabase() *databaseMock {
	return &databaseMock{}
}

func (d *databaseMock) CreateUser(ctx *context.Context, user *domain.User) (domain.User, error) {
	args := d.Called()
	return args.Get(0).(domain.User), args.Error(1)
}

func (d *databaseMock) ListAllUsers(ctx *context.Context) ([]domain.User, error) {
	args := d.Called()
	return args.Get(0).([]domain.User), args.Error(1)
}

func (d *databaseMock) GetUserById(ctx *context.Context, id int) (domain.User, error) {
	args := d.Called()
	return args.Get(0).(domain.User), args.Error(1)
}

func (d *databaseMock) UpdateUser(ctx *context.Context, user *domain.User) (domain.User, error) {
	args := d.Called()
	return args.Get(0).(domain.User), args.Error(1)
}

func (d *databaseMock) DeleteUser(ctx *context.Context, id int) error {
	args := d.Called()
	return args.Error(1)
}
