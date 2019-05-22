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

package crypto

import (
	"database/sql/driver"
	"encoding/hex"
	"encoding/json"
	"errors"

	"github.com/inn4science/exonum-go/types"
	"golang.org/x/crypto/ed25519"
)

// SecretKey wrapper on top of `exonum.SecretKey` and  `ed25519.PrivateKey`
type SecretKey struct {
	types.SecretKey
}

// New returns new `SecretKey`.
func (SecretKey) New(raw ed25519.PrivateKey) SecretKey {
	key := SecretKey{}
	key.Data = raw
	return key
}

// FromString returns new `SecretKey` decoded from hex string.
func (SecretKey) FromString(raw string) (SecretKey, error) {
	key := SecretKey{}
	err := key.Decode(raw)
	return key, err
}

// MarshalJSON convert `SecretKey` into hex string and than into json.
func (key *SecretKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(key.Encode())
}

// UnmarshalJSON unmarshal `SecretKey` from json as a hex string and `Decode`.
func (key *SecretKey) UnmarshalJSON(data []byte) error {
	var s string
	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}
	return key.Decode(s)
}

// Value is generated so SecretKey satisfies db row driver.Scanner.
func (key *SecretKey) Scan(src interface{}) error {
	switch v := src.(type) {
	case string:
		return key.Decode(v)
	}
	return errors.New("SecretKey: invalid type")
}

// Value is generated so SecretKey satisfies db row driver.Valuer.
func (key SecretKey) Value() (driver.Value, error) {
	return key.Encode(), nil
}

// Decode `SecretKey` from hex string.
func (key *SecretKey) Decode(str string) (err error) {
	key.Data, err = hex.DecodeString(str)
	return err
}

// Encode `SecretKey` into hex string.
func (key *SecretKey) Encode() string {
	return hex.EncodeToString(key.Data)
}

// ToPrivate cast `key` to `ed25519.PrivateKey`.
func (key *SecretKey) ToPrivate() ed25519.PrivateKey {
	return ed25519.PrivateKey(key.Data)
}

// GetPublic cast `SecretKey` to  `PublicKey`
func (key *SecretKey) GetPublic() PublicKey {
	pk, _ := key.ToPrivate().Public().(ed25519.PublicKey)
	return PublicKey{}.New(pk)
}

// Sign creates `Signature` of passed data by this `SecretKey`
func (key *SecretKey) Sign(data []byte) Signature {
	return ed25519.Sign(key.ToPrivate(), data)
}

// Encode `SecretKey` into hex string.
func (key *SecretKey) String() string {
	return key.Encode()
}

// PublicKey wrapper on top of `exonum.PublicKey` and  `ed25519.PublicKey`
type PublicKey struct {
	types.PublicKey
}

// New returns new `PublicKey`.
func (PublicKey) New(raw ed25519.PublicKey) PublicKey {
	key := PublicKey{}
	key.Data = raw
	return key
}

// FromString returns new `PublicKey` decoded from hex string.
func (PublicKey) FromString(raw string) (PublicKey, error) {
	key := PublicKey{}
	err := key.Decode(raw)
	return key, err
}

// UnmarshalJSON unmarshal `PublicKey` from json as a hex string and `Decode`.
func (key *PublicKey) UnmarshalJSON(data []byte) error {
	var s string
	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}
	return key.Decode(s)
}

// MarshalJSON convert `PublicKey` into hex string and than into json.
func (key *PublicKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(key.Encode())
}

// Value is generated so PublicKey satisfies db row driver.Scanner.
func (key *PublicKey) Scan(src interface{}) error {
	switch v := src.(type) {
	case string:
		return key.Decode(v)
	}
	return errors.New("PublicKey: invalid type")
}

// Value is generated so PublicKey satisfies db row driver.Valuer.
func (key PublicKey) Value() (driver.Value, error) {
	return key.Encode(), nil
}

// Decode `PublicKey` from hex string.
func (key *PublicKey) Decode(str string) (err error) {
	key.Data, err = hex.DecodeString(str)
	return err
}

// Encode `PublicKey` into hex string.
func (key *PublicKey) Encode() string {
	return hex.EncodeToString(key.Data)
}

// ToPublic cast `key` to `ed25519.PublicKey`.
func (key *PublicKey) ToPublic() ed25519.PublicKey {
	return ed25519.PublicKey(key.Data)
}

// Verify check is data was signed by secret from this `PublicKey`
func (key *PublicKey) Verify(data []byte, signature Signature) bool {
	return ed25519.Verify(key.ToPublic(), data, signature)
}

// Encode `PublicKey` into hex string.
func (key *PublicKey) String() string {
	return key.Encode()
}

type Signature []byte

// FromString returns new `Signature` decoded from hex string.
func (Signature) FromString(raw string) (Signature, error) {
	key := Signature{}
	err := key.Decode(raw)
	return key, err
}

// MarshalJSON convert `Signature` into hex string and than into json.
func (key *Signature) MarshalJSON() ([]byte, error) {
	return json.Marshal(key.Encode())
}

// UnmarshalJSON unmarshal `Signature` from json as a hex string and `Decode`.
func (key *Signature) UnmarshalJSON(data []byte) error {
	var s string
	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}
	return key.Decode(s)
}

// Value is generated so Signature satisfies db row driver.Scanner.
func (key *Signature) Scan(src interface{}) error {
	switch v := src.(type) {
	case string:
		return key.Decode(v)
	}
	return errors.New("Signature: invalid type")
}

// Value is generated so Signature satisfies db row driver.Valuer.
func (key Signature) Value() (driver.Value, error) {
	return key.Encode(), nil
}

// Decode `Signature` from hex string.
func (key *Signature) Decode(str string) error {
	b, err := hex.DecodeString(str)
	*key = Signature(b)
	return err
}

// Encode `Signature` into hex string.
func (key *Signature) Encode() string {
	return hex.EncodeToString([]byte(*key))
}

// Encode `Signature` into hex string.
func (key *Signature) String() string {
	return key.Encode()
}
