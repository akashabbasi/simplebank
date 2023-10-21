package db

import (
	"context"
	"testing"
	"time"

	"github.com/akashabbasi/simplebank/util"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	arg := CreateUserParams{
		Username:       util.RandomOwner(),
		HashedPassword: "secret",
		FullName:       util.RandomOwner(),
		Email:          util.RandomEmail(),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)

	require.True(t, user.PasswordChangedAt.IsZero())
	require.NotZero(t, user.CreatedAt)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	createdUser := createRandomUser(t)

	fetchedUser, err := testQueries.GetUser(
		context.Background(), createdUser.Username,
	)

	require.NoError(t, err)
	require.NotEmpty(t, fetchedUser)

	require.Equal(t, createdUser.Username, fetchedUser.Username)
	require.Equal(t, createdUser.FullName, fetchedUser.FullName)
	require.Equal(t, createdUser.Email, fetchedUser.Email)
	require.Equal(
		t, createdUser.HashedPassword,
		fetchedUser.HashedPassword,
	)

	require.WithinDuration(t, createdUser.CreatedAt, fetchedUser.CreatedAt, time.Second)
	require.WithinDuration(
		t, createdUser.PasswordChangedAt,
		fetchedUser.PasswordChangedAt, time.Second,
	)
}
