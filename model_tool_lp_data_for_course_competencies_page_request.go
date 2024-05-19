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

// checks if the ToolLpDataForCourseCompetenciesPageRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ToolLpDataForCourseCompetenciesPageRequest{}

// ToolLpDataForCourseCompetenciesPageRequest struct for ToolLpDataForCourseCompetenciesPageRequest
type ToolLpDataForCourseCompetenciesPageRequest struct {
	// The course id
	Courseid int32 `json:"courseid"`
	// The module id
	Moduleid *int32 `json:"moduleid,omitempty"`
}

type _ToolLpDataForCourseCompetenciesPageRequest ToolLpDataForCourseCompetenciesPageRequest

// NewToolLpDataForCourseCompetenciesPageRequest instantiates a new ToolLpDataForCourseCompetenciesPageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewToolLpDataForCourseCompetenciesPageRequest(courseid int32) *ToolLpDataForCourseCompetenciesPageRequest {
	this := ToolLpDataForCourseCompetenciesPageRequest{}
	this.Courseid = courseid
	var moduleid int32 = 0
	this.Moduleid = &moduleid
	return &this
}

// NewToolLpDataForCourseCompetenciesPageRequestWithDefaults instantiates a new ToolLpDataForCourseCompetenciesPageRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewToolLpDataForCourseCompetenciesPageRequestWithDefaults() *ToolLpDataForCourseCompetenciesPageRequest {
	this := ToolLpDataForCourseCompetenciesPageRequest{}
	var moduleid int32 = 0
	this.Moduleid = &moduleid
	return &this
}

// GetCourseid returns the Courseid field value
func (o *ToolLpDataForCourseCompetenciesPageRequest) GetCourseid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Courseid
}

// GetCourseidOk returns a tuple with the Courseid field value
// and a boolean to check if the value has been set.
func (o *ToolLpDataForCourseCompetenciesPageRequest) GetCourseidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Courseid, true
}

// SetCourseid sets field value
func (o *ToolLpDataForCourseCompetenciesPageRequest) SetCourseid(v int32) {
	o.Courseid = v
}

// GetModuleid returns the Moduleid field value if set, zero value otherwise.
func (o *ToolLpDataForCourseCompetenciesPageRequest) GetModuleid() int32 {
	if o == nil || IsNil(o.Moduleid) {
		var ret int32
		return ret
	}
	return *o.Moduleid
}

// GetModuleidOk returns a tuple with the Moduleid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ToolLpDataForCourseCompetenciesPageRequest) GetModuleidOk() (*int32, bool) {
	if o == nil || IsNil(o.Moduleid) {
		return nil, false
	}
	return o.Moduleid, true
}

// HasModuleid returns a boolean if a field has been set.
func (o *ToolLpDataForCourseCompetenciesPageRequest) HasModuleid() bool {
	if o != nil && !IsNil(o.Moduleid) {
		return true
	}

	return false
}

// SetModuleid gets a reference to the given int32 and assigns it to the Moduleid field.
func (o *ToolLpDataForCourseCompetenciesPageRequest) SetModuleid(v int32) {
	o.Moduleid = &v
}

func (o ToolLpDataForCourseCompetenciesPageRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ToolLpDataForCourseCompetenciesPageRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["courseid"] = o.Courseid
	if !IsNil(o.Moduleid) {
		toSerialize["moduleid"] = o.Moduleid
	}
	return toSerialize, nil
}

func (o *ToolLpDataForCourseCompetenciesPageRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"courseid",
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

	varToolLpDataForCourseCompetenciesPageRequest := _ToolLpDataForCourseCompetenciesPageRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varToolLpDataForCourseCompetenciesPageRequest)

	if err != nil {
		return err
	}

	*o = ToolLpDataForCourseCompetenciesPageRequest(varToolLpDataForCourseCompetenciesPageRequest)

	return err
}

type NullableToolLpDataForCourseCompetenciesPageRequest struct {
	value *ToolLpDataForCourseCompetenciesPageRequest
	isSet bool
}

func (v NullableToolLpDataForCourseCompetenciesPageRequest) Get() *ToolLpDataForCourseCompetenciesPageRequest {
	return v.value
}

func (v *NullableToolLpDataForCourseCompetenciesPageRequest) Set(val *ToolLpDataForCourseCompetenciesPageRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableToolLpDataForCourseCompetenciesPageRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableToolLpDataForCourseCompetenciesPageRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableToolLpDataForCourseCompetenciesPageRequest(val *ToolLpDataForCourseCompetenciesPageRequest) *NullableToolLpDataForCourseCompetenciesPageRequest {
	return &NullableToolLpDataForCourseCompetenciesPageRequest{value: val, isSet: true}
}

func (v NullableToolLpDataForCourseCompetenciesPageRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableToolLpDataForCourseCompetenciesPageRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


