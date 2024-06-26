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

// checks if the CoreCourseGetUpdatesSince200ResponseInstancesInnerUpdatesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreCourseGetUpdatesSince200ResponseInstancesInnerUpdatesInner{}

// CoreCourseGetUpdatesSince200ResponseInstancesInnerUpdatesInner struct for CoreCourseGetUpdatesSince200ResponseInstancesInnerUpdatesInner
type CoreCourseGetUpdatesSince200ResponseInstancesInnerUpdatesInner struct {
	Itemids []map[string]interface{} `json:"itemids,omitempty"`
	// Name of the area updated.
	Name *string `json:"name,omitempty"`
	// Last time was updated
	Timeupdated *int32 `json:"timeupdated,omitempty"`
}

// NewCoreCourseGetUpdatesSince200ResponseInstancesInnerUpdatesInner instantiates a new CoreCourseGetUpdatesSince200ResponseInstancesInnerUpdatesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreCourseGetUpdatesSince200ResponseInstancesInnerUpdatesInner() *CoreCourseGetUpdatesSince200ResponseInstancesInnerUpdatesInner {
	this := CoreCourseGetUpdatesSince200ResponseInstancesInnerUpdatesInner{}
	return &this
}

// NewCoreCourseGetUpdatesSince200ResponseInstancesInnerUpdatesInnerWithDefaults instantiates a new CoreCourseGetUpdatesSince200ResponseInstancesInnerUpdatesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreCourseGetUpdatesSince200ResponseInstancesInnerUpdatesInnerWithDefaults() *CoreCourseGetUpdatesSince200ResponseInstancesInnerUpdatesInner {
	this := CoreCourseGetUpdatesSince200ResponseInstancesInnerUpdatesInner{}
	return &this
}

// GetItemids returns the Itemids field value if set, zero value otherwise.
func (o *CoreCourseGetUpdatesSince200ResponseInstancesInnerUpdatesInner) GetItemids() []map[string]interface{} {
	if o == nil || IsNil(o.Itemids) {
		var ret []map[string]interface{}
		return ret
	}
	return o.Itemids
}

// GetItemidsOk returns a tuple with the Itemids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCourseGetUpdatesSince200ResponseInstancesInnerUpdatesInner) GetItemidsOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Itemids) {
		return nil, false
	}
	return o.Itemids, true
}

// HasItemids returns a boolean if a field has been set.
func (o *CoreCourseGetUpdatesSince200ResponseInstancesInnerUpdatesInner) HasItemids() bool {
	if o != nil && !IsNil(o.Itemids) {
		return true
	}

	return false
}

// SetItemids gets a reference to the given []map[string]interface{} and assigns it to the Itemids field.
func (o *CoreCourseGetUpdatesSince200ResponseInstancesInnerUpdatesInner) SetItemids(v []map[string]interface{}) {
	o.Itemids = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CoreCourseGetUpdatesSince200ResponseInstancesInnerUpdatesInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCourseGetUpdatesSince200ResponseInstancesInnerUpdatesInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CoreCourseGetUpdatesSince200ResponseInstancesInnerUpdatesInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CoreCourseGetUpdatesSince200ResponseInstancesInnerUpdatesInner) SetName(v string) {
	o.Name = &v
}

// GetTimeupdated returns the Timeupdated field value if set, zero value otherwise.
func (o *CoreCourseGetUpdatesSince200ResponseInstancesInnerUpdatesInner) GetTimeupdated() int32 {
	if o == nil || IsNil(o.Timeupdated) {
		var ret int32
		return ret
	}
	return *o.Timeupdated
}

// GetTimeupdatedOk returns a tuple with the Timeupdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCourseGetUpdatesSince200ResponseInstancesInnerUpdatesInner) GetTimeupdatedOk() (*int32, bool) {
	if o == nil || IsNil(o.Timeupdated) {
		return nil, false
	}
	return o.Timeupdated, true
}

// HasTimeupdated returns a boolean if a field has been set.
func (o *CoreCourseGetUpdatesSince200ResponseInstancesInnerUpdatesInner) HasTimeupdated() bool {
	if o != nil && !IsNil(o.Timeupdated) {
		return true
	}

	return false
}

// SetTimeupdated gets a reference to the given int32 and assigns it to the Timeupdated field.
func (o *CoreCourseGetUpdatesSince200ResponseInstancesInnerUpdatesInner) SetTimeupdated(v int32) {
	o.Timeupdated = &v
}

func (o CoreCourseGetUpdatesSince200ResponseInstancesInnerUpdatesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreCourseGetUpdatesSince200ResponseInstancesInnerUpdatesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Itemids) {
		toSerialize["itemids"] = o.Itemids
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Timeupdated) {
		toSerialize["timeupdated"] = o.Timeupdated
	}
	return toSerialize, nil
}

type NullableCoreCourseGetUpdatesSince200ResponseInstancesInnerUpdatesInner struct {
	value *CoreCourseGetUpdatesSince200ResponseInstancesInnerUpdatesInner
	isSet bool
}

func (v NullableCoreCourseGetUpdatesSince200ResponseInstancesInnerUpdatesInner) Get() *CoreCourseGetUpdatesSince200ResponseInstancesInnerUpdatesInner {
	return v.value
}

func (v *NullableCoreCourseGetUpdatesSince200ResponseInstancesInnerUpdatesInner) Set(val *CoreCourseGetUpdatesSince200ResponseInstancesInnerUpdatesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreCourseGetUpdatesSince200ResponseInstancesInnerUpdatesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreCourseGetUpdatesSince200ResponseInstancesInnerUpdatesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreCourseGetUpdatesSince200ResponseInstancesInnerUpdatesInner(val *CoreCourseGetUpdatesSince200ResponseInstancesInnerUpdatesInner) *NullableCoreCourseGetUpdatesSince200ResponseInstancesInnerUpdatesInner {
	return &NullableCoreCourseGetUpdatesSince200ResponseInstancesInnerUpdatesInner{value: val, isSet: true}
}

func (v NullableCoreCourseGetUpdatesSince200ResponseInstancesInnerUpdatesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreCourseGetUpdatesSince200ResponseInstancesInnerUpdatesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


