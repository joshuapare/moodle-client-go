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

// checks if the ModChatLoginUserRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModChatLoginUserRequest{}

// ModChatLoginUserRequest struct for ModChatLoginUserRequest
type ModChatLoginUserRequest struct {
	// chat instance id
	Chatid int32 `json:"chatid"`
	// group id, 0 means that the function will determine the user group
	Groupid *int32 `json:"groupid,omitempty"`
}

type _ModChatLoginUserRequest ModChatLoginUserRequest

// NewModChatLoginUserRequest instantiates a new ModChatLoginUserRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModChatLoginUserRequest(chatid int32) *ModChatLoginUserRequest {
	this := ModChatLoginUserRequest{}
	this.Chatid = chatid
	var groupid int32 = 0
	this.Groupid = &groupid
	return &this
}

// NewModChatLoginUserRequestWithDefaults instantiates a new ModChatLoginUserRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModChatLoginUserRequestWithDefaults() *ModChatLoginUserRequest {
	this := ModChatLoginUserRequest{}
	var chatid int32 = null
	this.Chatid = chatid
	var groupid int32 = 0
	this.Groupid = &groupid
	return &this
}

// GetChatid returns the Chatid field value
func (o *ModChatLoginUserRequest) GetChatid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Chatid
}

// GetChatidOk returns a tuple with the Chatid field value
// and a boolean to check if the value has been set.
func (o *ModChatLoginUserRequest) GetChatidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Chatid, true
}

// SetChatid sets field value
func (o *ModChatLoginUserRequest) SetChatid(v int32) {
	o.Chatid = v
}

// GetGroupid returns the Groupid field value if set, zero value otherwise.
func (o *ModChatLoginUserRequest) GetGroupid() int32 {
	if o == nil || IsNil(o.Groupid) {
		var ret int32
		return ret
	}
	return *o.Groupid
}

// GetGroupidOk returns a tuple with the Groupid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModChatLoginUserRequest) GetGroupidOk() (*int32, bool) {
	if o == nil || IsNil(o.Groupid) {
		return nil, false
	}
	return o.Groupid, true
}

// HasGroupid returns a boolean if a field has been set.
func (o *ModChatLoginUserRequest) HasGroupid() bool {
	if o != nil && !IsNil(o.Groupid) {
		return true
	}

	return false
}

// SetGroupid gets a reference to the given int32 and assigns it to the Groupid field.
func (o *ModChatLoginUserRequest) SetGroupid(v int32) {
	o.Groupid = &v
}

func (o ModChatLoginUserRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModChatLoginUserRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["chatid"] = o.Chatid
	if !IsNil(o.Groupid) {
		toSerialize["groupid"] = o.Groupid
	}
	return toSerialize, nil
}

func (o *ModChatLoginUserRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"chatid",
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

	varModChatLoginUserRequest := _ModChatLoginUserRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModChatLoginUserRequest)

	if err != nil {
		return err
	}

	*o = ModChatLoginUserRequest(varModChatLoginUserRequest)

	return err
}

type NullableModChatLoginUserRequest struct {
	value *ModChatLoginUserRequest
	isSet bool
}

func (v NullableModChatLoginUserRequest) Get() *ModChatLoginUserRequest {
	return v.value
}

func (v *NullableModChatLoginUserRequest) Set(val *ModChatLoginUserRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModChatLoginUserRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModChatLoginUserRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModChatLoginUserRequest(val *ModChatLoginUserRequest) *NullableModChatLoginUserRequest {
	return &NullableModChatLoginUserRequest{value: val, isSet: true}
}

func (v NullableModChatLoginUserRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModChatLoginUserRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


