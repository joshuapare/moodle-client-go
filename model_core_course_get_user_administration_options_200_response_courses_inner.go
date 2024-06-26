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

// checks if the CoreCourseGetUserAdministrationOptions200ResponseCoursesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreCourseGetUserAdministrationOptions200ResponseCoursesInner{}

// CoreCourseGetUserAdministrationOptions200ResponseCoursesInner struct for CoreCourseGetUserAdministrationOptions200ResponseCoursesInner
type CoreCourseGetUserAdministrationOptions200ResponseCoursesInner struct {
	// Course id
	Id *int32 `json:"id,omitempty"`
	Options []CoreCourseGetUserAdministrationOptions200ResponseCoursesInnerOptionsInner `json:"options,omitempty"`
}

// NewCoreCourseGetUserAdministrationOptions200ResponseCoursesInner instantiates a new CoreCourseGetUserAdministrationOptions200ResponseCoursesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreCourseGetUserAdministrationOptions200ResponseCoursesInner() *CoreCourseGetUserAdministrationOptions200ResponseCoursesInner {
	this := CoreCourseGetUserAdministrationOptions200ResponseCoursesInner{}
	return &this
}

// NewCoreCourseGetUserAdministrationOptions200ResponseCoursesInnerWithDefaults instantiates a new CoreCourseGetUserAdministrationOptions200ResponseCoursesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreCourseGetUserAdministrationOptions200ResponseCoursesInnerWithDefaults() *CoreCourseGetUserAdministrationOptions200ResponseCoursesInner {
	this := CoreCourseGetUserAdministrationOptions200ResponseCoursesInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CoreCourseGetUserAdministrationOptions200ResponseCoursesInner) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCourseGetUserAdministrationOptions200ResponseCoursesInner) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CoreCourseGetUserAdministrationOptions200ResponseCoursesInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *CoreCourseGetUserAdministrationOptions200ResponseCoursesInner) SetId(v int32) {
	o.Id = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *CoreCourseGetUserAdministrationOptions200ResponseCoursesInner) GetOptions() []CoreCourseGetUserAdministrationOptions200ResponseCoursesInnerOptionsInner {
	if o == nil || IsNil(o.Options) {
		var ret []CoreCourseGetUserAdministrationOptions200ResponseCoursesInnerOptionsInner
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCourseGetUserAdministrationOptions200ResponseCoursesInner) GetOptionsOk() ([]CoreCourseGetUserAdministrationOptions200ResponseCoursesInnerOptionsInner, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *CoreCourseGetUserAdministrationOptions200ResponseCoursesInner) HasOptions() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given []CoreCourseGetUserAdministrationOptions200ResponseCoursesInnerOptionsInner and assigns it to the Options field.
func (o *CoreCourseGetUserAdministrationOptions200ResponseCoursesInner) SetOptions(v []CoreCourseGetUserAdministrationOptions200ResponseCoursesInnerOptionsInner) {
	o.Options = v
}

func (o CoreCourseGetUserAdministrationOptions200ResponseCoursesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreCourseGetUserAdministrationOptions200ResponseCoursesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Options) {
		toSerialize["options"] = o.Options
	}
	return toSerialize, nil
}

type NullableCoreCourseGetUserAdministrationOptions200ResponseCoursesInner struct {
	value *CoreCourseGetUserAdministrationOptions200ResponseCoursesInner
	isSet bool
}

func (v NullableCoreCourseGetUserAdministrationOptions200ResponseCoursesInner) Get() *CoreCourseGetUserAdministrationOptions200ResponseCoursesInner {
	return v.value
}

func (v *NullableCoreCourseGetUserAdministrationOptions200ResponseCoursesInner) Set(val *CoreCourseGetUserAdministrationOptions200ResponseCoursesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreCourseGetUserAdministrationOptions200ResponseCoursesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreCourseGetUserAdministrationOptions200ResponseCoursesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreCourseGetUserAdministrationOptions200ResponseCoursesInner(val *CoreCourseGetUserAdministrationOptions200ResponseCoursesInner) *NullableCoreCourseGetUserAdministrationOptions200ResponseCoursesInner {
	return &NullableCoreCourseGetUserAdministrationOptions200ResponseCoursesInner{value: val, isSet: true}
}

func (v NullableCoreCourseGetUserAdministrationOptions200ResponseCoursesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreCourseGetUserAdministrationOptions200ResponseCoursesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


