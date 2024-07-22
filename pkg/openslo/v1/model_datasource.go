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

// checks if the Datasource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Datasource{}

// Datasource struct for Datasource
type Datasource struct {
	ApiVersion OpensloApiVersion `json:"apiVersion"`
	Kind       string            `json:"kind"`
	Metadata   Metadata          `json:"metadata"`
	Spec       *ServiceSpec      `json:"spec,omitempty"`
}

type _Datasource Datasource

// NewDatasource instantiates a new Datasource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatasource(apiVersion OpensloApiVersion, kind string, metadata Metadata) *Datasource {
	this := Datasource{}
	this.ApiVersion = apiVersion
	this.Kind = kind
	this.Metadata = metadata
	return &this
}

// NewDatasourceWithDefaults instantiates a new Datasource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatasourceWithDefaults() *Datasource {
	this := Datasource{}
	return &this
}

// GetApiVersion returns the ApiVersion field value
func (o *Datasource) GetApiVersion() OpensloApiVersion {
	if o == nil {
		var ret OpensloApiVersion
		return ret
	}

	return o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value
// and a boolean to check if the value has been set.
func (o *Datasource) GetApiVersionOk() (*OpensloApiVersion, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiVersion, true
}

// SetApiVersion sets field value
func (o *Datasource) SetApiVersion(v OpensloApiVersion) {
	o.ApiVersion = v
}

// GetKind returns the Kind field value
func (o *Datasource) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *Datasource) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *Datasource) SetKind(v string) {
	o.Kind = v
}

// GetMetadata returns the Metadata field value
func (o *Datasource) GetMetadata() Metadata {
	if o == nil {
		var ret Metadata
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *Datasource) GetMetadataOk() (*Metadata, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *Datasource) SetMetadata(v Metadata) {
	o.Metadata = v
}

// GetSpec returns the Spec field value if set, zero value otherwise.
func (o *Datasource) GetSpec() ServiceSpec {
	if o == nil || IsNil(o.Spec) {
		var ret ServiceSpec
		return ret
	}
	return *o.Spec
}

// GetSpecOk returns a tuple with the Spec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Datasource) GetSpecOk() (*ServiceSpec, bool) {
	if o == nil || IsNil(o.Spec) {
		return nil, false
	}
	return o.Spec, true
}

// HasSpec returns a boolean if a field has been set.
func (o *Datasource) HasSpec() bool {
	if o != nil && !IsNil(o.Spec) {
		return true
	}

	return false
}

// SetSpec gets a reference to the given ServiceSpec and assigns it to the Spec field.
func (o *Datasource) SetSpec(v ServiceSpec) {
	o.Spec = &v
}

func (o Datasource) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Datasource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["apiVersion"] = o.ApiVersion
	toSerialize["kind"] = o.Kind
	toSerialize["metadata"] = o.Metadata
	if !IsNil(o.Spec) {
		toSerialize["spec"] = o.Spec
	}
	return toSerialize, nil
}

func (o *Datasource) UnmarshalJSON(data []byte) (err error) {
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

	varDatasource := _Datasource{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDatasource)

	if err != nil {
		return err
	}

	*o = Datasource(varDatasource)

	return err
}

type NullableDatasource struct {
	value *Datasource
	isSet bool
}

func (v NullableDatasource) Get() *Datasource {
	return v.value
}

func (v *NullableDatasource) Set(val *Datasource) {
	v.value = val
	v.isSet = true
}

func (v NullableDatasource) IsSet() bool {
	return v.isSet
}

func (v *NullableDatasource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatasource(val *Datasource) *NullableDatasource {
	return &NullableDatasource{value: val, isSet: true}
}

func (v NullableDatasource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatasource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
