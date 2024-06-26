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

// checks if the ModForumGetDiscussionPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModForumGetDiscussionPostRequest{}

// ModForumGetDiscussionPostRequest struct for ModForumGetDiscussionPostRequest
type ModForumGetDiscussionPostRequest struct {
	// Post to fetch.
	Postid int32 `json:"postid"`
}

type _ModForumGetDiscussionPostRequest ModForumGetDiscussionPostRequest

// NewModForumGetDiscussionPostRequest instantiates a new ModForumGetDiscussionPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModForumGetDiscussionPostRequest(postid int32) *ModForumGetDiscussionPostRequest {
	this := ModForumGetDiscussionPostRequest{}
	this.Postid = postid
	return &this
}

// NewModForumGetDiscussionPostRequestWithDefaults instantiates a new ModForumGetDiscussionPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModForumGetDiscussionPostRequestWithDefaults() *ModForumGetDiscussionPostRequest {
	this := ModForumGetDiscussionPostRequest{}
	var postid int32 = null
	this.Postid = postid
	return &this
}

// GetPostid returns the Postid field value
func (o *ModForumGetDiscussionPostRequest) GetPostid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Postid
}

// GetPostidOk returns a tuple with the Postid field value
// and a boolean to check if the value has been set.
func (o *ModForumGetDiscussionPostRequest) GetPostidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Postid, true
}

// SetPostid sets field value
func (o *ModForumGetDiscussionPostRequest) SetPostid(v int32) {
	o.Postid = v
}

func (o ModForumGetDiscussionPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModForumGetDiscussionPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["postid"] = o.Postid
	return toSerialize, nil
}

func (o *ModForumGetDiscussionPostRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"postid",
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

	varModForumGetDiscussionPostRequest := _ModForumGetDiscussionPostRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModForumGetDiscussionPostRequest)

	if err != nil {
		return err
	}

	*o = ModForumGetDiscussionPostRequest(varModForumGetDiscussionPostRequest)

	return err
}

type NullableModForumGetDiscussionPostRequest struct {
	value *ModForumGetDiscussionPostRequest
	isSet bool
}

func (v NullableModForumGetDiscussionPostRequest) Get() *ModForumGetDiscussionPostRequest {
	return v.value
}

func (v *NullableModForumGetDiscussionPostRequest) Set(val *ModForumGetDiscussionPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModForumGetDiscussionPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModForumGetDiscussionPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModForumGetDiscussionPostRequest(val *ModForumGetDiscussionPostRequest) *NullableModForumGetDiscussionPostRequest {
	return &NullableModForumGetDiscussionPostRequest{value: val, isSet: true}
}

func (v NullableModForumGetDiscussionPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModForumGetDiscussionPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


