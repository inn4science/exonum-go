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
	"crypto/rand"
	"encoding/binary"

	"golang.org/x/crypto/ed25519"
)

type KeyPair struct {
	sk SecretKey
	pk PublicKey
}

func (KeyPair) New(key PublicKey, secretKey SecretKey) KeyPair {
	return KeyPair{
		sk: secretKey,
		pk: key,
	}
}

func (KeyPair) FromSecret(secretKey SecretKey) KeyPair {
	return KeyPair{
		sk: secretKey,
		pk: secretKey.GetPublic(),
	}
}

func (KeyPair) Random() KeyPair {
	public, secret, _ := ed25519.GenerateKey(rand.Reader)
	return KeyPair{
		sk: SecretKey{}.New(secret),
		pk: PublicKey{}.New(public),
	}
}

func (kp KeyPair) Sign(data []byte) Signature {
	return kp.sk.Sign(data)
}

func (kp KeyPair) Verify(data []byte, signature Signature) bool {
	return kp.pk.Verify(data, signature)
}

func RandomUint64() uint64 {
	buf := make([]byte, 8)
	_, _ = rand.Read(buf) // Always succeeds, no need to check error
	return binary.LittleEndian.Uint64(buf)
}
