//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2020 SeMI Technologies B.V. All rights reserved.
//
//  CONTACT: hello@semi.technology
//

// Code generated by go-swagger; DO NOT EDIT.

package actions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/semi-technologies/weaviate/entities/models"
)

// ActionsDeleteNoContentCode is the HTTP code returned for type ActionsDeleteNoContent
const ActionsDeleteNoContentCode int = 204

/*ActionsDeleteNoContent Successfully deleted.

swagger:response actionsDeleteNoContent
*/
type ActionsDeleteNoContent struct {
}

// NewActionsDeleteNoContent creates ActionsDeleteNoContent with default headers values
func NewActionsDeleteNoContent() *ActionsDeleteNoContent {

	return &ActionsDeleteNoContent{}
}

// WriteResponse to the client
func (o *ActionsDeleteNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// ActionsDeleteUnauthorizedCode is the HTTP code returned for type ActionsDeleteUnauthorized
const ActionsDeleteUnauthorizedCode int = 401

/*ActionsDeleteUnauthorized Unauthorized or invalid credentials.

swagger:response actionsDeleteUnauthorized
*/
type ActionsDeleteUnauthorized struct {
}

// NewActionsDeleteUnauthorized creates ActionsDeleteUnauthorized with default headers values
func NewActionsDeleteUnauthorized() *ActionsDeleteUnauthorized {

	return &ActionsDeleteUnauthorized{}
}

// WriteResponse to the client
func (o *ActionsDeleteUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// ActionsDeleteForbiddenCode is the HTTP code returned for type ActionsDeleteForbidden
const ActionsDeleteForbiddenCode int = 403

/*ActionsDeleteForbidden Forbidden

swagger:response actionsDeleteForbidden
*/
type ActionsDeleteForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewActionsDeleteForbidden creates ActionsDeleteForbidden with default headers values
func NewActionsDeleteForbidden() *ActionsDeleteForbidden {

	return &ActionsDeleteForbidden{}
}

// WithPayload adds the payload to the actions delete forbidden response
func (o *ActionsDeleteForbidden) WithPayload(payload *models.ErrorResponse) *ActionsDeleteForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the actions delete forbidden response
func (o *ActionsDeleteForbidden) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ActionsDeleteForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ActionsDeleteNotFoundCode is the HTTP code returned for type ActionsDeleteNotFound
const ActionsDeleteNotFoundCode int = 404

/*ActionsDeleteNotFound Successful query result but no resource was found.

swagger:response actionsDeleteNotFound
*/
type ActionsDeleteNotFound struct {
}

// NewActionsDeleteNotFound creates ActionsDeleteNotFound with default headers values
func NewActionsDeleteNotFound() *ActionsDeleteNotFound {

	return &ActionsDeleteNotFound{}
}

// WriteResponse to the client
func (o *ActionsDeleteNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// ActionsDeleteInternalServerErrorCode is the HTTP code returned for type ActionsDeleteInternalServerError
const ActionsDeleteInternalServerErrorCode int = 500

/*ActionsDeleteInternalServerError An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.

swagger:response actionsDeleteInternalServerError
*/
type ActionsDeleteInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewActionsDeleteInternalServerError creates ActionsDeleteInternalServerError with default headers values
func NewActionsDeleteInternalServerError() *ActionsDeleteInternalServerError {

	return &ActionsDeleteInternalServerError{}
}

// WithPayload adds the payload to the actions delete internal server error response
func (o *ActionsDeleteInternalServerError) WithPayload(payload *models.ErrorResponse) *ActionsDeleteInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the actions delete internal server error response
func (o *ActionsDeleteInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ActionsDeleteInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
