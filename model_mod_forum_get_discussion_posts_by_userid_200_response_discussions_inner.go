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

// checks if the ModForumGetDiscussionPostsByUserid200ResponseDiscussionsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModForumGetDiscussionPostsByUserid200ResponseDiscussionsInner{}

// ModForumGetDiscussionPostsByUserid200ResponseDiscussionsInner struct for ModForumGetDiscussionPostsByUserid200ResponseDiscussionsInner
type ModForumGetDiscussionPostsByUserid200ResponseDiscussionsInner struct {
	// Full name of the user that started the discussion
	Authorfullname *string `json:"authorfullname,omitempty"`
	// ID of the discussion
	Id *int32 `json:"id,omitempty"`
	// Name of the discussion
	Name *string `json:"name,omitempty"`
	Posts *ModForumGetDiscussionPostsByUserid200ResponseDiscussionsInnerPosts `json:"posts,omitempty"`
	// Timestamp of the discussion start
	Timecreated *int32 `json:"timecreated,omitempty"`
}

// NewModForumGetDiscussionPostsByUserid200ResponseDiscussionsInner instantiates a new ModForumGetDiscussionPostsByUserid200ResponseDiscussionsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModForumGetDiscussionPostsByUserid200ResponseDiscussionsInner() *ModForumGetDiscussionPostsByUserid200ResponseDiscussionsInner {
	this := ModForumGetDiscussionPostsByUserid200ResponseDiscussionsInner{}
	var authorfullname string = "null"
	this.Authorfullname = &authorfullname
	var id int32 = null
	this.Id = &id
	var name string = "null"
	this.Name = &name
	var timecreated int32 = null
	this.Timecreated = &timecreated
	return &this
}

// NewModForumGetDiscussionPostsByUserid200ResponseDiscussionsInnerWithDefaults instantiates a new ModForumGetDiscussionPostsByUserid200ResponseDiscussionsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModForumGetDiscussionPostsByUserid200ResponseDiscussionsInnerWithDefaults() *ModForumGetDiscussionPostsByUserid200ResponseDiscussionsInner {
	this := ModForumGetDiscussionPostsByUserid200ResponseDiscussionsInner{}
	var authorfullname string = "null"
	this.Authorfullname = &authorfullname
	var id int32 = null
	this.Id = &id
	var name string = "null"
	this.Name = &name
	var timecreated int32 = null
	this.Timecreated = &timecreated
	return &this
}

// GetAuthorfullname returns the Authorfullname field value if set, zero value otherwise.
func (o *ModForumGetDiscussionPostsByUserid200ResponseDiscussionsInner) GetAuthorfullname() string {
	if o == nil || IsNil(o.Authorfullname) {
		var ret string
		return ret
	}
	return *o.Authorfullname
}

// GetAuthorfullnameOk returns a tuple with the Authorfullname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModForumGetDiscussionPostsByUserid200ResponseDiscussionsInner) GetAuthorfullnameOk() (*string, bool) {
	if o == nil || IsNil(o.Authorfullname) {
		return nil, false
	}
	return o.Authorfullname, true
}

// HasAuthorfullname returns a boolean if a field has been set.
func (o *ModForumGetDiscussionPostsByUserid200ResponseDiscussionsInner) HasAuthorfullname() bool {
	if o != nil && !IsNil(o.Authorfullname) {
		return true
	}

	return false
}

// SetAuthorfullname gets a reference to the given string and assigns it to the Authorfullname field.
func (o *ModForumGetDiscussionPostsByUserid200ResponseDiscussionsInner) SetAuthorfullname(v string) {
	o.Authorfullname = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ModForumGetDiscussionPostsByUserid200ResponseDiscussionsInner) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModForumGetDiscussionPostsByUserid200ResponseDiscussionsInner) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ModForumGetDiscussionPostsByUserid200ResponseDiscussionsInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ModForumGetDiscussionPostsByUserid200ResponseDiscussionsInner) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ModForumGetDiscussionPostsByUserid200ResponseDiscussionsInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModForumGetDiscussionPostsByUserid200ResponseDiscussionsInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ModForumGetDiscussionPostsByUserid200ResponseDiscussionsInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ModForumGetDiscussionPostsByUserid200ResponseDiscussionsInner) SetName(v string) {
	o.Name = &v
}

// GetPosts returns the Posts field value if set, zero value otherwise.
func (o *ModForumGetDiscussionPostsByUserid200ResponseDiscussionsInner) GetPosts() ModForumGetDiscussionPostsByUserid200ResponseDiscussionsInnerPosts {
	if o == nil || IsNil(o.Posts) {
		var ret ModForumGetDiscussionPostsByUserid200ResponseDiscussionsInnerPosts
		return ret
	}
	return *o.Posts
}

// GetPostsOk returns a tuple with the Posts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModForumGetDiscussionPostsByUserid200ResponseDiscussionsInner) GetPostsOk() (*ModForumGetDiscussionPostsByUserid200ResponseDiscussionsInnerPosts, bool) {
	if o == nil || IsNil(o.Posts) {
		return nil, false
	}
	return o.Posts, true
}

// HasPosts returns a boolean if a field has been set.
func (o *ModForumGetDiscussionPostsByUserid200ResponseDiscussionsInner) HasPosts() bool {
	if o != nil && !IsNil(o.Posts) {
		return true
	}

	return false
}

// SetPosts gets a reference to the given ModForumGetDiscussionPostsByUserid200ResponseDiscussionsInnerPosts and assigns it to the Posts field.
func (o *ModForumGetDiscussionPostsByUserid200ResponseDiscussionsInner) SetPosts(v ModForumGetDiscussionPostsByUserid200ResponseDiscussionsInnerPosts) {
	o.Posts = &v
}

// GetTimecreated returns the Timecreated field value if set, zero value otherwise.
func (o *ModForumGetDiscussionPostsByUserid200ResponseDiscussionsInner) GetTimecreated() int32 {
	if o == nil || IsNil(o.Timecreated) {
		var ret int32
		return ret
	}
	return *o.Timecreated
}

// GetTimecreatedOk returns a tuple with the Timecreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModForumGetDiscussionPostsByUserid200ResponseDiscussionsInner) GetTimecreatedOk() (*int32, bool) {
	if o == nil || IsNil(o.Timecreated) {
		return nil, false
	}
	return o.Timecreated, true
}

// HasTimecreated returns a boolean if a field has been set.
func (o *ModForumGetDiscussionPostsByUserid200ResponseDiscussionsInner) HasTimecreated() bool {
	if o != nil && !IsNil(o.Timecreated) {
		return true
	}

	return false
}

// SetTimecreated gets a reference to the given int32 and assigns it to the Timecreated field.
func (o *ModForumGetDiscussionPostsByUserid200ResponseDiscussionsInner) SetTimecreated(v int32) {
	o.Timecreated = &v
}

func (o ModForumGetDiscussionPostsByUserid200ResponseDiscussionsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModForumGetDiscussionPostsByUserid200ResponseDiscussionsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Authorfullname) {
		toSerialize["authorfullname"] = o.Authorfullname
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Posts) {
		toSerialize["posts"] = o.Posts
	}
	if !IsNil(o.Timecreated) {
		toSerialize["timecreated"] = o.Timecreated
	}
	return toSerialize, nil
}

type NullableModForumGetDiscussionPostsByUserid200ResponseDiscussionsInner struct {
	value *ModForumGetDiscussionPostsByUserid200ResponseDiscussionsInner
	isSet bool
}

func (v NullableModForumGetDiscussionPostsByUserid200ResponseDiscussionsInner) Get() *ModForumGetDiscussionPostsByUserid200ResponseDiscussionsInner {
	return v.value
}

func (v *NullableModForumGetDiscussionPostsByUserid200ResponseDiscussionsInner) Set(val *ModForumGetDiscussionPostsByUserid200ResponseDiscussionsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableModForumGetDiscussionPostsByUserid200ResponseDiscussionsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableModForumGetDiscussionPostsByUserid200ResponseDiscussionsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModForumGetDiscussionPostsByUserid200ResponseDiscussionsInner(val *ModForumGetDiscussionPostsByUserid200ResponseDiscussionsInner) *NullableModForumGetDiscussionPostsByUserid200ResponseDiscussionsInner {
	return &NullableModForumGetDiscussionPostsByUserid200ResponseDiscussionsInner{value: val, isSet: true}
}

func (v NullableModForumGetDiscussionPostsByUserid200ResponseDiscussionsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModForumGetDiscussionPostsByUserid200ResponseDiscussionsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


