package test_version1

import (
	"testing"

	"github.com/pip-services3-go/pip-services3-commons-go/data"

	"github.com/pip-services-users/pip-clients-accounts-go/version1"
	"github.com/stretchr/testify/assert"
)

type AccountsClientFixtureV1 struct {
	Client version1.IAccountsClientV1
}

var ACCOUNT_ID1 = data.IdGenerator.NextLong()
var ACCOUNT_ID2 = data.IdGenerator.NextLong()
var ACCOUNT1 = version1.NewAccountV1(ACCOUNT_ID1, "Test Account "+ACCOUNT_ID1, ACCOUNT_ID1+"@conceptual.vision")
var ACCOUNT2 = version1.NewAccountV1(ACCOUNT_ID2, "Test Account "+ACCOUNT_ID2, ACCOUNT_ID2+"@conceptual.vision")

func NewAccountsClientFixtureV1(client version1.IAccountsClientV1) *AccountsClientFixtureV1 {
	return &AccountsClientFixtureV1{
		Client: client,
	}
}

func (c *AccountsClientFixtureV1) clear() {
	page, _ := c.Client.GetAccounts("", nil, nil)

	for _, v := range page.Data {
		account := v.(*version1.AccountV1)
		c.Client.DeleteAccountById("", account.Id)
	}
}

func (c *AccountsClientFixtureV1) TestCrudOperations(t *testing.T) {
	c.clear()
	defer c.clear()

	// Create one account
	account, err := c.Client.CreateAccount("", ACCOUNT1)
	assert.Nil(t, err)

	assert.NotNil(t, account)
	assert.Equal(t, account.Name, ACCOUNT1.Name)
	assert.Equal(t, account.Login, ACCOUNT1.Login)

	account1 := account

	// Create another account
	account, err = c.Client.CreateAccount("", ACCOUNT2)
	assert.Nil(t, err)

	assert.NotNil(t, account)
	assert.Equal(t, account.Name, ACCOUNT2.Name)
	assert.Equal(t, account.Login, ACCOUNT2.Login)

	//account2 := account

	// Get all accounts
	page, err1 := c.Client.GetAccounts("", nil, nil)
	assert.Nil(t, err1)

	assert.NotNil(t, page)
	assert.True(t, len(page.Data) >= 2)

	// Get account by login
	account, err = c.Client.GetAccountByIdOrLogin("", ACCOUNT1.Login)
	assert.Nil(t, err)

	assert.NotNil(t, account)

	// Update the account
	account1.Name = "Updated Account 1"
	account, err = c.Client.UpdateAccount("", account1)
	assert.Nil(t, err)

	assert.NotNil(t, account)
	assert.Equal(t, account.Name, "Updated Account 1")
	assert.Equal(t, account.Login, account1.Login)

	account1 = account

	// Delete account
	account, err = c.Client.DeleteAccountById("", account1.Id)
	assert.Nil(t, err)

	// Try to get deleted account
	account, err = c.Client.GetAccountById("", account1.Id)
	assert.Nil(t, err)

	assert.NotNil(t, account)
	assert.True(t, account.Deleted)
	//assert.Nil(t, account)
}
