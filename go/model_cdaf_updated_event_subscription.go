// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Ncdaf_EventExposure
 *
 * CDAF Event Exposure Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.3
 */

package openapi




// CdafUpdatedEventSubscription - Represents a successful update on an CDAF Event Subscription
type CdafUpdatedEventSubscription struct {

	Subscription CdafEventSubscription `json:"subscription"`

	ReportList []CdafEventReport `json:"reportList,omitempty"`
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
	for _, el := range obj.ReportList {
		if err := AssertCdafEventReportRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertCdafUpdatedEventSubscriptionConstraints checks if the values respects the defined constraints
func AssertCdafUpdatedEventSubscriptionConstraints(obj CdafUpdatedEventSubscription) error {
	return nil
}
