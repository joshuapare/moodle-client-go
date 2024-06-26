/*
Moodle Webservice API

Auto-generated OpenAPI spec for Moodle's Webservice API. Requires installation of the Catalyst RESTful API plugin: https://github.com/catalyst/moodle-webservice_restful.

API version: 4.3.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package moodleclient

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the CoreCourseGetModuleRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreCourseGetModuleRequest{}

// CoreCourseGetModuleRequest struct for CoreCourseGetModuleRequest
type CoreCourseGetModuleRequest struct {
	// course module id
	Id int32 `json:"id"`
	// section to return to
	Sectionreturn *int32 `json:"sectionreturn,omitempty"`
}

type _CoreCourseGetModuleRequest CoreCourseGetModuleRequest

// NewCoreCourseGetModuleRequest instantiates a new CoreCourseGetModuleRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreCourseGetModuleRequest(id int32) *CoreCourseGetModuleRequest {
	this := CoreCourseGetModuleRequest{}
	this.Id = id
	return &this
}

// NewCoreCourseGetModuleRequestWithDefaults instantiates a new CoreCourseGetModuleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreCourseGetModuleRequestWithDefaults() *CoreCourseGetModuleRequest {
	this := CoreCourseGetModuleRequest{}
	return &this
}

// GetId returns the Id field value
func (o *CoreCourseGetModuleRequest) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CoreCourseGetModuleRequest) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CoreCourseGetModuleRequest) SetId(v int32) {
	o.Id = v
}

// GetSectionreturn returns the Sectionreturn field value if set, zero value otherwise.
func (o *CoreCourseGetModuleRequest) GetSectionreturn() int32 {
	if o == nil || IsNil(o.Sectionreturn) {
		var ret int32
		return ret
	}
	return *o.Sectionreturn
}

// GetSectionreturnOk returns a tuple with the Sectionreturn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCourseGetModuleRequest) GetSectionreturnOk() (*int32, bool) {
	if o == nil || IsNil(o.Sectionreturn) {
		return nil, false
	}
	return o.Sectionreturn, true
}

// HasSectionreturn returns a boolean if a field has been set.
func (o *CoreCourseGetModuleRequest) HasSectionreturn() bool {
	if o != nil && !IsNil(o.Sectionreturn) {
		return true
	}

	return false
}

// SetSectionreturn gets a reference to the given int32 and assigns it to the Sectionreturn field.
func (o *CoreCourseGetModuleRequest) SetSectionreturn(v int32) {
	o.Sectionreturn = &v
}

func (o CoreCourseGetModuleRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreCourseGetModuleRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.Sectionreturn) {
		toSerialize["sectionreturn"] = o.Sectionreturn
	}
	return toSerialize, nil
}

func (o *CoreCourseGetModuleRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
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

	varCoreCourseGetModuleRequest := _CoreCourseGetModuleRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCoreCourseGetModuleRequest)

	if err != nil {
		return err
	}

	*o = CoreCourseGetModuleRequest(varCoreCourseGetModuleRequest)

	return err
}

type NullableCoreCourseGetModuleRequest struct {
	value *CoreCourseGetModuleRequest
	isSet bool
}

func (v NullableCoreCourseGetModuleRequest) Get() *CoreCourseGetModuleRequest {
	return v.value
}

func (v *NullableCoreCourseGetModuleRequest) Set(val *CoreCourseGetModuleRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreCourseGetModuleRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreCourseGetModuleRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreCourseGetModuleRequest(val *CoreCourseGetModuleRequest) *NullableCoreCourseGetModuleRequest {
	return &NullableCoreCourseGetModuleRequest{value: val, isSet: true}
}

func (v NullableCoreCourseGetModuleRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreCourseGetModuleRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


