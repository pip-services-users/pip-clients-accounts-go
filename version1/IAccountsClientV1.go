package version1

import (
	"github.com/pip-services3-go/pip-services3-commons-go/data"
)

type IAccountsClientV1 interface {
	GetAccounts(correlationId string, filter *data.FilterParams,
		paging *data.PagingParams) (result *data.DataPage, err error)

	GetAccountById(correlationId string, id string) (result *AccountV1, err error)

	GetAccountByLogin(correlationId string, login string) (result *AccountV1, err error)

	GetAccountByIdOrLogin(correlationId string, idOrLogin string) (result *AccountV1, err error)

	CreateAccount(correlationId string, account *AccountV1) (result *AccountV1, err error)

	UpdateAccount(correlationId string, account *AccountV1) (result *AccountV1, err error)

	DeleteAccountById(correlationId string, id string) (result *AccountV1, err error)
}
