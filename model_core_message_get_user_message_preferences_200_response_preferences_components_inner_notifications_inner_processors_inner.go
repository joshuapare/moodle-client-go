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

// checks if the CoreMessageGetUserMessagePreferences200ResponsePreferencesComponentsInnerNotificationsInnerProcessorsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreMessageGetUserMessagePreferences200ResponsePreferencesComponentsInnerNotificationsInnerProcessorsInner{}

// CoreMessageGetUserMessagePreferences200ResponsePreferencesComponentsInnerNotificationsInnerProcessorsInner struct for CoreMessageGetUserMessagePreferences200ResponsePreferencesComponentsInnerNotificationsInnerProcessorsInner
type CoreMessageGetUserMessagePreferences200ResponsePreferencesComponentsInnerNotificationsInnerProcessorsInner struct {
	// Display name
	Displayname *string `json:"displayname,omitempty"`
	// Is enabled?
	Enabled *bool `json:"enabled,omitempty"`
	// Is locked by admin?
	Locked *bool `json:"locked,omitempty"`
	// Text to display if locked
	Lockedmessage *string `json:"lockedmessage,omitempty"`
	Loggedin *CoreMessageGetUserMessagePreferences200ResponsePreferencesComponentsInnerNotificationsInnerProcessorsInnerLoggedin `json:"loggedin,omitempty"`
	Loggedoff *CoreMessageGetUserMessagePreferences200ResponsePreferencesComponentsInnerNotificationsInnerProcessorsInnerLoggedoff `json:"loggedoff,omitempty"`
	// Processor name
	Name *string `json:"name,omitempty"`
	// Is configured?
	Userconfigured *int32 `json:"userconfigured,omitempty"`
}

// NewCoreMessageGetUserMessagePreferences200ResponsePreferencesComponentsInnerNotificationsInnerProcessorsInner instantiates a new CoreMessageGetUserMessagePreferences200ResponsePreferencesComponentsInnerNotificationsInnerProcessorsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreMessageGetUserMessagePreferences200ResponsePreferencesComponentsInnerNotificationsInnerProcessorsInner() *CoreMessageGetUserMessagePreferences200ResponsePreferencesComponentsInnerNotificationsInnerProcessorsInner {
	this := CoreMessageGetUserMessagePreferences200ResponsePreferencesComponentsInnerNotificationsInnerProcessorsInner{}
	var enabled bool = null
	this.Enabled = &enabled
	var locked bool = null
	this.Locked = &locked
	var lockedmessage string = "null"
	this.Lockedmessage = &lockedmessage
	var name string = "null"
	this.Name = &name
	var userconfigured int32 = null
	this.Userconfigured = &userconfigured
	return &this
}

// NewCoreMessageGetUserMessagePreferences200ResponsePreferencesComponentsInnerNotificationsInnerProcessorsInnerWithDefaults instantiates a new CoreMessageGetUserMessagePreferences200ResponsePreferencesComponentsInnerNotificationsInnerProcessorsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreMessageGetUserMessagePreferences200ResponsePreferencesComponentsInnerNotificationsInnerProcessorsInnerWithDefaults() *CoreMessageGetUserMessagePreferences200ResponsePreferencesComponentsInnerNotificationsInnerProcessorsInner {
	this := CoreMessageGetUserMessagePreferences200ResponsePreferencesComponentsInnerNotificationsInnerProcessorsInner{}
	var enabled bool = null
	this.Enabled = &enabled
	var locked bool = null
	this.Locked = &locked
	var lockedmessage string = "null"
	this.Lockedmessage = &lockedmessage
	var name string = "null"
	this.Name = &name
	var userconfigured int32 = null
	this.Userconfigured = &userconfigured
	return &this
}

// GetDisplayname returns the Displayname field value if set, zero value otherwise.
func (o *CoreMessageGetUserMessagePreferences200ResponsePreferencesComponentsInnerNotificationsInnerProcessorsInner) GetDisplayname() string {
	if o == nil || IsNil(o.Displayname) {
		var ret string
		return ret
	}
	return *o.Displayname
}

// GetDisplaynameOk returns a tuple with the Displayname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreMessageGetUserMessagePreferences200ResponsePreferencesComponentsInnerNotificationsInnerProcessorsInner) GetDisplaynameOk() (*string, bool) {
	if o == nil || IsNil(o.Displayname) {
		return nil, false
	}
	return o.Displayname, true
}

// HasDisplayname returns a boolean if a field has been set.
func (o *CoreMessageGetUserMessagePreferences200ResponsePreferencesComponentsInnerNotificationsInnerProcessorsInner) HasDisplayname() bool {
	if o != nil && !IsNil(o.Displayname) {
		return true
	}

	return false
}

// SetDisplayname gets a reference to the given string and assigns it to the Displayname field.
func (o *CoreMessageGetUserMessagePreferences200ResponsePreferencesComponentsInnerNotificationsInnerProcessorsInner) SetDisplayname(v string) {
	o.Displayname = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *CoreMessageGetUserMessagePreferences200ResponsePreferencesComponentsInnerNotificationsInnerProcessorsInner) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreMessageGetUserMessagePreferences200ResponsePreferencesComponentsInnerNotificationsInnerProcessorsInner) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *CoreMessageGetUserMessagePreferences200ResponsePreferencesComponentsInnerNotificationsInnerProcessorsInner) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *CoreMessageGetUserMessagePreferences200ResponsePreferencesComponentsInnerNotificationsInnerProcessorsInner) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetLocked returns the Locked field value if set, zero value otherwise.
func (o *CoreMessageGetUserMessagePreferences200ResponsePreferencesComponentsInnerNotificationsInnerProcessorsInner) GetLocked() bool {
	if o == nil || IsNil(o.Locked) {
		var ret bool
		return ret
	}
	return *o.Locked
}

// GetLockedOk returns a tuple with the Locked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreMessageGetUserMessagePreferences200ResponsePreferencesComponentsInnerNotificationsInnerProcessorsInner) GetLockedOk() (*bool, bool) {
	if o == nil || IsNil(o.Locked) {
		return nil, false
	}
	return o.Locked, true
}

// HasLocked returns a boolean if a field has been set.
func (o *CoreMessageGetUserMessagePreferences200ResponsePreferencesComponentsInnerNotificationsInnerProcessorsInner) HasLocked() bool {
	if o != nil && !IsNil(o.Locked) {
		return true
	}

	return false
}

// SetLocked gets a reference to the given bool and assigns it to the Locked field.
func (o *CoreMessageGetUserMessagePreferences200ResponsePreferencesComponentsInnerNotificationsInnerProcessorsInner) SetLocked(v bool) {
	o.Locked = &v
}

// GetLockedmessage returns the Lockedmessage field value if set, zero value otherwise.
func (o *CoreMessageGetUserMessagePreferences200ResponsePreferencesComponentsInnerNotificationsInnerProcessorsInner) GetLockedmessage() string {
	if o == nil || IsNil(o.Lockedmessage) {
		var ret string
		return ret
	}
	return *o.Lockedmessage
}

// GetLockedmessageOk returns a tuple with the Lockedmessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreMessageGetUserMessagePreferences200ResponsePreferencesComponentsInnerNotificationsInnerProcessorsInner) GetLockedmessageOk() (*string, bool) {
	if o == nil || IsNil(o.Lockedmessage) {
		return nil, false
	}
	return o.Lockedmessage, true
}

// HasLockedmessage returns a boolean if a field has been set.
func (o *CoreMessageGetUserMessagePreferences200ResponsePreferencesComponentsInnerNotificationsInnerProcessorsInner) HasLockedmessage() bool {
	if o != nil && !IsNil(o.Lockedmessage) {
		return true
	}

	return false
}

// SetLockedmessage gets a reference to the given string and assigns it to the Lockedmessage field.
func (o *CoreMessageGetUserMessagePreferences200ResponsePreferencesComponentsInnerNotificationsInnerProcessorsInner) SetLockedmessage(v string) {
	o.Lockedmessage = &v
}

// GetLoggedin returns the Loggedin field value if set, zero value otherwise.
func (o *CoreMessageGetUserMessagePreferences200ResponsePreferencesComponentsInnerNotificationsInnerProcessorsInner) GetLoggedin() CoreMessageGetUserMessagePreferences200ResponsePreferencesComponentsInnerNotificationsInnerProcessorsInnerLoggedin {
	if o == nil || IsNil(o.Loggedin) {
		var ret CoreMessageGetUserMessagePreferences200ResponsePreferencesComponentsInnerNotificationsInnerProcessorsInnerLoggedin
		return ret
	}
	return *o.Loggedin
}

// GetLoggedinOk returns a tuple with the Loggedin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreMessageGetUserMessagePreferences200ResponsePreferencesComponentsInnerNotificationsInnerProcessorsInner) GetLoggedinOk() (*CoreMessageGetUserMessagePreferences200ResponsePreferencesComponentsInnerNotificationsInnerProcessorsInnerLoggedin, bool) {
	if o == nil || IsNil(o.Loggedin) {
		return nil, false
	}
	return o.Loggedin, true
}

// HasLoggedin returns a boolean if a field has been set.
func (o *CoreMessageGetUserMessagePreferences200ResponsePreferencesComponentsInnerNotificationsInnerProcessorsInner) HasLoggedin() bool {
	if o != nil && !IsNil(o.Loggedin) {
		return true
	}

	return false
}

// SetLoggedin gets a reference to the given CoreMessageGetUserMessagePreferences200ResponsePreferencesComponentsInnerNotificationsInnerProcessorsInnerLoggedin and assigns it to the Loggedin field.
func (o *CoreMessageGetUserMessagePreferences200ResponsePreferencesComponentsInnerNotificationsInnerProcessorsInner) SetLoggedin(v CoreMessageGetUserMessagePreferences200ResponsePreferencesComponentsInnerNotificationsInnerProcessorsInnerLoggedin) {
	o.Loggedin = &v
}

// GetLoggedoff returns the Loggedoff field value if set, zero value otherwise.
func (o *CoreMessageGetUserMessagePreferences200ResponsePreferencesComponentsInnerNotificationsInnerProcessorsInner) GetLoggedoff() CoreMessageGetUserMessagePreferences200ResponsePreferencesComponentsInnerNotificationsInnerProcessorsInnerLoggedoff {
	if o == nil || IsNil(o.Loggedoff) {
		var ret CoreMessageGetUserMessagePreferences200ResponsePreferencesComponentsInnerNotificationsInnerProcessorsInnerLoggedoff
		return ret
	}
	return *o.Loggedoff
}

// GetLoggedoffOk returns a tuple with the Loggedoff field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreMessageGetUserMessagePreferences200ResponsePreferencesComponentsInnerNotificationsInnerProcessorsInner) GetLoggedoffOk() (*CoreMessageGetUserMessagePreferences200ResponsePreferencesComponentsInnerNotificationsInnerProcessorsInnerLoggedoff, bool) {
	if o == nil || IsNil(o.Loggedoff) {
		return nil, false
	}
	return o.Loggedoff, true
}

// HasLoggedoff returns a boolean if a field has been set.
func (o *CoreMessageGetUserMessagePreferences200ResponsePreferencesComponentsInnerNotificationsInnerProcessorsInner) HasLoggedoff() bool {
	if o != nil && !IsNil(o.Loggedoff) {
		return true
	}

	return false
}

// SetLoggedoff gets a reference to the given CoreMessageGetUserMessagePreferences200ResponsePreferencesComponentsInnerNotificationsInnerProcessorsInnerLoggedoff and assigns it to the Loggedoff field.
func (o *CoreMessageGetUserMessagePreferences200ResponsePreferencesComponentsInnerNotificationsInnerProcessorsInner) SetLoggedoff(v CoreMessageGetUserMessagePreferences200ResponsePreferencesComponentsInnerNotificationsInnerProcessorsInnerLoggedoff) {
	o.Loggedoff = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CoreMessageGetUserMessagePreferences200ResponsePreferencesComponentsInnerNotificationsInnerProcessorsInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreMessageGetUserMessagePreferences200ResponsePreferencesComponentsInnerNotificationsInnerProcessorsInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CoreMessageGetUserMessagePreferences200ResponsePreferencesComponentsInnerNotificationsInnerProcessorsInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CoreMessageGetUserMessagePreferences200ResponsePreferencesComponentsInnerNotificationsInnerProcessorsInner) SetName(v string) {
	o.Name = &v
}

// GetUserconfigured returns the Userconfigured field value if set, zero value otherwise.
func (o *CoreMessageGetUserMessagePreferences200ResponsePreferencesComponentsInnerNotificationsInnerProcessorsInner) GetUserconfigured() int32 {
	if o == nil || IsNil(o.Userconfigured) {
		var ret int32
		return ret
	}
	return *o.Userconfigured
}

// GetUserconfiguredOk returns a tuple with the Userconfigured field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreMessageGetUserMessagePreferences200ResponsePreferencesComponentsInnerNotificationsInnerProcessorsInner) GetUserconfiguredOk() (*int32, bool) {
	if o == nil || IsNil(o.Userconfigured) {
		return nil, false
	}
	return o.Userconfigured, true
}

// HasUserconfigured returns a boolean if a field has been set.
func (o *CoreMessageGetUserMessagePreferences200ResponsePreferencesComponentsInnerNotificationsInnerProcessorsInner) HasUserconfigured() bool {
	if o != nil && !IsNil(o.Userconfigured) {
		return true
	}

	return false
}

// SetUserconfigured gets a reference to the given int32 and assigns it to the Userconfigured field.
func (o *CoreMessageGetUserMessagePreferences200ResponsePreferencesComponentsInnerNotificationsInnerProcessorsInner) SetUserconfigured(v int32) {
	o.Userconfigured = &v
}

func (o CoreMessageGetUserMessagePreferences200ResponsePreferencesComponentsInnerNotificationsInnerProcessorsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreMessageGetUserMessagePreferences200ResponsePreferencesComponentsInnerNotificationsInnerProcessorsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Displayname) {
		toSerialize["displayname"] = o.Displayname
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Locked) {
		toSerialize["locked"] = o.Locked
	}
	if !IsNil(o.Lockedmessage) {
		toSerialize["lockedmessage"] = o.Lockedmessage
	}
	if !IsNil(o.Loggedin) {
		toSerialize["loggedin"] = o.Loggedin
	}
	if !IsNil(o.Loggedoff) {
		toSerialize["loggedoff"] = o.Loggedoff
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Userconfigured) {
		toSerialize["userconfigured"] = o.Userconfigured
	}
	return toSerialize, nil
}

type NullableCoreMessageGetUserMessagePreferences200ResponsePreferencesComponentsInnerNotificationsInnerProcessorsInner struct {
	value *CoreMessageGetUserMessagePreferences200ResponsePreferencesComponentsInnerNotificationsInnerProcessorsInner
	isSet bool
}

func (v NullableCoreMessageGetUserMessagePreferences200ResponsePreferencesComponentsInnerNotificationsInnerProcessorsInner) Get() *CoreMessageGetUserMessagePreferences200ResponsePreferencesComponentsInnerNotificationsInnerProcessorsInner {
	return v.value
}

func (v *NullableCoreMessageGetUserMessagePreferences200ResponsePreferencesComponentsInnerNotificationsInnerProcessorsInner) Set(val *CoreMessageGetUserMessagePreferences200ResponsePreferencesComponentsInnerNotificationsInnerProcessorsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreMessageGetUserMessagePreferences200ResponsePreferencesComponentsInnerNotificationsInnerProcessorsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreMessageGetUserMessagePreferences200ResponsePreferencesComponentsInnerNotificationsInnerProcessorsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreMessageGetUserMessagePreferences200ResponsePreferencesComponentsInnerNotificationsInnerProcessorsInner(val *CoreMessageGetUserMessagePreferences200ResponsePreferencesComponentsInnerNotificationsInnerProcessorsInner) *NullableCoreMessageGetUserMessagePreferences200ResponsePreferencesComponentsInnerNotificationsInnerProcessorsInner {
	return &NullableCoreMessageGetUserMessagePreferences200ResponsePreferencesComponentsInnerNotificationsInnerProcessorsInner{value: val, isSet: true}
}

func (v NullableCoreMessageGetUserMessagePreferences200ResponsePreferencesComponentsInnerNotificationsInnerProcessorsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreMessageGetUserMessagePreferences200ResponsePreferencesComponentsInnerNotificationsInnerProcessorsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


