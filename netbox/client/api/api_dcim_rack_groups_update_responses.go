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
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/digitalocean/go-netbox/netbox/models"
)

// APIDcimRackGroupsUpdateReader is a Reader for the APIDcimRackGroupsUpdate structure.
type APIDcimRackGroupsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *APIDcimRackGroupsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewAPIDcimRackGroupsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAPIDcimRackGroupsUpdateOK creates a APIDcimRackGroupsUpdateOK with default headers values
func NewAPIDcimRackGroupsUpdateOK() *APIDcimRackGroupsUpdateOK {
	return &APIDcimRackGroupsUpdateOK{}
}

/*APIDcimRackGroupsUpdateOK handles this case with default header values.

APIDcimRackGroupsUpdateOK api dcim rack groups update o k
*/
type APIDcimRackGroupsUpdateOK struct {
	Payload *models.RackGroup
}

func (o *APIDcimRackGroupsUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /api/dcim/rack-groups/{id}/][%d] apiDcimRackGroupsUpdateOK  %+v", 200, o.Payload)
}

func (o *APIDcimRackGroupsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RackGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}