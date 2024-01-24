package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	gmock "github.com/golang/mock/gomock"
	mockdb "github.com/hagios2/simple-bank/db/mock"
	db "github.com/hagios2/simple-bank/db/sqlc"
	"github.com/hagios2/simple-bank/util"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAccountAPI(t *testing.T) {
	account := randomAccount()
	ctrl := gmock.NewController(t)
	defer ctrl.Finish()

	store := mockdb.NewMockStore(ctrl)

	store.EXPECT().
		GetAccount(gmock.Any(), gmock.Eq(account.ID)).
		Times(1).
		Return(account, nil)

	// start test server and send request
	server := NewServer(store)
	recorder := httptest.NewRecorder()

	url := fmt.Sprintf("/accounts/%d", account.ID)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	require.NoError(t, err)

	server.router.ServeHTTP(recorder, request)
	//check response
	require.Equal(t, http.StatusOK, recorder.Code)

	requireBodyMatchAccount(t, recorder.Body, account)
}

func randomAccount() db.Accounts {
	return db.Accounts{
		ID:       util.RandomInt(1, 1000),
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}
}

func requireBodyMatchAccount(t *testing.T, body *bytes.Buffer, account db.Accounts) {
	data, err := ioutil.ReadAll(body)
	require.NoError(t, err)

	var getAccount db.Accounts
	err = json.Unmarshal(data, &getAccount)
	require.NoError(t, err)
	require.Equal(t, account, getAccount)
}
