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

// checks if the ModLtiGetToolProxiesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModLtiGetToolProxiesRequest{}

// ModLtiGetToolProxiesRequest struct for ModLtiGetToolProxiesRequest
type ModLtiGetToolProxiesRequest struct {
	// Orphaned tool types only
	Orphanedonly *bool `json:"orphanedonly,omitempty"`
}

// NewModLtiGetToolProxiesRequest instantiates a new ModLtiGetToolProxiesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModLtiGetToolProxiesRequest() *ModLtiGetToolProxiesRequest {
	this := ModLtiGetToolProxiesRequest{}
	var orphanedonly bool = 0
	this.Orphanedonly = &orphanedonly
	return &this
}

// NewModLtiGetToolProxiesRequestWithDefaults instantiates a new ModLtiGetToolProxiesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModLtiGetToolProxiesRequestWithDefaults() *ModLtiGetToolProxiesRequest {
	this := ModLtiGetToolProxiesRequest{}
	var orphanedonly bool = 0
	this.Orphanedonly = &orphanedonly
	return &this
}

// GetOrphanedonly returns the Orphanedonly field value if set, zero value otherwise.
func (o *ModLtiGetToolProxiesRequest) GetOrphanedonly() bool {
	if o == nil || IsNil(o.Orphanedonly) {
		var ret bool
		return ret
	}
	return *o.Orphanedonly
}

// GetOrphanedonlyOk returns a tuple with the Orphanedonly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModLtiGetToolProxiesRequest) GetOrphanedonlyOk() (*bool, bool) {
	if o == nil || IsNil(o.Orphanedonly) {
		return nil, false
	}
	return o.Orphanedonly, true
}

// HasOrphanedonly returns a boolean if a field has been set.
func (o *ModLtiGetToolProxiesRequest) HasOrphanedonly() bool {
	if o != nil && !IsNil(o.Orphanedonly) {
		return true
	}

	return false
}

// SetOrphanedonly gets a reference to the given bool and assigns it to the Orphanedonly field.
func (o *ModLtiGetToolProxiesRequest) SetOrphanedonly(v bool) {
	o.Orphanedonly = &v
}

func (o ModLtiGetToolProxiesRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModLtiGetToolProxiesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Orphanedonly) {
		toSerialize["orphanedonly"] = o.Orphanedonly
	}
	return toSerialize, nil
}

type NullableModLtiGetToolProxiesRequest struct {
	value *ModLtiGetToolProxiesRequest
	isSet bool
}

func (v NullableModLtiGetToolProxiesRequest) Get() *ModLtiGetToolProxiesRequest {
	return v.value
}

func (v *NullableModLtiGetToolProxiesRequest) Set(val *ModLtiGetToolProxiesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModLtiGetToolProxiesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModLtiGetToolProxiesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModLtiGetToolProxiesRequest(val *ModLtiGetToolProxiesRequest) *NullableModLtiGetToolProxiesRequest {
	return &NullableModLtiGetToolProxiesRequest{value: val, isSet: true}
}

func (v NullableModLtiGetToolProxiesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModLtiGetToolProxiesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

