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

// checks if the ModForumToggleFavouriteStateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModForumToggleFavouriteStateRequest{}

// ModForumToggleFavouriteStateRequest struct for ModForumToggleFavouriteStateRequest
type ModForumToggleFavouriteStateRequest struct {
	// The discussion to subscribe or unsubscribe
	Discussionid int32 `json:"discussionid"`
	// The target state
	Targetstate bool `json:"targetstate"`
}

type _ModForumToggleFavouriteStateRequest ModForumToggleFavouriteStateRequest

// NewModForumToggleFavouriteStateRequest instantiates a new ModForumToggleFavouriteStateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModForumToggleFavouriteStateRequest(discussionid int32, targetstate bool) *ModForumToggleFavouriteStateRequest {
	this := ModForumToggleFavouriteStateRequest{}
	this.Discussionid = discussionid
	this.Targetstate = targetstate
	return &this
}

// NewModForumToggleFavouriteStateRequestWithDefaults instantiates a new ModForumToggleFavouriteStateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModForumToggleFavouriteStateRequestWithDefaults() *ModForumToggleFavouriteStateRequest {
	this := ModForumToggleFavouriteStateRequest{}
	return &this
}

// GetDiscussionid returns the Discussionid field value
func (o *ModForumToggleFavouriteStateRequest) GetDiscussionid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Discussionid
}

// GetDiscussionidOk returns a tuple with the Discussionid field value
// and a boolean to check if the value has been set.
func (o *ModForumToggleFavouriteStateRequest) GetDiscussionidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Discussionid, true
}

// SetDiscussionid sets field value
func (o *ModForumToggleFavouriteStateRequest) SetDiscussionid(v int32) {
	o.Discussionid = v
}

// GetTargetstate returns the Targetstate field value
func (o *ModForumToggleFavouriteStateRequest) GetTargetstate() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Targetstate
}

// GetTargetstateOk returns a tuple with the Targetstate field value
// and a boolean to check if the value has been set.
func (o *ModForumToggleFavouriteStateRequest) GetTargetstateOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Targetstate, true
}

// SetTargetstate sets field value
func (o *ModForumToggleFavouriteStateRequest) SetTargetstate(v bool) {
	o.Targetstate = v
}

func (o ModForumToggleFavouriteStateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModForumToggleFavouriteStateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["discussionid"] = o.Discussionid
	toSerialize["targetstate"] = o.Targetstate
	return toSerialize, nil
}

func (o *ModForumToggleFavouriteStateRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"discussionid",
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

	varModForumToggleFavouriteStateRequest := _ModForumToggleFavouriteStateRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModForumToggleFavouriteStateRequest)

	if err != nil {
		return err
	}

	*o = ModForumToggleFavouriteStateRequest(varModForumToggleFavouriteStateRequest)

	return err
}

type NullableModForumToggleFavouriteStateRequest struct {
	value *ModForumToggleFavouriteStateRequest
	isSet bool
}

func (v NullableModForumToggleFavouriteStateRequest) Get() *ModForumToggleFavouriteStateRequest {
	return v.value
}

func (v *NullableModForumToggleFavouriteStateRequest) Set(val *ModForumToggleFavouriteStateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModForumToggleFavouriteStateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModForumToggleFavouriteStateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModForumToggleFavouriteStateRequest(val *ModForumToggleFavouriteStateRequest) *NullableModForumToggleFavouriteStateRequest {
	return &NullableModForumToggleFavouriteStateRequest{value: val, isSet: true}
}

func (v NullableModForumToggleFavouriteStateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModForumToggleFavouriteStateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


