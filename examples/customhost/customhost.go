// Copyright 2018 The go-netbox Authors.
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

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/digitalocean/go-netbox/netbox/client"
	runtimeclient "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

func main() {
	token := os.Getenv("NETBOX_TOKEN")
	if token == "" {
		log.Fatal("no such token provided")
	}
	host := os.Getenv("NETBOX_HOST")
	if host == "" {
		log.Fatal("no such host provided")
	}
	t := runtimeclient.New(host, client.DefaultBasePath, []string{"https"})
	t.DefaultAuthentication = runtimeclient.APIKeyAuth("Authorization", "header", fmt.Sprintf("Token %v", token))
	c := client.New(t, strfmt.Default)

	// rs, err := c.API.APIIPAMIPAddressesList(nil, nil)
	// rs, err := c.API.APIDcimDevicesList(nil, nil)
	// rs, err := c.API.APIDcimInterfacesList(nil, nil)
	// rs, err := c.API.APIDcimInterfaceConnectionsList(nil, nil)
	rs, err := c.API.APICircuitsCircuitTerminationsList(nil, nil)
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	for _, d := range rs.Payload.Results {
		if d.ConnectedEndpoint != nil {
			fmt.Printf("%#v\n", d.ConnectedEndpoint)
		}
	}
}
