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
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"

	"github.com/inn4science/exonum-go/types"
)

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
func (key *Hash) MarshalJSON() ([]byte, error) {
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

// Decode `Hash` from hex string.
func (key *Hash) Decode(str string) (err error) {
	key.Data, err = hex.DecodeString(str)
	return err
}

// Encode `Hash` into hex string.
func (key *Hash) Encode() string {
	return hex.EncodeToString(key.Data)
}
