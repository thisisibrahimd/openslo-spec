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

// BudgetAdjustmentKind the model 'BudgetAdjustmentKind'
type BudgetAdjustmentKind string

// List of budgetAdjustment_kind
const (
	BUDGET_ADJUSTMENT BudgetAdjustmentKind = "BudgetAdjustment"
)

// All allowed values of BudgetAdjustmentKind enum
var AllowedBudgetAdjustmentKindEnumValues = []BudgetAdjustmentKind{
	"BudgetAdjustment",
}

func (v *BudgetAdjustmentKind) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BudgetAdjustmentKind(value)
	for _, existing := range AllowedBudgetAdjustmentKindEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BudgetAdjustmentKind", value)
}

// NewBudgetAdjustmentKindFromValue returns a pointer to a valid BudgetAdjustmentKind
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBudgetAdjustmentKindFromValue(v string) (*BudgetAdjustmentKind, error) {
	ev := BudgetAdjustmentKind(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BudgetAdjustmentKind: valid values are %v", v, AllowedBudgetAdjustmentKindEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BudgetAdjustmentKind) IsValid() bool {
	for _, existing := range AllowedBudgetAdjustmentKindEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to budgetAdjustment_kind value
func (v BudgetAdjustmentKind) Ptr() *BudgetAdjustmentKind {
	return &v
}

type NullableBudgetAdjustmentKind struct {
	value *BudgetAdjustmentKind
	isSet bool
}

func (v NullableBudgetAdjustmentKind) Get() *BudgetAdjustmentKind {
	return v.value
}

func (v *NullableBudgetAdjustmentKind) Set(val *BudgetAdjustmentKind) {
	v.value = val
	v.isSet = true
}

func (v NullableBudgetAdjustmentKind) IsSet() bool {
	return v.isSet
}

func (v *NullableBudgetAdjustmentKind) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBudgetAdjustmentKind(val *BudgetAdjustmentKind) *NullableBudgetAdjustmentKind {
	return &NullableBudgetAdjustmentKind{value: val, isSet: true}
}

func (v NullableBudgetAdjustmentKind) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBudgetAdjustmentKind) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
