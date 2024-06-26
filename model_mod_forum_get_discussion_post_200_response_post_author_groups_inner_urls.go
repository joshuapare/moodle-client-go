/*
Moodle Webservice API

Auto-generated OpenAPI spec for Moodle's Webservice API. Requires installation of the Catalyst RESTful API plugin: https://github.com/catalyst/moodle-webservice_restful.

API version: 4.3.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package moodleclient

import (
	"encoding/json"
)

// checks if the ModForumGetDiscussionPost200ResponsePostAuthorGroupsInnerUrls type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModForumGetDiscussionPost200ResponsePostAuthorGroupsInnerUrls{}

// ModForumGetDiscussionPost200ResponsePostAuthorGroupsInnerUrls struct for ModForumGetDiscussionPost200ResponsePostAuthorGroupsInnerUrls
type ModForumGetDiscussionPost200ResponsePostAuthorGroupsInnerUrls struct {
	// image
	Image *string `json:"image,omitempty"`
}

// NewModForumGetDiscussionPost200ResponsePostAuthorGroupsInnerUrls instantiates a new ModForumGetDiscussionPost200ResponsePostAuthorGroupsInnerUrls object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModForumGetDiscussionPost200ResponsePostAuthorGroupsInnerUrls() *ModForumGetDiscussionPost200ResponsePostAuthorGroupsInnerUrls {
	this := ModForumGetDiscussionPost200ResponsePostAuthorGroupsInnerUrls{}
	return &this
}

// NewModForumGetDiscussionPost200ResponsePostAuthorGroupsInnerUrlsWithDefaults instantiates a new ModForumGetDiscussionPost200ResponsePostAuthorGroupsInnerUrls object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModForumGetDiscussionPost200ResponsePostAuthorGroupsInnerUrlsWithDefaults() *ModForumGetDiscussionPost200ResponsePostAuthorGroupsInnerUrls {
	this := ModForumGetDiscussionPost200ResponsePostAuthorGroupsInnerUrls{}
	return &this
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *ModForumGetDiscussionPost200ResponsePostAuthorGroupsInnerUrls) GetImage() string {
	if o == nil || IsNil(o.Image) {
		var ret string
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModForumGetDiscussionPost200ResponsePostAuthorGroupsInnerUrls) GetImageOk() (*string, bool) {
	if o == nil || IsNil(o.Image) {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *ModForumGetDiscussionPost200ResponsePostAuthorGroupsInnerUrls) HasImage() bool {
	if o != nil && !IsNil(o.Image) {
		return true
	}

	return false
}

// SetImage gets a reference to the given string and assigns it to the Image field.
func (o *ModForumGetDiscussionPost200ResponsePostAuthorGroupsInnerUrls) SetImage(v string) {
	o.Image = &v
}

func (o ModForumGetDiscussionPost200ResponsePostAuthorGroupsInnerUrls) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModForumGetDiscussionPost200ResponsePostAuthorGroupsInnerUrls) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Image) {
		toSerialize["image"] = o.Image
	}
	return toSerialize, nil
}

type NullableModForumGetDiscussionPost200ResponsePostAuthorGroupsInnerUrls struct {
	value *ModForumGetDiscussionPost200ResponsePostAuthorGroupsInnerUrls
	isSet bool
}

func (v NullableModForumGetDiscussionPost200ResponsePostAuthorGroupsInnerUrls) Get() *ModForumGetDiscussionPost200ResponsePostAuthorGroupsInnerUrls {
	return v.value
}

func (v *NullableModForumGetDiscussionPost200ResponsePostAuthorGroupsInnerUrls) Set(val *ModForumGetDiscussionPost200ResponsePostAuthorGroupsInnerUrls) {
	v.value = val
	v.isSet = true
}

func (v NullableModForumGetDiscussionPost200ResponsePostAuthorGroupsInnerUrls) IsSet() bool {
	return v.isSet
}

func (v *NullableModForumGetDiscussionPost200ResponsePostAuthorGroupsInnerUrls) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModForumGetDiscussionPost200ResponsePostAuthorGroupsInnerUrls(val *ModForumGetDiscussionPost200ResponsePostAuthorGroupsInnerUrls) *NullableModForumGetDiscussionPost200ResponsePostAuthorGroupsInnerUrls {
	return &NullableModForumGetDiscussionPost200ResponsePostAuthorGroupsInnerUrls{value: val, isSet: true}
}

func (v NullableModForumGetDiscussionPost200ResponsePostAuthorGroupsInnerUrls) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModForumGetDiscussionPost200ResponsePostAuthorGroupsInnerUrls) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


