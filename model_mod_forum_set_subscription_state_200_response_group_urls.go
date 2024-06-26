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

// checks if the ModForumSetSubscriptionState200ResponseGroupUrls type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModForumSetSubscriptionState200ResponseGroupUrls{}

// ModForumSetSubscriptionState200ResponseGroupUrls struct for ModForumSetSubscriptionState200ResponseGroupUrls
type ModForumSetSubscriptionState200ResponseGroupUrls struct {
	// picture
	Picture *string `json:"picture,omitempty"`
	// userlist
	Userlist *string `json:"userlist,omitempty"`
}

// NewModForumSetSubscriptionState200ResponseGroupUrls instantiates a new ModForumSetSubscriptionState200ResponseGroupUrls object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModForumSetSubscriptionState200ResponseGroupUrls() *ModForumSetSubscriptionState200ResponseGroupUrls {
	this := ModForumSetSubscriptionState200ResponseGroupUrls{}
	return &this
}

// NewModForumSetSubscriptionState200ResponseGroupUrlsWithDefaults instantiates a new ModForumSetSubscriptionState200ResponseGroupUrls object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModForumSetSubscriptionState200ResponseGroupUrlsWithDefaults() *ModForumSetSubscriptionState200ResponseGroupUrls {
	this := ModForumSetSubscriptionState200ResponseGroupUrls{}
	return &this
}

// GetPicture returns the Picture field value if set, zero value otherwise.
func (o *ModForumSetSubscriptionState200ResponseGroupUrls) GetPicture() string {
	if o == nil || IsNil(o.Picture) {
		var ret string
		return ret
	}
	return *o.Picture
}

// GetPictureOk returns a tuple with the Picture field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModForumSetSubscriptionState200ResponseGroupUrls) GetPictureOk() (*string, bool) {
	if o == nil || IsNil(o.Picture) {
		return nil, false
	}
	return o.Picture, true
}

// HasPicture returns a boolean if a field has been set.
func (o *ModForumSetSubscriptionState200ResponseGroupUrls) HasPicture() bool {
	if o != nil && !IsNil(o.Picture) {
		return true
	}

	return false
}

// SetPicture gets a reference to the given string and assigns it to the Picture field.
func (o *ModForumSetSubscriptionState200ResponseGroupUrls) SetPicture(v string) {
	o.Picture = &v
}

// GetUserlist returns the Userlist field value if set, zero value otherwise.
func (o *ModForumSetSubscriptionState200ResponseGroupUrls) GetUserlist() string {
	if o == nil || IsNil(o.Userlist) {
		var ret string
		return ret
	}
	return *o.Userlist
}

// GetUserlistOk returns a tuple with the Userlist field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModForumSetSubscriptionState200ResponseGroupUrls) GetUserlistOk() (*string, bool) {
	if o == nil || IsNil(o.Userlist) {
		return nil, false
	}
	return o.Userlist, true
}

// HasUserlist returns a boolean if a field has been set.
func (o *ModForumSetSubscriptionState200ResponseGroupUrls) HasUserlist() bool {
	if o != nil && !IsNil(o.Userlist) {
		return true
	}

	return false
}

// SetUserlist gets a reference to the given string and assigns it to the Userlist field.
func (o *ModForumSetSubscriptionState200ResponseGroupUrls) SetUserlist(v string) {
	o.Userlist = &v
}

func (o ModForumSetSubscriptionState200ResponseGroupUrls) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModForumSetSubscriptionState200ResponseGroupUrls) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Picture) {
		toSerialize["picture"] = o.Picture
	}
	if !IsNil(o.Userlist) {
		toSerialize["userlist"] = o.Userlist
	}
	return toSerialize, nil
}

type NullableModForumSetSubscriptionState200ResponseGroupUrls struct {
	value *ModForumSetSubscriptionState200ResponseGroupUrls
	isSet bool
}

func (v NullableModForumSetSubscriptionState200ResponseGroupUrls) Get() *ModForumSetSubscriptionState200ResponseGroupUrls {
	return v.value
}

func (v *NullableModForumSetSubscriptionState200ResponseGroupUrls) Set(val *ModForumSetSubscriptionState200ResponseGroupUrls) {
	v.value = val
	v.isSet = true
}

func (v NullableModForumSetSubscriptionState200ResponseGroupUrls) IsSet() bool {
	return v.isSet
}

func (v *NullableModForumSetSubscriptionState200ResponseGroupUrls) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModForumSetSubscriptionState200ResponseGroupUrls(val *ModForumSetSubscriptionState200ResponseGroupUrls) *NullableModForumSetSubscriptionState200ResponseGroupUrls {
	return &NullableModForumSetSubscriptionState200ResponseGroupUrls{value: val, isSet: true}
}

func (v NullableModForumSetSubscriptionState200ResponseGroupUrls) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModForumSetSubscriptionState200ResponseGroupUrls) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


