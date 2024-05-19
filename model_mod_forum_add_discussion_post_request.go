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

// checks if the ModForumAddDiscussionPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModForumAddDiscussionPostRequest{}

// ModForumAddDiscussionPostRequest struct for ModForumAddDiscussionPostRequest
type ModForumAddDiscussionPostRequest struct {
	// new post message (html assumed if messageformat is not provided)
	Message string `json:"message"`
	// message format (1 = HTML, 0 = MOODLE, 2 = PLAIN, or 4 = MARKDOWN)
	Messageformat *int32 `json:"messageformat,omitempty"`
	Options []ModForumAddDiscussionPostRequestOptionsInner `json:"options,omitempty"`
	// the post id we are going to reply to                                                 (can be the initial discussion post
	Postid int32 `json:"postid"`
	// new post subject
	Subject string `json:"subject"`
}

type _ModForumAddDiscussionPostRequest ModForumAddDiscussionPostRequest

// NewModForumAddDiscussionPostRequest instantiates a new ModForumAddDiscussionPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModForumAddDiscussionPostRequest(message string, postid int32, subject string) *ModForumAddDiscussionPostRequest {
	this := ModForumAddDiscussionPostRequest{}
	this.Message = message
	var messageformat int32 = 1
	this.Messageformat = &messageformat
	this.Postid = postid
	this.Subject = subject
	return &this
}

// NewModForumAddDiscussionPostRequestWithDefaults instantiates a new ModForumAddDiscussionPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModForumAddDiscussionPostRequestWithDefaults() *ModForumAddDiscussionPostRequest {
	this := ModForumAddDiscussionPostRequest{}
	var message string = "null"
	this.Message = message
	var messageformat int32 = 1
	this.Messageformat = &messageformat
	var postid int32 = null
	this.Postid = postid
	var subject string = "null"
	this.Subject = subject
	return &this
}

// GetMessage returns the Message field value
func (o *ModForumAddDiscussionPostRequest) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *ModForumAddDiscussionPostRequest) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *ModForumAddDiscussionPostRequest) SetMessage(v string) {
	o.Message = v
}

// GetMessageformat returns the Messageformat field value if set, zero value otherwise.
func (o *ModForumAddDiscussionPostRequest) GetMessageformat() int32 {
	if o == nil || IsNil(o.Messageformat) {
		var ret int32
		return ret
	}
	return *o.Messageformat
}

// GetMessageformatOk returns a tuple with the Messageformat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModForumAddDiscussionPostRequest) GetMessageformatOk() (*int32, bool) {
	if o == nil || IsNil(o.Messageformat) {
		return nil, false
	}
	return o.Messageformat, true
}

// HasMessageformat returns a boolean if a field has been set.
func (o *ModForumAddDiscussionPostRequest) HasMessageformat() bool {
	if o != nil && !IsNil(o.Messageformat) {
		return true
	}

	return false
}

// SetMessageformat gets a reference to the given int32 and assigns it to the Messageformat field.
func (o *ModForumAddDiscussionPostRequest) SetMessageformat(v int32) {
	o.Messageformat = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *ModForumAddDiscussionPostRequest) GetOptions() []ModForumAddDiscussionPostRequestOptionsInner {
	if o == nil || IsNil(o.Options) {
		var ret []ModForumAddDiscussionPostRequestOptionsInner
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModForumAddDiscussionPostRequest) GetOptionsOk() ([]ModForumAddDiscussionPostRequestOptionsInner, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *ModForumAddDiscussionPostRequest) HasOptions() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given []ModForumAddDiscussionPostRequestOptionsInner and assigns it to the Options field.
func (o *ModForumAddDiscussionPostRequest) SetOptions(v []ModForumAddDiscussionPostRequestOptionsInner) {
	o.Options = v
}

// GetPostid returns the Postid field value
func (o *ModForumAddDiscussionPostRequest) GetPostid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Postid
}

// GetPostidOk returns a tuple with the Postid field value
// and a boolean to check if the value has been set.
func (o *ModForumAddDiscussionPostRequest) GetPostidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Postid, true
}

// SetPostid sets field value
func (o *ModForumAddDiscussionPostRequest) SetPostid(v int32) {
	o.Postid = v
}

// GetSubject returns the Subject field value
func (o *ModForumAddDiscussionPostRequest) GetSubject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value
// and a boolean to check if the value has been set.
func (o *ModForumAddDiscussionPostRequest) GetSubjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subject, true
}

// SetSubject sets field value
func (o *ModForumAddDiscussionPostRequest) SetSubject(v string) {
	o.Subject = v
}

func (o ModForumAddDiscussionPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModForumAddDiscussionPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["message"] = o.Message
	if !IsNil(o.Messageformat) {
		toSerialize["messageformat"] = o.Messageformat
	}
	if !IsNil(o.Options) {
		toSerialize["options"] = o.Options
	}
	toSerialize["postid"] = o.Postid
	toSerialize["subject"] = o.Subject
	return toSerialize, nil
}

func (o *ModForumAddDiscussionPostRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"message",
		"postid",
		"subject",
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

	varModForumAddDiscussionPostRequest := _ModForumAddDiscussionPostRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModForumAddDiscussionPostRequest)

	if err != nil {
		return err
	}

	*o = ModForumAddDiscussionPostRequest(varModForumAddDiscussionPostRequest)

	return err
}

type NullableModForumAddDiscussionPostRequest struct {
	value *ModForumAddDiscussionPostRequest
	isSet bool
}

func (v NullableModForumAddDiscussionPostRequest) Get() *ModForumAddDiscussionPostRequest {
	return v.value
}

func (v *NullableModForumAddDiscussionPostRequest) Set(val *ModForumAddDiscussionPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModForumAddDiscussionPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModForumAddDiscussionPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModForumAddDiscussionPostRequest(val *ModForumAddDiscussionPostRequest) *NullableModForumAddDiscussionPostRequest {
	return &NullableModForumAddDiscussionPostRequest{value: val, isSet: true}
}

func (v NullableModForumAddDiscussionPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModForumAddDiscussionPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

