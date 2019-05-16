package exonum

import (
	"testing"

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
