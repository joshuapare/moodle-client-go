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

// checks if the ModLessonGetLessonsByCourses200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModLessonGetLessonsByCourses200Response{}

// ModLessonGetLessonsByCourses200Response struct for ModLessonGetLessonsByCourses200Response
type ModLessonGetLessonsByCourses200Response struct {
	Lessons []ModLessonGetLessonsByCourses200ResponseLessonsInner `json:"lessons"`
	Warnings []AuthEmailSignupUser200ResponseWarningsInner `json:"warnings,omitempty"`
}

type _ModLessonGetLessonsByCourses200Response ModLessonGetLessonsByCourses200Response

// NewModLessonGetLessonsByCourses200Response instantiates a new ModLessonGetLessonsByCourses200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModLessonGetLessonsByCourses200Response(lessons []ModLessonGetLessonsByCourses200ResponseLessonsInner) *ModLessonGetLessonsByCourses200Response {
	this := ModLessonGetLessonsByCourses200Response{}
	this.Lessons = lessons
	return &this
}

// NewModLessonGetLessonsByCourses200ResponseWithDefaults instantiates a new ModLessonGetLessonsByCourses200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModLessonGetLessonsByCourses200ResponseWithDefaults() *ModLessonGetLessonsByCourses200Response {
	this := ModLessonGetLessonsByCourses200Response{}
	return &this
}

// GetLessons returns the Lessons field value
func (o *ModLessonGetLessonsByCourses200Response) GetLessons() []ModLessonGetLessonsByCourses200ResponseLessonsInner {
	if o == nil {
		var ret []ModLessonGetLessonsByCourses200ResponseLessonsInner
		return ret
	}

	return o.Lessons
}

// GetLessonsOk returns a tuple with the Lessons field value
// and a boolean to check if the value has been set.
func (o *ModLessonGetLessonsByCourses200Response) GetLessonsOk() ([]ModLessonGetLessonsByCourses200ResponseLessonsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Lessons, true
}

// SetLessons sets field value
func (o *ModLessonGetLessonsByCourses200Response) SetLessons(v []ModLessonGetLessonsByCourses200ResponseLessonsInner) {
	o.Lessons = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *ModLessonGetLessonsByCourses200Response) GetWarnings() []AuthEmailSignupUser200ResponseWarningsInner {
	if o == nil || IsNil(o.Warnings) {
		var ret []AuthEmailSignupUser200ResponseWarningsInner
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModLessonGetLessonsByCourses200Response) GetWarningsOk() ([]AuthEmailSignupUser200ResponseWarningsInner, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *ModLessonGetLessonsByCourses200Response) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []AuthEmailSignupUser200ResponseWarningsInner and assigns it to the Warnings field.
func (o *ModLessonGetLessonsByCourses200Response) SetWarnings(v []AuthEmailSignupUser200ResponseWarningsInner) {
	o.Warnings = v
}

func (o ModLessonGetLessonsByCourses200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModLessonGetLessonsByCourses200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["lessons"] = o.Lessons
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

func (o *ModLessonGetLessonsByCourses200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"lessons",
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

	varModLessonGetLessonsByCourses200Response := _ModLessonGetLessonsByCourses200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModLessonGetLessonsByCourses200Response)

	if err != nil {
		return err
	}

	*o = ModLessonGetLessonsByCourses200Response(varModLessonGetLessonsByCourses200Response)

	return err
}

type NullableModLessonGetLessonsByCourses200Response struct {
	value *ModLessonGetLessonsByCourses200Response
	isSet bool
}

func (v NullableModLessonGetLessonsByCourses200Response) Get() *ModLessonGetLessonsByCourses200Response {
	return v.value
}

func (v *NullableModLessonGetLessonsByCourses200Response) Set(val *ModLessonGetLessonsByCourses200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableModLessonGetLessonsByCourses200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableModLessonGetLessonsByCourses200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModLessonGetLessonsByCourses200Response(val *ModLessonGetLessonsByCourses200Response) *NullableModLessonGetLessonsByCourses200Response {
	return &NullableModLessonGetLessonsByCourses200Response{value: val, isSet: true}
}

func (v NullableModLessonGetLessonsByCourses200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModLessonGetLessonsByCourses200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


