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

// checks if the ModForumSetSubscriptionState200ResponseTimed type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModForumSetSubscriptionState200ResponseTimed{}

// ModForumSetSubscriptionState200ResponseTimed struct for ModForumSetSubscriptionState200ResponseTimed
type ModForumSetSubscriptionState200ResponseTimed struct {
	// istimed
	Istimed *bool `json:"istimed,omitempty"`
	// visible
	Visible *bool `json:"visible,omitempty"`
}

// NewModForumSetSubscriptionState200ResponseTimed instantiates a new ModForumSetSubscriptionState200ResponseTimed object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModForumSetSubscriptionState200ResponseTimed() *ModForumSetSubscriptionState200ResponseTimed {
	this := ModForumSetSubscriptionState200ResponseTimed{}
	return &this
}

// NewModForumSetSubscriptionState200ResponseTimedWithDefaults instantiates a new ModForumSetSubscriptionState200ResponseTimed object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModForumSetSubscriptionState200ResponseTimedWithDefaults() *ModForumSetSubscriptionState200ResponseTimed {
	this := ModForumSetSubscriptionState200ResponseTimed{}
	return &this
}

// GetIstimed returns the Istimed field value if set, zero value otherwise.
func (o *ModForumSetSubscriptionState200ResponseTimed) GetIstimed() bool {
	if o == nil || IsNil(o.Istimed) {
		var ret bool
		return ret
	}
	return *o.Istimed
}

// GetIstimedOk returns a tuple with the Istimed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModForumSetSubscriptionState200ResponseTimed) GetIstimedOk() (*bool, bool) {
	if o == nil || IsNil(o.Istimed) {
		return nil, false
	}
	return o.Istimed, true
}

// HasIstimed returns a boolean if a field has been set.
func (o *ModForumSetSubscriptionState200ResponseTimed) HasIstimed() bool {
	if o != nil && !IsNil(o.Istimed) {
		return true
	}

	return false
}

// SetIstimed gets a reference to the given bool and assigns it to the Istimed field.
func (o *ModForumSetSubscriptionState200ResponseTimed) SetIstimed(v bool) {
	o.Istimed = &v
}

// GetVisible returns the Visible field value if set, zero value otherwise.
func (o *ModForumSetSubscriptionState200ResponseTimed) GetVisible() bool {
	if o == nil || IsNil(o.Visible) {
		var ret bool
		return ret
	}
	return *o.Visible
}

// GetVisibleOk returns a tuple with the Visible field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModForumSetSubscriptionState200ResponseTimed) GetVisibleOk() (*bool, bool) {
	if o == nil || IsNil(o.Visible) {
		return nil, false
	}
	return o.Visible, true
}

// HasVisible returns a boolean if a field has been set.
func (o *ModForumSetSubscriptionState200ResponseTimed) HasVisible() bool {
	if o != nil && !IsNil(o.Visible) {
		return true
	}

	return false
}

// SetVisible gets a reference to the given bool and assigns it to the Visible field.
func (o *ModForumSetSubscriptionState200ResponseTimed) SetVisible(v bool) {
	o.Visible = &v
}

func (o ModForumSetSubscriptionState200ResponseTimed) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModForumSetSubscriptionState200ResponseTimed) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Istimed) {
		toSerialize["istimed"] = o.Istimed
	}
	if !IsNil(o.Visible) {
		toSerialize["visible"] = o.Visible
	}
	return toSerialize, nil
}

type NullableModForumSetSubscriptionState200ResponseTimed struct {
	value *ModForumSetSubscriptionState200ResponseTimed
	isSet bool
}

func (v NullableModForumSetSubscriptionState200ResponseTimed) Get() *ModForumSetSubscriptionState200ResponseTimed {
	return v.value
}

func (v *NullableModForumSetSubscriptionState200ResponseTimed) Set(val *ModForumSetSubscriptionState200ResponseTimed) {
	v.value = val
	v.isSet = true
}

func (v NullableModForumSetSubscriptionState200ResponseTimed) IsSet() bool {
	return v.isSet
}

func (v *NullableModForumSetSubscriptionState200ResponseTimed) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModForumSetSubscriptionState200ResponseTimed(val *ModForumSetSubscriptionState200ResponseTimed) *NullableModForumSetSubscriptionState200ResponseTimed {
	return &NullableModForumSetSubscriptionState200ResponseTimed{value: val, isSet: true}
}

func (v NullableModForumSetSubscriptionState200ResponseTimed) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModForumSetSubscriptionState200ResponseTimed) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


