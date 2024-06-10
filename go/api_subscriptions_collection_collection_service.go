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
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

var config MainConfig

type MainConfig struct {
	ExtractMetrics struct {
		Cadvisor      string `envconfig:"PROM_URL_CADVISOR"`
		Prometheus    string `envconfig:"PROM_URL_PROMETHEUS"`
		CpuPerPod     string `envconfig:"USED_CPU_PER_POD"`
		UsedRamPerPod string `envconfig:"USED_RAM_IN_BYTES_PER_POD"`
		Kube          bool   `envconfig:"KUBE"`
	}
}

// const KUBE = true

// const PROM_URL_CADVISOR = "http://192.168.70.1:8010/api/v1.3/docker/"

// const PROM_URL_PROMETHEUS = "http://192.168.14.139:30090/"
// const USED_CPU_PER_POD = "sum(eagle_pod_container_resource_usage_cpu_cores) by (pod, container, node, namespace, phase)"             //CPU cores in use by a specific container
// const USED_RAM_IN_BYTES_PER_POD = "sum(eagle_pod_container_resource_usage_memory_bytes) by (pod, container, node, namespace, phase)" //RAM bytes in use by a specific container

type ContainerInfo struct {
	Id      string   `json:"id"`
	Aliases []string `json:"aliases"`
	Name    string   `json:"name"`
	Stats   []Stats  `json:"stats"`
}

type Stats struct {
	Cpu    Cpu    `json:"cpu"`
	Memory Memory `json:"memory"`
}

type Cpu struct {
	Usage       Usage  `json:"usage"`
	LoadAverage uint64 `json:"load_average"`
}

type Usage struct {
	Total uint64 `json:"total"`
}

type Memory struct {
	Usage uint64 `json:"usage"`
}

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

	// var nwPerfInfo NfLoadLevelInformation
	// var err error
	switch eventSub.EventRequest.Type {

	case CDAFEVENT_REPORT_RESOURCE_USAGE:
		// LOGIC FOR OBTAIN CONTAINER METRICS
		// nwPerfInfo.NfType = "NFTYPE_AMF"
		// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.
		// nwPerfInfo.NfInstanceId = "3fa85f64-5717-4562-b3fc-2c963f66afa6"
		// NF Set Identifier (see clause 28.12 of 3GPP TS 23.003), formatted as the following string \"set<Set ID>.<nftype>set.5gc.mnc<MNC>.mcc<MCC>\", or  \"set<SetID>.<NFType>set.5gc.nid<NID>.mnc<MNC>.mcc<MCC>\" with  <MCC> encoded as defined in clause 5.4.2 (\"Mcc\" data type definition)  <MNC> encoding the Mobile Network Code part of the PLMN, comprising 3 digits.    If there are only 2 significant digits in the MNC, one \"0\" digit shall be inserted    at the left side to fill the 3 digits coding of MNC.  Pattern: '^[0-9]{3}$' <NFType> encoded as a value defined in Table 6.1.6.3.3-1 of 3GPP TS 29.510 but    with lower case characters <Set ID> encoded as a string of characters consisting of    alphabetic characters (A-Z and a-z), digits (0-9) and/or the hyphen (-) and that    shall end with either an alphabetic character or a digit.
		// nwPerfInfo.NfSetId = "1235657"
		// nwPerfInfo.NfCpuUsage = 15
		// nwPerfInfo.NfMemoryUsage = 10
		// nwPerfInfo.NfStorageUsage = 100
		// nwPerfInfo, err = requestNwPerfEngine(
		// 	eventSub,
		// 	config.Engine.Uri+config.Routes.NumOfUe,
		// )
		// if err != nil {
		// 	return nwPerfList, err
		// }
		if config.ExtractMetrics.Kube {
			use_cpu_per_pod := sendQuery(config.ExtractMetrics.Prometheus, config.ExtractMetrics.CpuPerPod, "pod/container/namespace/node/phase")     // Returns an array with [[podName, containerName, namespaceName, nodeName, status, used_cpu],[podName, containerName, namespaceName, nodeName, status, used_cpu]...])
			use_ram_per_pod := sendQuery(config.ExtractMetrics.Prometheus, config.ExtractMetrics.UsedRamPerPod, "pod/container/namespace/node/phase") // Returns an array with [[podName, containerName, namespaceName, nodeName, status, used_ram],[podName, containerName, namespaceName, nodeName, status, used_ram]...]

			nwPerfList = getMetricsForNWDAF(use_cpu_per_pod, use_ram_per_pod)
		} else {
			nwPerfList = getContainerMetrics()
		}

	default:
		// TODO - Implement other NwPerfTypes
		return nil, errors.New("invalid Network Performance Type")
	}
	// nwPerfInfo.NwPerfType = nwPerfReq.NwPerfType
	// nwPerfList = append(nwPerfList, nwPerfInfo)
	// }
	return nwPerfList, nil
}

func getContainerMetrics() []NfLoadLevelInformation {
	resp, err := http.Get(config.ExtractMetrics.Cadvisor)
	log.Printf("URL: %s\n", config.ExtractMetrics.Cadvisor)
	if err != nil {
		fmt.Printf("Error occurred while fetching container metrics: %v\n", err)
		return nil
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil
	}

	var response map[string]ContainerInfo
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil
	}

	var containerData []NfLoadLevelInformation

	for _, containerInfo := range response {
		containerId := containerInfo.Id
		containerName := ""
		containerType := "AF"
		if len(containerInfo.Aliases) > 0 {
			containerName = containerInfo.Aliases[0]
		} else {
			containerName = containerInfo.Name
		}
		cpuUsage := containerInfo.Stats[0].Cpu.Usage.Total
		memoryUsage := containerInfo.Stats[0].Memory.Usage
		//cpuLoad := containerInfo.Stats[0].Cpu.LoadAverage

		if strings.Contains(containerName, "nrf") {
			containerType = "NRF"
		} else if strings.Contains(containerName, "udm") {
			containerType = "UDM"
		} else if strings.Contains(containerName, "amf") {
			containerType = "AMF"
		} else if strings.Contains(containerName, "smf") {
			containerType = "SMF"
		} else if strings.Contains(containerName, "ausf") {
			containerType = "AUSF"
		} else if strings.Contains(containerName, "udr") {
			containerType = "UDR"
		} else if strings.Contains(containerName, "upf") {
			containerType = "UPF"
		} else if strings.Contains(containerName, "nwdaf") {
			containerType = "NWDAF"
		}

		//fmt.Println(containerInfo)
		item := NfLoadLevelInformation{
			NfType:        NfType(containerType),
			NfInstanceId:  containerId,
			NfSetId:       containerName,
			NfCpuUsage:    int32(cpuUsage),
			NfMemoryUsage: int32(memoryUsage),
			//"nfStorageUsage": cpuLoad,
		}

		containerData = append(containerData, item)
	}

	return containerData
}

func sendQuery(address, query, arguments string) [][]string {
	var res [][]string

	response, err := http.Get(address + "/api/v1/query?query=" + url.QueryEscape(query))
	if err != nil {
		return nil
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil
	}

	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil
	}

	results := data["data"].(map[string]interface{})["result"].([]interface{})

	if len(results) == 0 {
		return [][]string{{"0"}}
	}

	for _, item := range results {
		var itemVal []string
		arg := strings.Split(arguments, "/")
		for _, a := range arg {
			val := item.(map[string]interface{})["metric"].(map[string]interface{})[a]
			itemVal = append(itemVal, val.(string))
		}
		val := item.(map[string]interface{})["value"].([]interface{})[1].(string)
		itemVal = append(itemVal, val)
		res = append(res, itemVal)
	}

	return res
}

func getMetricsForNWDAF(use_cpu_per_pod [][]string, use_ram_per_pod [][]string) []NfLoadLevelInformation {
	var containerData []NfLoadLevelInformation
	containerType := "AF"

	for i := range use_cpu_per_pod {

		for j := range use_ram_per_pod {

			if (use_cpu_per_pod[i][1] == use_ram_per_pod[j][1]) && (use_cpu_per_pod[i][0] == use_ram_per_pod[j][0]) {
				nfCpuUsage, err := strconv.Atoi(use_cpu_per_pod[i][5])
				if err != nil {
					fmt.Println("Error converting nfCpuUsage string to int:", err)
					return nil
				}
				nfMemoryUsage, err := strconv.Atoi(use_ram_per_pod[j][5])
				if err != nil {
					fmt.Println("Error converting nfMemoryUsage string to int:", err)
					return nil
				}
				item := NfLoadLevelInformation{
					NfType:        NfType(containerType),
					NfInstanceId:  use_cpu_per_pod[i][0],
					NfSetId:       use_cpu_per_pod[i][1],
					NfCpuUsage:    int32(nfCpuUsage),
					NfMemoryUsage: int32(nfMemoryUsage),
					//"nfStorageUsage": cpuLoad,
				}

				containerData = append(containerData, item)
			}

		}
	}

	return containerData
}
