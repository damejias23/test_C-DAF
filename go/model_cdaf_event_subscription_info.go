/*
 * Ncdaf_EventExposure
 *
 * CDAF Event Exposure Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// CdafEventSubscriptionInfo - Individual CDAF Event Subscription Information
type CdafEventSubscriptionInfo struct {

	// String providing an URI formatted according to RFC 3986.
	SubId string `json:"subId"`

	// String providing an URI formatted according to RFC 3986.
	OldSubId string `json:"oldSubId,omitempty"`
}

// AssertCdafEventSubscriptionInfoRequired checks if the required fields are not zero-ed
func AssertCdafEventSubscriptionInfoRequired(obj CdafEventSubscriptionInfo) error {
	elements := map[string]interface{}{
		"subId": obj.SubId,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertRecurseCdafEventSubscriptionInfoRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of CdafEventSubscriptionInfo (e.g. [][]CdafEventSubscriptionInfo), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseCdafEventSubscriptionInfoRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aCdafEventSubscriptionInfo, ok := obj.(CdafEventSubscriptionInfo)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertCdafEventSubscriptionInfoRequired(aCdafEventSubscriptionInfo)
	})
}
