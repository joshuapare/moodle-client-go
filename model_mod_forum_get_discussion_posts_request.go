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

// checks if the ModForumGetDiscussionPostsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModForumGetDiscussionPostsRequest{}

// ModForumGetDiscussionPostsRequest struct for ModForumGetDiscussionPostsRequest
type ModForumGetDiscussionPostsRequest struct {
	// The ID of the discussion from which to fetch posts.
	Discussionid int32 `json:"discussionid"`
	// Whether inline attachments should be included or not
	Includeinlineattachments *bool `json:"includeinlineattachments,omitempty"`
	// Sort by this element: id, created or modified
	Sortby *string `json:"sortby,omitempty"`
	// Sort direction: ASC or DESC
	Sortdirection *string `json:"sortdirection,omitempty"`
}

type _ModForumGetDiscussionPostsRequest ModForumGetDiscussionPostsRequest

// NewModForumGetDiscussionPostsRequest instantiates a new ModForumGetDiscussionPostsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModForumGetDiscussionPostsRequest(discussionid int32) *ModForumGetDiscussionPostsRequest {
	this := ModForumGetDiscussionPostsRequest{}
	this.Discussionid = discussionid
	var includeinlineattachments bool = false
	this.Includeinlineattachments = &includeinlineattachments
	var sortby string = "created"
	this.Sortby = &sortby
	var sortdirection string = "DESC"
	this.Sortdirection = &sortdirection
	return &this
}

// NewModForumGetDiscussionPostsRequestWithDefaults instantiates a new ModForumGetDiscussionPostsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModForumGetDiscussionPostsRequestWithDefaults() *ModForumGetDiscussionPostsRequest {
	this := ModForumGetDiscussionPostsRequest{}
	var discussionid int32 = null
	this.Discussionid = discussionid
	var includeinlineattachments bool = false
	this.Includeinlineattachments = &includeinlineattachments
	var sortby string = "created"
	this.Sortby = &sortby
	var sortdirection string = "DESC"
	this.Sortdirection = &sortdirection
	return &this
}

// GetDiscussionid returns the Discussionid field value
func (o *ModForumGetDiscussionPostsRequest) GetDiscussionid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Discussionid
}

// GetDiscussionidOk returns a tuple with the Discussionid field value
// and a boolean to check if the value has been set.
func (o *ModForumGetDiscussionPostsRequest) GetDiscussionidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Discussionid, true
}

// SetDiscussionid sets field value
func (o *ModForumGetDiscussionPostsRequest) SetDiscussionid(v int32) {
	o.Discussionid = v
}

// GetIncludeinlineattachments returns the Includeinlineattachments field value if set, zero value otherwise.
func (o *ModForumGetDiscussionPostsRequest) GetIncludeinlineattachments() bool {
	if o == nil || IsNil(o.Includeinlineattachments) {
		var ret bool
		return ret
	}
	return *o.Includeinlineattachments
}

// GetIncludeinlineattachmentsOk returns a tuple with the Includeinlineattachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModForumGetDiscussionPostsRequest) GetIncludeinlineattachmentsOk() (*bool, bool) {
	if o == nil || IsNil(o.Includeinlineattachments) {
		return nil, false
	}
	return o.Includeinlineattachments, true
}

// HasIncludeinlineattachments returns a boolean if a field has been set.
func (o *ModForumGetDiscussionPostsRequest) HasIncludeinlineattachments() bool {
	if o != nil && !IsNil(o.Includeinlineattachments) {
		return true
	}

	return false
}

// SetIncludeinlineattachments gets a reference to the given bool and assigns it to the Includeinlineattachments field.
func (o *ModForumGetDiscussionPostsRequest) SetIncludeinlineattachments(v bool) {
	o.Includeinlineattachments = &v
}

// GetSortby returns the Sortby field value if set, zero value otherwise.
func (o *ModForumGetDiscussionPostsRequest) GetSortby() string {
	if o == nil || IsNil(o.Sortby) {
		var ret string
		return ret
	}
	return *o.Sortby
}

// GetSortbyOk returns a tuple with the Sortby field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModForumGetDiscussionPostsRequest) GetSortbyOk() (*string, bool) {
	if o == nil || IsNil(o.Sortby) {
		return nil, false
	}
	return o.Sortby, true
}

// HasSortby returns a boolean if a field has been set.
func (o *ModForumGetDiscussionPostsRequest) HasSortby() bool {
	if o != nil && !IsNil(o.Sortby) {
		return true
	}

	return false
}

// SetSortby gets a reference to the given string and assigns it to the Sortby field.
func (o *ModForumGetDiscussionPostsRequest) SetSortby(v string) {
	o.Sortby = &v
}

// GetSortdirection returns the Sortdirection field value if set, zero value otherwise.
func (o *ModForumGetDiscussionPostsRequest) GetSortdirection() string {
	if o == nil || IsNil(o.Sortdirection) {
		var ret string
		return ret
	}
	return *o.Sortdirection
}

// GetSortdirectionOk returns a tuple with the Sortdirection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModForumGetDiscussionPostsRequest) GetSortdirectionOk() (*string, bool) {
	if o == nil || IsNil(o.Sortdirection) {
		return nil, false
	}
	return o.Sortdirection, true
}

// HasSortdirection returns a boolean if a field has been set.
func (o *ModForumGetDiscussionPostsRequest) HasSortdirection() bool {
	if o != nil && !IsNil(o.Sortdirection) {
		return true
	}

	return false
}

// SetSortdirection gets a reference to the given string and assigns it to the Sortdirection field.
func (o *ModForumGetDiscussionPostsRequest) SetSortdirection(v string) {
	o.Sortdirection = &v
}

func (o ModForumGetDiscussionPostsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModForumGetDiscussionPostsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["discussionid"] = o.Discussionid
	if !IsNil(o.Includeinlineattachments) {
		toSerialize["includeinlineattachments"] = o.Includeinlineattachments
	}
	if !IsNil(o.Sortby) {
		toSerialize["sortby"] = o.Sortby
	}
	if !IsNil(o.Sortdirection) {
		toSerialize["sortdirection"] = o.Sortdirection
	}
	return toSerialize, nil
}

func (o *ModForumGetDiscussionPostsRequest) UnmarshalJSON(data []byte) (err error) {
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

	varModForumGetDiscussionPostsRequest := _ModForumGetDiscussionPostsRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModForumGetDiscussionPostsRequest)

	if err != nil {
		return err
	}

	*o = ModForumGetDiscussionPostsRequest(varModForumGetDiscussionPostsRequest)

	return err
}

type NullableModForumGetDiscussionPostsRequest struct {
	value *ModForumGetDiscussionPostsRequest
	isSet bool
}

func (v NullableModForumGetDiscussionPostsRequest) Get() *ModForumGetDiscussionPostsRequest {
	return v.value
}

func (v *NullableModForumGetDiscussionPostsRequest) Set(val *ModForumGetDiscussionPostsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModForumGetDiscussionPostsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModForumGetDiscussionPostsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModForumGetDiscussionPostsRequest(val *ModForumGetDiscussionPostsRequest) *NullableModForumGetDiscussionPostsRequest {
	return &NullableModForumGetDiscussionPostsRequest{value: val, isSet: true}
}

func (v NullableModForumGetDiscussionPostsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModForumGetDiscussionPostsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


