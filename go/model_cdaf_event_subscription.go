/*
 * Ncdaf_EventExposure
 *
 * CDAF Event Exposure Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// CdafEventSubscription - Represents an individual event subscription resource on CDAF
type CdafEventSubscription struct {

	EventRequest CdafEvent `json:"eventRequest"`

	// String providing an URI formatted according to RFC 3986.
	EventNotifyUri string `json:"eventNotifyUri"`

	Options CdafEventMode `json:"options,omitempty"`
}

// AssertCdafEventSubscriptionRequired checks if the required fields are not zero-ed
func AssertCdafEventSubscriptionRequired(obj CdafEventSubscription) error {
	elements := map[string]interface{}{
		"eventRequest": obj.EventRequest,
		"eventNotifyUri": obj.EventNotifyUri,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertCdafEventRequired(obj.EventRequest); err != nil {
		return err
	}
	if err := AssertCdafEventModeRequired(obj.Options); err != nil {
		return err
	}
	return nil
}

// AssertRecurseCdafEventSubscriptionRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of CdafEventSubscription (e.g. [][]CdafEventSubscription), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseCdafEventSubscriptionRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aCdafEventSubscription, ok := obj.(CdafEventSubscription)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertCdafEventSubscriptionRequired(aCdafEventSubscription)
	})
}
