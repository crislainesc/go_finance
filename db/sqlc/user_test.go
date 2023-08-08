package db

import (
	"context"
	"testing"

	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	fake := faker.New()

	arg := CreateUserParams{
		Username: fake.Person().Name(),
		Password: fake.Internet().Password(),
		Email:    fake.Internet().Email(),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.Password, user.Password)
	require.Equal(t, arg.Email, user.Email)
	require.NotEmpty(t, user.CreatedAt)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	createdUser := createRandomUser(t)
	returnedUser, err := testQueries.GetUser(context.Background(), createdUser.Username)

	require.NoError(t, err)
	require.NotEmpty(t, returnedUser)

	require.Equal(t, returnedUser.Username, createdUser.Username)
	require.Equal(t, returnedUser.Password, createdUser.Password)
	require.Equal(t, returnedUser.Email, createdUser.Email)
}
