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

// checks if the CoreCourseGetUserNavigationOptions200ResponseCoursesInnerOptionsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreCourseGetUserNavigationOptions200ResponseCoursesInnerOptionsInner{}

// CoreCourseGetUserNavigationOptions200ResponseCoursesInnerOptionsInner struct for CoreCourseGetUserNavigationOptions200ResponseCoursesInnerOptionsInner
type CoreCourseGetUserNavigationOptions200ResponseCoursesInnerOptionsInner struct {
	// Whether the option is available or not
	Available *bool `json:"available,omitempty"`
	// Option name
	Name *string `json:"name,omitempty"`
}

// NewCoreCourseGetUserNavigationOptions200ResponseCoursesInnerOptionsInner instantiates a new CoreCourseGetUserNavigationOptions200ResponseCoursesInnerOptionsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreCourseGetUserNavigationOptions200ResponseCoursesInnerOptionsInner() *CoreCourseGetUserNavigationOptions200ResponseCoursesInnerOptionsInner {
	this := CoreCourseGetUserNavigationOptions200ResponseCoursesInnerOptionsInner{}
	return &this
}

// NewCoreCourseGetUserNavigationOptions200ResponseCoursesInnerOptionsInnerWithDefaults instantiates a new CoreCourseGetUserNavigationOptions200ResponseCoursesInnerOptionsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreCourseGetUserNavigationOptions200ResponseCoursesInnerOptionsInnerWithDefaults() *CoreCourseGetUserNavigationOptions200ResponseCoursesInnerOptionsInner {
	this := CoreCourseGetUserNavigationOptions200ResponseCoursesInnerOptionsInner{}
	return &this
}

// GetAvailable returns the Available field value if set, zero value otherwise.
func (o *CoreCourseGetUserNavigationOptions200ResponseCoursesInnerOptionsInner) GetAvailable() bool {
	if o == nil || IsNil(o.Available) {
		var ret bool
		return ret
	}
	return *o.Available
}

// GetAvailableOk returns a tuple with the Available field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCourseGetUserNavigationOptions200ResponseCoursesInnerOptionsInner) GetAvailableOk() (*bool, bool) {
	if o == nil || IsNil(o.Available) {
		return nil, false
	}
	return o.Available, true
}

// HasAvailable returns a boolean if a field has been set.
func (o *CoreCourseGetUserNavigationOptions200ResponseCoursesInnerOptionsInner) HasAvailable() bool {
	if o != nil && !IsNil(o.Available) {
		return true
	}

	return false
}

// SetAvailable gets a reference to the given bool and assigns it to the Available field.
func (o *CoreCourseGetUserNavigationOptions200ResponseCoursesInnerOptionsInner) SetAvailable(v bool) {
	o.Available = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CoreCourseGetUserNavigationOptions200ResponseCoursesInnerOptionsInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCourseGetUserNavigationOptions200ResponseCoursesInnerOptionsInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CoreCourseGetUserNavigationOptions200ResponseCoursesInnerOptionsInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CoreCourseGetUserNavigationOptions200ResponseCoursesInnerOptionsInner) SetName(v string) {
	o.Name = &v
}

func (o CoreCourseGetUserNavigationOptions200ResponseCoursesInnerOptionsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreCourseGetUserNavigationOptions200ResponseCoursesInnerOptionsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Available) {
		toSerialize["available"] = o.Available
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableCoreCourseGetUserNavigationOptions200ResponseCoursesInnerOptionsInner struct {
	value *CoreCourseGetUserNavigationOptions200ResponseCoursesInnerOptionsInner
	isSet bool
}

func (v NullableCoreCourseGetUserNavigationOptions200ResponseCoursesInnerOptionsInner) Get() *CoreCourseGetUserNavigationOptions200ResponseCoursesInnerOptionsInner {
	return v.value
}

func (v *NullableCoreCourseGetUserNavigationOptions200ResponseCoursesInnerOptionsInner) Set(val *CoreCourseGetUserNavigationOptions200ResponseCoursesInnerOptionsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreCourseGetUserNavigationOptions200ResponseCoursesInnerOptionsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreCourseGetUserNavigationOptions200ResponseCoursesInnerOptionsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreCourseGetUserNavigationOptions200ResponseCoursesInnerOptionsInner(val *CoreCourseGetUserNavigationOptions200ResponseCoursesInnerOptionsInner) *NullableCoreCourseGetUserNavigationOptions200ResponseCoursesInnerOptionsInner {
	return &NullableCoreCourseGetUserNavigationOptions200ResponseCoursesInnerOptionsInner{value: val, isSet: true}
}

func (v NullableCoreCourseGetUserNavigationOptions200ResponseCoursesInnerOptionsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreCourseGetUserNavigationOptions200ResponseCoursesInnerOptionsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


