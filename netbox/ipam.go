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

// An IPAMService is used in a Client to access NetBox's IPAM API methods.
type IPAMService struct {
	c        *Client
	VRFs     *VRFsService
	Prefixes *PrefixesService
}

// NewIPAMService returns a IPAMService initialized with all sub-services.
func NewIPAMService(client *Client) *IPAMService {
	return &IPAMService{
		c: client,
		VRFs: &VRFsService{
			c: client,
		},
		Prefixes: &PrefixesService{
			c: client,
		},
	}
}
