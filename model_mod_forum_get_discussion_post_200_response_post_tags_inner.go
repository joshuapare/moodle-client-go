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

// checks if the ModForumGetDiscussionPost200ResponsePostTagsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModForumGetDiscussionPost200ResponsePostTagsInner{}

// ModForumGetDiscussionPost200ResponsePostTagsInner struct for ModForumGetDiscussionPost200ResponsePostTagsInner
type ModForumGetDiscussionPost200ResponsePostTagsInner struct {
	// The display name of the tag
	Displayname *string `json:"displayname,omitempty"`
	// Wehther this tag is flagged
	Flag *bool `json:"flag,omitempty"`
	// The ID of the Tag
	Id *int32 `json:"id,omitempty"`
	// Whether this is a standard tag
	Isstandard *bool `json:"isstandard,omitempty"`
	// The tagid
	Tagid *int32 `json:"tagid,omitempty"`
	Urls *ModForumGetDiscussionPost200ResponsePostTagsInnerUrls `json:"urls,omitempty"`
}

// NewModForumGetDiscussionPost200ResponsePostTagsInner instantiates a new ModForumGetDiscussionPost200ResponsePostTagsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModForumGetDiscussionPost200ResponsePostTagsInner() *ModForumGetDiscussionPost200ResponsePostTagsInner {
	this := ModForumGetDiscussionPost200ResponsePostTagsInner{}
	return &this
}

// NewModForumGetDiscussionPost200ResponsePostTagsInnerWithDefaults instantiates a new ModForumGetDiscussionPost200ResponsePostTagsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModForumGetDiscussionPost200ResponsePostTagsInnerWithDefaults() *ModForumGetDiscussionPost200ResponsePostTagsInner {
	this := ModForumGetDiscussionPost200ResponsePostTagsInner{}
	return &this
}

// GetDisplayname returns the Displayname field value if set, zero value otherwise.
func (o *ModForumGetDiscussionPost200ResponsePostTagsInner) GetDisplayname() string {
	if o == nil || IsNil(o.Displayname) {
		var ret string
		return ret
	}
	return *o.Displayname
}

// GetDisplaynameOk returns a tuple with the Displayname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModForumGetDiscussionPost200ResponsePostTagsInner) GetDisplaynameOk() (*string, bool) {
	if o == nil || IsNil(o.Displayname) {
		return nil, false
	}
	return o.Displayname, true
}

// HasDisplayname returns a boolean if a field has been set.
func (o *ModForumGetDiscussionPost200ResponsePostTagsInner) HasDisplayname() bool {
	if o != nil && !IsNil(o.Displayname) {
		return true
	}

	return false
}

// SetDisplayname gets a reference to the given string and assigns it to the Displayname field.
func (o *ModForumGetDiscussionPost200ResponsePostTagsInner) SetDisplayname(v string) {
	o.Displayname = &v
}

// GetFlag returns the Flag field value if set, zero value otherwise.
func (o *ModForumGetDiscussionPost200ResponsePostTagsInner) GetFlag() bool {
	if o == nil || IsNil(o.Flag) {
		var ret bool
		return ret
	}
	return *o.Flag
}

// GetFlagOk returns a tuple with the Flag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModForumGetDiscussionPost200ResponsePostTagsInner) GetFlagOk() (*bool, bool) {
	if o == nil || IsNil(o.Flag) {
		return nil, false
	}
	return o.Flag, true
}

// HasFlag returns a boolean if a field has been set.
func (o *ModForumGetDiscussionPost200ResponsePostTagsInner) HasFlag() bool {
	if o != nil && !IsNil(o.Flag) {
		return true
	}

	return false
}

// SetFlag gets a reference to the given bool and assigns it to the Flag field.
func (o *ModForumGetDiscussionPost200ResponsePostTagsInner) SetFlag(v bool) {
	o.Flag = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ModForumGetDiscussionPost200ResponsePostTagsInner) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModForumGetDiscussionPost200ResponsePostTagsInner) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ModForumGetDiscussionPost200ResponsePostTagsInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ModForumGetDiscussionPost200ResponsePostTagsInner) SetId(v int32) {
	o.Id = &v
}

// GetIsstandard returns the Isstandard field value if set, zero value otherwise.
func (o *ModForumGetDiscussionPost200ResponsePostTagsInner) GetIsstandard() bool {
	if o == nil || IsNil(o.Isstandard) {
		var ret bool
		return ret
	}
	return *o.Isstandard
}

// GetIsstandardOk returns a tuple with the Isstandard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModForumGetDiscussionPost200ResponsePostTagsInner) GetIsstandardOk() (*bool, bool) {
	if o == nil || IsNil(o.Isstandard) {
		return nil, false
	}
	return o.Isstandard, true
}

// HasIsstandard returns a boolean if a field has been set.
func (o *ModForumGetDiscussionPost200ResponsePostTagsInner) HasIsstandard() bool {
	if o != nil && !IsNil(o.Isstandard) {
		return true
	}

	return false
}

// SetIsstandard gets a reference to the given bool and assigns it to the Isstandard field.
func (o *ModForumGetDiscussionPost200ResponsePostTagsInner) SetIsstandard(v bool) {
	o.Isstandard = &v
}

// GetTagid returns the Tagid field value if set, zero value otherwise.
func (o *ModForumGetDiscussionPost200ResponsePostTagsInner) GetTagid() int32 {
	if o == nil || IsNil(o.Tagid) {
		var ret int32
		return ret
	}
	return *o.Tagid
}

// GetTagidOk returns a tuple with the Tagid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModForumGetDiscussionPost200ResponsePostTagsInner) GetTagidOk() (*int32, bool) {
	if o == nil || IsNil(o.Tagid) {
		return nil, false
	}
	return o.Tagid, true
}

// HasTagid returns a boolean if a field has been set.
func (o *ModForumGetDiscussionPost200ResponsePostTagsInner) HasTagid() bool {
	if o != nil && !IsNil(o.Tagid) {
		return true
	}

	return false
}

// SetTagid gets a reference to the given int32 and assigns it to the Tagid field.
func (o *ModForumGetDiscussionPost200ResponsePostTagsInner) SetTagid(v int32) {
	o.Tagid = &v
}

// GetUrls returns the Urls field value if set, zero value otherwise.
func (o *ModForumGetDiscussionPost200ResponsePostTagsInner) GetUrls() ModForumGetDiscussionPost200ResponsePostTagsInnerUrls {
	if o == nil || IsNil(o.Urls) {
		var ret ModForumGetDiscussionPost200ResponsePostTagsInnerUrls
		return ret
	}
	return *o.Urls
}

// GetUrlsOk returns a tuple with the Urls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModForumGetDiscussionPost200ResponsePostTagsInner) GetUrlsOk() (*ModForumGetDiscussionPost200ResponsePostTagsInnerUrls, bool) {
	if o == nil || IsNil(o.Urls) {
		return nil, false
	}
	return o.Urls, true
}

// HasUrls returns a boolean if a field has been set.
func (o *ModForumGetDiscussionPost200ResponsePostTagsInner) HasUrls() bool {
	if o != nil && !IsNil(o.Urls) {
		return true
	}

	return false
}

// SetUrls gets a reference to the given ModForumGetDiscussionPost200ResponsePostTagsInnerUrls and assigns it to the Urls field.
func (o *ModForumGetDiscussionPost200ResponsePostTagsInner) SetUrls(v ModForumGetDiscussionPost200ResponsePostTagsInnerUrls) {
	o.Urls = &v
}

func (o ModForumGetDiscussionPost200ResponsePostTagsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModForumGetDiscussionPost200ResponsePostTagsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Displayname) {
		toSerialize["displayname"] = o.Displayname
	}
	if !IsNil(o.Flag) {
		toSerialize["flag"] = o.Flag
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Isstandard) {
		toSerialize["isstandard"] = o.Isstandard
	}
	if !IsNil(o.Tagid) {
		toSerialize["tagid"] = o.Tagid
	}
	if !IsNil(o.Urls) {
		toSerialize["urls"] = o.Urls
	}
	return toSerialize, nil
}

type NullableModForumGetDiscussionPost200ResponsePostTagsInner struct {
	value *ModForumGetDiscussionPost200ResponsePostTagsInner
	isSet bool
}

func (v NullableModForumGetDiscussionPost200ResponsePostTagsInner) Get() *ModForumGetDiscussionPost200ResponsePostTagsInner {
	return v.value
}

func (v *NullableModForumGetDiscussionPost200ResponsePostTagsInner) Set(val *ModForumGetDiscussionPost200ResponsePostTagsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableModForumGetDiscussionPost200ResponsePostTagsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableModForumGetDiscussionPost200ResponsePostTagsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModForumGetDiscussionPost200ResponsePostTagsInner(val *ModForumGetDiscussionPost200ResponsePostTagsInner) *NullableModForumGetDiscussionPost200ResponsePostTagsInner {
	return &NullableModForumGetDiscussionPost200ResponsePostTagsInner{value: val, isSet: true}
}

func (v NullableModForumGetDiscussionPost200ResponsePostTagsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModForumGetDiscussionPost200ResponsePostTagsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


