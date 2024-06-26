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

// checks if the ModForumGetDiscussionPosts200ResponseRatinginfoScalesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModForumGetDiscussionPosts200ResponseRatinginfoScalesInner{}

// ModForumGetDiscussionPosts200ResponseRatinginfoScalesInner Scale information
type ModForumGetDiscussionPosts200ResponseRatinginfoScalesInner struct {
	// Course id.
	Courseid *int32 `json:"courseid,omitempty"`
	// Scale id.
	Id *int32 `json:"id,omitempty"`
	// Whether is a numeric scale.
	Isnumeric *bool `json:"isnumeric,omitempty"`
	Items []ModForumGetDiscussionPosts200ResponseRatinginfoScalesInnerItemsInner `json:"items,omitempty"`
	// Max value for the scale.
	Max *int32 `json:"max,omitempty"`
	// Scale name (when a real scale is used).
	Name *string `json:"name,omitempty"`
}

// NewModForumGetDiscussionPosts200ResponseRatinginfoScalesInner instantiates a new ModForumGetDiscussionPosts200ResponseRatinginfoScalesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModForumGetDiscussionPosts200ResponseRatinginfoScalesInner() *ModForumGetDiscussionPosts200ResponseRatinginfoScalesInner {
	this := ModForumGetDiscussionPosts200ResponseRatinginfoScalesInner{}
	return &this
}

// NewModForumGetDiscussionPosts200ResponseRatinginfoScalesInnerWithDefaults instantiates a new ModForumGetDiscussionPosts200ResponseRatinginfoScalesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModForumGetDiscussionPosts200ResponseRatinginfoScalesInnerWithDefaults() *ModForumGetDiscussionPosts200ResponseRatinginfoScalesInner {
	this := ModForumGetDiscussionPosts200ResponseRatinginfoScalesInner{}
	return &this
}

// GetCourseid returns the Courseid field value if set, zero value otherwise.
func (o *ModForumGetDiscussionPosts200ResponseRatinginfoScalesInner) GetCourseid() int32 {
	if o == nil || IsNil(o.Courseid) {
		var ret int32
		return ret
	}
	return *o.Courseid
}

// GetCourseidOk returns a tuple with the Courseid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModForumGetDiscussionPosts200ResponseRatinginfoScalesInner) GetCourseidOk() (*int32, bool) {
	if o == nil || IsNil(o.Courseid) {
		return nil, false
	}
	return o.Courseid, true
}

// HasCourseid returns a boolean if a field has been set.
func (o *ModForumGetDiscussionPosts200ResponseRatinginfoScalesInner) HasCourseid() bool {
	if o != nil && !IsNil(o.Courseid) {
		return true
	}

	return false
}

// SetCourseid gets a reference to the given int32 and assigns it to the Courseid field.
func (o *ModForumGetDiscussionPosts200ResponseRatinginfoScalesInner) SetCourseid(v int32) {
	o.Courseid = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ModForumGetDiscussionPosts200ResponseRatinginfoScalesInner) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModForumGetDiscussionPosts200ResponseRatinginfoScalesInner) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ModForumGetDiscussionPosts200ResponseRatinginfoScalesInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ModForumGetDiscussionPosts200ResponseRatinginfoScalesInner) SetId(v int32) {
	o.Id = &v
}

// GetIsnumeric returns the Isnumeric field value if set, zero value otherwise.
func (o *ModForumGetDiscussionPosts200ResponseRatinginfoScalesInner) GetIsnumeric() bool {
	if o == nil || IsNil(o.Isnumeric) {
		var ret bool
		return ret
	}
	return *o.Isnumeric
}

// GetIsnumericOk returns a tuple with the Isnumeric field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModForumGetDiscussionPosts200ResponseRatinginfoScalesInner) GetIsnumericOk() (*bool, bool) {
	if o == nil || IsNil(o.Isnumeric) {
		return nil, false
	}
	return o.Isnumeric, true
}

// HasIsnumeric returns a boolean if a field has been set.
func (o *ModForumGetDiscussionPosts200ResponseRatinginfoScalesInner) HasIsnumeric() bool {
	if o != nil && !IsNil(o.Isnumeric) {
		return true
	}

	return false
}

// SetIsnumeric gets a reference to the given bool and assigns it to the Isnumeric field.
func (o *ModForumGetDiscussionPosts200ResponseRatinginfoScalesInner) SetIsnumeric(v bool) {
	o.Isnumeric = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ModForumGetDiscussionPosts200ResponseRatinginfoScalesInner) GetItems() []ModForumGetDiscussionPosts200ResponseRatinginfoScalesInnerItemsInner {
	if o == nil || IsNil(o.Items) {
		var ret []ModForumGetDiscussionPosts200ResponseRatinginfoScalesInnerItemsInner
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModForumGetDiscussionPosts200ResponseRatinginfoScalesInner) GetItemsOk() ([]ModForumGetDiscussionPosts200ResponseRatinginfoScalesInnerItemsInner, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ModForumGetDiscussionPosts200ResponseRatinginfoScalesInner) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []ModForumGetDiscussionPosts200ResponseRatinginfoScalesInnerItemsInner and assigns it to the Items field.
func (o *ModForumGetDiscussionPosts200ResponseRatinginfoScalesInner) SetItems(v []ModForumGetDiscussionPosts200ResponseRatinginfoScalesInnerItemsInner) {
	o.Items = v
}

// GetMax returns the Max field value if set, zero value otherwise.
func (o *ModForumGetDiscussionPosts200ResponseRatinginfoScalesInner) GetMax() int32 {
	if o == nil || IsNil(o.Max) {
		var ret int32
		return ret
	}
	return *o.Max
}

// GetMaxOk returns a tuple with the Max field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModForumGetDiscussionPosts200ResponseRatinginfoScalesInner) GetMaxOk() (*int32, bool) {
	if o == nil || IsNil(o.Max) {
		return nil, false
	}
	return o.Max, true
}

// HasMax returns a boolean if a field has been set.
func (o *ModForumGetDiscussionPosts200ResponseRatinginfoScalesInner) HasMax() bool {
	if o != nil && !IsNil(o.Max) {
		return true
	}

	return false
}

// SetMax gets a reference to the given int32 and assigns it to the Max field.
func (o *ModForumGetDiscussionPosts200ResponseRatinginfoScalesInner) SetMax(v int32) {
	o.Max = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ModForumGetDiscussionPosts200ResponseRatinginfoScalesInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModForumGetDiscussionPosts200ResponseRatinginfoScalesInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ModForumGetDiscussionPosts200ResponseRatinginfoScalesInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ModForumGetDiscussionPosts200ResponseRatinginfoScalesInner) SetName(v string) {
	o.Name = &v
}

func (o ModForumGetDiscussionPosts200ResponseRatinginfoScalesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModForumGetDiscussionPosts200ResponseRatinginfoScalesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Courseid) {
		toSerialize["courseid"] = o.Courseid
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Isnumeric) {
		toSerialize["isnumeric"] = o.Isnumeric
	}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !IsNil(o.Max) {
		toSerialize["max"] = o.Max
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableModForumGetDiscussionPosts200ResponseRatinginfoScalesInner struct {
	value *ModForumGetDiscussionPosts200ResponseRatinginfoScalesInner
	isSet bool
}

func (v NullableModForumGetDiscussionPosts200ResponseRatinginfoScalesInner) Get() *ModForumGetDiscussionPosts200ResponseRatinginfoScalesInner {
	return v.value
}

func (v *NullableModForumGetDiscussionPosts200ResponseRatinginfoScalesInner) Set(val *ModForumGetDiscussionPosts200ResponseRatinginfoScalesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableModForumGetDiscussionPosts200ResponseRatinginfoScalesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableModForumGetDiscussionPosts200ResponseRatinginfoScalesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModForumGetDiscussionPosts200ResponseRatinginfoScalesInner(val *ModForumGetDiscussionPosts200ResponseRatinginfoScalesInner) *NullableModForumGetDiscussionPosts200ResponseRatinginfoScalesInner {
	return &NullableModForumGetDiscussionPosts200ResponseRatinginfoScalesInner{value: val, isSet: true}
}

func (v NullableModForumGetDiscussionPosts200ResponseRatinginfoScalesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModForumGetDiscussionPosts200ResponseRatinginfoScalesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


