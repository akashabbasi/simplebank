package util

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestPassword(t *testing.T) {
	password := RandomString(8)

	hashedPasswd1, err := HashPassword(password)

	require.NoError(t, err)
	require.NotEmpty(t, hashedPasswd1)

	err = CheckPassword(password, hashedPasswd1)
	require.NoError(t, err)

	hashedPasswd2, err := HashPassword(password)

	require.NoError(t, err)
	require.NotEmpty(t, hashedPasswd2)
	require.NotEqual(t, hashedPasswd1, hashedPasswd2)

	wrongPasswd := RandomString(8)
	err = CheckPassword(wrongPasswd, hashedPasswd1)
	require.Equal(t, err.Error(), bcrypt.ErrMismatchedHashAndPassword.Error())
}