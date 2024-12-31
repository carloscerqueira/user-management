package usecase

import (
	"context"
	"testing"

	"github.com/carloscerqueira/user-management/core/domain"
	"github.com/carloscerqueira/user-management/core/usecase/input"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetUserByIdExecute(t *testing.T) {
	userResp := domain.User{
		Id:    1,
		Name:  "João",
		Email: "joao@email.com",
		Age:   20,
	}

	testCases := []struct {
		name     string
		input    *input.GetUserByIdInput
		expected domain.User
		err      error
	}{
		{
			name:     "Success",
			input:    &input.GetUserByIdInput{Id: 1},
			expected: userResp,
			err:      nil,
		},
		{
			name:     "Error",
			input:    &input.GetUserByIdInput{Id: 1},
			expected: domain.User{},
			err:      databaseErr,
		},
	}

	for _, sc := range testCases {
		t.Run(sc.name, func(t *testing.T) {
			r := NewMockDatabase()
			uc := NewGetUserById(r)

			r.On("GetUserById", mock.Anything, mock.Anything).Return(sc.expected, sc.err)

			ctx := context.TODO()
			out, err := uc.Execute(&ctx, sc.input)

			assert.Equal(t, sc.expected, *out)
			assert.Equal(t, sc.err, err)
		})
	}
}
