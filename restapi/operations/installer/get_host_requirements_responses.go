// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openshift/assisted-service/models"
)

// GetHostRequirementsOKCode is the HTTP code returned for type GetHostRequirementsOK
const GetHostRequirementsOKCode int = 200

/*GetHostRequirementsOK Success.

swagger:response getHostRequirementsOK
*/
type GetHostRequirementsOK struct {

	/*
	  In: Body
	*/
	Payload *models.HostRequirements `json:"body,omitempty"`
}

// NewGetHostRequirementsOK creates GetHostRequirementsOK with default headers values
func NewGetHostRequirementsOK() *GetHostRequirementsOK {

	return &GetHostRequirementsOK{}
}

// WithPayload adds the payload to the get host requirements o k response
func (o *GetHostRequirementsOK) WithPayload(payload *models.HostRequirements) *GetHostRequirementsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get host requirements o k response
func (o *GetHostRequirementsOK) SetPayload(payload *models.HostRequirements) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetHostRequirementsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetHostRequirementsUnauthorizedCode is the HTTP code returned for type GetHostRequirementsUnauthorized
const GetHostRequirementsUnauthorizedCode int = 401

/*GetHostRequirementsUnauthorized Unauthorized.

swagger:response getHostRequirementsUnauthorized
*/
type GetHostRequirementsUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.InfraError `json:"body,omitempty"`
}

// NewGetHostRequirementsUnauthorized creates GetHostRequirementsUnauthorized with default headers values
func NewGetHostRequirementsUnauthorized() *GetHostRequirementsUnauthorized {

	return &GetHostRequirementsUnauthorized{}
}

// WithPayload adds the payload to the get host requirements unauthorized response
func (o *GetHostRequirementsUnauthorized) WithPayload(payload *models.InfraError) *GetHostRequirementsUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get host requirements unauthorized response
func (o *GetHostRequirementsUnauthorized) SetPayload(payload *models.InfraError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetHostRequirementsUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetHostRequirementsForbiddenCode is the HTTP code returned for type GetHostRequirementsForbidden
const GetHostRequirementsForbiddenCode int = 403

/*GetHostRequirementsForbidden Forbidden.

swagger:response getHostRequirementsForbidden
*/
type GetHostRequirementsForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.InfraError `json:"body,omitempty"`
}

// NewGetHostRequirementsForbidden creates GetHostRequirementsForbidden with default headers values
func NewGetHostRequirementsForbidden() *GetHostRequirementsForbidden {

	return &GetHostRequirementsForbidden{}
}

// WithPayload adds the payload to the get host requirements forbidden response
func (o *GetHostRequirementsForbidden) WithPayload(payload *models.InfraError) *GetHostRequirementsForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get host requirements forbidden response
func (o *GetHostRequirementsForbidden) SetPayload(payload *models.InfraError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetHostRequirementsForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetHostRequirementsMethodNotAllowedCode is the HTTP code returned for type GetHostRequirementsMethodNotAllowed
const GetHostRequirementsMethodNotAllowedCode int = 405

/*GetHostRequirementsMethodNotAllowed Method Not Allowed.

swagger:response getHostRequirementsMethodNotAllowed
*/
type GetHostRequirementsMethodNotAllowed struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetHostRequirementsMethodNotAllowed creates GetHostRequirementsMethodNotAllowed with default headers values
func NewGetHostRequirementsMethodNotAllowed() *GetHostRequirementsMethodNotAllowed {

	return &GetHostRequirementsMethodNotAllowed{}
}

// WithPayload adds the payload to the get host requirements method not allowed response
func (o *GetHostRequirementsMethodNotAllowed) WithPayload(payload *models.Error) *GetHostRequirementsMethodNotAllowed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get host requirements method not allowed response
func (o *GetHostRequirementsMethodNotAllowed) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetHostRequirementsMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(405)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetHostRequirementsInternalServerErrorCode is the HTTP code returned for type GetHostRequirementsInternalServerError
const GetHostRequirementsInternalServerErrorCode int = 500

/*GetHostRequirementsInternalServerError Error.

swagger:response getHostRequirementsInternalServerError
*/
type GetHostRequirementsInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetHostRequirementsInternalServerError creates GetHostRequirementsInternalServerError with default headers values
func NewGetHostRequirementsInternalServerError() *GetHostRequirementsInternalServerError {

	return &GetHostRequirementsInternalServerError{}
}

// WithPayload adds the payload to the get host requirements internal server error response
func (o *GetHostRequirementsInternalServerError) WithPayload(payload *models.Error) *GetHostRequirementsInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get host requirements internal server error response
func (o *GetHostRequirementsInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetHostRequirementsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
