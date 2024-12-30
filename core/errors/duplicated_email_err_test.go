package errors

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDuplicatedEmailErr(t *testing.T) {
	test := struct {
		name  string
		email string
		err   string
	}{
		name:  "Test duplicated email error",
		email: "joao@email.com",
		err:   fmt.Sprintf("Email %v has aready been taken.", "joao@email.com"),
	}
	t.Run(test.name, func(t *testing.T) {
		err := NewDuplicatedEmailErr(test.email)
		assert.Equal(t, test.err, err.Error())
	})
}
