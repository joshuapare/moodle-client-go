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

// checks if the ModForumGetDiscussionPostsByUseridRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModForumGetDiscussionPostsByUseridRequest{}

// ModForumGetDiscussionPostsByUseridRequest struct for ModForumGetDiscussionPostsByUseridRequest
type ModForumGetDiscussionPostsByUseridRequest struct {
	// The ID of the module of which to fetch items.
	Cmid int32 `json:"cmid"`
	// Sort by this element: id, created or modified
	Sortby *string `json:"sortby,omitempty"`
	// Sort direction: ASC or DESC
	Sortdirection *string `json:"sortdirection,omitempty"`
	// The ID of the user of whom to fetch posts.
	Userid int32 `json:"userid"`
}

type _ModForumGetDiscussionPostsByUseridRequest ModForumGetDiscussionPostsByUseridRequest

// NewModForumGetDiscussionPostsByUseridRequest instantiates a new ModForumGetDiscussionPostsByUseridRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModForumGetDiscussionPostsByUseridRequest(cmid int32, userid int32) *ModForumGetDiscussionPostsByUseridRequest {
	this := ModForumGetDiscussionPostsByUseridRequest{}
	this.Cmid = cmid
	var sortby string = "created"
	this.Sortby = &sortby
	var sortdirection string = "DESC"
	this.Sortdirection = &sortdirection
	this.Userid = userid
	return &this
}

// NewModForumGetDiscussionPostsByUseridRequestWithDefaults instantiates a new ModForumGetDiscussionPostsByUseridRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModForumGetDiscussionPostsByUseridRequestWithDefaults() *ModForumGetDiscussionPostsByUseridRequest {
	this := ModForumGetDiscussionPostsByUseridRequest{}
	var cmid int32 = null
	this.Cmid = cmid
	var sortby string = "created"
	this.Sortby = &sortby
	var sortdirection string = "DESC"
	this.Sortdirection = &sortdirection
	var userid int32 = null
	this.Userid = userid
	return &this
}

// GetCmid returns the Cmid field value
func (o *ModForumGetDiscussionPostsByUseridRequest) GetCmid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Cmid
}

// GetCmidOk returns a tuple with the Cmid field value
// and a boolean to check if the value has been set.
func (o *ModForumGetDiscussionPostsByUseridRequest) GetCmidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cmid, true
}

// SetCmid sets field value
func (o *ModForumGetDiscussionPostsByUseridRequest) SetCmid(v int32) {
	o.Cmid = v
}

// GetSortby returns the Sortby field value if set, zero value otherwise.
func (o *ModForumGetDiscussionPostsByUseridRequest) GetSortby() string {
	if o == nil || IsNil(o.Sortby) {
		var ret string
		return ret
	}
	return *o.Sortby
}

// GetSortbyOk returns a tuple with the Sortby field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModForumGetDiscussionPostsByUseridRequest) GetSortbyOk() (*string, bool) {
	if o == nil || IsNil(o.Sortby) {
		return nil, false
	}
	return o.Sortby, true
}

// HasSortby returns a boolean if a field has been set.
func (o *ModForumGetDiscussionPostsByUseridRequest) HasSortby() bool {
	if o != nil && !IsNil(o.Sortby) {
		return true
	}

	return false
}

// SetSortby gets a reference to the given string and assigns it to the Sortby field.
func (o *ModForumGetDiscussionPostsByUseridRequest) SetSortby(v string) {
	o.Sortby = &v
}

// GetSortdirection returns the Sortdirection field value if set, zero value otherwise.
func (o *ModForumGetDiscussionPostsByUseridRequest) GetSortdirection() string {
	if o == nil || IsNil(o.Sortdirection) {
		var ret string
		return ret
	}
	return *o.Sortdirection
}

// GetSortdirectionOk returns a tuple with the Sortdirection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModForumGetDiscussionPostsByUseridRequest) GetSortdirectionOk() (*string, bool) {
	if o == nil || IsNil(o.Sortdirection) {
		return nil, false
	}
	return o.Sortdirection, true
}

// HasSortdirection returns a boolean if a field has been set.
func (o *ModForumGetDiscussionPostsByUseridRequest) HasSortdirection() bool {
	if o != nil && !IsNil(o.Sortdirection) {
		return true
	}

	return false
}

// SetSortdirection gets a reference to the given string and assigns it to the Sortdirection field.
func (o *ModForumGetDiscussionPostsByUseridRequest) SetSortdirection(v string) {
	o.Sortdirection = &v
}

// GetUserid returns the Userid field value
func (o *ModForumGetDiscussionPostsByUseridRequest) GetUserid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Userid
}

// GetUseridOk returns a tuple with the Userid field value
// and a boolean to check if the value has been set.
func (o *ModForumGetDiscussionPostsByUseridRequest) GetUseridOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Userid, true
}

// SetUserid sets field value
func (o *ModForumGetDiscussionPostsByUseridRequest) SetUserid(v int32) {
	o.Userid = v
}

func (o ModForumGetDiscussionPostsByUseridRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModForumGetDiscussionPostsByUseridRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cmid"] = o.Cmid
	if !IsNil(o.Sortby) {
		toSerialize["sortby"] = o.Sortby
	}
	if !IsNil(o.Sortdirection) {
		toSerialize["sortdirection"] = o.Sortdirection
	}
	toSerialize["userid"] = o.Userid
	return toSerialize, nil
}

func (o *ModForumGetDiscussionPostsByUseridRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"cmid",
		"userid",
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

	varModForumGetDiscussionPostsByUseridRequest := _ModForumGetDiscussionPostsByUseridRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModForumGetDiscussionPostsByUseridRequest)

	if err != nil {
		return err
	}

	*o = ModForumGetDiscussionPostsByUseridRequest(varModForumGetDiscussionPostsByUseridRequest)

	return err
}

type NullableModForumGetDiscussionPostsByUseridRequest struct {
	value *ModForumGetDiscussionPostsByUseridRequest
	isSet bool
}

func (v NullableModForumGetDiscussionPostsByUseridRequest) Get() *ModForumGetDiscussionPostsByUseridRequest {
	return v.value
}

func (v *NullableModForumGetDiscussionPostsByUseridRequest) Set(val *ModForumGetDiscussionPostsByUseridRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModForumGetDiscussionPostsByUseridRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModForumGetDiscussionPostsByUseridRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModForumGetDiscussionPostsByUseridRequest(val *ModForumGetDiscussionPostsByUseridRequest) *NullableModForumGetDiscussionPostsByUseridRequest {
	return &NullableModForumGetDiscussionPostsByUseridRequest{value: val, isSet: true}
}

func (v NullableModForumGetDiscussionPostsByUseridRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModForumGetDiscussionPostsByUseridRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

