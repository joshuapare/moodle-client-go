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

// checks if the ModBigbluebuttonbnGetRecordingsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModBigbluebuttonbnGetRecordingsRequest{}

// ModBigbluebuttonbnGetRecordingsRequest struct for ModBigbluebuttonbnGetRecordingsRequest
type ModBigbluebuttonbnGetRecordingsRequest struct {
	// bigbluebuttonbn instance id
	Bigbluebuttonbnid int32 `json:"bigbluebuttonbnid"`
	// Group ID
	Groupid *int32 `json:"groupid,omitempty"`
	// a set of enabled tools
	Tools *string `json:"tools,omitempty"`
}

type _ModBigbluebuttonbnGetRecordingsRequest ModBigbluebuttonbnGetRecordingsRequest

// NewModBigbluebuttonbnGetRecordingsRequest instantiates a new ModBigbluebuttonbnGetRecordingsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModBigbluebuttonbnGetRecordingsRequest(bigbluebuttonbnid int32) *ModBigbluebuttonbnGetRecordingsRequest {
	this := ModBigbluebuttonbnGetRecordingsRequest{}
	this.Bigbluebuttonbnid = bigbluebuttonbnid
	var groupid int32 = null
	this.Groupid = &groupid
	var tools string = "protect,unprotect,publish,unpublish,delete"
	this.Tools = &tools
	return &this
}

// NewModBigbluebuttonbnGetRecordingsRequestWithDefaults instantiates a new ModBigbluebuttonbnGetRecordingsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModBigbluebuttonbnGetRecordingsRequestWithDefaults() *ModBigbluebuttonbnGetRecordingsRequest {
	this := ModBigbluebuttonbnGetRecordingsRequest{}
	var groupid int32 = null
	this.Groupid = &groupid
	var tools string = "protect,unprotect,publish,unpublish,delete"
	this.Tools = &tools
	return &this
}

// GetBigbluebuttonbnid returns the Bigbluebuttonbnid field value
func (o *ModBigbluebuttonbnGetRecordingsRequest) GetBigbluebuttonbnid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Bigbluebuttonbnid
}

// GetBigbluebuttonbnidOk returns a tuple with the Bigbluebuttonbnid field value
// and a boolean to check if the value has been set.
func (o *ModBigbluebuttonbnGetRecordingsRequest) GetBigbluebuttonbnidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Bigbluebuttonbnid, true
}

// SetBigbluebuttonbnid sets field value
func (o *ModBigbluebuttonbnGetRecordingsRequest) SetBigbluebuttonbnid(v int32) {
	o.Bigbluebuttonbnid = v
}

// GetGroupid returns the Groupid field value if set, zero value otherwise.
func (o *ModBigbluebuttonbnGetRecordingsRequest) GetGroupid() int32 {
	if o == nil || IsNil(o.Groupid) {
		var ret int32
		return ret
	}
	return *o.Groupid
}

// GetGroupidOk returns a tuple with the Groupid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModBigbluebuttonbnGetRecordingsRequest) GetGroupidOk() (*int32, bool) {
	if o == nil || IsNil(o.Groupid) {
		return nil, false
	}
	return o.Groupid, true
}

// HasGroupid returns a boolean if a field has been set.
func (o *ModBigbluebuttonbnGetRecordingsRequest) HasGroupid() bool {
	if o != nil && !IsNil(o.Groupid) {
		return true
	}

	return false
}

// SetGroupid gets a reference to the given int32 and assigns it to the Groupid field.
func (o *ModBigbluebuttonbnGetRecordingsRequest) SetGroupid(v int32) {
	o.Groupid = &v
}

// GetTools returns the Tools field value if set, zero value otherwise.
func (o *ModBigbluebuttonbnGetRecordingsRequest) GetTools() string {
	if o == nil || IsNil(o.Tools) {
		var ret string
		return ret
	}
	return *o.Tools
}

// GetToolsOk returns a tuple with the Tools field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModBigbluebuttonbnGetRecordingsRequest) GetToolsOk() (*string, bool) {
	if o == nil || IsNil(o.Tools) {
		return nil, false
	}
	return o.Tools, true
}

// HasTools returns a boolean if a field has been set.
func (o *ModBigbluebuttonbnGetRecordingsRequest) HasTools() bool {
	if o != nil && !IsNil(o.Tools) {
		return true
	}

	return false
}

// SetTools gets a reference to the given string and assigns it to the Tools field.
func (o *ModBigbluebuttonbnGetRecordingsRequest) SetTools(v string) {
	o.Tools = &v
}

func (o ModBigbluebuttonbnGetRecordingsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModBigbluebuttonbnGetRecordingsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["bigbluebuttonbnid"] = o.Bigbluebuttonbnid
	if !IsNil(o.Groupid) {
		toSerialize["groupid"] = o.Groupid
	}
	if !IsNil(o.Tools) {
		toSerialize["tools"] = o.Tools
	}
	return toSerialize, nil
}

func (o *ModBigbluebuttonbnGetRecordingsRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"bigbluebuttonbnid",
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

	varModBigbluebuttonbnGetRecordingsRequest := _ModBigbluebuttonbnGetRecordingsRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModBigbluebuttonbnGetRecordingsRequest)

	if err != nil {
		return err
	}

	*o = ModBigbluebuttonbnGetRecordingsRequest(varModBigbluebuttonbnGetRecordingsRequest)

	return err
}

type NullableModBigbluebuttonbnGetRecordingsRequest struct {
	value *ModBigbluebuttonbnGetRecordingsRequest
	isSet bool
}

func (v NullableModBigbluebuttonbnGetRecordingsRequest) Get() *ModBigbluebuttonbnGetRecordingsRequest {
	return v.value
}

func (v *NullableModBigbluebuttonbnGetRecordingsRequest) Set(val *ModBigbluebuttonbnGetRecordingsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModBigbluebuttonbnGetRecordingsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModBigbluebuttonbnGetRecordingsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModBigbluebuttonbnGetRecordingsRequest(val *ModBigbluebuttonbnGetRecordingsRequest) *NullableModBigbluebuttonbnGetRecordingsRequest {
	return &NullableModBigbluebuttonbnGetRecordingsRequest{value: val, isSet: true}
}

func (v NullableModBigbluebuttonbnGetRecordingsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModBigbluebuttonbnGetRecordingsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

