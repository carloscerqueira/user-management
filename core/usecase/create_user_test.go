package usecase

import (
	"context"
	"testing"

	"github.com/carloscerqueira/user-management/core/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateUserExecute(t *testing.T) {
	defaultUser := domain.NewUser(
		domain.WithName("João"),
		domain.WithEmail("joao@email.com"),
		domain.WithAge(20),
	)

	testCases := []struct {
		name     string
		user     domain.User
		expected domain.User
		err      error
	}{
		{
			name:     "Success",
			user:     *defaultUser,
			expected: *defaultUser,
			err:      nil,
		},
		{
			name:     "Error",
			user:     *defaultUser,
			expected: domain.User{},
			err:      databaseErr,
		},
	}

	for _, sc := range testCases {
		t.Run(sc.name, func(t *testing.T) {
			r := NewMockDatabase()
			uc := NewCreateUser(r)

			r.On("CreateUser", mock.Anything, mock.Anything).Return(sc.expected, sc.err)

			ctx := context.TODO()
			out, err := uc.Execute(&ctx, &sc.user)

			assert.Equal(t, sc.expected, *out)
			assert.Equal(t, sc.err, err)
		})
	}
}
