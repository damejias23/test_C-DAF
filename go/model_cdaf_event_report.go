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
	"time"
)

// CdafEventReport - Represents a report triggered by a subscribed event type
type CdafEventReport struct {

	Type CdafEventType `json:"type"`

	State CdafEventState `json:"state"`

	// string with format 'date-time' as defined in OpenAPI.
	TimeStamp time.Time `json:"timeStamp"`

	// String providing an URI formatted according to RFC 3986.
	SubscriptionId string `json:"subscriptionId,omitempty"`

	NfLoadLevelInfos []NfLoadLevelInformation `json:"nfLoadLevelInfos,omitempty"`
}

// AssertCdafEventReportRequired checks if the required fields are not zero-ed
func AssertCdafEventReportRequired(obj CdafEventReport) error {
	elements := map[string]interface{}{
		"type": obj.Type,
		"state": obj.State,
		"timeStamp": obj.TimeStamp,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertCdafEventTypeRequired(obj.Type); err != nil {
		return err
	}
	if err := AssertCdafEventStateRequired(obj.State); err != nil {
		return err
	}
	for _, el := range obj.NfLoadLevelInfos {
		if err := AssertNfLoadLevelInformationRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertRecurseCdafEventReportRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of CdafEventReport (e.g. [][]CdafEventReport), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseCdafEventReportRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aCdafEventReport, ok := obj.(CdafEventReport)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertCdafEventReportRequired(aCdafEventReport)
	})
}
