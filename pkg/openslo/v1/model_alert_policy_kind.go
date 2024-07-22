/*
OpenSLO V1 Schema

OpenSLO V1 Schema

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openslo_v1

import (
	"encoding/json"
	"fmt"
)

// AlertPolicyKind the model 'AlertPolicyKind'
type AlertPolicyKind string

// List of alertPolicy_kind
const (
	ALERT_POLICY AlertPolicyKind = "AlertPolicy"
)

// All allowed values of AlertPolicyKind enum
var AllowedAlertPolicyKindEnumValues = []AlertPolicyKind{
	"AlertPolicy",
}

func (v *AlertPolicyKind) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AlertPolicyKind(value)
	for _, existing := range AllowedAlertPolicyKindEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AlertPolicyKind", value)
}

// NewAlertPolicyKindFromValue returns a pointer to a valid AlertPolicyKind
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAlertPolicyKindFromValue(v string) (*AlertPolicyKind, error) {
	ev := AlertPolicyKind(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AlertPolicyKind: valid values are %v", v, AllowedAlertPolicyKindEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AlertPolicyKind) IsValid() bool {
	for _, existing := range AllowedAlertPolicyKindEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to alertPolicy_kind value
func (v AlertPolicyKind) Ptr() *AlertPolicyKind {
	return &v
}

type NullableAlertPolicyKind struct {
	value *AlertPolicyKind
	isSet bool
}

func (v NullableAlertPolicyKind) Get() *AlertPolicyKind {
	return v.value
}

func (v *NullableAlertPolicyKind) Set(val *AlertPolicyKind) {
	v.value = val
	v.isSet = true
}

func (v NullableAlertPolicyKind) IsSet() bool {
	return v.isSet
}

func (v *NullableAlertPolicyKind) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlertPolicyKind(val *AlertPolicyKind) *NullableAlertPolicyKind {
	return &NullableAlertPolicyKind{value: val, isSet: true}
}

func (v NullableAlertPolicyKind) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlertPolicyKind) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
