/*
 * Ncdaf_EventExposure
 *
 * CDAF Event Exposure Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// CdafUpdatedEventSubscription - Represents a successful update on an CDAF Event Subscription
type CdafUpdatedEventSubscription struct {

	Subscription CdafEventSubscription `json:"subscription"`

	ReportEvent CdafEventReport `json:"reportEvent,omitempty"`
}

// AssertCdafUpdatedEventSubscriptionRequired checks if the required fields are not zero-ed
func AssertCdafUpdatedEventSubscriptionRequired(obj CdafUpdatedEventSubscription) error {
	elements := map[string]interface{}{
		"subscription": obj.Subscription,
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

// AssertRecurseCdafUpdatedEventSubscriptionRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of CdafUpdatedEventSubscription (e.g. [][]CdafUpdatedEventSubscription), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseCdafUpdatedEventSubscriptionRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aCdafUpdatedEventSubscription, ok := obj.(CdafUpdatedEventSubscription)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertCdafUpdatedEventSubscriptionRequired(aCdafUpdatedEventSubscription)
	})
}
