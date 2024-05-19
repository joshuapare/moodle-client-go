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

// checks if the ModChatGetSessionMessagesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModChatGetSessionMessagesRequest{}

// ModChatGetSessionMessagesRequest struct for ModChatGetSessionMessagesRequest
type ModChatGetSessionMessagesRequest struct {
	// Chat instance id.
	Chatid int32 `json:"chatid"`
	// Get messages from users in this group.                                                 0 means that the function will determine the user group
	Groupid *int32 `json:"groupid,omitempty"`
	// The session end time (timestamp).
	Sessionend int32 `json:"sessionend"`
	// The session start time (timestamp).
	Sessionstart int32 `json:"sessionstart"`
}

type _ModChatGetSessionMessagesRequest ModChatGetSessionMessagesRequest

// NewModChatGetSessionMessagesRequest instantiates a new ModChatGetSessionMessagesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModChatGetSessionMessagesRequest(chatid int32, sessionend int32, sessionstart int32) *ModChatGetSessionMessagesRequest {
	this := ModChatGetSessionMessagesRequest{}
	this.Chatid = chatid
	var groupid int32 = 0
	this.Groupid = &groupid
	this.Sessionend = sessionend
	this.Sessionstart = sessionstart
	return &this
}

// NewModChatGetSessionMessagesRequestWithDefaults instantiates a new ModChatGetSessionMessagesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModChatGetSessionMessagesRequestWithDefaults() *ModChatGetSessionMessagesRequest {
	this := ModChatGetSessionMessagesRequest{}
	var chatid int32 = null
	this.Chatid = chatid
	var groupid int32 = 0
	this.Groupid = &groupid
	var sessionend int32 = null
	this.Sessionend = sessionend
	var sessionstart int32 = null
	this.Sessionstart = sessionstart
	return &this
}

// GetChatid returns the Chatid field value
func (o *ModChatGetSessionMessagesRequest) GetChatid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Chatid
}

// GetChatidOk returns a tuple with the Chatid field value
// and a boolean to check if the value has been set.
func (o *ModChatGetSessionMessagesRequest) GetChatidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Chatid, true
}

// SetChatid sets field value
func (o *ModChatGetSessionMessagesRequest) SetChatid(v int32) {
	o.Chatid = v
}

// GetGroupid returns the Groupid field value if set, zero value otherwise.
func (o *ModChatGetSessionMessagesRequest) GetGroupid() int32 {
	if o == nil || IsNil(o.Groupid) {
		var ret int32
		return ret
	}
	return *o.Groupid
}

// GetGroupidOk returns a tuple with the Groupid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModChatGetSessionMessagesRequest) GetGroupidOk() (*int32, bool) {
	if o == nil || IsNil(o.Groupid) {
		return nil, false
	}
	return o.Groupid, true
}

// HasGroupid returns a boolean if a field has been set.
func (o *ModChatGetSessionMessagesRequest) HasGroupid() bool {
	if o != nil && !IsNil(o.Groupid) {
		return true
	}

	return false
}

// SetGroupid gets a reference to the given int32 and assigns it to the Groupid field.
func (o *ModChatGetSessionMessagesRequest) SetGroupid(v int32) {
	o.Groupid = &v
}

// GetSessionend returns the Sessionend field value
func (o *ModChatGetSessionMessagesRequest) GetSessionend() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Sessionend
}

// GetSessionendOk returns a tuple with the Sessionend field value
// and a boolean to check if the value has been set.
func (o *ModChatGetSessionMessagesRequest) GetSessionendOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sessionend, true
}

// SetSessionend sets field value
func (o *ModChatGetSessionMessagesRequest) SetSessionend(v int32) {
	o.Sessionend = v
}

// GetSessionstart returns the Sessionstart field value
func (o *ModChatGetSessionMessagesRequest) GetSessionstart() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Sessionstart
}

// GetSessionstartOk returns a tuple with the Sessionstart field value
// and a boolean to check if the value has been set.
func (o *ModChatGetSessionMessagesRequest) GetSessionstartOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sessionstart, true
}

// SetSessionstart sets field value
func (o *ModChatGetSessionMessagesRequest) SetSessionstart(v int32) {
	o.Sessionstart = v
}

func (o ModChatGetSessionMessagesRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModChatGetSessionMessagesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["chatid"] = o.Chatid
	if !IsNil(o.Groupid) {
		toSerialize["groupid"] = o.Groupid
	}
	toSerialize["sessionend"] = o.Sessionend
	toSerialize["sessionstart"] = o.Sessionstart
	return toSerialize, nil
}

func (o *ModChatGetSessionMessagesRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"chatid",
		"sessionend",
		"sessionstart",
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

	varModChatGetSessionMessagesRequest := _ModChatGetSessionMessagesRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModChatGetSessionMessagesRequest)

	if err != nil {
		return err
	}

	*o = ModChatGetSessionMessagesRequest(varModChatGetSessionMessagesRequest)

	return err
}

type NullableModChatGetSessionMessagesRequest struct {
	value *ModChatGetSessionMessagesRequest
	isSet bool
}

func (v NullableModChatGetSessionMessagesRequest) Get() *ModChatGetSessionMessagesRequest {
	return v.value
}

func (v *NullableModChatGetSessionMessagesRequest) Set(val *ModChatGetSessionMessagesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModChatGetSessionMessagesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModChatGetSessionMessagesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModChatGetSessionMessagesRequest(val *ModChatGetSessionMessagesRequest) *NullableModChatGetSessionMessagesRequest {
	return &NullableModChatGetSessionMessagesRequest{value: val, isSet: true}
}

func (v NullableModChatGetSessionMessagesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModChatGetSessionMessagesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


