package db

import (
	"context"
	"math/rand"
	"testing"
	"time"

	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/require"
)

func createRandomCategory(t *testing.T) Category {
	fake := faker.New()

	createdUser := createRandomUser(t)

	categoryType := []string{"credit", "debit"}
	customRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomNumber := customRand.Intn(2)

	arg := CreateCategoryParams{
		UserID:      createdUser.ID,
		Title:       fake.RandomStringWithLength(10),
		Type:        categoryType[randomNumber],
		Description: fake.Lorem().Text(20),
	}

	category, err := testQueries.CreateCategory(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, category)

	require.Equal(t, arg.Title, category.Title)
	require.Equal(t, arg.Description, category.Description)
	require.Equal(t, arg.Type, category.Type)
	require.Equal(t, arg.UserID, category.UserID)
	require.NotEmpty(t, category.CreatedAt)

	return category
}

func TestCreateCategory(t *testing.T) {
	createRandomCategory(t)
}

func TestGetCategory(t *testing.T) {
	createdCategory := createRandomCategory(t)
	returnedCategory, err := testQueries.GetCategoryById(context.Background(), createdCategory.ID)

	require.NoError(t, err)
	require.NotEmpty(t, returnedCategory)

	require.Equal(t, returnedCategory.Title, createdCategory.Title)
	require.Equal(t, returnedCategory.Description, createdCategory.Description)
	require.Equal(t, returnedCategory.Type, createdCategory.Type)
	require.Equal(t, returnedCategory.UserID, createdCategory.UserID)
}

func TestDeleteCategory(t *testing.T) {
	createdCategory := createRandomCategory(t)
	err := testQueries.DeleteCategory(context.Background(), createdCategory.ID)

	require.NoError(t, err)
}

func TestUpdateCategory(t *testing.T) {
	fake := faker.New()

	createdCategory := createRandomCategory(t)

	arg := UpdateCategoryParams{
		ID:          createdCategory.ID,
		Title:       fake.RandomStringWithLength(10),
		Description: fake.Lorem().Text(20),
	}

	category, err := testQueries.UpdateCategory(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, category)

	require.Equal(t, arg.Title, category.Title)
	require.Equal(t, arg.Description, category.Description)
	require.NotEmpty(t, category.CreatedAt)
}

func TestListCategories(t *testing.T) {
	lastCreatedCategory := createRandomCategory(t)

	arg := GetCategoriesParams{
		UserID:      lastCreatedCategory.UserID,
		Type:        lastCreatedCategory.Type,
		Title:       lastCreatedCategory.Title,
		Description: lastCreatedCategory.Description,
	}

	categories, err := testQueries.GetCategories(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, categories)

	for _, category := range categories {
		require.Equal(t, lastCreatedCategory.ID, category.ID)
		require.Equal(t, lastCreatedCategory.UserID, category.UserID)
		require.Equal(t, lastCreatedCategory.Title, category.Title)
		require.Equal(t, lastCreatedCategory.Description, category.Description)
		require.NotEmpty(t, lastCreatedCategory.CreatedAt)
	}
}
