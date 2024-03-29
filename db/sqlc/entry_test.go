package db

import (
	"context"
	"github.com/hagios2/simple-bank/util"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func createTestAccount(t *testing.T) Accounts {
	user := createRandomUser(t)

	arg := CreateAccountParams{
		Owner:    user.Username,
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	return account
}

func createRandomEntry(t *testing.T) Entries {
	account1 := createTestAccount(t)
	amount := util.RandomMoney()

	arg := CreateEntryParams{
		AccountID: account1.ID,
		Amount:    amount,
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, arg.AccountID, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)
	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)

	return entry
}

func TestCreateEntry(t *testing.T) {
	createRandomEntry(t)
}

func TestGetEntry(t *testing.T) {
	entry1 := createRandomEntry(t)
	entry2, err := testQueries.GetEntry(context.Background(), entry1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, entry2)

	require.Equal(t, entry1.ID, entry2.ID)
	require.Equal(t, entry1.AccountID, entry2.AccountID)
	require.Equal(t, entry1.Amount, entry2.Amount)
	require.WithinDuration(t, entry1.CreatedAt, entry2.CreatedAt, time.Second)
}

func TestListEntries(t *testing.T) {
	account1 := createTestAccount(t)
	amount := util.RandomMoney()

	for i := 0; i < 10; i++ {
		arg := CreateEntryParams{
			AccountID: account1.ID,
			Amount:    amount,
		}

		entry, _ := testQueries.CreateEntry(context.Background(), arg)
		require.NotEmpty(t, entry)
	}

	for i := 0; i < 3; i++ {
		createRandomEntry(t)
	}

	arg := ListEntriesParams{
		AccountID: account1.ID,
		Limit:     5,
		Offset:    5,
	}

	entries, err := testQueries.ListEntries(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, entries, 5)

	for _, entry := range entries {
		require.NotEmpty(t, entry)
	}
}
