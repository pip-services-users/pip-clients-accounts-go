package test_version1

import (
	"os"
	"testing"

	"github.com/pip-services-users/pip-clients-accounts-go/version1"
	"github.com/pip-services3-go/pip-services3-commons-go/config"
)

type accountsHttpCommandableClientV1Test struct {
	client  *version1.AccountsHttpCommandableClientV1
	fixture *AccountsClientFixtureV1
}

func newAccountsHttpCommandableClientV1Test() *accountsHttpCommandableClientV1Test {
	return &accountsHttpCommandableClientV1Test{}
}

func (c *accountsHttpCommandableClientV1Test) setup(t *testing.T) *AccountsClientFixtureV1 {
	var HTTP_HOST = os.Getenv("HTTP_HOST")
	if HTTP_HOST == "" {
		HTTP_HOST = "localhost"
	}
	var HTTP_PORT = os.Getenv("HTTP_PORT")
	if HTTP_PORT == "" {
		HTTP_PORT = "8080"
	}

	var httpConfig = config.NewConfigParamsFromTuples(
		"connection.protocol", "http",
		"connection.host", HTTP_HOST,
		"connection.port", HTTP_PORT,
	)

	c.client = version1.NewAccountsHttpCommandableClientV1()
	c.client.Configure(httpConfig)
	c.client.Open("")

	c.fixture = NewAccountsClientFixtureV1(c.client)

	return c.fixture
}

func (c *accountsHttpCommandableClientV1Test) teardown(t *testing.T) {
	c.client.Close("")
}

func TestHttpCrudOperations(t *testing.T) {
	c := newAccountsHttpCommandableClientV1Test()
	fixture := c.setup(t)
	defer c.teardown(t)

	fixture.TestCrudOperations(t)
}
