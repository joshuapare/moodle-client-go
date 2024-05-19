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

// checks if the CoreCourseGetCourseModuleByInstance200ResponseCmOutcomesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreCourseGetCourseModuleByInstance200ResponseCmOutcomesInner{}

// CoreCourseGetCourseModuleByInstance200ResponseCmOutcomesInner struct for CoreCourseGetCourseModuleByInstance200ResponseCmOutcomesInner
type CoreCourseGetCourseModuleByInstance200ResponseCmOutcomesInner struct {
	// Outcome id
	Id *string `json:"id,omitempty"`
	// Outcome full name
	Name *string `json:"name,omitempty"`
	// Scale items
	Scale *string `json:"scale,omitempty"`
}

// NewCoreCourseGetCourseModuleByInstance200ResponseCmOutcomesInner instantiates a new CoreCourseGetCourseModuleByInstance200ResponseCmOutcomesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreCourseGetCourseModuleByInstance200ResponseCmOutcomesInner() *CoreCourseGetCourseModuleByInstance200ResponseCmOutcomesInner {
	this := CoreCourseGetCourseModuleByInstance200ResponseCmOutcomesInner{}
	return &this
}

// NewCoreCourseGetCourseModuleByInstance200ResponseCmOutcomesInnerWithDefaults instantiates a new CoreCourseGetCourseModuleByInstance200ResponseCmOutcomesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreCourseGetCourseModuleByInstance200ResponseCmOutcomesInnerWithDefaults() *CoreCourseGetCourseModuleByInstance200ResponseCmOutcomesInner {
	this := CoreCourseGetCourseModuleByInstance200ResponseCmOutcomesInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CoreCourseGetCourseModuleByInstance200ResponseCmOutcomesInner) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCourseGetCourseModuleByInstance200ResponseCmOutcomesInner) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CoreCourseGetCourseModuleByInstance200ResponseCmOutcomesInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CoreCourseGetCourseModuleByInstance200ResponseCmOutcomesInner) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CoreCourseGetCourseModuleByInstance200ResponseCmOutcomesInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCourseGetCourseModuleByInstance200ResponseCmOutcomesInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CoreCourseGetCourseModuleByInstance200ResponseCmOutcomesInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CoreCourseGetCourseModuleByInstance200ResponseCmOutcomesInner) SetName(v string) {
	o.Name = &v
}

// GetScale returns the Scale field value if set, zero value otherwise.
func (o *CoreCourseGetCourseModuleByInstance200ResponseCmOutcomesInner) GetScale() string {
	if o == nil || IsNil(o.Scale) {
		var ret string
		return ret
	}
	return *o.Scale
}

// GetScaleOk returns a tuple with the Scale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCourseGetCourseModuleByInstance200ResponseCmOutcomesInner) GetScaleOk() (*string, bool) {
	if o == nil || IsNil(o.Scale) {
		return nil, false
	}
	return o.Scale, true
}

// HasScale returns a boolean if a field has been set.
func (o *CoreCourseGetCourseModuleByInstance200ResponseCmOutcomesInner) HasScale() bool {
	if o != nil && !IsNil(o.Scale) {
		return true
	}

	return false
}

// SetScale gets a reference to the given string and assigns it to the Scale field.
func (o *CoreCourseGetCourseModuleByInstance200ResponseCmOutcomesInner) SetScale(v string) {
	o.Scale = &v
}

func (o CoreCourseGetCourseModuleByInstance200ResponseCmOutcomesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreCourseGetCourseModuleByInstance200ResponseCmOutcomesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Scale) {
		toSerialize["scale"] = o.Scale
	}
	return toSerialize, nil
}

type NullableCoreCourseGetCourseModuleByInstance200ResponseCmOutcomesInner struct {
	value *CoreCourseGetCourseModuleByInstance200ResponseCmOutcomesInner
	isSet bool
}

func (v NullableCoreCourseGetCourseModuleByInstance200ResponseCmOutcomesInner) Get() *CoreCourseGetCourseModuleByInstance200ResponseCmOutcomesInner {
	return v.value
}

func (v *NullableCoreCourseGetCourseModuleByInstance200ResponseCmOutcomesInner) Set(val *CoreCourseGetCourseModuleByInstance200ResponseCmOutcomesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreCourseGetCourseModuleByInstance200ResponseCmOutcomesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreCourseGetCourseModuleByInstance200ResponseCmOutcomesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreCourseGetCourseModuleByInstance200ResponseCmOutcomesInner(val *CoreCourseGetCourseModuleByInstance200ResponseCmOutcomesInner) *NullableCoreCourseGetCourseModuleByInstance200ResponseCmOutcomesInner {
	return &NullableCoreCourseGetCourseModuleByInstance200ResponseCmOutcomesInner{value: val, isSet: true}
}

func (v NullableCoreCourseGetCourseModuleByInstance200ResponseCmOutcomesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreCourseGetCourseModuleByInstance200ResponseCmOutcomesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


