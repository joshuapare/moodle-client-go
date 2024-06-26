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

// checks if the CoreMessageGetUserMessagePreferences200ResponsePreferencesComponentsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreMessageGetUserMessagePreferences200ResponsePreferencesComponentsInner{}

// CoreMessageGetUserMessagePreferences200ResponsePreferencesComponentsInner struct for CoreMessageGetUserMessagePreferences200ResponsePreferencesComponentsInner
type CoreMessageGetUserMessagePreferences200ResponsePreferencesComponentsInner struct {
	// Display name
	Displayname *string `json:"displayname,omitempty"`
	Notifications []CoreMessageGetUserMessagePreferences200ResponsePreferencesComponentsInnerNotificationsInner `json:"notifications,omitempty"`
}

// NewCoreMessageGetUserMessagePreferences200ResponsePreferencesComponentsInner instantiates a new CoreMessageGetUserMessagePreferences200ResponsePreferencesComponentsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreMessageGetUserMessagePreferences200ResponsePreferencesComponentsInner() *CoreMessageGetUserMessagePreferences200ResponsePreferencesComponentsInner {
	this := CoreMessageGetUserMessagePreferences200ResponsePreferencesComponentsInner{}
	var displayname string = "null"
	this.Displayname = &displayname
	return &this
}

// NewCoreMessageGetUserMessagePreferences200ResponsePreferencesComponentsInnerWithDefaults instantiates a new CoreMessageGetUserMessagePreferences200ResponsePreferencesComponentsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreMessageGetUserMessagePreferences200ResponsePreferencesComponentsInnerWithDefaults() *CoreMessageGetUserMessagePreferences200ResponsePreferencesComponentsInner {
	this := CoreMessageGetUserMessagePreferences200ResponsePreferencesComponentsInner{}
	var displayname string = "null"
	this.Displayname = &displayname
	return &this
}

// GetDisplayname returns the Displayname field value if set, zero value otherwise.
func (o *CoreMessageGetUserMessagePreferences200ResponsePreferencesComponentsInner) GetDisplayname() string {
	if o == nil || IsNil(o.Displayname) {
		var ret string
		return ret
	}
	return *o.Displayname
}

// GetDisplaynameOk returns a tuple with the Displayname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreMessageGetUserMessagePreferences200ResponsePreferencesComponentsInner) GetDisplaynameOk() (*string, bool) {
	if o == nil || IsNil(o.Displayname) {
		return nil, false
	}
	return o.Displayname, true
}

// HasDisplayname returns a boolean if a field has been set.
func (o *CoreMessageGetUserMessagePreferences200ResponsePreferencesComponentsInner) HasDisplayname() bool {
	if o != nil && !IsNil(o.Displayname) {
		return true
	}

	return false
}

// SetDisplayname gets a reference to the given string and assigns it to the Displayname field.
func (o *CoreMessageGetUserMessagePreferences200ResponsePreferencesComponentsInner) SetDisplayname(v string) {
	o.Displayname = &v
}

// GetNotifications returns the Notifications field value if set, zero value otherwise.
func (o *CoreMessageGetUserMessagePreferences200ResponsePreferencesComponentsInner) GetNotifications() []CoreMessageGetUserMessagePreferences200ResponsePreferencesComponentsInnerNotificationsInner {
	if o == nil || IsNil(o.Notifications) {
		var ret []CoreMessageGetUserMessagePreferences200ResponsePreferencesComponentsInnerNotificationsInner
		return ret
	}
	return o.Notifications
}

// GetNotificationsOk returns a tuple with the Notifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreMessageGetUserMessagePreferences200ResponsePreferencesComponentsInner) GetNotificationsOk() ([]CoreMessageGetUserMessagePreferences200ResponsePreferencesComponentsInnerNotificationsInner, bool) {
	if o == nil || IsNil(o.Notifications) {
		return nil, false
	}
	return o.Notifications, true
}

// HasNotifications returns a boolean if a field has been set.
func (o *CoreMessageGetUserMessagePreferences200ResponsePreferencesComponentsInner) HasNotifications() bool {
	if o != nil && !IsNil(o.Notifications) {
		return true
	}

	return false
}

// SetNotifications gets a reference to the given []CoreMessageGetUserMessagePreferences200ResponsePreferencesComponentsInnerNotificationsInner and assigns it to the Notifications field.
func (o *CoreMessageGetUserMessagePreferences200ResponsePreferencesComponentsInner) SetNotifications(v []CoreMessageGetUserMessagePreferences200ResponsePreferencesComponentsInnerNotificationsInner) {
	o.Notifications = v
}

func (o CoreMessageGetUserMessagePreferences200ResponsePreferencesComponentsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreMessageGetUserMessagePreferences200ResponsePreferencesComponentsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Displayname) {
		toSerialize["displayname"] = o.Displayname
	}
	if !IsNil(o.Notifications) {
		toSerialize["notifications"] = o.Notifications
	}
	return toSerialize, nil
}

type NullableCoreMessageGetUserMessagePreferences200ResponsePreferencesComponentsInner struct {
	value *CoreMessageGetUserMessagePreferences200ResponsePreferencesComponentsInner
	isSet bool
}

func (v NullableCoreMessageGetUserMessagePreferences200ResponsePreferencesComponentsInner) Get() *CoreMessageGetUserMessagePreferences200ResponsePreferencesComponentsInner {
	return v.value
}

func (v *NullableCoreMessageGetUserMessagePreferences200ResponsePreferencesComponentsInner) Set(val *CoreMessageGetUserMessagePreferences200ResponsePreferencesComponentsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreMessageGetUserMessagePreferences200ResponsePreferencesComponentsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreMessageGetUserMessagePreferences200ResponsePreferencesComponentsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreMessageGetUserMessagePreferences200ResponsePreferencesComponentsInner(val *CoreMessageGetUserMessagePreferences200ResponsePreferencesComponentsInner) *NullableCoreMessageGetUserMessagePreferences200ResponsePreferencesComponentsInner {
	return &NullableCoreMessageGetUserMessagePreferences200ResponsePreferencesComponentsInner{value: val, isSet: true}
}

func (v NullableCoreMessageGetUserMessagePreferences200ResponsePreferencesComponentsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreMessageGetUserMessagePreferences200ResponsePreferencesComponentsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


