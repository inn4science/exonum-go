package exonum

import (
	"encoding/hex"
	"testing"

	"github.com/inn4science/exonum-go/crypto"
	"github.com/inn4science/exonum-go/internal/testTypes"
	"github.com/inn4science/exonum-go/types"
	"github.com/stretchr/testify/assert"
)

const ServiceId = 4269

const (
	//CreateAccount::SYSTEM
	SYSTEM_UID     = "bbc5e661e106c6dcd8dc6dd186454c2fcba3c710fb4d8e71a60c93eaf077f073"
	SYSTEM_TX      = "671067a235785e9150fddac0f2f5644b298b992d76c0eae213cd2a8e7894af5d0000ad1001000a0608011001180112220a20bbc5e661e106c6dcd8dc6dd186454c2fcba3c710fb4d8e71a60c93eaf077f0731801e4a046d0e042656497dc1db67d406940cb714f7c5196f5e0e62056db7303922542e5ae1926472f625a8185a2d1fdf11493b6e362b69ce317ec36c036dbc94906"
	SYSTEM_TX_HASH = "6742b95779e92fba07e02c1aebdafb09d1e7521a0180fed062eab53f8aeef1dd"
	//CreateAccount::GENERAL
	GENERAL_UID     = "2bd806c97f0e00af1a1fc3328fa763a9269723c8db8fac4f93af71db186d6e90"
	GENERAL_TX      = "671067a235785e9150fddac0f2f5644b298b992d76c0eae213cd2a8e7894af5d0000ad1001000a0608011001180112220a202bd806c97f0e00af1a1fc3328fa763a9269723c8db8fac4f93af71db186d6e9018020a445c3785679233f0eda16b9bef09d3583d593778dcaad8e7392b1b98ba8d657a17d4546698511754a065b3fabd0f95053d28a06704f66889eb4b9fbd4ad90a"
	GENERAL_TX_HASH = "9235585db4893b4aec761601db26fe124612fd4af634f99404948dfbe28dbf63"
	//CreateAccount::MERCHANT
	MERCHANT        = "81b637d8fcd2c6da6359e6963113a1170de795e4b725b84d1e0b4cfd9ec58ce9"
	MERCHANT_TX     = "671067a235785e9150fddac0f2f5644b298b992d76c0eae213cd2a8e7894af5d0000ad1001000a0608011001180112220a2081b637d8fcd2c6da6359e6963113a1170de795e4b725b84d1e0b4cfd9ec58ce91803a30034003003bb1c7005e08c2c1c3c3a5dab36976dad1b62b6b1a09a058fb0b2e2197d6194c18eace2df450b0a1b2702e9a9d66b899ddf4ef8c72fedfeb15f06"
	MERCHANT_TXHASH = "d42a39415fd80333edfa23d4b2290f3e797cead3c8b5228d3d6b6a33cbf4a29e"
)

func TestServiceTx_IntoSignedTx(t *testing.T) {
	var systemKp, err = crypto.SecretKey{}.FromString("b3bae303e3c13d33305d3ca6c1f55d76b80bc517d9f131dc3bc05fc584a5441d671067a235785e9150fddac0f2f5644b298b992d76c0eae213cd2a8e7894af5d")
	assert.NoError(t, err)

	accountUidData, err := hex.DecodeString(SYSTEM_UID)
	assert.NoError(t, err)

	schema := &testTypes.CreateAccount{
		Meta:        &testTypes.TxMeta{ProtocolVer: 1, TxType: testTypes.TransactionType_TxCreateAccount, TxVer: 1},
		AccountUid:  &types.Hash{Data: accountUidData},
		AccountType: testTypes.AccountType_SYSTEM,
	}

	exonumTx := ServiceTx{}.New(schema, systemKp.GetPublic(), ServiceId, uint16(testTypes.TransactionType_TxCreateAccount))
	signedTx, err := exonumTx.IntoSignedTx(systemKp)
	assert.NoError(t, err)

	h, err := exonumTx.Hash()
	assert.NoError(t, err)
	println("TX:", signedTx)
	println("HASH:", hex.EncodeToString(h.Data))

	assert.Equal(t, SYSTEM_TX, signedTx)
	assert.Equal(t, SYSTEM_TX_HASH, hex.EncodeToString(h.Data))
}

func TestServiceTx_DecodeSignedTx(t *testing.T) {
	schema := &testTypes.CreateAccount{}
	exonumTx, err := ServiceTx{}.DecodeSignedTx(SYSTEM_TX, schema)
	assert.NoError(t, err)

	// expected data
	systemKp, err := crypto.SecretKey{}.FromString("b3bae303e3c13d33305d3ca6c1f55d76b80bc517d9f131dc3bc05fc584a5441d671067a235785e9150fddac0f2f5644b298b992d76c0eae213cd2a8e7894af5d")
	assert.NoError(t, err)
	systemPk := systemKp.GetPublic()

	accountUIDData, err := hex.DecodeString(SYSTEM_UID)
	assert.NoError(t, err)

	expectedSchema := &testTypes.CreateAccount{
		Meta:        &testTypes.TxMeta{ProtocolVer: 1, TxType: testTypes.TransactionType_TxCreateAccount, TxVer: 1},
		AccountUid:  &types.Hash{Data: accountUIDData},
		AccountType: testTypes.AccountType_SYSTEM,
	}

	assert.Equal(t, ServiceId, int(exonumTx.ServiceID))
	assert.Equal(t, int(testTypes.TransactionType_TxCreateAccount), int(exonumTx.MessageID))

	assert.Equal(t, expectedSchema, exonumTx.Message.schema)
	assert.Equal(t, systemPk.String(), exonumTx.Message.author.String())
	assert.Equal(t, TransactionClass, int(exonumTx.Message.class))
	assert.Equal(t, TransactionClass, int(exonumTx.Message.messageType))

	_, err = ServiceTx{}.DecodeSignedTx(SYSTEM_TX[0:12], schema)
	assert.Error(t, err)

}
