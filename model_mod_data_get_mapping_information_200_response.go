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

// checks if the ModDataGetMappingInformation200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModDataGetMappingInformation200Response{}

// ModDataGetMappingInformation200Response struct for ModDataGetMappingInformation200Response
type ModDataGetMappingInformation200Response struct {
	Data *ModDataGetMappingInformation200ResponseData `json:"data,omitempty"`
	Warnings []AuthEmailSignupUser200ResponseWarningsInner `json:"warnings,omitempty"`
}

// NewModDataGetMappingInformation200Response instantiates a new ModDataGetMappingInformation200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModDataGetMappingInformation200Response() *ModDataGetMappingInformation200Response {
	this := ModDataGetMappingInformation200Response{}
	return &this
}

// NewModDataGetMappingInformation200ResponseWithDefaults instantiates a new ModDataGetMappingInformation200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModDataGetMappingInformation200ResponseWithDefaults() *ModDataGetMappingInformation200Response {
	this := ModDataGetMappingInformation200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ModDataGetMappingInformation200Response) GetData() ModDataGetMappingInformation200ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret ModDataGetMappingInformation200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModDataGetMappingInformation200Response) GetDataOk() (*ModDataGetMappingInformation200ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ModDataGetMappingInformation200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given ModDataGetMappingInformation200ResponseData and assigns it to the Data field.
func (o *ModDataGetMappingInformation200Response) SetData(v ModDataGetMappingInformation200ResponseData) {
	o.Data = &v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *ModDataGetMappingInformation200Response) GetWarnings() []AuthEmailSignupUser200ResponseWarningsInner {
	if o == nil || IsNil(o.Warnings) {
		var ret []AuthEmailSignupUser200ResponseWarningsInner
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModDataGetMappingInformation200Response) GetWarningsOk() ([]AuthEmailSignupUser200ResponseWarningsInner, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *ModDataGetMappingInformation200Response) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []AuthEmailSignupUser200ResponseWarningsInner and assigns it to the Warnings field.
func (o *ModDataGetMappingInformation200Response) SetWarnings(v []AuthEmailSignupUser200ResponseWarningsInner) {
	o.Warnings = v
}

func (o ModDataGetMappingInformation200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModDataGetMappingInformation200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullableModDataGetMappingInformation200Response struct {
	value *ModDataGetMappingInformation200Response
	isSet bool
}

func (v NullableModDataGetMappingInformation200Response) Get() *ModDataGetMappingInformation200Response {
	return v.value
}

func (v *NullableModDataGetMappingInformation200Response) Set(val *ModDataGetMappingInformation200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableModDataGetMappingInformation200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableModDataGetMappingInformation200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModDataGetMappingInformation200Response(val *ModDataGetMappingInformation200Response) *NullableModDataGetMappingInformation200Response {
	return &NullableModDataGetMappingInformation200Response{value: val, isSet: true}
}

func (v NullableModDataGetMappingInformation200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModDataGetMappingInformation200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


