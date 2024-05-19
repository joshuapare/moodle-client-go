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

// checks if the ToolMoodlenetSearchCourses200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ToolMoodlenetSearchCourses200Response{}

// ToolMoodlenetSearchCourses200Response struct for ToolMoodlenetSearchCourses200Response
type ToolMoodlenetSearchCourses200Response struct {
	Courses []ToolMoodlenetSearchCourses200ResponseCoursesInner `json:"courses"`
}

type _ToolMoodlenetSearchCourses200Response ToolMoodlenetSearchCourses200Response

// NewToolMoodlenetSearchCourses200Response instantiates a new ToolMoodlenetSearchCourses200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewToolMoodlenetSearchCourses200Response(courses []ToolMoodlenetSearchCourses200ResponseCoursesInner) *ToolMoodlenetSearchCourses200Response {
	this := ToolMoodlenetSearchCourses200Response{}
	this.Courses = courses
	return &this
}

// NewToolMoodlenetSearchCourses200ResponseWithDefaults instantiates a new ToolMoodlenetSearchCourses200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewToolMoodlenetSearchCourses200ResponseWithDefaults() *ToolMoodlenetSearchCourses200Response {
	this := ToolMoodlenetSearchCourses200Response{}
	return &this
}

// GetCourses returns the Courses field value
func (o *ToolMoodlenetSearchCourses200Response) GetCourses() []ToolMoodlenetSearchCourses200ResponseCoursesInner {
	if o == nil {
		var ret []ToolMoodlenetSearchCourses200ResponseCoursesInner
		return ret
	}

	return o.Courses
}

// GetCoursesOk returns a tuple with the Courses field value
// and a boolean to check if the value has been set.
func (o *ToolMoodlenetSearchCourses200Response) GetCoursesOk() ([]ToolMoodlenetSearchCourses200ResponseCoursesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Courses, true
}

// SetCourses sets field value
func (o *ToolMoodlenetSearchCourses200Response) SetCourses(v []ToolMoodlenetSearchCourses200ResponseCoursesInner) {
	o.Courses = v
}

func (o ToolMoodlenetSearchCourses200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ToolMoodlenetSearchCourses200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["courses"] = o.Courses
	return toSerialize, nil
}

func (o *ToolMoodlenetSearchCourses200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"courses",
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

	varToolMoodlenetSearchCourses200Response := _ToolMoodlenetSearchCourses200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varToolMoodlenetSearchCourses200Response)

	if err != nil {
		return err
	}

	*o = ToolMoodlenetSearchCourses200Response(varToolMoodlenetSearchCourses200Response)

	return err
}

type NullableToolMoodlenetSearchCourses200Response struct {
	value *ToolMoodlenetSearchCourses200Response
	isSet bool
}

func (v NullableToolMoodlenetSearchCourses200Response) Get() *ToolMoodlenetSearchCourses200Response {
	return v.value
}

func (v *NullableToolMoodlenetSearchCourses200Response) Set(val *ToolMoodlenetSearchCourses200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableToolMoodlenetSearchCourses200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableToolMoodlenetSearchCourses200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableToolMoodlenetSearchCourses200Response(val *ToolMoodlenetSearchCourses200Response) *NullableToolMoodlenetSearchCourses200Response {
	return &NullableToolMoodlenetSearchCourses200Response{value: val, isSet: true}
}

func (v NullableToolMoodlenetSearchCourses200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableToolMoodlenetSearchCourses200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

