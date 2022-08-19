// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2021 Security Scorecard Authors
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

package results

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/ossf/scorecard-webapp/app/generated/models"
)

// PostResultCreatedCode is the HTTP code returned for type PostResultCreated
const PostResultCreatedCode int = 201

/*PostResultCreated Successfully updated ScorecardResult

swagger:response postResultCreated
*/
type PostResultCreated struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewPostResultCreated creates PostResultCreated with default headers values
func NewPostResultCreated() *PostResultCreated {

	return &PostResultCreated{}
}

// WithPayload adds the payload to the post result created response
func (o *PostResultCreated) WithPayload(payload string) *PostResultCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post result created response
func (o *PostResultCreated) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostResultCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// PostResultBadRequestCode is the HTTP code returned for type PostResultBadRequest
const PostResultBadRequestCode int = 400

/*PostResultBadRequest The request provided to the server was invalid

swagger:response postResultBadRequest
*/
type PostResultBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostResultBadRequest creates PostResultBadRequest with default headers values
func NewPostResultBadRequest() *PostResultBadRequest {

	return &PostResultBadRequest{}
}

// WithPayload adds the payload to the post result bad request response
func (o *PostResultBadRequest) WithPayload(payload *models.Error) *PostResultBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post result bad request response
func (o *PostResultBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostResultBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*PostResultDefault There was an internal error in the server while processing the request

swagger:response postResultDefault
*/
type PostResultDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostResultDefault creates PostResultDefault with default headers values
func NewPostResultDefault(code int) *PostResultDefault {
	if code <= 0 {
		code = 500
	}

	return &PostResultDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the post result default response
func (o *PostResultDefault) WithStatusCode(code int) *PostResultDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the post result default response
func (o *PostResultDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the post result default response
func (o *PostResultDefault) WithPayload(payload *models.Error) *PostResultDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post result default response
func (o *PostResultDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostResultDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}