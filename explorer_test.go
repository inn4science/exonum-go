package exonum

import (
	"testing"

	"github.com/inn4science/exonum-go/crypto"
	"github.com/stretchr/testify/assert"
)

const TestURL = "http://127.0.0.1:8201"

func TestExplorerApi_New(t *testing.T) {
	url, err := URL{}.New(TestURL)
	assert.NoError(t, err)

	explorerApi := NewExplorerApi(url)

	expectedUrl := TestURL + "/api/explorer/v1/transaction"
	assert.Equal(t, expectedUrl, explorerApi.ExplorerPath("transaction").URL.String())
}

func TestGetTx(t *testing.T) {
	url, err := URL{}.New(TestURL)
	assert.NoError(t, err)

	explorerApi := NewExplorerApi(url)

	hash, err := crypto.Hash{}.FromString("22e641b583c1c03e9645f8817b1552b046d91883a9e376dda3b08fb644a611f2")
	assert.NoError(t, err)

	fullTx, err := explorerApi.GetTx(hash)
	assert.NoError(t, err)

	assert.Equal(t, "success", fullTx.Status.Type)
}
