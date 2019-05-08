/*
 * Copyright (c)  2019. The Inn4Science Team
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

	XXX_Unmarshal(b []byte) error
	XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
	XXX_Merge(src proto.Message)
	XXX_Size() int
	XXX_DiscardUnknown()
}

type Message struct {
	schema      Schema
	author      crypto.PublicKey
	class       uint8
	messageType uint8
}

func newTxMessage(schema Schema, author crypto.PublicKey) Message {
	return Message{schema: schema, author: author, class: TransactionClass, messageType: TransactionType}
}

func newPrecomitMessage(schema Schema, author crypto.PublicKey) Message {
	return Message{schema: schema, author: author, class: PreCommitClass, messageType: PreCommitType}
}

type ServiceTx struct {
	Message
	ServiceID uint16           `json:"service_id"`
	MessageID uint16           `json:"message_id"`
	Signature crypto.Signature `json:"signature"`
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
	var err error

	sidBytes := make([]byte, 2)
	binary.LittleEndian.PutUint16(sidBytes, tx.ServiceID)
	midBytes := make([]byte, 2)
	binary.LittleEndian.PutUint16(midBytes, tx.MessageID)

	buf := make([]byte, 0)
	buf = append(tx.author.Data, tx.class)
	buf = append(buf, tx.class)
	buf = append(buf, tx.class)
	buf = append(buf, tx.messageType)
	buf = append(buf, sidBytes...)
	buf = append(buf, midBytes...)
	buf, err = tx.schema.XXX_Marshal(buf, true)
	if err != nil {
		return nil, err
	}

	if tx.Signature != nil {
		buf = append(buf, tx.Signature...)
	}

	return tx.schema.XXX_Marshal(buf, true)
}

// Sign serialized `ServiceTx` with passed key.
func (tx *ServiceTx) Sign(key crypto.SecretKey) (crypto.Signature, error) {
	data, err := tx.Serialize()
	if err != nil {
		return nil, err
	}

	return key.Sign(data), nil
}

// Hash creates SHA256 of `ServiceTx`.
func (tx ServiceTx) Hash() (crypto.Hash, error) {
	if tx.signedTx != nil {
		return crypto.Hash{}.FromData(tx.signedTx), nil
	}

	data, err := tx.Serialize()
	if err != nil {
		return crypto.Hash{}, err
	}

	return crypto.Hash{}.FromData(data), nil
}

// IntoSignedTx signs serialized `ServiceTx` with passed key, attach signature and encode to hex.
func (tx *ServiceTx) IntoSignedTx(key crypto.SecretKey) (string, error) {
	data, err := tx.Serialize()
	if err != nil {
		return "", err
	}

	data = append(data, key.Sign(data)...)
	tx.signedTx = data
	return hex.EncodeToString(data), nil
}
