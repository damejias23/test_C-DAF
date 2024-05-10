/*
 * Ncdaf_EventExposure
 *
 * CDAF Event Exposure Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"encoding/json"
	"net/http"
	"strings"
)

// IndividualSubscriptionDocumentApiController binds http requests to an api service and writes the service results to the http response
type IndividualSubscriptionDocumentApiController struct {
	service IndividualSubscriptionDocumentApiServicer
	errorHandler ErrorHandler
}

// IndividualSubscriptionDocumentApiOption for how the controller is set up.
type IndividualSubscriptionDocumentApiOption func(*IndividualSubscriptionDocumentApiController)

// WithIndividualSubscriptionDocumentApiErrorHandler inject ErrorHandler into controller
func WithIndividualSubscriptionDocumentApiErrorHandler(h ErrorHandler) IndividualSubscriptionDocumentApiOption {
	return func(c *IndividualSubscriptionDocumentApiController) {
		c.errorHandler = h
	}
}

// NewIndividualSubscriptionDocumentApiController creates a default api controller
func NewIndividualSubscriptionDocumentApiController(s IndividualSubscriptionDocumentApiServicer, opts ...IndividualSubscriptionDocumentApiOption) Router {
	controller := &IndividualSubscriptionDocumentApiController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the IndividualSubscriptionDocumentApiController
func (c *IndividualSubscriptionDocumentApiController) Routes() Routes {
	return Routes{ 
		{
			"DeleteSubscription",
			strings.ToUpper("Delete"),
			"/ncdaf-evts/v1/subscriptions/{subscriptionId}",
			c.DeleteSubscription,
		},
		{
			"ModifySubscription",
			strings.ToUpper("Patch"),
			"/ncdaf-evts/v1/subscriptions/{subscriptionId}",
			c.ModifySubscription,
		},
	}
}

// DeleteSubscription - Ncdaf_EventExposure Unsubscribe service Operation
func (c *IndividualSubscriptionDocumentApiController) DeleteSubscription(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	subscriptionIdParam := params["subscriptionId"]
	
	result, err := c.service.DeleteSubscription(r.Context(), subscriptionIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// ModifySubscription - Ncdaf_EventExposure Subscribe Modify service Operation
func (c *IndividualSubscriptionDocumentApiController) ModifySubscription(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	subscriptionIdParam := params["subscriptionId"]
	
	modifySubscriptionRequestParam := ModifySubscriptionRequest{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&modifySubscriptionRequestParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertModifySubscriptionRequestRequired(modifySubscriptionRequestParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.ModifySubscription(r.Context(), subscriptionIdParam, modifySubscriptionRequestParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}
