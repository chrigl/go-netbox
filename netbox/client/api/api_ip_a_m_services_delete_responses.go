// Code generated by go-swagger; DO NOT EDIT.

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

package api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// APIIPAMServicesDeleteReader is a Reader for the APIIPAMServicesDelete structure.
type APIIPAMServicesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *APIIPAMServicesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewAPIIPAMServicesDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAPIIPAMServicesDeleteNoContent creates a APIIPAMServicesDeleteNoContent with default headers values
func NewAPIIPAMServicesDeleteNoContent() *APIIPAMServicesDeleteNoContent {
	return &APIIPAMServicesDeleteNoContent{}
}

/*APIIPAMServicesDeleteNoContent handles this case with default header values.

APIIPAMServicesDeleteNoContent api ipam services delete no content
*/
type APIIPAMServicesDeleteNoContent struct {
}

func (o *APIIPAMServicesDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/ipam/services/{id}/][%d] apiIpamServicesDeleteNoContent ", 204)
}

func (o *APIIPAMServicesDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}