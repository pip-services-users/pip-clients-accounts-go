package build

import (
	clients1 "github.com/pip-services-users/pip-clients-accounts-go/version1"
	cref "github.com/pip-services3-go/pip-services3-commons-go/refer"
	cbuild "github.com/pip-services3-go/pip-services3-components-go/build"
)

type AccountsClientFactory struct {
	cbuild.Factory
}

func NewAccountsClientFactory() *AccountsClientFactory {
	c := &AccountsClientFactory{
		Factory: *cbuild.NewFactory(),
	}

	// nullClientDescriptor := cref.NewDescriptor("pip-services-sasswords", "client", "null", "*", "1.0")
	// directClientDescriptor := cref.NewDescriptor("pip-services-sasswords", "client", "direct", "*", "1.0")
	memoryClientDescriptor := cref.NewDescriptor("pip-services-accounts", "client", "memory", "*", "1.0")
	cmdHttpClientDescriptor := cref.NewDescriptor("pip-services-accounts", "client", "commandable-http", "*", "1.0")
	grpcClientDescriptor := cref.NewDescriptor("pip-services-accounts", "client", "grpc", "*", "1.0")

	// c.RegisterType(nullClientDescriptor, clients1.NewAccountsNullClientV1)
	// c.RegisterType(directClientDescriptor, clients1.NewAccountsDirectClientV1)
	c.RegisterType(memoryClientDescriptor, clients1.NewEmptyAccountsMemoryClientV1)
	c.RegisterType(cmdHttpClientDescriptor, clients1.NewAccountsHttpCommandableClientV1)
	c.RegisterType(grpcClientDescriptor, clients1.NewAccountGrpcClientV1)

	return c
}
