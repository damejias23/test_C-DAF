// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Ncdaf_EventExposure
 *
 * CDAF Event Exposure Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.3
 */

package openapi




// CdafCreatedEventSubscription - Data within a create CDAF event subscription response
type CdafCreatedEventSubscription struct {

	Subscription CdafEventSubscription `json:"subscription"`

	// String providing an URI formatted according to RFC 3986.
	SubscriptionId string `json:"subscriptionId"`

	ReportList []CdafEventReport `json:"reportList,omitempty"`

	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SupportedFeatures string `json:"supportedFeatures,omitempty"`
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
	for _, el := range obj.ReportList {
		if err := AssertCdafEventReportRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertCdafCreatedEventSubscriptionConstraints checks if the values respects the defined constraints
func AssertCdafCreatedEventSubscriptionConstraints(obj CdafCreatedEventSubscription) error {
	return nil
}
