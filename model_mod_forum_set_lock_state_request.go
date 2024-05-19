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

// checks if the ModForumSetLockStateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModForumSetLockStateRequest{}

// ModForumSetLockStateRequest struct for ModForumSetLockStateRequest
type ModForumSetLockStateRequest struct {
	// The discussion to lock / unlock
	Discussionid int32 `json:"discussionid"`
	// Forum that the discussion is in
	Forumid int32 `json:"forumid"`
	// The timestamp for the lock state
	Targetstate int32 `json:"targetstate"`
}

type _ModForumSetLockStateRequest ModForumSetLockStateRequest

// NewModForumSetLockStateRequest instantiates a new ModForumSetLockStateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModForumSetLockStateRequest(discussionid int32, forumid int32, targetstate int32) *ModForumSetLockStateRequest {
	this := ModForumSetLockStateRequest{}
	this.Discussionid = discussionid
	this.Forumid = forumid
	this.Targetstate = targetstate
	return &this
}

// NewModForumSetLockStateRequestWithDefaults instantiates a new ModForumSetLockStateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModForumSetLockStateRequestWithDefaults() *ModForumSetLockStateRequest {
	this := ModForumSetLockStateRequest{}
	var discussionid int32 = null
	this.Discussionid = discussionid
	var forumid int32 = null
	this.Forumid = forumid
	var targetstate int32 = null
	this.Targetstate = targetstate
	return &this
}

// GetDiscussionid returns the Discussionid field value
func (o *ModForumSetLockStateRequest) GetDiscussionid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Discussionid
}

// GetDiscussionidOk returns a tuple with the Discussionid field value
// and a boolean to check if the value has been set.
func (o *ModForumSetLockStateRequest) GetDiscussionidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Discussionid, true
}

// SetDiscussionid sets field value
func (o *ModForumSetLockStateRequest) SetDiscussionid(v int32) {
	o.Discussionid = v
}

// GetForumid returns the Forumid field value
func (o *ModForumSetLockStateRequest) GetForumid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Forumid
}

// GetForumidOk returns a tuple with the Forumid field value
// and a boolean to check if the value has been set.
func (o *ModForumSetLockStateRequest) GetForumidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Forumid, true
}

// SetForumid sets field value
func (o *ModForumSetLockStateRequest) SetForumid(v int32) {
	o.Forumid = v
}

// GetTargetstate returns the Targetstate field value
func (o *ModForumSetLockStateRequest) GetTargetstate() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Targetstate
}

// GetTargetstateOk returns a tuple with the Targetstate field value
// and a boolean to check if the value has been set.
func (o *ModForumSetLockStateRequest) GetTargetstateOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Targetstate, true
}

// SetTargetstate sets field value
func (o *ModForumSetLockStateRequest) SetTargetstate(v int32) {
	o.Targetstate = v
}

func (o ModForumSetLockStateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModForumSetLockStateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["discussionid"] = o.Discussionid
	toSerialize["forumid"] = o.Forumid
	toSerialize["targetstate"] = o.Targetstate
	return toSerialize, nil
}

func (o *ModForumSetLockStateRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"discussionid",
		"forumid",
		"targetstate",
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

	varModForumSetLockStateRequest := _ModForumSetLockStateRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModForumSetLockStateRequest)

	if err != nil {
		return err
	}

	*o = ModForumSetLockStateRequest(varModForumSetLockStateRequest)

	return err
}

type NullableModForumSetLockStateRequest struct {
	value *ModForumSetLockStateRequest
	isSet bool
}

func (v NullableModForumSetLockStateRequest) Get() *ModForumSetLockStateRequest {
	return v.value
}

func (v *NullableModForumSetLockStateRequest) Set(val *ModForumSetLockStateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModForumSetLockStateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModForumSetLockStateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModForumSetLockStateRequest(val *ModForumSetLockStateRequest) *NullableModForumSetLockStateRequest {
	return &NullableModForumSetLockStateRequest{value: val, isSet: true}
}

func (v NullableModForumSetLockStateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModForumSetLockStateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

