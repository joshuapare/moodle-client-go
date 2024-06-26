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

// checks if the CoreReportbuilderConditionsDelete200ResponseAvailableconditionsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreReportbuilderConditionsDelete200ResponseAvailableconditionsInner{}

// CoreReportbuilderConditionsDelete200ResponseAvailableconditionsInner struct for CoreReportbuilderConditionsDelete200ResponseAvailableconditionsInner
type CoreReportbuilderConditionsDelete200ResponseAvailableconditionsInner struct {
	Optiongroup *CoreReportbuilderConditionsDelete200ResponseAvailableconditionsInnerOptiongroup `json:"optiongroup,omitempty"`
}

// NewCoreReportbuilderConditionsDelete200ResponseAvailableconditionsInner instantiates a new CoreReportbuilderConditionsDelete200ResponseAvailableconditionsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreReportbuilderConditionsDelete200ResponseAvailableconditionsInner() *CoreReportbuilderConditionsDelete200ResponseAvailableconditionsInner {
	this := CoreReportbuilderConditionsDelete200ResponseAvailableconditionsInner{}
	return &this
}

// NewCoreReportbuilderConditionsDelete200ResponseAvailableconditionsInnerWithDefaults instantiates a new CoreReportbuilderConditionsDelete200ResponseAvailableconditionsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreReportbuilderConditionsDelete200ResponseAvailableconditionsInnerWithDefaults() *CoreReportbuilderConditionsDelete200ResponseAvailableconditionsInner {
	this := CoreReportbuilderConditionsDelete200ResponseAvailableconditionsInner{}
	return &this
}

// GetOptiongroup returns the Optiongroup field value if set, zero value otherwise.
func (o *CoreReportbuilderConditionsDelete200ResponseAvailableconditionsInner) GetOptiongroup() CoreReportbuilderConditionsDelete200ResponseAvailableconditionsInnerOptiongroup {
	if o == nil || IsNil(o.Optiongroup) {
		var ret CoreReportbuilderConditionsDelete200ResponseAvailableconditionsInnerOptiongroup
		return ret
	}
	return *o.Optiongroup
}

// GetOptiongroupOk returns a tuple with the Optiongroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreReportbuilderConditionsDelete200ResponseAvailableconditionsInner) GetOptiongroupOk() (*CoreReportbuilderConditionsDelete200ResponseAvailableconditionsInnerOptiongroup, bool) {
	if o == nil || IsNil(o.Optiongroup) {
		return nil, false
	}
	return o.Optiongroup, true
}

// HasOptiongroup returns a boolean if a field has been set.
func (o *CoreReportbuilderConditionsDelete200ResponseAvailableconditionsInner) HasOptiongroup() bool {
	if o != nil && !IsNil(o.Optiongroup) {
		return true
	}

	return false
}

// SetOptiongroup gets a reference to the given CoreReportbuilderConditionsDelete200ResponseAvailableconditionsInnerOptiongroup and assigns it to the Optiongroup field.
func (o *CoreReportbuilderConditionsDelete200ResponseAvailableconditionsInner) SetOptiongroup(v CoreReportbuilderConditionsDelete200ResponseAvailableconditionsInnerOptiongroup) {
	o.Optiongroup = &v
}

func (o CoreReportbuilderConditionsDelete200ResponseAvailableconditionsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreReportbuilderConditionsDelete200ResponseAvailableconditionsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Optiongroup) {
		toSerialize["optiongroup"] = o.Optiongroup
	}
	return toSerialize, nil
}

type NullableCoreReportbuilderConditionsDelete200ResponseAvailableconditionsInner struct {
	value *CoreReportbuilderConditionsDelete200ResponseAvailableconditionsInner
	isSet bool
}

func (v NullableCoreReportbuilderConditionsDelete200ResponseAvailableconditionsInner) Get() *CoreReportbuilderConditionsDelete200ResponseAvailableconditionsInner {
	return v.value
}

func (v *NullableCoreReportbuilderConditionsDelete200ResponseAvailableconditionsInner) Set(val *CoreReportbuilderConditionsDelete200ResponseAvailableconditionsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreReportbuilderConditionsDelete200ResponseAvailableconditionsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreReportbuilderConditionsDelete200ResponseAvailableconditionsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreReportbuilderConditionsDelete200ResponseAvailableconditionsInner(val *CoreReportbuilderConditionsDelete200ResponseAvailableconditionsInner) *NullableCoreReportbuilderConditionsDelete200ResponseAvailableconditionsInner {
	return &NullableCoreReportbuilderConditionsDelete200ResponseAvailableconditionsInner{value: val, isSet: true}
}

func (v NullableCoreReportbuilderConditionsDelete200ResponseAvailableconditionsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreReportbuilderConditionsDelete200ResponseAvailableconditionsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


