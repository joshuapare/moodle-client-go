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

// checks if the CoreCourseGetCoursesByField200ResponseCoursesInnerCourseformatoptionsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreCourseGetCoursesByField200ResponseCoursesInnerCourseformatoptionsInner{}

// CoreCourseGetCoursesByField200ResponseCoursesInnerCourseformatoptionsInner struct for CoreCourseGetCoursesByField200ResponseCoursesInnerCourseformatoptionsInner
type CoreCourseGetCoursesByField200ResponseCoursesInnerCourseformatoptionsInner struct {
	// Course format option name.
	Name *string `json:"name,omitempty"`
	// Course format option value.
	Value *string `json:"value,omitempty"`
}

// NewCoreCourseGetCoursesByField200ResponseCoursesInnerCourseformatoptionsInner instantiates a new CoreCourseGetCoursesByField200ResponseCoursesInnerCourseformatoptionsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreCourseGetCoursesByField200ResponseCoursesInnerCourseformatoptionsInner() *CoreCourseGetCoursesByField200ResponseCoursesInnerCourseformatoptionsInner {
	this := CoreCourseGetCoursesByField200ResponseCoursesInnerCourseformatoptionsInner{}
	var name string = "null"
	this.Name = &name
	var value string = "null"
	this.Value = &value
	return &this
}

// NewCoreCourseGetCoursesByField200ResponseCoursesInnerCourseformatoptionsInnerWithDefaults instantiates a new CoreCourseGetCoursesByField200ResponseCoursesInnerCourseformatoptionsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreCourseGetCoursesByField200ResponseCoursesInnerCourseformatoptionsInnerWithDefaults() *CoreCourseGetCoursesByField200ResponseCoursesInnerCourseformatoptionsInner {
	this := CoreCourseGetCoursesByField200ResponseCoursesInnerCourseformatoptionsInner{}
	var name string = "null"
	this.Name = &name
	var value string = "null"
	this.Value = &value
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CoreCourseGetCoursesByField200ResponseCoursesInnerCourseformatoptionsInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCourseGetCoursesByField200ResponseCoursesInnerCourseformatoptionsInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CoreCourseGetCoursesByField200ResponseCoursesInnerCourseformatoptionsInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CoreCourseGetCoursesByField200ResponseCoursesInnerCourseformatoptionsInner) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CoreCourseGetCoursesByField200ResponseCoursesInnerCourseformatoptionsInner) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCourseGetCoursesByField200ResponseCoursesInnerCourseformatoptionsInner) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CoreCourseGetCoursesByField200ResponseCoursesInnerCourseformatoptionsInner) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *CoreCourseGetCoursesByField200ResponseCoursesInnerCourseformatoptionsInner) SetValue(v string) {
	o.Value = &v
}

func (o CoreCourseGetCoursesByField200ResponseCoursesInnerCourseformatoptionsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreCourseGetCoursesByField200ResponseCoursesInnerCourseformatoptionsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableCoreCourseGetCoursesByField200ResponseCoursesInnerCourseformatoptionsInner struct {
	value *CoreCourseGetCoursesByField200ResponseCoursesInnerCourseformatoptionsInner
	isSet bool
}

func (v NullableCoreCourseGetCoursesByField200ResponseCoursesInnerCourseformatoptionsInner) Get() *CoreCourseGetCoursesByField200ResponseCoursesInnerCourseformatoptionsInner {
	return v.value
}

func (v *NullableCoreCourseGetCoursesByField200ResponseCoursesInnerCourseformatoptionsInner) Set(val *CoreCourseGetCoursesByField200ResponseCoursesInnerCourseformatoptionsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreCourseGetCoursesByField200ResponseCoursesInnerCourseformatoptionsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreCourseGetCoursesByField200ResponseCoursesInnerCourseformatoptionsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreCourseGetCoursesByField200ResponseCoursesInnerCourseformatoptionsInner(val *CoreCourseGetCoursesByField200ResponseCoursesInnerCourseformatoptionsInner) *NullableCoreCourseGetCoursesByField200ResponseCoursesInnerCourseformatoptionsInner {
	return &NullableCoreCourseGetCoursesByField200ResponseCoursesInnerCourseformatoptionsInner{value: val, isSet: true}
}

func (v NullableCoreCourseGetCoursesByField200ResponseCoursesInnerCourseformatoptionsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreCourseGetCoursesByField200ResponseCoursesInnerCourseformatoptionsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


