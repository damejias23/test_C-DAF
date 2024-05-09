// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Ncdaf_EventExposure
 *
 * CDAF Event Exposure Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.3
 */

package openapi




// CdafEventNotification - Data within a CADF Event Notification request
type CdafEventNotification struct {

	NotifyCorrelationId string `json:"notifyCorrelationId,omitempty"`

	SubsChangeNotifyCorrelationId string `json:"subsChangeNotifyCorrelationId,omitempty"`

	ReportList []CdafEventReport `json:"reportList,omitempty"`

	EventSubsSyncInfo CdafEventSubsSyncInfo `json:"eventSubsSyncInfo,omitempty"`
}

// AssertCdafEventNotificationRequired checks if the required fields are not zero-ed
func AssertCdafEventNotificationRequired(obj CdafEventNotification) error {
	for _, el := range obj.ReportList {
		if err := AssertCdafEventReportRequired(el); err != nil {
			return err
		}
	}
	if err := AssertCdafEventSubsSyncInfoRequired(obj.EventSubsSyncInfo); err != nil {
		return err
	}
	return nil
}

// AssertCdafEventNotificationConstraints checks if the values respects the defined constraints
func AssertCdafEventNotificationConstraints(obj CdafEventNotification) error {
	return nil
}
