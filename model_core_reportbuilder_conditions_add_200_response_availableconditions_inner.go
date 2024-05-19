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

// checks if the CoreReportbuilderConditionsAdd200ResponseAvailableconditionsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreReportbuilderConditionsAdd200ResponseAvailableconditionsInner{}

// CoreReportbuilderConditionsAdd200ResponseAvailableconditionsInner struct for CoreReportbuilderConditionsAdd200ResponseAvailableconditionsInner
type CoreReportbuilderConditionsAdd200ResponseAvailableconditionsInner struct {
	Optiongroup *CoreReportbuilderConditionsAdd200ResponseAvailableconditionsInnerOptiongroup `json:"optiongroup,omitempty"`
}

// NewCoreReportbuilderConditionsAdd200ResponseAvailableconditionsInner instantiates a new CoreReportbuilderConditionsAdd200ResponseAvailableconditionsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreReportbuilderConditionsAdd200ResponseAvailableconditionsInner() *CoreReportbuilderConditionsAdd200ResponseAvailableconditionsInner {
	this := CoreReportbuilderConditionsAdd200ResponseAvailableconditionsInner{}
	return &this
}

// NewCoreReportbuilderConditionsAdd200ResponseAvailableconditionsInnerWithDefaults instantiates a new CoreReportbuilderConditionsAdd200ResponseAvailableconditionsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreReportbuilderConditionsAdd200ResponseAvailableconditionsInnerWithDefaults() *CoreReportbuilderConditionsAdd200ResponseAvailableconditionsInner {
	this := CoreReportbuilderConditionsAdd200ResponseAvailableconditionsInner{}
	return &this
}

// GetOptiongroup returns the Optiongroup field value if set, zero value otherwise.
func (o *CoreReportbuilderConditionsAdd200ResponseAvailableconditionsInner) GetOptiongroup() CoreReportbuilderConditionsAdd200ResponseAvailableconditionsInnerOptiongroup {
	if o == nil || IsNil(o.Optiongroup) {
		var ret CoreReportbuilderConditionsAdd200ResponseAvailableconditionsInnerOptiongroup
		return ret
	}
	return *o.Optiongroup
}

// GetOptiongroupOk returns a tuple with the Optiongroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreReportbuilderConditionsAdd200ResponseAvailableconditionsInner) GetOptiongroupOk() (*CoreReportbuilderConditionsAdd200ResponseAvailableconditionsInnerOptiongroup, bool) {
	if o == nil || IsNil(o.Optiongroup) {
		return nil, false
	}
	return o.Optiongroup, true
}

// HasOptiongroup returns a boolean if a field has been set.
func (o *CoreReportbuilderConditionsAdd200ResponseAvailableconditionsInner) HasOptiongroup() bool {
	if o != nil && !IsNil(o.Optiongroup) {
		return true
	}

	return false
}

// SetOptiongroup gets a reference to the given CoreReportbuilderConditionsAdd200ResponseAvailableconditionsInnerOptiongroup and assigns it to the Optiongroup field.
func (o *CoreReportbuilderConditionsAdd200ResponseAvailableconditionsInner) SetOptiongroup(v CoreReportbuilderConditionsAdd200ResponseAvailableconditionsInnerOptiongroup) {
	o.Optiongroup = &v
}

func (o CoreReportbuilderConditionsAdd200ResponseAvailableconditionsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreReportbuilderConditionsAdd200ResponseAvailableconditionsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Optiongroup) {
		toSerialize["optiongroup"] = o.Optiongroup
	}
	return toSerialize, nil
}

type NullableCoreReportbuilderConditionsAdd200ResponseAvailableconditionsInner struct {
	value *CoreReportbuilderConditionsAdd200ResponseAvailableconditionsInner
	isSet bool
}

func (v NullableCoreReportbuilderConditionsAdd200ResponseAvailableconditionsInner) Get() *CoreReportbuilderConditionsAdd200ResponseAvailableconditionsInner {
	return v.value
}

func (v *NullableCoreReportbuilderConditionsAdd200ResponseAvailableconditionsInner) Set(val *CoreReportbuilderConditionsAdd200ResponseAvailableconditionsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreReportbuilderConditionsAdd200ResponseAvailableconditionsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreReportbuilderConditionsAdd200ResponseAvailableconditionsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreReportbuilderConditionsAdd200ResponseAvailableconditionsInner(val *CoreReportbuilderConditionsAdd200ResponseAvailableconditionsInner) *NullableCoreReportbuilderConditionsAdd200ResponseAvailableconditionsInner {
	return &NullableCoreReportbuilderConditionsAdd200ResponseAvailableconditionsInner{value: val, isSet: true}
}

func (v NullableCoreReportbuilderConditionsAdd200ResponseAvailableconditionsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreReportbuilderConditionsAdd200ResponseAvailableconditionsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


