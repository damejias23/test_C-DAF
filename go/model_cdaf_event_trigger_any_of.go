/*
 * Ncdaf_EventExposure
 *
 * CDAF Event Exposure Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type CdafEventTriggerAnyOf string

// List of CdafEventTriggerAnyOf
const (
	ONE_TIME CdafEventTriggerAnyOf = "ONE_TIME"
	CONTINUOUS CdafEventTriggerAnyOf = "CONTINUOUS"
	PERIODIC CdafEventTriggerAnyOf = "PERIODIC"
)

// AssertCdafEventTriggerAnyOfRequired checks if the required fields are not zero-ed
func AssertCdafEventTriggerAnyOfRequired(obj CdafEventTriggerAnyOf) error {
	return nil
}

// AssertRecurseCdafEventTriggerAnyOfRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of CdafEventTriggerAnyOf (e.g. [][]CdafEventTriggerAnyOf), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseCdafEventTriggerAnyOfRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aCdafEventTriggerAnyOf, ok := obj.(CdafEventTriggerAnyOf)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertCdafEventTriggerAnyOfRequired(aCdafEventTriggerAnyOf)
	})
}
