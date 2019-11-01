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

package crypto

import (
	"crypto/sha256"
	"database/sql/driver"
	"encoding/hex"
	"encoding/json"
	"errors"

	"github.com/inn4science/exonum-go/types"
)

// Hash is a type-wrapper for exonum.Hash; hex-encoded result of the SHA256.
// Wrappers implements
type Hash struct {
	types.Hash
}

// FromData returns the SHA256 checksum of the data as `Hash`
func (Hash) FromData(data []byte) Hash {
	key := Hash{}
	h := sha256.Sum256(data)
	key.Data = h[:]
	return key
}

// FromString returns new `Hash` decoded from hex string.
func (Hash) FromString(raw string) (Hash, error) {
	key := Hash{}
	err := key.Decode(raw)
	return key, err
}

// MarshalJSON convert `Hash` into hex string and than into json.
func (key Hash) MarshalJSON() ([]byte, error) {
	return json.Marshal(key.Encode())
}

// UnmarshalJSON unmarshal `Hash` from json as a hex string and `Decode`.
func (key *Hash) UnmarshalJSON(data []byte) error {
	var s string
	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}
	return key.Decode(s)
}

// Value is generated so Hash satisfies db row driver.Scanner.
func (key *Hash) Scan(src interface{}) error {
	switch v := src.(type) {
	case string:
		return key.Decode(v)
	}
	return errors.New("hash: invalid type")
}

// Value is generated so Hash satisfies db row driver.Valuer.
func (key Hash) Value() (driver.Value, error) {
	return key.Encode(), nil
}

// MarshalYAML convert `Hash` into hex string and than into yaml.
func (key *Hash) MarshalYAML() (interface{}, error) {
	return key.Encode(), nil
}

// UnmarshalYAML unmarshal `Hash` from yaml as a hex string and `Decode`.
func (key *Hash) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var stringValue string
	err := unmarshal(&stringValue)
	if err != nil {
		return err
	}
	return key.Decode(stringValue)
}

// MarshalText convert `Hash` into hex string and than into textual representation.
func (key *Hash) MarshalText() (text []byte, err error) {
	return []byte(key.Encode()), nil
}

// UnmarshalText unmarshal `Hash` from textual representation as a hex string and `Decode`.
func (key *Hash) UnmarshalText(text []byte) error {
	return key.Decode(string(text))
}

// Decode `Hash` from hex string.
func (key *Hash) Decode(str string) (err error) {
	key.Data, err = hex.DecodeString(str)
	return err
}

// Encode `Hash` into hex string.
func (key *Hash) Encode() string {
	return hex.EncodeToString(key.Data)
}

// Encode `Hash` into hex string.
func (key *Hash) String() string {
	return key.Encode()
}
