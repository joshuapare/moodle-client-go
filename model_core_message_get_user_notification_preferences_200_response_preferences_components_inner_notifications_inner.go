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

// checks if the CoreMessageGetUserNotificationPreferences200ResponsePreferencesComponentsInnerNotificationsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreMessageGetUserNotificationPreferences200ResponsePreferencesComponentsInnerNotificationsInner{}

// CoreMessageGetUserNotificationPreferences200ResponsePreferencesComponentsInnerNotificationsInner struct for CoreMessageGetUserNotificationPreferences200ResponsePreferencesComponentsInnerNotificationsInner
type CoreMessageGetUserNotificationPreferences200ResponsePreferencesComponentsInnerNotificationsInner struct {
	// Display name
	Displayname *string `json:"displayname,omitempty"`
	// Preference key
	Preferencekey *string `json:"preferencekey,omitempty"`
	Processors []CoreMessageGetUserNotificationPreferences200ResponsePreferencesComponentsInnerNotificationsInnerProcessorsInner `json:"processors,omitempty"`
}

// NewCoreMessageGetUserNotificationPreferences200ResponsePreferencesComponentsInnerNotificationsInner instantiates a new CoreMessageGetUserNotificationPreferences200ResponsePreferencesComponentsInnerNotificationsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreMessageGetUserNotificationPreferences200ResponsePreferencesComponentsInnerNotificationsInner() *CoreMessageGetUserNotificationPreferences200ResponsePreferencesComponentsInnerNotificationsInner {
	this := CoreMessageGetUserNotificationPreferences200ResponsePreferencesComponentsInnerNotificationsInner{}
	return &this
}

// NewCoreMessageGetUserNotificationPreferences200ResponsePreferencesComponentsInnerNotificationsInnerWithDefaults instantiates a new CoreMessageGetUserNotificationPreferences200ResponsePreferencesComponentsInnerNotificationsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreMessageGetUserNotificationPreferences200ResponsePreferencesComponentsInnerNotificationsInnerWithDefaults() *CoreMessageGetUserNotificationPreferences200ResponsePreferencesComponentsInnerNotificationsInner {
	this := CoreMessageGetUserNotificationPreferences200ResponsePreferencesComponentsInnerNotificationsInner{}
	return &this
}

// GetDisplayname returns the Displayname field value if set, zero value otherwise.
func (o *CoreMessageGetUserNotificationPreferences200ResponsePreferencesComponentsInnerNotificationsInner) GetDisplayname() string {
	if o == nil || IsNil(o.Displayname) {
		var ret string
		return ret
	}
	return *o.Displayname
}

// GetDisplaynameOk returns a tuple with the Displayname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreMessageGetUserNotificationPreferences200ResponsePreferencesComponentsInnerNotificationsInner) GetDisplaynameOk() (*string, bool) {
	if o == nil || IsNil(o.Displayname) {
		return nil, false
	}
	return o.Displayname, true
}

// HasDisplayname returns a boolean if a field has been set.
func (o *CoreMessageGetUserNotificationPreferences200ResponsePreferencesComponentsInnerNotificationsInner) HasDisplayname() bool {
	if o != nil && !IsNil(o.Displayname) {
		return true
	}

	return false
}

// SetDisplayname gets a reference to the given string and assigns it to the Displayname field.
func (o *CoreMessageGetUserNotificationPreferences200ResponsePreferencesComponentsInnerNotificationsInner) SetDisplayname(v string) {
	o.Displayname = &v
}

// GetPreferencekey returns the Preferencekey field value if set, zero value otherwise.
func (o *CoreMessageGetUserNotificationPreferences200ResponsePreferencesComponentsInnerNotificationsInner) GetPreferencekey() string {
	if o == nil || IsNil(o.Preferencekey) {
		var ret string
		return ret
	}
	return *o.Preferencekey
}

// GetPreferencekeyOk returns a tuple with the Preferencekey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreMessageGetUserNotificationPreferences200ResponsePreferencesComponentsInnerNotificationsInner) GetPreferencekeyOk() (*string, bool) {
	if o == nil || IsNil(o.Preferencekey) {
		return nil, false
	}
	return o.Preferencekey, true
}

// HasPreferencekey returns a boolean if a field has been set.
func (o *CoreMessageGetUserNotificationPreferences200ResponsePreferencesComponentsInnerNotificationsInner) HasPreferencekey() bool {
	if o != nil && !IsNil(o.Preferencekey) {
		return true
	}

	return false
}

// SetPreferencekey gets a reference to the given string and assigns it to the Preferencekey field.
func (o *CoreMessageGetUserNotificationPreferences200ResponsePreferencesComponentsInnerNotificationsInner) SetPreferencekey(v string) {
	o.Preferencekey = &v
}

// GetProcessors returns the Processors field value if set, zero value otherwise.
func (o *CoreMessageGetUserNotificationPreferences200ResponsePreferencesComponentsInnerNotificationsInner) GetProcessors() []CoreMessageGetUserNotificationPreferences200ResponsePreferencesComponentsInnerNotificationsInnerProcessorsInner {
	if o == nil || IsNil(o.Processors) {
		var ret []CoreMessageGetUserNotificationPreferences200ResponsePreferencesComponentsInnerNotificationsInnerProcessorsInner
		return ret
	}
	return o.Processors
}

// GetProcessorsOk returns a tuple with the Processors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreMessageGetUserNotificationPreferences200ResponsePreferencesComponentsInnerNotificationsInner) GetProcessorsOk() ([]CoreMessageGetUserNotificationPreferences200ResponsePreferencesComponentsInnerNotificationsInnerProcessorsInner, bool) {
	if o == nil || IsNil(o.Processors) {
		return nil, false
	}
	return o.Processors, true
}

// HasProcessors returns a boolean if a field has been set.
func (o *CoreMessageGetUserNotificationPreferences200ResponsePreferencesComponentsInnerNotificationsInner) HasProcessors() bool {
	if o != nil && !IsNil(o.Processors) {
		return true
	}

	return false
}

// SetProcessors gets a reference to the given []CoreMessageGetUserNotificationPreferences200ResponsePreferencesComponentsInnerNotificationsInnerProcessorsInner and assigns it to the Processors field.
func (o *CoreMessageGetUserNotificationPreferences200ResponsePreferencesComponentsInnerNotificationsInner) SetProcessors(v []CoreMessageGetUserNotificationPreferences200ResponsePreferencesComponentsInnerNotificationsInnerProcessorsInner) {
	o.Processors = v
}

func (o CoreMessageGetUserNotificationPreferences200ResponsePreferencesComponentsInnerNotificationsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreMessageGetUserNotificationPreferences200ResponsePreferencesComponentsInnerNotificationsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Displayname) {
		toSerialize["displayname"] = o.Displayname
	}
	if !IsNil(o.Preferencekey) {
		toSerialize["preferencekey"] = o.Preferencekey
	}
	if !IsNil(o.Processors) {
		toSerialize["processors"] = o.Processors
	}
	return toSerialize, nil
}

type NullableCoreMessageGetUserNotificationPreferences200ResponsePreferencesComponentsInnerNotificationsInner struct {
	value *CoreMessageGetUserNotificationPreferences200ResponsePreferencesComponentsInnerNotificationsInner
	isSet bool
}

func (v NullableCoreMessageGetUserNotificationPreferences200ResponsePreferencesComponentsInnerNotificationsInner) Get() *CoreMessageGetUserNotificationPreferences200ResponsePreferencesComponentsInnerNotificationsInner {
	return v.value
}

func (v *NullableCoreMessageGetUserNotificationPreferences200ResponsePreferencesComponentsInnerNotificationsInner) Set(val *CoreMessageGetUserNotificationPreferences200ResponsePreferencesComponentsInnerNotificationsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreMessageGetUserNotificationPreferences200ResponsePreferencesComponentsInnerNotificationsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreMessageGetUserNotificationPreferences200ResponsePreferencesComponentsInnerNotificationsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreMessageGetUserNotificationPreferences200ResponsePreferencesComponentsInnerNotificationsInner(val *CoreMessageGetUserNotificationPreferences200ResponsePreferencesComponentsInnerNotificationsInner) *NullableCoreMessageGetUserNotificationPreferences200ResponsePreferencesComponentsInnerNotificationsInner {
	return &NullableCoreMessageGetUserNotificationPreferences200ResponsePreferencesComponentsInnerNotificationsInner{value: val, isSet: true}
}

func (v NullableCoreMessageGetUserNotificationPreferences200ResponsePreferencesComponentsInnerNotificationsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreMessageGetUserNotificationPreferences200ResponsePreferencesComponentsInnerNotificationsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


