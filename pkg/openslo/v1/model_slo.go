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

// checks if the Slo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Slo{}

// Slo struct for Slo
type Slo struct {
	ApiVersion OpensloApiVersion `json:"apiVersion"`
	Kind       SloKind           `json:"kind"`
	Metadata   Metadata          `json:"metadata"`
	Spec       SloSpec           `json:"spec"`
}

type _Slo Slo

// NewSlo instantiates a new Slo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSlo(apiVersion OpensloApiVersion, kind SloKind, metadata Metadata, spec SloSpec) *Slo {
	this := Slo{}
	this.ApiVersion = apiVersion
	this.Kind = kind
	this.Metadata = metadata
	this.Spec = spec
	return &this
}

// NewSloWithDefaults instantiates a new Slo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSloWithDefaults() *Slo {
	this := Slo{}
	return &this
}

// GetApiVersion returns the ApiVersion field value
func (o *Slo) GetApiVersion() OpensloApiVersion {
	if o == nil {
		var ret OpensloApiVersion
		return ret
	}

	return o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value
// and a boolean to check if the value has been set.
func (o *Slo) GetApiVersionOk() (*OpensloApiVersion, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiVersion, true
}

// SetApiVersion sets field value
func (o *Slo) SetApiVersion(v OpensloApiVersion) {
	o.ApiVersion = v
}

// GetKind returns the Kind field value
func (o *Slo) GetKind() SloKind {
	if o == nil {
		var ret SloKind
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *Slo) GetKindOk() (*SloKind, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *Slo) SetKind(v SloKind) {
	o.Kind = v
}

// GetMetadata returns the Metadata field value
func (o *Slo) GetMetadata() Metadata {
	if o == nil {
		var ret Metadata
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *Slo) GetMetadataOk() (*Metadata, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *Slo) SetMetadata(v Metadata) {
	o.Metadata = v
}

// GetSpec returns the Spec field value
func (o *Slo) GetSpec() SloSpec {
	if o == nil {
		var ret SloSpec
		return ret
	}

	return o.Spec
}

// GetSpecOk returns a tuple with the Spec field value
// and a boolean to check if the value has been set.
func (o *Slo) GetSpecOk() (*SloSpec, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Spec, true
}

// SetSpec sets field value
func (o *Slo) SetSpec(v SloSpec) {
	o.Spec = v
}

func (o Slo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Slo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["apiVersion"] = o.ApiVersion
	toSerialize["kind"] = o.Kind
	toSerialize["metadata"] = o.Metadata
	toSerialize["spec"] = o.Spec
	return toSerialize, nil
}

func (o *Slo) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"apiVersion",
		"kind",
		"metadata",
		"spec",
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

	varSlo := _Slo{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSlo)

	if err != nil {
		return err
	}

	*o = Slo(varSlo)

	return err
}

type NullableSlo struct {
	value *Slo
	isSet bool
}

func (v NullableSlo) Get() *Slo {
	return v.value
}

func (v *NullableSlo) Set(val *Slo) {
	v.value = val
	v.isSet = true
}

func (v NullableSlo) IsSet() bool {
	return v.isSet
}

func (v *NullableSlo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSlo(val *Slo) *NullableSlo {
	return &NullableSlo{value: val, isSet: true}
}

func (v NullableSlo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSlo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
