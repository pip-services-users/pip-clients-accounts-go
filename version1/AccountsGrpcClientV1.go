package version1

import (
	"github.com/pip-services-users/pip-clients-accounts-go/protos"
	"github.com/pip-services3-go/pip-services3-commons-go/data"
)

type AccountGrpcClientV1 struct {
	GrpcClient
}

func NewAccountGrpcClientV1() *AccountGrpcClientV1 {
	return &AccountGrpcClientV1{
		GrpcClient: *NewGrpcClient("accounts.Accounts"),
	}
}

func (c *AccountGrpcClientV1) GetAccounts(correlationId string, filter *data.FilterParams,
	paging *data.PagingParams) (result *data.DataPage, err error) {
	req := &protos.AccountPageRequest{
		CorrelationId: correlationId,
	}
	if filter != nil {
		req.Filter = filter.Value()
	}
	if paging != nil {
		req.Paging = &protos.PagingParams{
			Skip:  paging.GetSkip(0),
			Take:  (int32)(paging.GetTake(100)),
			Total: paging.Total,
		}
	}

	reply := new(protos.AccountPageReply)
	err = c.Call("get_accounts", "", req, reply)
	if err != nil {
		return nil, err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return nil, err
	}

	result = toAccountPage(reply.Page)

	return result, nil
}

func (c *AccountGrpcClientV1) GetAccountById(correlationId string, id string) (result *AccountV1, err error) {
	req := &protos.AccountIdRequest{
		CorrelationId: correlationId,
		AccountId:     id,
	}

	reply := new(protos.AccountObjectReply)
	err = c.Call("get_account_by_id", "", req, reply)
	if err != nil {
		return nil, err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return nil, err
	}

	result = toAccount(reply.Account)

	return result, nil
}

func (c *AccountGrpcClientV1) GetAccountByLogin(correlationId string, login string) (result *AccountV1, err error) {
	req := &protos.AccountLoginRequest{
		CorrelationId: correlationId,
		Login:         login,
	}

	reply := new(protos.AccountObjectReply)
	err = c.Call("get_account_by_login", "", req, reply)
	if err != nil {
		return nil, err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return nil, err
	}

	result = toAccount(reply.Account)

	return result, nil
}

func (c *AccountGrpcClientV1) GetAccountByIdOrLogin(correlationId string, idOrLogin string) (result *AccountV1, err error) {
	req := &protos.AccountLoginRequest{
		CorrelationId: correlationId,
		Login:         idOrLogin,
	}

	reply := new(protos.AccountObjectReply)
	err = c.Call("get_account_by_id_or_login", "", req, reply)
	if err != nil {
		return nil, err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return nil, err
	}

	result = toAccount(reply.Account)

	return result, nil
}

func (c *AccountGrpcClientV1) CreateAccount(correlationId string, account *AccountV1) (result *AccountV1, err error) {
	req := &protos.AccountObjectRequest{
		CorrelationId: correlationId,
		Account:       fromAccount(account),
	}

	reply := new(protos.AccountObjectReply)
	err = c.Call("create_account", "", req, reply)
	if err != nil {
		return nil, err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return nil, err
	}

	result = toAccount(reply.Account)

	return result, nil
}

func (c *AccountGrpcClientV1) UpdateAccount(correlationId string, account *AccountV1) (result *AccountV1, err error) {
	req := &protos.AccountObjectRequest{
		CorrelationId: correlationId,
		Account:       fromAccount(account),
	}

	reply := new(protos.AccountObjectReply)
	err = c.Call("update_account", "", req, reply)
	if err != nil {
		return nil, err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return nil, err
	}

	result = toAccount(reply.Account)

	return result, nil
}

func (c *AccountGrpcClientV1) DeleteAccountById(correlationId string, id string) (result *AccountV1, err error) {
	req := &protos.AccountIdRequest{
		CorrelationId: correlationId,
		AccountId:     id,
	}

	reply := new(protos.AccountObjectReply)
	err = c.Call("delete_account_by_id", "", req, reply)
	if err != nil {
		return nil, err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return nil, err
	}

	result = toAccount(reply.Account)

	return result, nil
}
