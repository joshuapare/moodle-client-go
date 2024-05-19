/*
Moodle Webservice API

Auto-generated OpenAPI spec for Moodle's Webservice API. Requires installation of the Catalyst RESTful API plugin: https://github.com/catalyst/moodle-webservice_restful.

API version: 4.3.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package moodleclient

import (
	"encoding/json"
)

// checks if the CoreReportbuilderCanViewSystemReportRequestParametersInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreReportbuilderCanViewSystemReportRequestParametersInner{}

// CoreReportbuilderCanViewSystemReportRequestParametersInner struct for CoreReportbuilderCanViewSystemReportRequestParametersInner
type CoreReportbuilderCanViewSystemReportRequestParametersInner struct {
	Name *string `json:"name,omitempty"`
	Value *string `json:"value,omitempty"`
}

// NewCoreReportbuilderCanViewSystemReportRequestParametersInner instantiates a new CoreReportbuilderCanViewSystemReportRequestParametersInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreReportbuilderCanViewSystemReportRequestParametersInner() *CoreReportbuilderCanViewSystemReportRequestParametersInner {
	this := CoreReportbuilderCanViewSystemReportRequestParametersInner{}
	return &this
}

// NewCoreReportbuilderCanViewSystemReportRequestParametersInnerWithDefaults instantiates a new CoreReportbuilderCanViewSystemReportRequestParametersInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreReportbuilderCanViewSystemReportRequestParametersInnerWithDefaults() *CoreReportbuilderCanViewSystemReportRequestParametersInner {
	this := CoreReportbuilderCanViewSystemReportRequestParametersInner{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CoreReportbuilderCanViewSystemReportRequestParametersInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreReportbuilderCanViewSystemReportRequestParametersInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CoreReportbuilderCanViewSystemReportRequestParametersInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CoreReportbuilderCanViewSystemReportRequestParametersInner) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CoreReportbuilderCanViewSystemReportRequestParametersInner) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreReportbuilderCanViewSystemReportRequestParametersInner) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CoreReportbuilderCanViewSystemReportRequestParametersInner) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *CoreReportbuilderCanViewSystemReportRequestParametersInner) SetValue(v string) {
	o.Value = &v
}

func (o CoreReportbuilderCanViewSystemReportRequestParametersInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreReportbuilderCanViewSystemReportRequestParametersInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableCoreReportbuilderCanViewSystemReportRequestParametersInner struct {
	value *CoreReportbuilderCanViewSystemReportRequestParametersInner
	isSet bool
}

func (v NullableCoreReportbuilderCanViewSystemReportRequestParametersInner) Get() *CoreReportbuilderCanViewSystemReportRequestParametersInner {
	return v.value
}

func (v *NullableCoreReportbuilderCanViewSystemReportRequestParametersInner) Set(val *CoreReportbuilderCanViewSystemReportRequestParametersInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreReportbuilderCanViewSystemReportRequestParametersInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreReportbuilderCanViewSystemReportRequestParametersInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreReportbuilderCanViewSystemReportRequestParametersInner(val *CoreReportbuilderCanViewSystemReportRequestParametersInner) *NullableCoreReportbuilderCanViewSystemReportRequestParametersInner {
	return &NullableCoreReportbuilderCanViewSystemReportRequestParametersInner{value: val, isSet: true}
}

func (v NullableCoreReportbuilderCanViewSystemReportRequestParametersInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreReportbuilderCanViewSystemReportRequestParametersInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


