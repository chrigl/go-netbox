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

// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by robots at
// 2017-05-18 14:20:17.557582817 +0200 CEST

package netbox

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// TenantsService is used in a Client to access NetBox's tenancy/tenants API methods.
type TenantsService struct {
	c *Client
}

// Get retrieves an Tenant object from NetBox by its ID.
func (s *TenantsService) Get(id int) (*Tenant, error) {
	req, err := s.c.NewRequest(
		http.MethodGet,
		fmt.Sprintf("api/tenancy/tenants/%d/", id),
		nil,
	)
	if err != nil {
		return nil, err
	}

	t := new(Tenant)
	err = s.c.Do(req, t)
	if err != nil {
		return nil, err
	}
	return t, nil
}

// List returns a Page associated with an NetBox API Endpoint.
func (s *TenantsService) List(options *ListTenantOptions) *Page {
	return NewPage(s.c, "api/tenancy/tenants/", options)
}

// Extract retrives a list of Tenant objects from page.
func (s *TenantsService) Extract(page *Page) ([]*Tenant, error) {
	if err := page.Err(); err != nil {
		return nil, err
	}

	var groups []*Tenant
	if err := json.Unmarshal(page.data.Results, &groups); err != nil {
		return nil, err
	}
	return groups, nil
}

// Create creates a new Tenant object in NetBox and returns the ID of the new object.
func (s *TenantsService) Create(data *Tenant) (int, error) {
	req, err := s.c.NewJSONRequest(http.MethodPost, "api/tenancy/tenants/", nil, data)
	if err != nil {
		return 0, err
	}

	g := new(updateTenant)
	err = s.c.Do(req, g)
	if err != nil {
		return 0, err
	}
	return g.ID, nil
}

// Update changes an existing Tenant object in NetBox, and returns the ID of the new object.
func (s *TenantsService) Update(data *Tenant) (int, error) {
	req, err := s.c.NewJSONRequest(
		http.MethodPatch,
		fmt.Sprintf("api/tenancy/tenants/%d/", data.ID),
		nil,
		data)
	if err != nil {
		return 0, err
	}

	// g is just used to verify correct api result.
	// data is not changed, because the g is not the full representation that one would
	// get with Get. But if the response was unmarshaled into updateTenant corretcly,
	// everything went fine, and we do not need to update data.
	g := new(updateTenant)
	err = s.c.Do(req, g)
	if err != nil {
		return 0, err
	}
	return g.ID, nil
}

// Delete deletes an existing Tenant object from NetBox.
func (s *TenantsService) Delete(data *Tenant) error {
	req, err := s.c.NewRequest(
		http.MethodDelete,
		fmt.Sprintf("api/tenancy/tenants/%d/", data.ID),
		nil)
	if err != nil {
		return err
	}

	return s.c.Do(req, nil)
}