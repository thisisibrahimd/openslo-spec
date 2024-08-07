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

// AlertNotificationTargetKind the model 'AlertNotificationTargetKind'
type AlertNotificationTargetKind string

// List of alertNotificationTarget_kind
const (
	ALERT_NOTIFICATION_TARGET AlertNotificationTargetKind = "AlertNotificationTarget"
)

// All allowed values of AlertNotificationTargetKind enum
var AllowedAlertNotificationTargetKindEnumValues = []AlertNotificationTargetKind{
	"AlertNotificationTarget",
}

func (v *AlertNotificationTargetKind) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AlertNotificationTargetKind(value)
	for _, existing := range AllowedAlertNotificationTargetKindEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AlertNotificationTargetKind", value)
}

// NewAlertNotificationTargetKindFromValue returns a pointer to a valid AlertNotificationTargetKind
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAlertNotificationTargetKindFromValue(v string) (*AlertNotificationTargetKind, error) {
	ev := AlertNotificationTargetKind(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AlertNotificationTargetKind: valid values are %v", v, AllowedAlertNotificationTargetKindEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AlertNotificationTargetKind) IsValid() bool {
	for _, existing := range AllowedAlertNotificationTargetKindEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to alertNotificationTarget_kind value
func (v AlertNotificationTargetKind) Ptr() *AlertNotificationTargetKind {
	return &v
}

type NullableAlertNotificationTargetKind struct {
	value *AlertNotificationTargetKind
	isSet bool
}

func (v NullableAlertNotificationTargetKind) Get() *AlertNotificationTargetKind {
	return v.value
}

func (v *NullableAlertNotificationTargetKind) Set(val *AlertNotificationTargetKind) {
	v.value = val
	v.isSet = true
}

func (v NullableAlertNotificationTargetKind) IsSet() bool {
	return v.isSet
}

func (v *NullableAlertNotificationTargetKind) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlertNotificationTargetKind(val *AlertNotificationTargetKind) *NullableAlertNotificationTargetKind {
	return &NullableAlertNotificationTargetKind{value: val, isSet: true}
}

func (v NullableAlertNotificationTargetKind) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlertNotificationTargetKind) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
