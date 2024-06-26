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

// checks if the ModChatGetChatsByCoursesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModChatGetChatsByCoursesRequest{}

// ModChatGetChatsByCoursesRequest struct for ModChatGetChatsByCoursesRequest
type ModChatGetChatsByCoursesRequest struct {
	Courseids []map[string]interface{} `json:"courseids,omitempty"`
}

// NewModChatGetChatsByCoursesRequest instantiates a new ModChatGetChatsByCoursesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModChatGetChatsByCoursesRequest() *ModChatGetChatsByCoursesRequest {
	this := ModChatGetChatsByCoursesRequest{}
	return &this
}

// NewModChatGetChatsByCoursesRequestWithDefaults instantiates a new ModChatGetChatsByCoursesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModChatGetChatsByCoursesRequestWithDefaults() *ModChatGetChatsByCoursesRequest {
	this := ModChatGetChatsByCoursesRequest{}
	return &this
}

// GetCourseids returns the Courseids field value if set, zero value otherwise.
func (o *ModChatGetChatsByCoursesRequest) GetCourseids() []map[string]interface{} {
	if o == nil || IsNil(o.Courseids) {
		var ret []map[string]interface{}
		return ret
	}
	return o.Courseids
}

// GetCourseidsOk returns a tuple with the Courseids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModChatGetChatsByCoursesRequest) GetCourseidsOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Courseids) {
		return nil, false
	}
	return o.Courseids, true
}

// HasCourseids returns a boolean if a field has been set.
func (o *ModChatGetChatsByCoursesRequest) HasCourseids() bool {
	if o != nil && !IsNil(o.Courseids) {
		return true
	}

	return false
}

// SetCourseids gets a reference to the given []map[string]interface{} and assigns it to the Courseids field.
func (o *ModChatGetChatsByCoursesRequest) SetCourseids(v []map[string]interface{}) {
	o.Courseids = v
}

func (o ModChatGetChatsByCoursesRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModChatGetChatsByCoursesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Courseids) {
		toSerialize["courseids"] = o.Courseids
	}
	return toSerialize, nil
}

type NullableModChatGetChatsByCoursesRequest struct {
	value *ModChatGetChatsByCoursesRequest
	isSet bool
}

func (v NullableModChatGetChatsByCoursesRequest) Get() *ModChatGetChatsByCoursesRequest {
	return v.value
}

func (v *NullableModChatGetChatsByCoursesRequest) Set(val *ModChatGetChatsByCoursesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModChatGetChatsByCoursesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModChatGetChatsByCoursesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModChatGetChatsByCoursesRequest(val *ModChatGetChatsByCoursesRequest) *NullableModChatGetChatsByCoursesRequest {
	return &NullableModChatGetChatsByCoursesRequest{value: val, isSet: true}
}

func (v NullableModChatGetChatsByCoursesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModChatGetChatsByCoursesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


