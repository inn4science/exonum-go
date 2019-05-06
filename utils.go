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
	"net/url"
	"strings"
)

type URL struct {
	URL      *url.URL
	basePath string
}

const slash = "/"

func (URL) New(rawURL string) (URL, error) {
	var s string
	j := URL{}
	u, err := url.Parse(s)
	if err != nil {
		return j, err
	}

	j.URL = u
	j.basePath = strings.TrimSuffix(u.Path, slash)
	return j, err
}

func (j *URL) SetBasePath(path string) *URL {
	j.basePath = path
	return j
}

func (j *URL) SetPath(path string) *URL {
	j.URL.Path = j.basePath + slash + strings.TrimPrefix(path, slash)
	return j
}

func (j *URL) WithPath(path string) string {
	ur := *j.URL
	ur.Path = j.basePath + slash + strings.TrimPrefix(path, slash)

	return ur.String()
}

func (j *URL) WithQuery(values url.Values) string {
	ur := *j.URL
	ur.RawQuery = values.Encode()
	return ur.String()
}

func (j *URL) WithPathURL(path string) url.URL {
	ur := *j.URL
	ur.Path = j.basePath + slash + strings.TrimPrefix(path, slash)

	return ur
}

func (j *URL) String() string {
	return j.URL.String()
}
