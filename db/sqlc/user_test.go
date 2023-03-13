package db

import (
	"context"
	"github.com/stretchr/testify/require"
	"roma/go-finance/util"
	"testing"
)

func createRandomUser(t *testing.T) User {

	arg := CreateUserParams{
		Username: util.GetRandomName(),
		Email:    util.GetRandomEmail(),
		Password: util.GetRandomPassword(),
	}
	user, err := queries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.Email, user.Email)
	require.NotEmpty(t, user.CreatedAt)
	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetRandomUser(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := queries.GetUser(context.Background(), user1.Username)
	require.NoError(t, err)
	require.NotEmpty(t, user2)
	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.Email, user2.Email)
	require.NotEmpty(t, user2.CreatedAt)
}

func TestGetRandomUserById(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := queries.GetUserById(context.Background(), user1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, user2)
	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.Email, user2.Email)
	require.NotEmpty(t, user2.CreatedAt)
}
