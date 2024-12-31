package usecase

import (
	"context"
	"testing"

	"github.com/carloscerqueira/user-management/core/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestListAllUsers(t *testing.T) {
	userResp := []domain.User{
		{
			Id:    1,
			Name:  "João",
			Email: "joao@email.com",
			Age:   20,
		},
		{
			Id:    2,
			Name:  "Maria",
			Email: "maria@email.com",
			Age:   21,
		},
	}

	testCases := []struct {
		name     string
		expected []domain.User
		err      error
	}{
		{
			name:     "Success",
			expected: userResp,
			err:      nil,
		},
		{
			name:     "Error",
			expected: []domain.User{},
			err:      databaseErr,
		},
	}

	for _, sc := range testCases {
		t.Run(sc.name, func(t *testing.T) {
			r := NewMockDatabase()
			uc := NewListAllUsers(r)

			r.On("ListAllUsers", mock.Anything, mock.Anything).Return(sc.expected, sc.err)

			ctx := context.TODO()
			out, err := uc.Execute(&ctx)

			assert.Equal(t, sc.expected, *out)
			assert.Equal(t, sc.err, err)
		})
	}
}
