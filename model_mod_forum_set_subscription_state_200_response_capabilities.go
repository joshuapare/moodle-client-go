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

// checks if the ModForumSetSubscriptionState200ResponseCapabilities type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModForumSetSubscriptionState200ResponseCapabilities{}

// ModForumSetSubscriptionState200ResponseCapabilities struct for ModForumSetSubscriptionState200ResponseCapabilities
type ModForumSetSubscriptionState200ResponseCapabilities struct {
	// favourite
	Favourite bool `json:"favourite"`
	// manage
	Manage bool `json:"manage"`
	// move
	Move bool `json:"move"`
	// pin
	Pin bool `json:"pin"`
	// post
	Post bool `json:"post"`
	// subscribe
	Subscribe bool `json:"subscribe"`
}

type _ModForumSetSubscriptionState200ResponseCapabilities ModForumSetSubscriptionState200ResponseCapabilities

// NewModForumSetSubscriptionState200ResponseCapabilities instantiates a new ModForumSetSubscriptionState200ResponseCapabilities object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModForumSetSubscriptionState200ResponseCapabilities(favourite bool, manage bool, move bool, pin bool, post bool, subscribe bool) *ModForumSetSubscriptionState200ResponseCapabilities {
	this := ModForumSetSubscriptionState200ResponseCapabilities{}
	this.Favourite = favourite
	this.Manage = manage
	this.Move = move
	this.Pin = pin
	this.Post = post
	this.Subscribe = subscribe
	return &this
}

// NewModForumSetSubscriptionState200ResponseCapabilitiesWithDefaults instantiates a new ModForumSetSubscriptionState200ResponseCapabilities object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModForumSetSubscriptionState200ResponseCapabilitiesWithDefaults() *ModForumSetSubscriptionState200ResponseCapabilities {
	this := ModForumSetSubscriptionState200ResponseCapabilities{}
	return &this
}

// GetFavourite returns the Favourite field value
func (o *ModForumSetSubscriptionState200ResponseCapabilities) GetFavourite() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Favourite
}

// GetFavouriteOk returns a tuple with the Favourite field value
// and a boolean to check if the value has been set.
func (o *ModForumSetSubscriptionState200ResponseCapabilities) GetFavouriteOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Favourite, true
}

// SetFavourite sets field value
func (o *ModForumSetSubscriptionState200ResponseCapabilities) SetFavourite(v bool) {
	o.Favourite = v
}

// GetManage returns the Manage field value
func (o *ModForumSetSubscriptionState200ResponseCapabilities) GetManage() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Manage
}

// GetManageOk returns a tuple with the Manage field value
// and a boolean to check if the value has been set.
func (o *ModForumSetSubscriptionState200ResponseCapabilities) GetManageOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Manage, true
}

// SetManage sets field value
func (o *ModForumSetSubscriptionState200ResponseCapabilities) SetManage(v bool) {
	o.Manage = v
}

// GetMove returns the Move field value
func (o *ModForumSetSubscriptionState200ResponseCapabilities) GetMove() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Move
}

// GetMoveOk returns a tuple with the Move field value
// and a boolean to check if the value has been set.
func (o *ModForumSetSubscriptionState200ResponseCapabilities) GetMoveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Move, true
}

// SetMove sets field value
func (o *ModForumSetSubscriptionState200ResponseCapabilities) SetMove(v bool) {
	o.Move = v
}

// GetPin returns the Pin field value
func (o *ModForumSetSubscriptionState200ResponseCapabilities) GetPin() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Pin
}

// GetPinOk returns a tuple with the Pin field value
// and a boolean to check if the value has been set.
func (o *ModForumSetSubscriptionState200ResponseCapabilities) GetPinOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pin, true
}

// SetPin sets field value
func (o *ModForumSetSubscriptionState200ResponseCapabilities) SetPin(v bool) {
	o.Pin = v
}

// GetPost returns the Post field value
func (o *ModForumSetSubscriptionState200ResponseCapabilities) GetPost() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Post
}

// GetPostOk returns a tuple with the Post field value
// and a boolean to check if the value has been set.
func (o *ModForumSetSubscriptionState200ResponseCapabilities) GetPostOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Post, true
}

// SetPost sets field value
func (o *ModForumSetSubscriptionState200ResponseCapabilities) SetPost(v bool) {
	o.Post = v
}

// GetSubscribe returns the Subscribe field value
func (o *ModForumSetSubscriptionState200ResponseCapabilities) GetSubscribe() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Subscribe
}

// GetSubscribeOk returns a tuple with the Subscribe field value
// and a boolean to check if the value has been set.
func (o *ModForumSetSubscriptionState200ResponseCapabilities) GetSubscribeOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subscribe, true
}

// SetSubscribe sets field value
func (o *ModForumSetSubscriptionState200ResponseCapabilities) SetSubscribe(v bool) {
	o.Subscribe = v
}

func (o ModForumSetSubscriptionState200ResponseCapabilities) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModForumSetSubscriptionState200ResponseCapabilities) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["favourite"] = o.Favourite
	toSerialize["manage"] = o.Manage
	toSerialize["move"] = o.Move
	toSerialize["pin"] = o.Pin
	toSerialize["post"] = o.Post
	toSerialize["subscribe"] = o.Subscribe
	return toSerialize, nil
}

func (o *ModForumSetSubscriptionState200ResponseCapabilities) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"favourite",
		"manage",
		"move",
		"pin",
		"post",
		"subscribe",
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

	varModForumSetSubscriptionState200ResponseCapabilities := _ModForumSetSubscriptionState200ResponseCapabilities{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModForumSetSubscriptionState200ResponseCapabilities)

	if err != nil {
		return err
	}

	*o = ModForumSetSubscriptionState200ResponseCapabilities(varModForumSetSubscriptionState200ResponseCapabilities)

	return err
}

type NullableModForumSetSubscriptionState200ResponseCapabilities struct {
	value *ModForumSetSubscriptionState200ResponseCapabilities
	isSet bool
}

func (v NullableModForumSetSubscriptionState200ResponseCapabilities) Get() *ModForumSetSubscriptionState200ResponseCapabilities {
	return v.value
}

func (v *NullableModForumSetSubscriptionState200ResponseCapabilities) Set(val *ModForumSetSubscriptionState200ResponseCapabilities) {
	v.value = val
	v.isSet = true
}

func (v NullableModForumSetSubscriptionState200ResponseCapabilities) IsSet() bool {
	return v.isSet
}

func (v *NullableModForumSetSubscriptionState200ResponseCapabilities) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModForumSetSubscriptionState200ResponseCapabilities(val *ModForumSetSubscriptionState200ResponseCapabilities) *NullableModForumSetSubscriptionState200ResponseCapabilities {
	return &NullableModForumSetSubscriptionState200ResponseCapabilities{value: val, isSet: true}
}

func (v NullableModForumSetSubscriptionState200ResponseCapabilities) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModForumSetSubscriptionState200ResponseCapabilities) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

