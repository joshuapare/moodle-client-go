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

// checks if the CoreCourseGetCoursesByField200ResponseCoursesInnerCustomfieldsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreCourseGetCoursesByField200ResponseCoursesInnerCustomfieldsInner{}

// CoreCourseGetCoursesByField200ResponseCoursesInnerCustomfieldsInner struct for CoreCourseGetCoursesByField200ResponseCoursesInnerCustomfieldsInner
type CoreCourseGetCoursesByField200ResponseCoursesInnerCustomfieldsInner struct {
	// The name of the custom field
	Name *string `json:"name,omitempty"`
	// The shortname of the custom field - to be able to build the field class in the code
	Shortname *string `json:"shortname,omitempty"`
	// The type of the custom field - text field, checkbox...
	Type *string `json:"type,omitempty"`
	// The value of the custom field
	Value *string `json:"value,omitempty"`
	// The raw value of the custom field
	Valueraw *string `json:"valueraw,omitempty"`
}

// NewCoreCourseGetCoursesByField200ResponseCoursesInnerCustomfieldsInner instantiates a new CoreCourseGetCoursesByField200ResponseCoursesInnerCustomfieldsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreCourseGetCoursesByField200ResponseCoursesInnerCustomfieldsInner() *CoreCourseGetCoursesByField200ResponseCoursesInnerCustomfieldsInner {
	this := CoreCourseGetCoursesByField200ResponseCoursesInnerCustomfieldsInner{}
	return &this
}

// NewCoreCourseGetCoursesByField200ResponseCoursesInnerCustomfieldsInnerWithDefaults instantiates a new CoreCourseGetCoursesByField200ResponseCoursesInnerCustomfieldsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreCourseGetCoursesByField200ResponseCoursesInnerCustomfieldsInnerWithDefaults() *CoreCourseGetCoursesByField200ResponseCoursesInnerCustomfieldsInner {
	this := CoreCourseGetCoursesByField200ResponseCoursesInnerCustomfieldsInner{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CoreCourseGetCoursesByField200ResponseCoursesInnerCustomfieldsInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCourseGetCoursesByField200ResponseCoursesInnerCustomfieldsInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CoreCourseGetCoursesByField200ResponseCoursesInnerCustomfieldsInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CoreCourseGetCoursesByField200ResponseCoursesInnerCustomfieldsInner) SetName(v string) {
	o.Name = &v
}

// GetShortname returns the Shortname field value if set, zero value otherwise.
func (o *CoreCourseGetCoursesByField200ResponseCoursesInnerCustomfieldsInner) GetShortname() string {
	if o == nil || IsNil(o.Shortname) {
		var ret string
		return ret
	}
	return *o.Shortname
}

// GetShortnameOk returns a tuple with the Shortname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCourseGetCoursesByField200ResponseCoursesInnerCustomfieldsInner) GetShortnameOk() (*string, bool) {
	if o == nil || IsNil(o.Shortname) {
		return nil, false
	}
	return o.Shortname, true
}

// HasShortname returns a boolean if a field has been set.
func (o *CoreCourseGetCoursesByField200ResponseCoursesInnerCustomfieldsInner) HasShortname() bool {
	if o != nil && !IsNil(o.Shortname) {
		return true
	}

	return false
}

// SetShortname gets a reference to the given string and assigns it to the Shortname field.
func (o *CoreCourseGetCoursesByField200ResponseCoursesInnerCustomfieldsInner) SetShortname(v string) {
	o.Shortname = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CoreCourseGetCoursesByField200ResponseCoursesInnerCustomfieldsInner) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCourseGetCoursesByField200ResponseCoursesInnerCustomfieldsInner) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CoreCourseGetCoursesByField200ResponseCoursesInnerCustomfieldsInner) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *CoreCourseGetCoursesByField200ResponseCoursesInnerCustomfieldsInner) SetType(v string) {
	o.Type = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CoreCourseGetCoursesByField200ResponseCoursesInnerCustomfieldsInner) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCourseGetCoursesByField200ResponseCoursesInnerCustomfieldsInner) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CoreCourseGetCoursesByField200ResponseCoursesInnerCustomfieldsInner) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *CoreCourseGetCoursesByField200ResponseCoursesInnerCustomfieldsInner) SetValue(v string) {
	o.Value = &v
}

// GetValueraw returns the Valueraw field value if set, zero value otherwise.
func (o *CoreCourseGetCoursesByField200ResponseCoursesInnerCustomfieldsInner) GetValueraw() string {
	if o == nil || IsNil(o.Valueraw) {
		var ret string
		return ret
	}
	return *o.Valueraw
}

// GetValuerawOk returns a tuple with the Valueraw field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCourseGetCoursesByField200ResponseCoursesInnerCustomfieldsInner) GetValuerawOk() (*string, bool) {
	if o == nil || IsNil(o.Valueraw) {
		return nil, false
	}
	return o.Valueraw, true
}

// HasValueraw returns a boolean if a field has been set.
func (o *CoreCourseGetCoursesByField200ResponseCoursesInnerCustomfieldsInner) HasValueraw() bool {
	if o != nil && !IsNil(o.Valueraw) {
		return true
	}

	return false
}

// SetValueraw gets a reference to the given string and assigns it to the Valueraw field.
func (o *CoreCourseGetCoursesByField200ResponseCoursesInnerCustomfieldsInner) SetValueraw(v string) {
	o.Valueraw = &v
}

func (o CoreCourseGetCoursesByField200ResponseCoursesInnerCustomfieldsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreCourseGetCoursesByField200ResponseCoursesInnerCustomfieldsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Shortname) {
		toSerialize["shortname"] = o.Shortname
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.Valueraw) {
		toSerialize["valueraw"] = o.Valueraw
	}
	return toSerialize, nil
}

type NullableCoreCourseGetCoursesByField200ResponseCoursesInnerCustomfieldsInner struct {
	value *CoreCourseGetCoursesByField200ResponseCoursesInnerCustomfieldsInner
	isSet bool
}

func (v NullableCoreCourseGetCoursesByField200ResponseCoursesInnerCustomfieldsInner) Get() *CoreCourseGetCoursesByField200ResponseCoursesInnerCustomfieldsInner {
	return v.value
}

func (v *NullableCoreCourseGetCoursesByField200ResponseCoursesInnerCustomfieldsInner) Set(val *CoreCourseGetCoursesByField200ResponseCoursesInnerCustomfieldsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreCourseGetCoursesByField200ResponseCoursesInnerCustomfieldsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreCourseGetCoursesByField200ResponseCoursesInnerCustomfieldsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreCourseGetCoursesByField200ResponseCoursesInnerCustomfieldsInner(val *CoreCourseGetCoursesByField200ResponseCoursesInnerCustomfieldsInner) *NullableCoreCourseGetCoursesByField200ResponseCoursesInnerCustomfieldsInner {
	return &NullableCoreCourseGetCoursesByField200ResponseCoursesInnerCustomfieldsInner{value: val, isSet: true}
}

func (v NullableCoreCourseGetCoursesByField200ResponseCoursesInnerCustomfieldsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreCourseGetCoursesByField200ResponseCoursesInnerCustomfieldsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


