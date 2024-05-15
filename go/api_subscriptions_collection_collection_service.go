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
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
)

// SubscriptionsCollectionCollectionApiService is a service that implements the logic for the SubscriptionsCollectionCollectionApiServicer
// This service should implement the business logic for every endpoint for the SubscriptionsCollectionCollectionApi API.
// Include any external packages or services that will be required by this service.
type SubscriptionsCollectionCollectionApiService struct {
}

// NewSubscriptionsCollectionCollectionApiService creates a default api service
func NewSubscriptionsCollectionCollectionApiService() SubscriptionsCollectionCollectionApiServicer {
	return &SubscriptionsCollectionCollectionApiService{}
}

// CreateSubscription - Ncdaf_EventExposure Subscribe service Operation
func (s *SubscriptionsCollectionCollectionApiService) CreateSubscription(ctx context.Context, cdafCreateEventSubscription CdafCreateEventSubscription) (ImplResponse, error) {
	// TODO - update CreateSubscription with the required logic for this service method.
	// Add api_subscriptions_collection_collection_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(201, CdafCreatedEventSubscription{}) or use other options such as http.Ok ...
	//return Response(201, CdafCreatedEventSubscription{}), nil

	//TODO: Uncomment the next line to return response Response(0, {}) or use other options such as http.Ok ...
	//return Response(0, nil),nil
	subscriptionId := uuid.New().String()
	subscriptionCh := make(chan string)
	// add subscription channel to mapping table
	// subscriptionTable[subscriptionId] = subscriptionCh
	// start go routine to handle subscription
	// for _, eventSub := range cdafCreateEventSubscription.Subscription {
	go handleSubscriptionEvent(ctx,
		cdafCreateEventSubscription.Subscription,
		cdafCreateEventSubscription.Subscription.EventNotifyUri,
		subscriptionId,
		subscriptionCh,
	)
	// }
	// Add location header in http response
	// respHeaders := make(map[string][]string)
	// respHeaders["Location"] = []string{
	// 	config.Events.Uri + urlBasePath + "/" + subscriptionId,
	// }
	// TODO - implement FailEventReports when events aren't accepted
	eventSubInfo := CdafCreateEventSubscription{
		Subscription: cdafCreateEventSubscription.Subscription,
	}
	// return ResponseWithHeaders(201, respHeaders, eventSubInfo), nil
	return Response(201, eventSubInfo), errors.New("CreateSubscription method not implemented")
	// return Response(http.StatusNotImplemented, nil), errors.New("CreateSubscription method not implemented")
}

// ------------------------------------------------------------------------------
// send only accepted events to handle subscription after feasibility check
func handleSubscriptionEvent(
	ctx context.Context,
	eventSub CdafEventSubscription,
	notificationURI string,
	subscriptionId string,
	subscriptionCh <-chan string,
) {
	log.Print("Handling subscription to ", eventSub.EventRequest.Type,
		" with subscription id ", subscriptionId)
loop:
	for {
		// check if channel is closed, break loop if true.
		select {
		case <-subscriptionCh:
			break loop

		default:
			switch eventSub.Options.Trigger {

			case EVENTTRIGGER_PERIODIC:
				// fill the event subscription information with data from mongoDB
				eventNotif, err := fillEventNotification(ctx, eventSub)
				if err != nil {
					log.Print(err)
					break loop
				}
				// send notification to client
				err = sendNotification(ctx, eventNotif, notificationURI)
				if err != nil {
					log.Print(err)
					break loop
				}
				// sleep periodically
				time.Sleep(time.Duration(eventSub.Options.RepPeriod) * time.Second)

			default:
				// TODO - implement THRESHOLD case
				log.Print("Not implemented yet")
				break loop
			}
		}
	}
	log.Print("subscription to ", eventSub.EventRequest.Type,
		" with subscription id ", subscriptionId, " is closed.")
}

// fillEventNotification - return event notification information
func fillEventNotification(ctx context.Context,
	eventSub CdafEventSubscription,
) (CdafEventNotification, error) {
	// only NETWORK_PERFORMACE - NUM_OF_UE is implemented for the moment
	var eventNotif CdafEventNotification
	switch eventSub.EventRequest.Type {

	case CDAFEVENT_REPORT_RESOURCE_USAGE:

		nwPerfNotifData, err := getContainerData(eventSub)

		if err != nil {
			return eventNotif, err
		}
		eventNotif.ReportEvent.NfLoadLevelInfos = nwPerfNotifData
		// eventNotif.ReportEvent.NfLoadLevelInfos = []NfLoadLevelInformation{
		// 	NfType: NFTYPE_AMF,
		// 	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.
		// 	NfInstanceId: "3fa85f64-5717-4562-b3fc-2c963f66afa6",
		// 	// NF Set Identifier (see clause 28.12 of 3GPP TS 23.003), formatted as the following string \"set<Set ID>.<nftype>set.5gc.mnc<MNC>.mcc<MCC>\", or  \"set<SetID>.<NFType>set.5gc.nid<NID>.mnc<MNC>.mcc<MCC>\" with  <MCC> encoded as defined in clause 5.4.2 (\"Mcc\" data type definition)  <MNC> encoding the Mobile Network Code part of the PLMN, comprising 3 digits.    If there are only 2 significant digits in the MNC, one \"0\" digit shall be inserted    at the left side to fill the 3 digits coding of MNC.  Pattern: '^[0-9]{3}$' <NFType> encoded as a value defined in Table 6.1.6.3.3-1 of 3GPP TS 29.510 but    with lower case characters <Set ID> encoded as a string of characters consisting of    alphabetic characters (A-Z and a-z), digits (0-9) and/or the hyphen (-) and that    shall end with either an alphabetic character or a digit.
		// 	NfSetId:        "1235657",
		// 	NfCpuUsage:     15,
		// 	NfMemoryUsage:  10,
		// 	NfStorageUsage: 100,
		// }
	default:
		// Implement others
		log.Print("Not implemented yet")
	}
	eventNotif.ReportEvent.Type = eventSub.EventRequest.Type
	// eventNotif.AnaMetaInfo.DataWindow.StartTime = eventSub.ExtraReportReq.StartTs
	// eventNotif.AnaMetaInfo.DataWindow.StopTime = eventSub.ExtraReportReq.EndTs
	return eventNotif, nil
}

// ------------------------------------------------------------------------------
func sendNotification(
	ctx context.Context,
	eventNotif CdafEventNotification,
	notificationURI string,
) error {
	log.Print("Sending notification to client")
	jsonStr, _ := json.Marshal(eventNotif)
	_, err := http.Post(notificationURI, "application/json", bytes.NewBuffer(jsonStr))
	return err
}

// ------------------------------------------------------------------------------
// getNwPerfAnalytics - Get list of NetworkPerfInfo
func getContainerData(eventSub CdafEventSubscription) ([]NfLoadLevelInformation, error) {
	log.Printf("Getting NW Performance Notification Data")
	var nwPerfList []NfLoadLevelInformation
	// for _, nwPerfReq := range eventSub.NwPerfRequs {

	var nwPerfInfo NfLoadLevelInformation
	// var err error
	switch eventSub.EventRequest.Type {

	case CDAFEVENT_REPORT_RESOURCE_USAGE:
		nwPerfInfo.NfType = "NFTYPE_AMF"
		// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.
		nwPerfInfo.NfInstanceId = "3fa85f64-5717-4562-b3fc-2c963f66afa6"
		// NF Set Identifier (see clause 28.12 of 3GPP TS 23.003), formatted as the following string \"set<Set ID>.<nftype>set.5gc.mnc<MNC>.mcc<MCC>\", or  \"set<SetID>.<NFType>set.5gc.nid<NID>.mnc<MNC>.mcc<MCC>\" with  <MCC> encoded as defined in clause 5.4.2 (\"Mcc\" data type definition)  <MNC> encoding the Mobile Network Code part of the PLMN, comprising 3 digits.    If there are only 2 significant digits in the MNC, one \"0\" digit shall be inserted    at the left side to fill the 3 digits coding of MNC.  Pattern: '^[0-9]{3}$' <NFType> encoded as a value defined in Table 6.1.6.3.3-1 of 3GPP TS 29.510 but    with lower case characters <Set ID> encoded as a string of characters consisting of    alphabetic characters (A-Z and a-z), digits (0-9) and/or the hyphen (-) and that    shall end with either an alphabetic character or a digit.
		nwPerfInfo.NfSetId = "1235657"
		nwPerfInfo.NfCpuUsage = 15
		nwPerfInfo.NfMemoryUsage = 10
		nwPerfInfo.NfStorageUsage = 100
		// nwPerfInfo, err = requestNwPerfEngine(
		// 	eventSub,
		// 	config.Engine.Uri+config.Routes.NumOfUe,
		// )
		// if err != nil {
		// 	return nwPerfList, err
		// }

	default:
		// TODO - Implement other NwPerfTypes
		return nil, errors.New("invalid Network Performance Type")
	}
	// nwPerfInfo.NwPerfType = nwPerfReq.NwPerfType
	nwPerfList = append(nwPerfList, nwPerfInfo)
	// }
	return nwPerfList, nil
}
