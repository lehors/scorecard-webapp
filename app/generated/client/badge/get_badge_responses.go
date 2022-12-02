// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2021 OpenSSF Scorecard Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package badge

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ossf/scorecard-webapp/app/generated/models"
)

// GetBadgeReader is a Reader for the GetBadge structure.
type GetBadgeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBadgeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 302:
		result := NewGetBadgeFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetBadgeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetBadgeFound creates a GetBadgeFound with default headers values
func NewGetBadgeFound() *GetBadgeFound {
	return &GetBadgeFound{}
}

/*
GetBadgeFound describes a response with status code 302, with default header values.

Scorecard badge for the repository
*/
type GetBadgeFound struct {
}

// IsSuccess returns true when this get badge found response has a 2xx status code
func (o *GetBadgeFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get badge found response has a 3xx status code
func (o *GetBadgeFound) IsRedirect() bool {
	return true
}

// IsClientError returns true when this get badge found response has a 4xx status code
func (o *GetBadgeFound) IsClientError() bool {
	return false
}

// IsServerError returns true when this get badge found response has a 5xx status code
func (o *GetBadgeFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get badge found response a status code equal to that given
func (o *GetBadgeFound) IsCode(code int) bool {
	return code == 302
}

func (o *GetBadgeFound) Error() string {
	return fmt.Sprintf("[GET /projects/{platform}/{org}/{repo}/badge][%d] getBadgeFound ", 302)
}

func (o *GetBadgeFound) String() string {
	return fmt.Sprintf("[GET /projects/{platform}/{org}/{repo}/badge][%d] getBadgeFound ", 302)
}

func (o *GetBadgeFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetBadgeDefault creates a GetBadgeDefault with default headers values
func NewGetBadgeDefault(code int) *GetBadgeDefault {
	return &GetBadgeDefault{
		_statusCode: code,
	}
}

/*
GetBadgeDefault describes a response with status code -1, with default header values.

There was an internal error in the server while processing the request
*/
type GetBadgeDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get badge default response
func (o *GetBadgeDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this get badge default response has a 2xx status code
func (o *GetBadgeDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get badge default response has a 3xx status code
func (o *GetBadgeDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get badge default response has a 4xx status code
func (o *GetBadgeDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get badge default response has a 5xx status code
func (o *GetBadgeDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get badge default response a status code equal to that given
func (o *GetBadgeDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *GetBadgeDefault) Error() string {
	return fmt.Sprintf("[GET /projects/{platform}/{org}/{repo}/badge][%d] getBadge default  %+v", o._statusCode, o.Payload)
}

func (o *GetBadgeDefault) String() string {
	return fmt.Sprintf("[GET /projects/{platform}/{org}/{repo}/badge][%d] getBadge default  %+v", o._statusCode, o.Payload)
}

func (o *GetBadgeDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetBadgeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
