/*
 * Ncdaf_EventExposure
 *
 * CDAF Event Exposure Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type NfTypeAnyOf string

// List of NfTypeAnyOf
const (
	NRF NfTypeAnyOf = "NRF"
	UDM NfTypeAnyOf = "UDM"
	AMF NfTypeAnyOf = "AMF"
	SMF NfTypeAnyOf = "SMF"
	AUSF NfTypeAnyOf = "AUSF"
	NEF NfTypeAnyOf = "NEF"
	PCF NfTypeAnyOf = "PCF"
	SMSF NfTypeAnyOf = "SMSF"
	NSSF NfTypeAnyOf = "NSSF"
	UDR NfTypeAnyOf = "UDR"
	LMF NfTypeAnyOf = "LMF"
	GMLC NfTypeAnyOf = "GMLC"
	_5_G_EIR NfTypeAnyOf = "5G_EIR"
	SEPP NfTypeAnyOf = "SEPP"
	UPF NfTypeAnyOf = "UPF"
	N3_IWF NfTypeAnyOf = "N3IWF"
	AF NfTypeAnyOf = "AF"
	UDSF NfTypeAnyOf = "UDSF"
	BSF NfTypeAnyOf = "BSF"
	CHF NfTypeAnyOf = "CHF"
	NWDAF NfTypeAnyOf = "NWDAF"
	PCSCF NfTypeAnyOf = "PCSCF"
	CBCF NfTypeAnyOf = "CBCF"
	HSS NfTypeAnyOf = "HSS"
	UCMF NfTypeAnyOf = "UCMF"
	SOR_AF NfTypeAnyOf = "SOR_AF"
	SPAF NfTypeAnyOf = "SPAF"
	MME NfTypeAnyOf = "MME"
	SCSAS NfTypeAnyOf = "SCSAS"
	SCEF NfTypeAnyOf = "SCEF"
	SCP NfTypeAnyOf = "SCP"
	NSSAAF NfTypeAnyOf = "NSSAAF"
	ICSCF NfTypeAnyOf = "ICSCF"
	SCSCF NfTypeAnyOf = "SCSCF"
	DRA NfTypeAnyOf = "DRA"
	IMS_AS NfTypeAnyOf = "IMS_AS"
	AANF NfTypeAnyOf = "AANF"
	_5_G_DDNMF NfTypeAnyOf = "5G_DDNMF"
	NSACF NfTypeAnyOf = "NSACF"
	MFAF NfTypeAnyOf = "MFAF"
	EASDF NfTypeAnyOf = "EASDF"
	DCCF NfTypeAnyOf = "DCCF"
	MB_SMF NfTypeAnyOf = "MB_SMF"
	TSCTSF NfTypeAnyOf = "TSCTSF"
	ADRF NfTypeAnyOf = "ADRF"
	GBA_BSF NfTypeAnyOf = "GBA_BSF"
	CEF NfTypeAnyOf = "CEF"
	MB_UPF NfTypeAnyOf = "MB_UPF"
	NSWOF NfTypeAnyOf = "NSWOF"
	PKMF NfTypeAnyOf = "PKMF"
	MNPF NfTypeAnyOf = "MNPF"
	SMS_GMSC NfTypeAnyOf = "SMS_GMSC"
	SMS_IWMSC NfTypeAnyOf = "SMS_IWMSC"
	MBSF NfTypeAnyOf = "MBSF"
	MBSTF NfTypeAnyOf = "MBSTF"
	PANF NfTypeAnyOf = "PANF"
	IP_SM_GW NfTypeAnyOf = "IP_SM_GW"
	SMS_ROUTER NfTypeAnyOf = "SMS_ROUTER"
)

// AssertNfTypeAnyOfRequired checks if the required fields are not zero-ed
func AssertNfTypeAnyOfRequired(obj NfTypeAnyOf) error {
	return nil
}

// AssertRecurseNfTypeAnyOfRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of NfTypeAnyOf (e.g. [][]NfTypeAnyOf), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseNfTypeAnyOfRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aNfTypeAnyOf, ok := obj.(NfTypeAnyOf)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertNfTypeAnyOfRequired(aNfTypeAnyOf)
	})
}
