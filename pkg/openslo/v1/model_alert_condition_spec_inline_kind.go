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

// AlertConditionSpecInlineKind the model 'AlertConditionSpecInlineKind'
type AlertConditionSpecInlineKind string

// List of AlertConditionSpecInline_kind
const (
	BURNRATE AlertConditionSpecInlineKind = "burnrate"
)

// All allowed values of AlertConditionSpecInlineKind enum
var AllowedAlertConditionSpecInlineKindEnumValues = []AlertConditionSpecInlineKind{
	"burnrate",
}

func (v *AlertConditionSpecInlineKind) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AlertConditionSpecInlineKind(value)
	for _, existing := range AllowedAlertConditionSpecInlineKindEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AlertConditionSpecInlineKind", value)
}

// NewAlertConditionSpecInlineKindFromValue returns a pointer to a valid AlertConditionSpecInlineKind
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAlertConditionSpecInlineKindFromValue(v string) (*AlertConditionSpecInlineKind, error) {
	ev := AlertConditionSpecInlineKind(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AlertConditionSpecInlineKind: valid values are %v", v, AllowedAlertConditionSpecInlineKindEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AlertConditionSpecInlineKind) IsValid() bool {
	for _, existing := range AllowedAlertConditionSpecInlineKindEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AlertConditionSpecInline_kind value
func (v AlertConditionSpecInlineKind) Ptr() *AlertConditionSpecInlineKind {
	return &v
}

type NullableAlertConditionSpecInlineKind struct {
	value *AlertConditionSpecInlineKind
	isSet bool
}

func (v NullableAlertConditionSpecInlineKind) Get() *AlertConditionSpecInlineKind {
	return v.value
}

func (v *NullableAlertConditionSpecInlineKind) Set(val *AlertConditionSpecInlineKind) {
	v.value = val
	v.isSet = true
}

func (v NullableAlertConditionSpecInlineKind) IsSet() bool {
	return v.isSet
}

func (v *NullableAlertConditionSpecInlineKind) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlertConditionSpecInlineKind(val *AlertConditionSpecInlineKind) *NullableAlertConditionSpecInlineKind {
	return &NullableAlertConditionSpecInlineKind{value: val, isSet: true}
}

func (v NullableAlertConditionSpecInlineKind) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlertConditionSpecInlineKind) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
