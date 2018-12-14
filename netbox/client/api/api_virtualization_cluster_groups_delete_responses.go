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

// APIVirtualizationClusterGroupsDeleteReader is a Reader for the APIVirtualizationClusterGroupsDelete structure.
type APIVirtualizationClusterGroupsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *APIVirtualizationClusterGroupsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewAPIVirtualizationClusterGroupsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAPIVirtualizationClusterGroupsDeleteNoContent creates a APIVirtualizationClusterGroupsDeleteNoContent with default headers values
func NewAPIVirtualizationClusterGroupsDeleteNoContent() *APIVirtualizationClusterGroupsDeleteNoContent {
	return &APIVirtualizationClusterGroupsDeleteNoContent{}
}

/*APIVirtualizationClusterGroupsDeleteNoContent handles this case with default header values.

APIVirtualizationClusterGroupsDeleteNoContent api virtualization cluster groups delete no content
*/
type APIVirtualizationClusterGroupsDeleteNoContent struct {
}

func (o *APIVirtualizationClusterGroupsDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/virtualization/cluster-groups/{id}/][%d] apiVirtualizationClusterGroupsDeleteNoContent ", 204)
}

func (o *APIVirtualizationClusterGroupsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}