package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	name := "João"
	email := "joao@email.com"
	age := "20"

	user := NewUser(
		WithName(name),
		WithEmail(email),
		WithAge(age),
	)

	assert.Equal(t, user.Name, name)
	assert.Equal(t, user.Email, email)
	assert.Equal(t, user.Age, age)
}
