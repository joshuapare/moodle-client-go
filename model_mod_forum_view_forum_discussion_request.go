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

// checks if the ModForumViewForumDiscussionRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModForumViewForumDiscussionRequest{}

// ModForumViewForumDiscussionRequest struct for ModForumViewForumDiscussionRequest
type ModForumViewForumDiscussionRequest struct {
	// discussion id
	Discussionid int32 `json:"discussionid"`
}

type _ModForumViewForumDiscussionRequest ModForumViewForumDiscussionRequest

// NewModForumViewForumDiscussionRequest instantiates a new ModForumViewForumDiscussionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModForumViewForumDiscussionRequest(discussionid int32) *ModForumViewForumDiscussionRequest {
	this := ModForumViewForumDiscussionRequest{}
	this.Discussionid = discussionid
	return &this
}

// NewModForumViewForumDiscussionRequestWithDefaults instantiates a new ModForumViewForumDiscussionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModForumViewForumDiscussionRequestWithDefaults() *ModForumViewForumDiscussionRequest {
	this := ModForumViewForumDiscussionRequest{}
	var discussionid int32 = null
	this.Discussionid = discussionid
	return &this
}

// GetDiscussionid returns the Discussionid field value
func (o *ModForumViewForumDiscussionRequest) GetDiscussionid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Discussionid
}

// GetDiscussionidOk returns a tuple with the Discussionid field value
// and a boolean to check if the value has been set.
func (o *ModForumViewForumDiscussionRequest) GetDiscussionidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Discussionid, true
}

// SetDiscussionid sets field value
func (o *ModForumViewForumDiscussionRequest) SetDiscussionid(v int32) {
	o.Discussionid = v
}

func (o ModForumViewForumDiscussionRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModForumViewForumDiscussionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["discussionid"] = o.Discussionid
	return toSerialize, nil
}

func (o *ModForumViewForumDiscussionRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"discussionid",
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

	varModForumViewForumDiscussionRequest := _ModForumViewForumDiscussionRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModForumViewForumDiscussionRequest)

	if err != nil {
		return err
	}

	*o = ModForumViewForumDiscussionRequest(varModForumViewForumDiscussionRequest)

	return err
}

type NullableModForumViewForumDiscussionRequest struct {
	value *ModForumViewForumDiscussionRequest
	isSet bool
}

func (v NullableModForumViewForumDiscussionRequest) Get() *ModForumViewForumDiscussionRequest {
	return v.value
}

func (v *NullableModForumViewForumDiscussionRequest) Set(val *ModForumViewForumDiscussionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModForumViewForumDiscussionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModForumViewForumDiscussionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModForumViewForumDiscussionRequest(val *ModForumViewForumDiscussionRequest) *NullableModForumViewForumDiscussionRequest {
	return &NullableModForumViewForumDiscussionRequest{value: val, isSet: true}
}

func (v NullableModForumViewForumDiscussionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModForumViewForumDiscussionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


