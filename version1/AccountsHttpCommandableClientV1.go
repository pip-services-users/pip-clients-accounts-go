package version1

import (
	"reflect"

	"github.com/pip-services3-go/pip-services3-commons-go/data"
	cdata "github.com/pip-services3-go/pip-services3-commons-go/data"
	cclients "github.com/pip-services3-go/pip-services3-rpc-go/clients"
)

type AccountsHttpCommandableClientV1 struct {
	*cclients.CommandableHttpClient
	dataPageType  reflect.Type
	accountV1Type reflect.Type
}

func NewAccountsHttpCommandableClientV1() *AccountsHttpCommandableClientV1 {
	c := &AccountsHttpCommandableClientV1{
		CommandableHttpClient: cclients.NewCommandableHttpClient("v1/accounts"),
		dataPageType:          reflect.TypeOf(&data.DataPage{}),
		accountV1Type:         reflect.TypeOf(&AccountV1{}),
	}
	return c
}

func (c *AccountsHttpCommandableClientV1) GetAccounts(correlationId string, filter *data.FilterParams,
	paging *data.PagingParams) (result *data.DataPage, err error) {

	params := cdata.NewAnyValueMapFromTuples(
		"filter", filter,
		"paging", paging,
	)

	res, err := c.CallCommand(c.dataPageType, "get_accounts", correlationId, params)
	if err != nil {
		return nil, err
	}

	result, _ = res.(*cdata.DataPage)
	return result, nil
}

func (c *AccountsHttpCommandableClientV1) GetAccountById(correlationId string, id string) (result *AccountV1, err error) {

	params := cdata.NewAnyValueMapFromTuples(
		"account_id", id,
	)

	res, err := c.CallCommand(c.accountV1Type, "get_account_by_id", correlationId, params)
	if err != nil {
		return nil, err
	}

	result, _ = res.(*AccountV1)
	return result, nil
}

func (c *AccountsHttpCommandableClientV1) GetAccountByLogin(correlationId string, login string) (result *AccountV1, err error) {

	params := cdata.NewAnyValueMapFromTuples(
		"login", login,
	)

	res, err := c.CallCommand(c.accountV1Type, "get_account_by_login", correlationId, params)
	if err != nil {
		return nil, err
	}

	result, _ = res.(*AccountV1)
	return result, nil
}

func (c *AccountsHttpCommandableClientV1) GetAccountByIdOrLogin(correlationId string, idOrLogin string) (result *AccountV1, err error) {

	params := cdata.NewAnyValueMapFromTuples(
		"id_or_login", idOrLogin,
	)

	res, err := c.CallCommand(c.accountV1Type, "get_account_by_id_or_login", correlationId, params)
	if err != nil {
		return nil, err
	}

	result, _ = res.(*AccountV1)
	return result, nil
}

func (c *AccountsHttpCommandableClientV1) CreateAccount(correlationId string, account *AccountV1) (result *AccountV1, err error) {

	params := cdata.NewAnyValueMapFromTuples(
		"account", account,
	)

	res, err := c.CallCommand(c.accountV1Type, "create_account", correlationId, params)
	if err != nil {
		return nil, err
	}

	result, _ = res.(*AccountV1)
	return result, nil
}

func (c *AccountsHttpCommandableClientV1) UpdateAccount(correlationId string, account *AccountV1) (result *AccountV1, err error) {

	params := cdata.NewAnyValueMapFromTuples(
		"account", account,
	)

	res, err := c.CallCommand(c.accountV1Type, "update_account", correlationId, params)
	if err != nil {
		return nil, err
	}

	result, _ = res.(*AccountV1)
	return result, nil
}

func (c *AccountsHttpCommandableClientV1) DeleteAccountById(correlationId string, id string) (result *AccountV1, err error) {

	params := cdata.NewAnyValueMapFromTuples(
		"account_id", id,
	)

	res, err := c.CallCommand(c.accountV1Type, "delete_account_by_id", correlationId, params)
	if err != nil {
		return nil, err
	}

	result, _ = res.(*AccountV1)
	return result, nil
}
