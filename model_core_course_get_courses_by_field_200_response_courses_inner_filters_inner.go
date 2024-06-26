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

// checks if the CoreCourseGetCoursesByField200ResponseCoursesInnerFiltersInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreCourseGetCoursesByField200ResponseCoursesInnerFiltersInner{}

// CoreCourseGetCoursesByField200ResponseCoursesInnerFiltersInner struct for CoreCourseGetCoursesByField200ResponseCoursesInnerFiltersInner
type CoreCourseGetCoursesByField200ResponseCoursesInnerFiltersInner struct {
	// Filter plugin name
	Filter *string `json:"filter,omitempty"`
	// 1 or 0 to use when localstate is set to inherit
	Inheritedstate *int32 `json:"inheritedstate,omitempty"`
	// Filter state: 1 for on, -1 for off, 0 if inherit
	Localstate *int32 `json:"localstate,omitempty"`
}

// NewCoreCourseGetCoursesByField200ResponseCoursesInnerFiltersInner instantiates a new CoreCourseGetCoursesByField200ResponseCoursesInnerFiltersInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreCourseGetCoursesByField200ResponseCoursesInnerFiltersInner() *CoreCourseGetCoursesByField200ResponseCoursesInnerFiltersInner {
	this := CoreCourseGetCoursesByField200ResponseCoursesInnerFiltersInner{}
	var filter string = "null"
	this.Filter = &filter
	var inheritedstate int32 = null
	this.Inheritedstate = &inheritedstate
	var localstate int32 = null
	this.Localstate = &localstate
	return &this
}

// NewCoreCourseGetCoursesByField200ResponseCoursesInnerFiltersInnerWithDefaults instantiates a new CoreCourseGetCoursesByField200ResponseCoursesInnerFiltersInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreCourseGetCoursesByField200ResponseCoursesInnerFiltersInnerWithDefaults() *CoreCourseGetCoursesByField200ResponseCoursesInnerFiltersInner {
	this := CoreCourseGetCoursesByField200ResponseCoursesInnerFiltersInner{}
	var filter string = "null"
	this.Filter = &filter
	var inheritedstate int32 = null
	this.Inheritedstate = &inheritedstate
	var localstate int32 = null
	this.Localstate = &localstate
	return &this
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *CoreCourseGetCoursesByField200ResponseCoursesInnerFiltersInner) GetFilter() string {
	if o == nil || IsNil(o.Filter) {
		var ret string
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCourseGetCoursesByField200ResponseCoursesInnerFiltersInner) GetFilterOk() (*string, bool) {
	if o == nil || IsNil(o.Filter) {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *CoreCourseGetCoursesByField200ResponseCoursesInnerFiltersInner) HasFilter() bool {
	if o != nil && !IsNil(o.Filter) {
		return true
	}

	return false
}

// SetFilter gets a reference to the given string and assigns it to the Filter field.
func (o *CoreCourseGetCoursesByField200ResponseCoursesInnerFiltersInner) SetFilter(v string) {
	o.Filter = &v
}

// GetInheritedstate returns the Inheritedstate field value if set, zero value otherwise.
func (o *CoreCourseGetCoursesByField200ResponseCoursesInnerFiltersInner) GetInheritedstate() int32 {
	if o == nil || IsNil(o.Inheritedstate) {
		var ret int32
		return ret
	}
	return *o.Inheritedstate
}

// GetInheritedstateOk returns a tuple with the Inheritedstate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCourseGetCoursesByField200ResponseCoursesInnerFiltersInner) GetInheritedstateOk() (*int32, bool) {
	if o == nil || IsNil(o.Inheritedstate) {
		return nil, false
	}
	return o.Inheritedstate, true
}

// HasInheritedstate returns a boolean if a field has been set.
func (o *CoreCourseGetCoursesByField200ResponseCoursesInnerFiltersInner) HasInheritedstate() bool {
	if o != nil && !IsNil(o.Inheritedstate) {
		return true
	}

	return false
}

// SetInheritedstate gets a reference to the given int32 and assigns it to the Inheritedstate field.
func (o *CoreCourseGetCoursesByField200ResponseCoursesInnerFiltersInner) SetInheritedstate(v int32) {
	o.Inheritedstate = &v
}

// GetLocalstate returns the Localstate field value if set, zero value otherwise.
func (o *CoreCourseGetCoursesByField200ResponseCoursesInnerFiltersInner) GetLocalstate() int32 {
	if o == nil || IsNil(o.Localstate) {
		var ret int32
		return ret
	}
	return *o.Localstate
}

// GetLocalstateOk returns a tuple with the Localstate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCourseGetCoursesByField200ResponseCoursesInnerFiltersInner) GetLocalstateOk() (*int32, bool) {
	if o == nil || IsNil(o.Localstate) {
		return nil, false
	}
	return o.Localstate, true
}

// HasLocalstate returns a boolean if a field has been set.
func (o *CoreCourseGetCoursesByField200ResponseCoursesInnerFiltersInner) HasLocalstate() bool {
	if o != nil && !IsNil(o.Localstate) {
		return true
	}

	return false
}

// SetLocalstate gets a reference to the given int32 and assigns it to the Localstate field.
func (o *CoreCourseGetCoursesByField200ResponseCoursesInnerFiltersInner) SetLocalstate(v int32) {
	o.Localstate = &v
}

func (o CoreCourseGetCoursesByField200ResponseCoursesInnerFiltersInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreCourseGetCoursesByField200ResponseCoursesInnerFiltersInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Filter) {
		toSerialize["filter"] = o.Filter
	}
	if !IsNil(o.Inheritedstate) {
		toSerialize["inheritedstate"] = o.Inheritedstate
	}
	if !IsNil(o.Localstate) {
		toSerialize["localstate"] = o.Localstate
	}
	return toSerialize, nil
}

type NullableCoreCourseGetCoursesByField200ResponseCoursesInnerFiltersInner struct {
	value *CoreCourseGetCoursesByField200ResponseCoursesInnerFiltersInner
	isSet bool
}

func (v NullableCoreCourseGetCoursesByField200ResponseCoursesInnerFiltersInner) Get() *CoreCourseGetCoursesByField200ResponseCoursesInnerFiltersInner {
	return v.value
}

func (v *NullableCoreCourseGetCoursesByField200ResponseCoursesInnerFiltersInner) Set(val *CoreCourseGetCoursesByField200ResponseCoursesInnerFiltersInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreCourseGetCoursesByField200ResponseCoursesInnerFiltersInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreCourseGetCoursesByField200ResponseCoursesInnerFiltersInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreCourseGetCoursesByField200ResponseCoursesInnerFiltersInner(val *CoreCourseGetCoursesByField200ResponseCoursesInnerFiltersInner) *NullableCoreCourseGetCoursesByField200ResponseCoursesInnerFiltersInner {
	return &NullableCoreCourseGetCoursesByField200ResponseCoursesInnerFiltersInner{value: val, isSet: true}
}

func (v NullableCoreCourseGetCoursesByField200ResponseCoursesInnerFiltersInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreCourseGetCoursesByField200ResponseCoursesInnerFiltersInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


