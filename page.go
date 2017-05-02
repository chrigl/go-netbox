// Copyright 2017 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package netbox

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strconv"
)

type Page struct {
	limit    int
	offset   int
	done     bool
	c        *Client
	data     PageData
	endpoint string
	options  Valuer
	err      error
}

func NewPage(client *Client, endpoint string, options Valuer) *Page {
	return &Page{limit: 0,
		offset:   0,
		c:        client,
		done:     false,
		data:     PageData{},
		endpoint: endpoint,
		options:  options,
		err:      nil}
}

type PageData struct {
	Count       int             `json:"count"`
	NextURL     string          `json:"next"`
	PreviousURL string          `json:"previous"`
	Results     json.RawMessage `json:"results"`
}

func (p *Page) values() (url.Values, error) {
	v, err := p.options.values()
	if err != nil {
		return nil, err
	}
	if v == nil {
		v = url.Values{}
	}
	v.Set("limit", strconv.Itoa(p.limit))
	v.Set("offset", strconv.Itoa(p.offset))

	return v, nil
}

func (p *Page) Next() bool {
	if p.done {
		return false
	}

	req, err := p.c.NewRequest(http.MethodGet, p.endpoint, p)
	if err != nil {
		p.err = err
		return false
	}

	// Clear NextURL, because it will not be overriden if it is empty in the
	// resultset. I(cglaubitz) do not want to allocate a new PageData on each call
	// to Next() to reduce the need for garbage collection. No idea if this still
	// works when I have to "empty" NextURL here.
	p.data.NextURL = ""
	err = p.c.Do(req, &p.data)
	if err != nil {
		p.err = err
		return false
	}

	if p.data.NextURL == "" {
		// We are on the last page, so we still need to return true to
		// indicate that there is still data to process. But we set
		// p.done to true, so the next "Next()" returns false.
		p.done = true
	} else {
		p.SetNextURL(p.data.NextURL)
	}
	return true

}

func (p *Page) setNext(limit int, offset int) {
	p.limit = limit
	p.offset = offset
}

func (p *Page) SetNextURL(urlStr string) {
	nextUrl, err := url.Parse(urlStr)
	if err != nil {
		// We dont want to cancel this run, since there is data,
		// but do not want to run into Next
		p.SetErr(err)
	}

	query, err := url.ParseQuery(nextUrl.RawQuery)
	if err != nil {
		// Same like above
		p.SetErr(err)
		return
	}

	limits := query["limit"]
	offsets := query["offset"]
	if len(limits) == 0 {
		p.SetErr(errors.New("No such query parameter limit."))
		return
	}
	if len(offsets) == 0 {
		p.SetErr(errors.New("No such query parameter offset."))
		return
	}

	limit, err := strconv.Atoi(limits[0])
	if err != nil {
		p.SetErr(err)
		return
	}

	offset, err := strconv.Atoi(offsets[0])
	if err != nil {
		p.SetErr(err)
		return
	}
	p.setNext(limit, offset)
}

func (p *Page) Err() error {
	return p.err
}

func (p *Page) SetErr(err error) {
	if p.err == nil {
		p.err = err
	}
}
