package exonum

import (
	"encoding/hex"
	"testing"

	"github.com/inn4science/exonum-go/crypto"
	"github.com/inn4science/exonum-go/internal/testTypes"
	"github.com/stretchr/testify/assert"
)

const ServiceId = 4269

const (
	SystemUid    = "bbc5e661e106c6dcd8dc6dd186454c2fcba3c710fb4d8e71a60c93eaf077f073"
	SystemTx     = "671067a235785e9150fddac0f2f5644b298b992d76c0eae213cd2a8e7894af5d0000ad1001000a4062626335653636316531303663366463643864633664643138363435346332666362613363373130666234643865373161363063393365616630373766303733124032626438303663393766306530306166316131666333333238666137363361393236393732336338646238666163346639336166373164623138366436653930182a20a7db87a62e045fa60f79c1dbd7ff7b152a8633d91c119f538177433a6497b82c25847f0ab686a7c50cea05659324475f7dfd369b0febcc773f94c9c8ee8320ab99ba66fc0a"
	SystemTxHash = "2f63444bb2efc3ba460bd2c12feceef61fd1d1ca372e18074798f02d5ac63153"
	GeneralUid   = "2bd806c97f0e00af1a1fc3328fa763a9269723c8db8fac4f93af71db186d6e90"
)

func TestServiceTx_IntoSignedTx(t *testing.T) {
	var systemKp, err = crypto.SecretKey{}.FromString("b3bae303e3c13d33305d3ca6c1f55d76b80bc517d9f131dc3bc05fc584a5441d671067a235785e9150fddac0f2f5644b298b992d76c0eae213cd2a8e7894af5d")
	assert.NoError(t, err)

	schema := &testTypes.Transfer{
		From:   SystemUid,
		To:     GeneralUid,
		Amount: 42,
		Seed:   12427849127,
	}

	exonumTx := ServiceTx{}.New(schema, systemKp.GetPublic(), ServiceId, uint16(1))
	signedTx, err := exonumTx.IntoSignedTx(systemKp)
	assert.NoError(t, err)

	h, err := exonumTx.Hash()
	assert.NoError(t, err)
	println("TX:", signedTx)
	println("HASH:", hex.EncodeToString(h.Data))

	assert.Equal(t, SystemTx, signedTx)
	assert.Equal(t, SystemTxHash, hex.EncodeToString(h.Data))
}

func TestServiceTx_DecodeSignedTx(t *testing.T) {
	schema := &testTypes.Transfer{}
	exonumTx, err := ServiceTx{}.DecodeSignedTx(SystemTx, schema)
	assert.NoError(t, err)

	// expected data
	systemKp, err := crypto.SecretKey{}.FromString("b3bae303e3c13d33305d3ca6c1f55d76b80bc517d9f131dc3bc05fc584a5441d671067a235785e9150fddac0f2f5644b298b992d76c0eae213cd2a8e7894af5d")
	assert.NoError(t, err)
	systemPk := systemKp.GetPublic()

	expectedSchema := &testTypes.Transfer{
		From:   SystemUid,
		To:     GeneralUid,
		Amount: 42,
		Seed:   12427849127,
	}
	assert.Equal(t, uint16(ServiceId), exonumTx.ServiceID)
	assert.Equal(t, uint16(1), exonumTx.MessageID)

	assert.Equal(t, expectedSchema, exonumTx.Message.schema)
	assert.Equal(t, systemPk.String(), exonumTx.Message.author.String())
	assert.Equal(t, TransactionClass, int(exonumTx.Message.class))
	assert.Equal(t, TransactionClass, int(exonumTx.Message.messageType))

	_, err = ServiceTx{}.DecodeSignedTx(SystemTx[0:12], schema)
	assert.Error(t, err)

}
