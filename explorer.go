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
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/inn4science/exonum-go/crypto"
)

type ExplorerApi interface {
	New(baseURL URL) ExplorerApi
	SetURL(url URL) ExplorerApi
	SetHeader(header string, value string) ExplorerApi
	ExplorerPath(prefix string) *URL
	ServicePath(serviceName, prefix string) *URL
	GetBlocks(count uint32, latest uint64, skipEmptyBlocks bool, addTime bool) (*BlocksResponse, error)
	GetBlock(height uint64) (*Block, error)
	GetTx(hash crypto.Hash) (*FullTx, error)
	SubmitTx(signedTx string) (*TxResult, error)
	GetJSON(fullURL string, dest interface{}) error
}

const explorerPathPrefix = "/api/explorer/v1"

var servicePathTemplate = func(service string) string {
	return fmt.Sprintf("/api/services/%s/v1", service)
}

type explorerApi struct {
	baseURL URL
	client  http.Client
	header  http.Header
}

func NewExplorerApi(baseURL URL) ExplorerApi {
	return explorerApi{}.New(baseURL)
}

func (explorerApi) New(baseURL URL) ExplorerApi {
	return &explorerApi{
		baseURL: baseURL,
		client:  http.Client{Timeout: 15 * time.Second},
		header:  make(http.Header),
	}
}

func (api *explorerApi) SetURL(baseURL URL) ExplorerApi {
	api.baseURL = baseURL
	return api
}

func (api *explorerApi) SetHeader(h string, v string) ExplorerApi {
	api.header.Set(h, v)
	return api
}

func (api *explorerApi) ExplorerPath(prefix string) *URL {
	return api.baseURL.SetBasePath(explorerPathPrefix).SetPath(prefix)
}

func (api *explorerApi) ServicePath(serviceName, prefix string) *URL {
	return api.baseURL.SetBasePath(servicePathTemplate(serviceName)).SetPath(prefix)
}

func (api *explorerApi) GetBlocks(count uint32, latest uint64, skipEmptyBlocks bool, addTime bool) (*BlocksResponse, error) {
	val := make(url.Values)
	val.Set("count", fmt.Sprintf("%v", count))
	val.Set("latest", fmt.Sprintf("%v", latest))
	val.Set("skip_empty_blocks", fmt.Sprintf("%v", skipEmptyBlocks))
	val.Set("add_blocks_time", fmt.Sprintf("%v", addTime))
	reqURL := api.ExplorerPath("/blocks").WithQuery(val)

	result := new(BlocksResponse)
	err := api.GetJSON(reqURL, result)
	return result, err
}

func (api *explorerApi) GetBlock(height uint64) (*Block, error) {
	val := make(url.Values)
	val.Set("height", fmt.Sprintf("%v", height))
	reqURL := api.ExplorerPath("/block").WithQuery(val)

	result := new(Block)
	err := api.GetJSON(reqURL, result)
	return result, err
}

func (api *explorerApi) GetTx(hash crypto.Hash) (*FullTx, error) {
	val := make(url.Values)
	val.Set("hash", fmt.Sprintf("%v", hash))
	reqURL := api.ExplorerPath("/transactions").WithQuery(val)

	result := new(FullTx)
	err := api.GetJSON(reqURL, result)
	return result, err
}

func (api *explorerApi) SubmitTx(signedTx string) (*TxResult, error) {
	fullURL := api.ExplorerPath("/transactions").String()

	rawData, err := json.Marshal(map[string]string{
		"tx_body": signedTx,
	})
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fullURL, bytes.NewBuffer(rawData))
	if err != nil {
		return nil, err
	}

	req.Header = api.header
	req.Header.Set("Content-Type", "application/json")
	resp, err := api.client.Do(req)
	if err != nil {
		return nil, err
	}

	result := new(TxResult)
	err = json.NewDecoder(resp.Body).Decode(result)
	return result, err
}

func (api *explorerApi) GetJSON(fullURL string, dest interface{}) error {
	resp, err := api.client.Get(fullURL)
	if err != nil {
		return err
	}

	return json.NewDecoder(resp.Body).Decode(dest)
}
