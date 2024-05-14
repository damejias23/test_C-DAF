/*
 * Ncdaf_EventExposure
 *
 * CDAF Event Exposure Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// CdafCreatedEventSubscription - Data within a create CDAF event subscription response
type CdafCreatedEventSubscription struct {

	Subscription CdafEventSubscription `json:"subscription"`

	// String providing an URI formatted according to RFC 3986.
	SubscriptionId string `json:"subscriptionId"`

	ReportEvent CdafEventReport `json:"reportEvent,omitempty"`
}

// AssertCdafCreatedEventSubscriptionRequired checks if the required fields are not zero-ed
func AssertCdafCreatedEventSubscriptionRequired(obj CdafCreatedEventSubscription) error {
	elements := map[string]interface{}{
		"subscription": obj.Subscription,
		"subscriptionId": obj.SubscriptionId,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertCdafEventSubscriptionRequired(obj.Subscription); err != nil {
		return err
	}
	if err := AssertCdafEventReportRequired(obj.ReportEvent); err != nil {
		return err
	}
	return nil
}

// AssertRecurseCdafCreatedEventSubscriptionRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of CdafCreatedEventSubscription (e.g. [][]CdafCreatedEventSubscription), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseCdafCreatedEventSubscriptionRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aCdafCreatedEventSubscription, ok := obj.(CdafCreatedEventSubscription)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertCdafCreatedEventSubscriptionRequired(aCdafCreatedEventSubscription)
	})
}
