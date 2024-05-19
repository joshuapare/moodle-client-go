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

// checks if the ModScormGetScormUserDataRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModScormGetScormUserDataRequest{}

// ModScormGetScormUserDataRequest struct for ModScormGetScormUserDataRequest
type ModScormGetScormUserDataRequest struct {
	// attempt number
	Attempt int32 `json:"attempt"`
	// scorm instance id
	Scormid int32 `json:"scormid"`
}

type _ModScormGetScormUserDataRequest ModScormGetScormUserDataRequest

// NewModScormGetScormUserDataRequest instantiates a new ModScormGetScormUserDataRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModScormGetScormUserDataRequest(attempt int32, scormid int32) *ModScormGetScormUserDataRequest {
	this := ModScormGetScormUserDataRequest{}
	this.Attempt = attempt
	this.Scormid = scormid
	return &this
}

// NewModScormGetScormUserDataRequestWithDefaults instantiates a new ModScormGetScormUserDataRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModScormGetScormUserDataRequestWithDefaults() *ModScormGetScormUserDataRequest {
	this := ModScormGetScormUserDataRequest{}
	return &this
}

// GetAttempt returns the Attempt field value
func (o *ModScormGetScormUserDataRequest) GetAttempt() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Attempt
}

// GetAttemptOk returns a tuple with the Attempt field value
// and a boolean to check if the value has been set.
func (o *ModScormGetScormUserDataRequest) GetAttemptOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attempt, true
}

// SetAttempt sets field value
func (o *ModScormGetScormUserDataRequest) SetAttempt(v int32) {
	o.Attempt = v
}

// GetScormid returns the Scormid field value
func (o *ModScormGetScormUserDataRequest) GetScormid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Scormid
}

// GetScormidOk returns a tuple with the Scormid field value
// and a boolean to check if the value has been set.
func (o *ModScormGetScormUserDataRequest) GetScormidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Scormid, true
}

// SetScormid sets field value
func (o *ModScormGetScormUserDataRequest) SetScormid(v int32) {
	o.Scormid = v
}

func (o ModScormGetScormUserDataRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModScormGetScormUserDataRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["attempt"] = o.Attempt
	toSerialize["scormid"] = o.Scormid
	return toSerialize, nil
}

func (o *ModScormGetScormUserDataRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"attempt",
		"scormid",
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

	varModScormGetScormUserDataRequest := _ModScormGetScormUserDataRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModScormGetScormUserDataRequest)

	if err != nil {
		return err
	}

	*o = ModScormGetScormUserDataRequest(varModScormGetScormUserDataRequest)

	return err
}

type NullableModScormGetScormUserDataRequest struct {
	value *ModScormGetScormUserDataRequest
	isSet bool
}

func (v NullableModScormGetScormUserDataRequest) Get() *ModScormGetScormUserDataRequest {
	return v.value
}

func (v *NullableModScormGetScormUserDataRequest) Set(val *ModScormGetScormUserDataRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModScormGetScormUserDataRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModScormGetScormUserDataRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModScormGetScormUserDataRequest(val *ModScormGetScormUserDataRequest) *NullableModScormGetScormUserDataRequest {
	return &NullableModScormGetScormUserDataRequest{value: val, isSet: true}
}

func (v NullableModScormGetScormUserDataRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModScormGetScormUserDataRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

