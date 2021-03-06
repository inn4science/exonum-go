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
	"bytes"
	"encoding/json"
	"errors"
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

	SystemPath(prefix string) *URL
	ExplorerPath(prefix string) *URL
	ServicePath(serviceName, prefix string) *URL

	Stats() (*Stats, ExplorerApiError)
	Services() (*ServiceList, ExplorerApiError)
	HealthCheck() (*HealthCheck, ExplorerApiError)

	GetBlocks(count uint32, latest uint64, skipEmptyBlocks bool, addTime bool) (*BlocksResponse, ExplorerApiError)
	GetBlock(height uint64) (*FullBlock, ExplorerApiError)
	LastBlock() (*FullBlock, ExplorerApiError)

	GetTx(hash crypto.Hash) (*FullTx, ExplorerApiError)
	SubmitTx(signedTx string) (*TxResult, ExplorerApiError)

	GetJSON(fullURL string, dest interface{}) ExplorerApiError
	PostJSON(fullURL string, body []byte, dest interface{}) ExplorerApiError
}

type ExplorerApiError interface {
	Error() string
	StatusCode() int
	Wrap(err error) ExplorerApiError
	Unwrap() error
	AsError() error
}

const explorerPathPrefix = "/api/explorer/v1"
const systemPathPrefix = "/api/system/v1"
const invalidCode = "invalid status code"

var servicePathTemplate = func(service string) string {
	return fmt.Sprintf("/api/services/%s/v1", service)
}

type explorerApi struct {
	baseURL URL
	client  http.Client
	header  http.Header
}

type explorerApiError struct {
	message string
	err     error
	code    int
}

func (e *explorerApiError) Error() string {
	if e.err != nil {
		if len(e.message) > 0 {
			return e.message + ": " + e.err.Error()
		}
		return e.err.Error()
	}
	return e.message
}

func (e *explorerApiError) StatusCode() int {
	return e.code
}

func (e *explorerApiError) Wrap(err error) ExplorerApiError {
	e.err = err
	return e
}

func (e *explorerApiError) Unwrap() error {
	return e.err
}

func (e *explorerApiError) AsError() error {
	return errors.New(e.Error())
}

func NewError(msg string, statusCode int) ExplorerApiError {
	return &explorerApiError{
		message: msg,
		code:    statusCode,
	}
}

func WrapError(err error) ExplorerApiError {
	if err == nil {
		return nil
	}
	return &explorerApiError{
		err: err,
	}
}

func NewExplorerApi(baseURL URL) ExplorerApi {
	return explorerApi{}.New(baseURL)
}

func (explorerApi) New(baseURL URL) ExplorerApi {
	header := make(http.Header)
	header.Set("Content-Type", "application/json")
	return &explorerApi{
		baseURL: baseURL,
		client:  http.Client{Timeout: 15 * time.Second},
		header:  header,
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

func (api *explorerApi) SystemPath(prefix string) *URL {
	return api.baseURL.SetBasePath(systemPathPrefix).SetPath(prefix)
}

func (api *explorerApi) ExplorerPath(prefix string) *URL {
	return api.baseURL.SetBasePath(explorerPathPrefix).SetPath(prefix)
}

func (api *explorerApi) ServicePath(serviceName, prefix string) *URL {
	return api.baseURL.SetBasePath(servicePathTemplate(serviceName)).SetPath(prefix)
}

func (api *explorerApi) Stats() (*Stats, ExplorerApiError) {
	reqURL := api.SystemPath("/stats").String()

	result := new(Stats)
	err := api.GetJSON(reqURL, result)
	return result, err
}

func (api *explorerApi) Services() (*ServiceList, ExplorerApiError) {
	reqURL := api.SystemPath("/services").String()

	result := new(ServiceList)
	err := api.GetJSON(reqURL, result)
	return result, err
}

func (api *explorerApi) HealthCheck() (*HealthCheck, ExplorerApiError) {
	reqURL := api.SystemPath("/healthcheck").String()

	result := new(HealthCheck)
	err := api.GetJSON(reqURL, result)
	return result, err
}

func (api *explorerApi) LastBlock() (*FullBlock, ExplorerApiError) {
	val := make(url.Values)
	val.Set("count", fmt.Sprintf("%v", 1))
	val.Set("add_blocks_time", "true")
	reqURL := api.ExplorerPath("/blocks").WithQuery(val)

	result := new(BlocksResponse)
	err := api.GetJSON(reqURL, result)
	if err != nil {
		return nil, err
	}

	if len(result.Blocks) < 1 {
		result.Blocks = []Block{Block{}}
	}

	height := result.Blocks[0].Height
	block, err := api.GetBlock(height)

	return block, err
}

func (api *explorerApi) GetBlocks(count uint32, latest uint64, skipEmptyBlocks bool, addTime bool) (*BlocksResponse, ExplorerApiError) {
	val := make(url.Values)
	val.Set("count", fmt.Sprintf("%v", count))

	if latest > 0 {
		val.Set("latest", fmt.Sprintf("%v", latest))
	}

	val.Set("skip_empty_blocks", fmt.Sprintf("%v", skipEmptyBlocks))
	val.Set("add_blocks_time", fmt.Sprintf("%v", addTime))
	reqURL := api.ExplorerPath("/blocks").WithQuery(val)

	result := new(BlocksResponse)
	err := api.GetJSON(reqURL, result)
	return result, err
}

func (api *explorerApi) GetBlock(height uint64) (*FullBlock, ExplorerApiError) {
	val := make(url.Values)
	val.Set("height", fmt.Sprintf("%v", height))
	reqURL := api.ExplorerPath("/block").WithQuery(val)

	result := new(FullBlock)
	err := api.GetJSON(reqURL, result)
	return result, err
}

func (api *explorerApi) GetTx(hash crypto.Hash) (*FullTx, ExplorerApiError) {
	val := make(url.Values)
	val.Set("hash", hash.String())
	reqURL := api.ExplorerPath("/transactions").WithQuery(val)

	result := new(FullTx)
	err := api.GetJSON(reqURL, result)
	return result, err
}

func (api *explorerApi) SubmitTx(signedTx string) (*TxResult, ExplorerApiError) {
	fullURL := api.ExplorerPath("/transactions").String()

	rawData, err := json.Marshal(map[string]string{
		"tx_body": signedTx,
	})
	if err != nil {
		return nil, WrapError(err)
	}

	req, err := http.NewRequest("POST", fullURL, bytes.NewBuffer(rawData))
	if err != nil {
		return nil, WrapError(err)
	}
	req.Close = true

	req.Header = api.header
	resp, err := api.client.Do(req)
	if err != nil {
		return nil, WrapError(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return nil, NewError(invalidCode, resp.StatusCode)
	}

	result := new(TxResult)
	err = json.NewDecoder(resp.Body).Decode(result)
	return result, WrapError(err)
}

func (api *explorerApi) GetJSON(fullURL string, dest interface{}) ExplorerApiError {
	resp, err := api.client.Get(fullURL)
	if err != nil {
		return WrapError(err)
	}
	if resp.StatusCode >= 400 {
		return NewError(invalidCode, resp.StatusCode)
	}
	return WrapError(json.NewDecoder(resp.Body).Decode(dest))
}

func (api *explorerApi) PostJSON(fullURL string, body []byte, dest interface{}) ExplorerApiError {
	resp, err := api.client.Post(fullURL, "application/json", bytes.NewReader(body))
	if err != nil {
		return WrapError(err)
	}
	if resp.StatusCode >= 400 {
		return NewError(invalidCode, resp.StatusCode)
	}
	return WrapError(json.NewDecoder(resp.Body).Decode(dest))
}
