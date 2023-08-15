package db

import (
	"context"
	"database/sql"
	"math/rand"
	"testing"
	"time"

	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {
	fake := faker.New()

	createdCategory := createRandomCategory(t)

	accountType := []string{"credit", "debit"}
	customRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomNumber := customRand.Intn(2)

	arg := CreateAccountParams{
		UserID:      createdCategory.UserID,
		Title:       fake.RandomStringWithLength(10),
		Description: fake.Lorem().Text(20),
		CategoryID:  createdCategory.ID,
		Type:        accountType[randomNumber],
		Value:       10,
		Date:        time.Now(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.UserID, account.UserID)
	require.Equal(t, arg.CategoryID, account.CategoryID)
	require.Equal(t, arg.Value, account.Value)
	require.Equal(t, arg.Title, account.Title)
	require.Equal(t, arg.Type, account.Type)
	require.Equal(t, arg.Description, account.Description)

	require.NotEmpty(t, account.CreatedAt)
	require.NotEmpty(t, account.Date)

	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	createdAccount := createRandomAccount(t)
	returnedAccount, err := testQueries.GetAccountById(context.Background(), createdAccount.ID)

	require.NoError(t, err)
	require.NotEmpty(t, returnedAccount)

	require.Equal(t, createdAccount.UserID, returnedAccount.UserID)
	require.Equal(t, createdAccount.CategoryID, returnedAccount.CategoryID)
	require.Equal(t, createdAccount.Value, returnedAccount.Value)
	require.Equal(t, createdAccount.Title, returnedAccount.Title)
	require.Equal(t, createdAccount.Type, returnedAccount.Type)
	require.Equal(t, createdAccount.Description, returnedAccount.Description)

	require.NotEmpty(t, returnedAccount.Date)
	require.NotEmpty(t, returnedAccount.CreatedAt)
}

func TestDeleteAccount(t *testing.T) {
	createdAccount := createRandomAccount(t)
	err := testQueries.DeleteAccount(context.Background(), createdAccount.ID)

	require.NoError(t, err)
}

func TestUpdateAccount(t *testing.T) {
	fake := faker.New()

	createdAccount := createRandomAccount(t)

	arg := UpdateAccountParams{
		ID:          createdAccount.ID,
		Title:       fake.RandomStringWithLength(10),
		Description: fake.Lorem().Text(20),
		Value:       15,
	}

	account, err := testQueries.UpdateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, createdAccount.ID, account.ID)
	require.Equal(t, arg.Title, account.Title)
	require.Equal(t, arg.Description, account.Description)
	require.Equal(t, arg.Value, account.Value)
	require.Equal(t, createdAccount.CreatedAt, account.CreatedAt)

}

func TestListAccounts(t *testing.T) {
	lastCreatedAccount := createRandomAccount(t)

	arg := GetAccountsParams{
		UserID: lastCreatedAccount.UserID,
		Type:   lastCreatedAccount.Type,
		CategoryID: sql.NullInt32{
			Valid: true,
			Int32: lastCreatedAccount.CategoryID,
		},
		Date: sql.NullTime{
			Valid: true,
			Time:  lastCreatedAccount.Date,
		},
		Title:       lastCreatedAccount.Title,
		Description: lastCreatedAccount.Description,
	}

	accounts, err := testQueries.GetAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, accounts)

	for _, account := range accounts {
		require.Equal(t, lastCreatedAccount.ID, account.ID)
		require.Equal(t, lastCreatedAccount.UserID, account.UserID)
		require.Equal(t, lastCreatedAccount.Title, account.Title)
		require.Equal(t, lastCreatedAccount.Description, account.Description)
		require.Equal(t, lastCreatedAccount.Value, account.Value)
		require.NotEmpty(t, lastCreatedAccount.CreatedAt)
		require.NotEmpty(t, lastCreatedAccount.Date)
	}
}

func TestListGetReports(t *testing.T) {
	lastCreatedAccount := createRandomAccount(t)

	arg := GetAccountsReportsParams{
		UserID: lastCreatedAccount.UserID,
		Type:   lastCreatedAccount.Type,
	}

	sumValue, err := testQueries.GetAccountsReports(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, sumValue)
}

func TestListGetGraph(t *testing.T) {
	lastCreatedAccount := createRandomAccount(t)

	arg := GetAccountsGraphParams{
		UserID: lastCreatedAccount.UserID,
		Type:   lastCreatedAccount.Type,
	}

	graphValue, err := testQueries.GetAccountsGraph(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, graphValue)
}
