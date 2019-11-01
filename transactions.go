/*
 * Copyright (c) 2018 - 2019. The Inn4Science Team
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package exonum

import (
	"encoding/binary"
	"encoding/hex"
	"errors"

	"github.com/golang/protobuf/proto"
	"github.com/inn4science/exonum-go/crypto"
)

const (
	TransactionClass = 0
	TransactionType  = 0
	PreCommitClass   = 1
	PreCommitType    = 0
)

type Schema interface {
	Reset()
	String() string
	ProtoMessage()
	Descriptor() ([]byte, []int)
}

type Message struct {
	Schema      Schema
	Author      crypto.PublicKey
	Class       uint8
	MessageType uint8
}

func newTxMessage(schema Schema, author crypto.PublicKey) Message {
	return Message{Schema: schema, Author: author, Class: TransactionClass, MessageType: TransactionType}
}

func newPrecomitMessage(schema Schema, author crypto.PublicKey) Message {
	return Message{Schema: schema, Author: author, Class: PreCommitClass, MessageType: PreCommitType}
}

type ServiceTx struct {
	Message
	ServiceID uint16           `json:"service_id"`
	MessageID uint16           `json:"message_id"`
	Signature crypto.Signature `json:"signature"`
	signed    bool
	signedTx  []byte
}

func (ServiceTx) New(schema Schema, author crypto.PublicKey, serviceID uint16, messageID uint16) ServiceTx {
	return ServiceTx{
		Message:   newTxMessage(schema, author),
		ServiceID: serviceID,
		MessageID: messageID,
	}
}

func (tx *ServiceTx) Serialize() ([]byte, error) {
	sidBytes := make([]byte, 2)
	binary.LittleEndian.PutUint16(sidBytes, tx.ServiceID)
	midBytes := make([]byte, 2)
	binary.LittleEndian.PutUint16(midBytes, tx.MessageID)

	buf := make([]byte, 0)
	buf = append(tx.Author.Data, tx.Class)
	buf = append(buf, tx.MessageType)
	buf = append(buf, sidBytes...)
	buf = append(buf, midBytes...)
	bytes, err := proto.Marshal(tx.Schema)
	if err != nil {
		return nil, err
	}
	buf = append(buf, bytes...)

	if tx.Signature != nil {
		buf = append(buf, tx.Signature...)
	}

	return buf, nil
}

func (ServiceTx) DecodeSignedTx(rawTx string, schema Schema) (ServiceTx, error) {
	return ServiceTx{}.DecodeSignedWSchemaProvider(rawTx, func(uint16, uint16) (schema Schema, e error) {
		return schema, nil
	})
}

// DecodeSignedWSchemaProvider ...
// SchemaProvider should return Schema implementation based on serviceID and messageID.
func (ServiceTx) DecodeSignedWSchemaProvider(rawTx string, provider func(uint16, uint16) (Schema, error)) (ServiceTx, error) {
	txBytes, err := hex.DecodeString(rawTx)
	if err != nil {
		return ServiceTx{}, err
	}

	if len(txBytes) <= 102 {
		return ServiceTx{}, errors.New("raw transaction is not valid")
	}

	signatureByte := txBytes[len(txBytes)-64:]
	data := txBytes[:len(txBytes)-64]

	authorPk, err := crypto.PublicKey{}.FromString(hex.EncodeToString(data[:32]))
	if err != nil {
		return ServiceTx{}, err
	}

	class := data[32:33]
	messageType := data[33:34]

	serviceID := binary.LittleEndian.Uint16(data[34:36])
	messageID := binary.LittleEndian.Uint16(data[36:38])

	schema, err := provider(serviceID, messageID)
	if err != nil {
		return ServiceTx{}, err
	}

	err = proto.Unmarshal(data[38:], schema)
	if err != nil {
		return ServiceTx{}, err
	}

	message := Message{
		Schema:      schema,
		Author:      authorPk,
		Class:       uint8(class[0]),
		MessageType: uint8(messageType[0]),
	}

	signature, err := crypto.Signature{}.FromString(hex.EncodeToString(signatureByte))
	if err != nil {
		return ServiceTx{}, err
	}

	return ServiceTx{
		Message:   message,
		ServiceID: serviceID,
		MessageID: messageID,
		Signature: signature,
	}, nil
}

// Sign serialized `ServiceTx` with passed key.
func (tx *ServiceTx) Sign(key crypto.SecretKey) (crypto.Signature, error) {
	tx.signed = false
	tx.Signature = nil

	data, err := tx.Serialize()
	if err != nil {
		return nil, err
	}

	signature := key.Sign(data)
	tx.signedTx = append(data, signature...)
	tx.signed = true
	return signature, nil
}

// Hash creates SHA256 of `ServiceTx`.
func (tx ServiceTx) Hash() (crypto.Hash, error) {
	if tx.signedTx != nil {
		return crypto.Hash{}.FromData(tx.signedTx), nil
	}

	if tx.Signature == nil {
		return crypto.Hash{}, errors.New("transaction not signed")
	}

	data, err := tx.Serialize()
	if err != nil {
		return crypto.Hash{}, err
	}

	return crypto.Hash{}.FromData(data), nil
}

// IntoSignedTx signs serialized `ServiceTx` with passed key, attach signature and encode to hex.
func (tx *ServiceTx) IntoSignedTx(key crypto.SecretKey) (string, error) {
	_, err := tx.Sign(key)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(tx.signedTx), nil
}
