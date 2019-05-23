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
	"encoding/json"
	"time"

	"github.com/inn4science/exonum-go/crypto"
)

type Block struct {
	ProposerID int    `json:"proposer_id"`
	Height     int64  `json:"height"`
	TxCount    int64  `json:"tx_count"`
	PrevHash   string `json:"prev_hash"`
	TxHash     string `json:"tx_hash"`
	StateHash  string `json:"state_hash"`
}

type BlockResponse struct {
	Block      Block         `json:"block"`
	PreCommits []string      `json:"precommits"`
	Txs        []interface{} `json:"txs"`
	Time       time.Time     `json:"time"`
}

type BlocksResponse struct {
	Range struct {
		Start int `json:"start"`
		End   int `json:"end"`
	} `json:"range"`
	Blocks []Block     `json:"blocks"`
	Times  []time.Time `json:"times"`
}

type FullTx struct {
	Type    string `json:"type"`
	Content struct {
		Debug   json.RawMessage `json:"debug"`
		Message string          `json:"message"`
	} `json:"content"`
	Location struct {
		BlockHeight     int `json:"block_height"`
		PositionInBlock int `json:"position_in_block"`
	} `json:"location"`
	LocationProof struct {
		Val string `json:"val"`
	} `json:"location_proof"`
	Status struct {
		Type        string `json:"type"`
		Code        int    `json:"code"`
		Description int    `json:"description"`
	} `json:"status"`
}

type TxResult struct {
	TxHash crypto.Hash `json:"tx_hash"`
}
