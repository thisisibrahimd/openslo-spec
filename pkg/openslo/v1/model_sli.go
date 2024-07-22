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

// checks if the Sli type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Sli{}

// Sli struct for Sli
type Sli struct {
	ApiVersion OpensloApiVersion `json:"apiVersion"`
	Kind       SliKind           `json:"kind"`
	Metadata   Metadata          `json:"metadata"`
	Spec       *SliSpec          `json:"spec,omitempty"`
}

type _Sli Sli

// NewSli instantiates a new Sli object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSli(apiVersion OpensloApiVersion, kind SliKind, metadata Metadata) *Sli {
	this := Sli{}
	this.ApiVersion = apiVersion
	this.Kind = kind
	this.Metadata = metadata
	return &this
}

// NewSliWithDefaults instantiates a new Sli object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSliWithDefaults() *Sli {
	this := Sli{}
	return &this
}

// GetApiVersion returns the ApiVersion field value
func (o *Sli) GetApiVersion() OpensloApiVersion {
	if o == nil {
		var ret OpensloApiVersion
		return ret
	}

	return o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value
// and a boolean to check if the value has been set.
func (o *Sli) GetApiVersionOk() (*OpensloApiVersion, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiVersion, true
}

// SetApiVersion sets field value
func (o *Sli) SetApiVersion(v OpensloApiVersion) {
	o.ApiVersion = v
}

// GetKind returns the Kind field value
func (o *Sli) GetKind() SliKind {
	if o == nil {
		var ret SliKind
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *Sli) GetKindOk() (*SliKind, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *Sli) SetKind(v SliKind) {
	o.Kind = v
}

// GetMetadata returns the Metadata field value
func (o *Sli) GetMetadata() Metadata {
	if o == nil {
		var ret Metadata
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *Sli) GetMetadataOk() (*Metadata, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *Sli) SetMetadata(v Metadata) {
	o.Metadata = v
}

// GetSpec returns the Spec field value if set, zero value otherwise.
func (o *Sli) GetSpec() SliSpec {
	if o == nil || IsNil(o.Spec) {
		var ret SliSpec
		return ret
	}
	return *o.Spec
}

// GetSpecOk returns a tuple with the Spec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Sli) GetSpecOk() (*SliSpec, bool) {
	if o == nil || IsNil(o.Spec) {
		return nil, false
	}
	return o.Spec, true
}

// HasSpec returns a boolean if a field has been set.
func (o *Sli) HasSpec() bool {
	if o != nil && !IsNil(o.Spec) {
		return true
	}

	return false
}

// SetSpec gets a reference to the given SliSpec and assigns it to the Spec field.
func (o *Sli) SetSpec(v SliSpec) {
	o.Spec = &v
}

func (o Sli) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Sli) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["apiVersion"] = o.ApiVersion
	toSerialize["kind"] = o.Kind
	toSerialize["metadata"] = o.Metadata
	if !IsNil(o.Spec) {
		toSerialize["spec"] = o.Spec
	}
	return toSerialize, nil
}

func (o *Sli) UnmarshalJSON(data []byte) (err error) {
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

	varSli := _Sli{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSli)

	if err != nil {
		return err
	}

	*o = Sli(varSli)

	return err
}

type NullableSli struct {
	value *Sli
	isSet bool
}

func (v NullableSli) Get() *Sli {
	return v.value
}

func (v *NullableSli) Set(val *Sli) {
	v.value = val
	v.isSet = true
}

func (v NullableSli) IsSet() bool {
	return v.isSet
}

func (v *NullableSli) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSli(val *Sli) *NullableSli {
	return &NullableSli{value: val, isSet: true}
}

func (v NullableSli) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSli) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
