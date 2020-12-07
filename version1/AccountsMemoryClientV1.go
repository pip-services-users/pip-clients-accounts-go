package version1

import (
	"strings"

	"github.com/pip-services3-go/pip-services3-commons-go/data"
	cdata "github.com/pip-services3-go/pip-services3-commons-go/data"
	cerr "github.com/pip-services3-go/pip-services3-commons-go/errors"
	mdata "github.com/pip-services3-go/pip-services3-data-go/persistence"
)

type AccountsMemoryClientV1 struct {
	accounts    []AccountV1
	maxPageSize int
}

func NewAccountsMemoryClientV1(accounts []AccountV1) *AccountsMemoryClientV1 {

	c := AccountsMemoryClientV1{
		accounts:    make([]AccountV1, 0),
		maxPageSize: 100,
	}
	if accounts != nil {
		c.accounts = append(c.accounts, accounts...)
	}
	return &c
}

func (c *AccountsMemoryClientV1) GetAccounts(correlationId string, filter *data.FilterParams,
	paging *data.PagingParams) (result *data.DataPage, err error) {

	var total int64 = (int64)(len(c.accounts))
	items := make([]interface{}, 0)
	for _, v := range c.accounts {
		item := v
		items = append(items, &item)
	}
	return cdata.NewDataPage(&total, items), nil
}

func (c *AccountsMemoryClientV1) GetAccountById(correlationId string, id string) (result *AccountV1, err error) {

	for _, v := range c.accounts {
		if v.Id == id {
			buf := v
			result = &buf
			break
		}
	}
	return result, nil
}

func (c *AccountsMemoryClientV1) GetAccountByLogin(correlationId string, login string) (result *AccountV1, err error) {
	for _, v := range c.accounts {
		if v.Login == login {
			buf := v
			result = &buf
			break
		}
	}
	return result, nil
}

func (c *AccountsMemoryClientV1) GetAccountByIdOrLogin(correlationId string, idOrLogin string) (result *AccountV1, err error) {
	for _, v := range c.accounts {
		if v.Id == idOrLogin || v.Login == idOrLogin {
			buf := v
			result = &buf
			break
		}
	}
	return result, nil
}

func (c *AccountsMemoryClientV1) CreateAccount(correlationId string, account *AccountV1) (result *AccountV1, err error) {
	if account == nil {
		return nil, nil
	}

	var index = -1
	for i, v := range c.accounts {
		if v.Id == account.Id {
			index = i
			break
		}
	}

	if index >= 0 {
		err := cerr.NewBadRequestError(correlationId, "ACCOUNT_ALREADY_EXIST", "Account "+account.Login+" already exists")
		return nil, err
	}

	newItem := mdata.CloneObject(account)
	item, _ := newItem.(AccountV1)
	mdata.GenerateObjectId(&newItem)

	c.accounts = append(c.accounts, item)

	return &item, nil
}

func (c *AccountsMemoryClientV1) UpdateAccount(correlationId string, account *AccountV1) (result *AccountV1, err error) {

	if account == nil {
		return nil, nil
	}

	var index = -1
	for i, v := range c.accounts {
		if v.Id == account.Id {
			index = i
			break
		}
	}

	if index < 0 {
		return nil, nil
	}

	newItem := mdata.CloneObject(account)
	item, _ := newItem.(AccountV1)
	c.accounts[index] = item
	return &item, nil

}

func (c *AccountsMemoryClientV1) DeleteAccountById(correlationId string, id string) (result *AccountV1, err error) {
	var index = -1
	for i, v := range c.accounts {
		if v.Id == id {
			index = i
			break
		}
	}

	if index < 0 {
		return nil, nil
	}
	c.accounts[index].Deleted = true
	var item = c.accounts[index]
	return &item, nil
}

func (c *AccountsMemoryClientV1) matchString(value string, search string) bool {
	if value == "" && search == "" {
		return true
	}
	if value == "" || search == "" {
		return false
	}
	return strings.Index(strings.ToLower(value), strings.ToLower(search)) >= 0
}

func (c *AccountsMemoryClientV1) matchSearch(item AccountV1, search string) bool {
	search = strings.ToLower(search)
	if c.matchString(item.Name, search) {
		return true
	}
	return false
}

func (c *AccountsMemoryClientV1) composeFilter(filter *cdata.FilterParams) func(item AccountV1) bool {
	if filter == nil {
		filter = cdata.NewEmptyFilterParams()
	}

	search := filter.GetAsString("search")
	id := filter.GetAsString("id")
	name := filter.GetAsString("name")
	login := filter.GetAsString("login")
	active := filter.GetAsNullableBoolean("active")
	fromCreateTime := filter.GetAsNullableDateTime("from_create_time")
	toCreateTime := filter.GetAsNullableDateTime("to_create_time")
	deleted := filter.GetAsBooleanWithDefault("deleted", false)

	return func(item AccountV1) bool {
		if search != "" && !c.matchSearch(item, search) {
			return false
		}
		if id != "" && id != item.Id {
			return false
		}
		if name != "" && name != item.Name {
			return false
		}
		if login != "" && login != item.Login {
			return false
		}
		if active != nil && *active != item.Active {
			return false
		}
		if fromCreateTime != nil && item.CreateTime.Nanosecond() >= fromCreateTime.Nanosecond() {
			return false
		}
		if toCreateTime != nil && item.CreateTime.Nanosecond() < toCreateTime.Nanosecond() {
			return false
		}
		if !deleted && item.Deleted {
			return false
		}
		return true
	}
}
