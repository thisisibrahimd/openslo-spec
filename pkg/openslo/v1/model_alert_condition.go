/*
OpenSLO V1 Schema

OpenSLO V1 Schema

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openslo_v1

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the AlertCondition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlertCondition{}

// AlertCondition struct for AlertCondition
type AlertCondition struct {
	ApiVersion OpensloApiVersion   `json:"apiVersion"`
	Kind       string              `json:"kind"`
	Metadata   Metadata            `json:"metadata"`
	Spec       *AlertConditionSpec `json:"spec,omitempty"`
}

type _AlertCondition AlertCondition

// NewAlertCondition instantiates a new AlertCondition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlertCondition(apiVersion OpensloApiVersion, kind string, metadata Metadata) *AlertCondition {
	this := AlertCondition{}
	this.ApiVersion = apiVersion
	this.Kind = kind
	this.Metadata = metadata
	return &this
}

// NewAlertConditionWithDefaults instantiates a new AlertCondition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlertConditionWithDefaults() *AlertCondition {
	this := AlertCondition{}
	return &this
}

// GetApiVersion returns the ApiVersion field value
func (o *AlertCondition) GetApiVersion() OpensloApiVersion {
	if o == nil {
		var ret OpensloApiVersion
		return ret
	}

	return o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value
// and a boolean to check if the value has been set.
func (o *AlertCondition) GetApiVersionOk() (*OpensloApiVersion, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiVersion, true
}

// SetApiVersion sets field value
func (o *AlertCondition) SetApiVersion(v OpensloApiVersion) {
	o.ApiVersion = v
}

// GetKind returns the Kind field value
func (o *AlertCondition) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *AlertCondition) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *AlertCondition) SetKind(v string) {
	o.Kind = v
}

// GetMetadata returns the Metadata field value
func (o *AlertCondition) GetMetadata() Metadata {
	if o == nil {
		var ret Metadata
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *AlertCondition) GetMetadataOk() (*Metadata, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *AlertCondition) SetMetadata(v Metadata) {
	o.Metadata = v
}

// GetSpec returns the Spec field value if set, zero value otherwise.
func (o *AlertCondition) GetSpec() AlertConditionSpec {
	if o == nil || IsNil(o.Spec) {
		var ret AlertConditionSpec
		return ret
	}
	return *o.Spec
}

// GetSpecOk returns a tuple with the Spec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertCondition) GetSpecOk() (*AlertConditionSpec, bool) {
	if o == nil || IsNil(o.Spec) {
		return nil, false
	}
	return o.Spec, true
}

// HasSpec returns a boolean if a field has been set.
func (o *AlertCondition) HasSpec() bool {
	if o != nil && !IsNil(o.Spec) {
		return true
	}

	return false
}

// SetSpec gets a reference to the given AlertConditionSpec and assigns it to the Spec field.
func (o *AlertCondition) SetSpec(v AlertConditionSpec) {
	o.Spec = &v
}

func (o AlertCondition) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlertCondition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["apiVersion"] = o.ApiVersion
	toSerialize["kind"] = o.Kind
	toSerialize["metadata"] = o.Metadata
	if !IsNil(o.Spec) {
		toSerialize["spec"] = o.Spec
	}
	return toSerialize, nil
}

func (o *AlertCondition) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"apiVersion",
		"kind",
		"metadata",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varAlertCondition := _AlertCondition{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAlertCondition)

	if err != nil {
		return err
	}

	*o = AlertCondition(varAlertCondition)

	return err
}

type NullableAlertCondition struct {
	value *AlertCondition
	isSet bool
}

func (v NullableAlertCondition) Get() *AlertCondition {
	return v.value
}

func (v *NullableAlertCondition) Set(val *AlertCondition) {
	v.value = val
	v.isSet = true
}

func (v NullableAlertCondition) IsSet() bool {
	return v.isSet
}

func (v *NullableAlertCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlertCondition(val *AlertCondition) *NullableAlertCondition {
	return &NullableAlertCondition{value: val, isSet: true}
}

func (v NullableAlertCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlertCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
