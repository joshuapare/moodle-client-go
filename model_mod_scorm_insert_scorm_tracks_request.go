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

// checks if the ModScormInsertScormTracksRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModScormInsertScormTracksRequest{}

// ModScormInsertScormTracksRequest struct for ModScormInsertScormTracksRequest
type ModScormInsertScormTracksRequest struct {
	// attempt number
	Attempt int32 `json:"attempt"`
	// SCO id
	Scoid int32 `json:"scoid"`
	Tracks []ModScormGetScormUserData200ResponseDataInnerDefaultdataInner `json:"tracks"`
}

type _ModScormInsertScormTracksRequest ModScormInsertScormTracksRequest

// NewModScormInsertScormTracksRequest instantiates a new ModScormInsertScormTracksRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModScormInsertScormTracksRequest(attempt int32, scoid int32, tracks []ModScormGetScormUserData200ResponseDataInnerDefaultdataInner) *ModScormInsertScormTracksRequest {
	this := ModScormInsertScormTracksRequest{}
	this.Attempt = attempt
	this.Scoid = scoid
	this.Tracks = tracks
	return &this
}

// NewModScormInsertScormTracksRequestWithDefaults instantiates a new ModScormInsertScormTracksRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModScormInsertScormTracksRequestWithDefaults() *ModScormInsertScormTracksRequest {
	this := ModScormInsertScormTracksRequest{}
	var scoid int32 = null
	this.Scoid = scoid
	return &this
}

// GetAttempt returns the Attempt field value
func (o *ModScormInsertScormTracksRequest) GetAttempt() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Attempt
}

// GetAttemptOk returns a tuple with the Attempt field value
// and a boolean to check if the value has been set.
func (o *ModScormInsertScormTracksRequest) GetAttemptOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attempt, true
}

// SetAttempt sets field value
func (o *ModScormInsertScormTracksRequest) SetAttempt(v int32) {
	o.Attempt = v
}

// GetScoid returns the Scoid field value
func (o *ModScormInsertScormTracksRequest) GetScoid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Scoid
}

// GetScoidOk returns a tuple with the Scoid field value
// and a boolean to check if the value has been set.
func (o *ModScormInsertScormTracksRequest) GetScoidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Scoid, true
}

// SetScoid sets field value
func (o *ModScormInsertScormTracksRequest) SetScoid(v int32) {
	o.Scoid = v
}

// GetTracks returns the Tracks field value
func (o *ModScormInsertScormTracksRequest) GetTracks() []ModScormGetScormUserData200ResponseDataInnerDefaultdataInner {
	if o == nil {
		var ret []ModScormGetScormUserData200ResponseDataInnerDefaultdataInner
		return ret
	}

	return o.Tracks
}

// GetTracksOk returns a tuple with the Tracks field value
// and a boolean to check if the value has been set.
func (o *ModScormInsertScormTracksRequest) GetTracksOk() ([]ModScormGetScormUserData200ResponseDataInnerDefaultdataInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tracks, true
}

// SetTracks sets field value
func (o *ModScormInsertScormTracksRequest) SetTracks(v []ModScormGetScormUserData200ResponseDataInnerDefaultdataInner) {
	o.Tracks = v
}

func (o ModScormInsertScormTracksRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModScormInsertScormTracksRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["attempt"] = o.Attempt
	toSerialize["scoid"] = o.Scoid
	toSerialize["tracks"] = o.Tracks
	return toSerialize, nil
}

func (o *ModScormInsertScormTracksRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"attempt",
		"scoid",
		"tracks",
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

	varModScormInsertScormTracksRequest := _ModScormInsertScormTracksRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModScormInsertScormTracksRequest)

	if err != nil {
		return err
	}

	*o = ModScormInsertScormTracksRequest(varModScormInsertScormTracksRequest)

	return err
}

type NullableModScormInsertScormTracksRequest struct {
	value *ModScormInsertScormTracksRequest
	isSet bool
}

func (v NullableModScormInsertScormTracksRequest) Get() *ModScormInsertScormTracksRequest {
	return v.value
}

func (v *NullableModScormInsertScormTracksRequest) Set(val *ModScormInsertScormTracksRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModScormInsertScormTracksRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModScormInsertScormTracksRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModScormInsertScormTracksRequest(val *ModScormInsertScormTracksRequest) *NullableModScormInsertScormTracksRequest {
	return &NullableModScormInsertScormTracksRequest{value: val, isSet: true}
}

func (v NullableModScormInsertScormTracksRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModScormInsertScormTracksRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


