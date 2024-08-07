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

// SliKind the model 'SliKind'
type SliKind string

// List of sli_kind
const (
	SLI SliKind = "SLI"
)

// All allowed values of SliKind enum
var AllowedSliKindEnumValues = []SliKind{
	"SLI",
}

func (v *SliKind) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SliKind(value)
	for _, existing := range AllowedSliKindEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SliKind", value)
}

// NewSliKindFromValue returns a pointer to a valid SliKind
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSliKindFromValue(v string) (*SliKind, error) {
	ev := SliKind(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SliKind: valid values are %v", v, AllowedSliKindEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SliKind) IsValid() bool {
	for _, existing := range AllowedSliKindEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to sli_kind value
func (v SliKind) Ptr() *SliKind {
	return &v
}

type NullableSliKind struct {
	value *SliKind
	isSet bool
}

func (v NullableSliKind) Get() *SliKind {
	return v.value
}

func (v *NullableSliKind) Set(val *SliKind) {
	v.value = val
	v.isSet = true
}

func (v NullableSliKind) IsSet() bool {
	return v.isSet
}

func (v *NullableSliKind) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSliKind(val *SliKind) *NullableSliKind {
	return &NullableSliKind{value: val, isSet: true}
}

func (v NullableSliKind) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSliKind) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
