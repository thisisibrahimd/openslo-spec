/*
OpenSLO V1 Schema

OpenSLO V1 Schema

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openslo_v1

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the BudgetAdjustmentSpec type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BudgetAdjustmentSpec{}

// BudgetAdjustmentSpec struct for BudgetAdjustmentSpec
type BudgetAdjustmentSpec struct {
	Description *string `json:"description,omitempty"`
	Service string `json:"service" validate:"regexp=^[a-z0-9][a-z0-9.|\\/\\\\\\\\-]*[a-z0-9]*$"`
	IndicatorRef string `json:"indicatorRef" validate:"regexp=^[a-z0-9][a-z0-9.|\\/\\\\\\\\-]*[a-z0-9]*$"`
	StartTime time.Time `json:"startTime"`
	EndTime *string `json:"endTime,omitempty"`
	Duration *string `json:"duration,omitempty"`
}

type _BudgetAdjustmentSpec BudgetAdjustmentSpec

// NewBudgetAdjustmentSpec instantiates a new BudgetAdjustmentSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBudgetAdjustmentSpec(service string, indicatorRef string, startTime time.Time) *BudgetAdjustmentSpec {
	this := BudgetAdjustmentSpec{}
	this.Service = service
	this.IndicatorRef = indicatorRef
	this.StartTime = startTime
	return &this
}

// NewBudgetAdjustmentSpecWithDefaults instantiates a new BudgetAdjustmentSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBudgetAdjustmentSpecWithDefaults() *BudgetAdjustmentSpec {
	this := BudgetAdjustmentSpec{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BudgetAdjustmentSpec) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BudgetAdjustmentSpec) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BudgetAdjustmentSpec) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BudgetAdjustmentSpec) SetDescription(v string) {
	o.Description = &v
}

// GetService returns the Service field value
func (o *BudgetAdjustmentSpec) GetService() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Service
}

// GetServiceOk returns a tuple with the Service field value
// and a boolean to check if the value has been set.
func (o *BudgetAdjustmentSpec) GetServiceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Service, true
}

// SetService sets field value
func (o *BudgetAdjustmentSpec) SetService(v string) {
	o.Service = v
}

// GetIndicatorRef returns the IndicatorRef field value
func (o *BudgetAdjustmentSpec) GetIndicatorRef() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IndicatorRef
}

// GetIndicatorRefOk returns a tuple with the IndicatorRef field value
// and a boolean to check if the value has been set.
func (o *BudgetAdjustmentSpec) GetIndicatorRefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IndicatorRef, true
}

// SetIndicatorRef sets field value
func (o *BudgetAdjustmentSpec) SetIndicatorRef(v string) {
	o.IndicatorRef = v
}

// GetStartTime returns the StartTime field value
func (o *BudgetAdjustmentSpec) GetStartTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value
// and a boolean to check if the value has been set.
func (o *BudgetAdjustmentSpec) GetStartTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartTime, true
}

// SetStartTime sets field value
func (o *BudgetAdjustmentSpec) SetStartTime(v time.Time) {
	o.StartTime = v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *BudgetAdjustmentSpec) GetEndTime() string {
	if o == nil || IsNil(o.EndTime) {
		var ret string
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BudgetAdjustmentSpec) GetEndTimeOk() (*string, bool) {
	if o == nil || IsNil(o.EndTime) {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *BudgetAdjustmentSpec) HasEndTime() bool {
	if o != nil && !IsNil(o.EndTime) {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given string and assigns it to the EndTime field.
func (o *BudgetAdjustmentSpec) SetEndTime(v string) {
	o.EndTime = &v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *BudgetAdjustmentSpec) GetDuration() string {
	if o == nil || IsNil(o.Duration) {
		var ret string
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BudgetAdjustmentSpec) GetDurationOk() (*string, bool) {
	if o == nil || IsNil(o.Duration) {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *BudgetAdjustmentSpec) HasDuration() bool {
	if o != nil && !IsNil(o.Duration) {
		return true
	}

	return false
}

// SetDuration gets a reference to the given string and assigns it to the Duration field.
func (o *BudgetAdjustmentSpec) SetDuration(v string) {
	o.Duration = &v
}

func (o BudgetAdjustmentSpec) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BudgetAdjustmentSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["service"] = o.Service
	toSerialize["indicatorRef"] = o.IndicatorRef
	toSerialize["startTime"] = o.StartTime
	if !IsNil(o.EndTime) {
		toSerialize["endTime"] = o.EndTime
	}
	if !IsNil(o.Duration) {
		toSerialize["duration"] = o.Duration
	}
	return toSerialize, nil
}

func (o *BudgetAdjustmentSpec) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"service",
		"indicatorRef",
		"startTime",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varBudgetAdjustmentSpec := _BudgetAdjustmentSpec{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBudgetAdjustmentSpec)

	if err != nil {
		return err
	}

	*o = BudgetAdjustmentSpec(varBudgetAdjustmentSpec)

	return err
}

type NullableBudgetAdjustmentSpec struct {
	value *BudgetAdjustmentSpec
	isSet bool
}

func (v NullableBudgetAdjustmentSpec) Get() *BudgetAdjustmentSpec {
	return v.value
}

func (v *NullableBudgetAdjustmentSpec) Set(val *BudgetAdjustmentSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableBudgetAdjustmentSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableBudgetAdjustmentSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBudgetAdjustmentSpec(val *BudgetAdjustmentSpec) *NullableBudgetAdjustmentSpec {
	return &NullableBudgetAdjustmentSpec{value: val, isSet: true}
}

func (v NullableBudgetAdjustmentSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBudgetAdjustmentSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

