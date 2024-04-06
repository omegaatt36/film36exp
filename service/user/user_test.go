package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncryptPassword(t *testing.T) {
	hashedPassword, err := encryptPassword("P@ssw0rd")
	t.Log(hashedPassword)
	assert.New(t).NoError(err)
}
