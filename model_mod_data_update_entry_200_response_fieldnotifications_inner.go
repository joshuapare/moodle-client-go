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

// checks if the ModDataUpdateEntry200ResponseFieldnotificationsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModDataUpdateEntry200ResponseFieldnotificationsInner{}

// ModDataUpdateEntry200ResponseFieldnotificationsInner struct for ModDataUpdateEntry200ResponseFieldnotificationsInner
type ModDataUpdateEntry200ResponseFieldnotificationsInner struct {
	// The field name.
	Fieldname *string `json:"fieldname,omitempty"`
	// The notification for the field.
	Notification *string `json:"notification,omitempty"`
}

// NewModDataUpdateEntry200ResponseFieldnotificationsInner instantiates a new ModDataUpdateEntry200ResponseFieldnotificationsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModDataUpdateEntry200ResponseFieldnotificationsInner() *ModDataUpdateEntry200ResponseFieldnotificationsInner {
	this := ModDataUpdateEntry200ResponseFieldnotificationsInner{}
	return &this
}

// NewModDataUpdateEntry200ResponseFieldnotificationsInnerWithDefaults instantiates a new ModDataUpdateEntry200ResponseFieldnotificationsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModDataUpdateEntry200ResponseFieldnotificationsInnerWithDefaults() *ModDataUpdateEntry200ResponseFieldnotificationsInner {
	this := ModDataUpdateEntry200ResponseFieldnotificationsInner{}
	return &this
}

// GetFieldname returns the Fieldname field value if set, zero value otherwise.
func (o *ModDataUpdateEntry200ResponseFieldnotificationsInner) GetFieldname() string {
	if o == nil || IsNil(o.Fieldname) {
		var ret string
		return ret
	}
	return *o.Fieldname
}

// GetFieldnameOk returns a tuple with the Fieldname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModDataUpdateEntry200ResponseFieldnotificationsInner) GetFieldnameOk() (*string, bool) {
	if o == nil || IsNil(o.Fieldname) {
		return nil, false
	}
	return o.Fieldname, true
}

// HasFieldname returns a boolean if a field has been set.
func (o *ModDataUpdateEntry200ResponseFieldnotificationsInner) HasFieldname() bool {
	if o != nil && !IsNil(o.Fieldname) {
		return true
	}

	return false
}

// SetFieldname gets a reference to the given string and assigns it to the Fieldname field.
func (o *ModDataUpdateEntry200ResponseFieldnotificationsInner) SetFieldname(v string) {
	o.Fieldname = &v
}

// GetNotification returns the Notification field value if set, zero value otherwise.
func (o *ModDataUpdateEntry200ResponseFieldnotificationsInner) GetNotification() string {
	if o == nil || IsNil(o.Notification) {
		var ret string
		return ret
	}
	return *o.Notification
}

// GetNotificationOk returns a tuple with the Notification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModDataUpdateEntry200ResponseFieldnotificationsInner) GetNotificationOk() (*string, bool) {
	if o == nil || IsNil(o.Notification) {
		return nil, false
	}
	return o.Notification, true
}

// HasNotification returns a boolean if a field has been set.
func (o *ModDataUpdateEntry200ResponseFieldnotificationsInner) HasNotification() bool {
	if o != nil && !IsNil(o.Notification) {
		return true
	}

	return false
}

// SetNotification gets a reference to the given string and assigns it to the Notification field.
func (o *ModDataUpdateEntry200ResponseFieldnotificationsInner) SetNotification(v string) {
	o.Notification = &v
}

func (o ModDataUpdateEntry200ResponseFieldnotificationsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModDataUpdateEntry200ResponseFieldnotificationsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Fieldname) {
		toSerialize["fieldname"] = o.Fieldname
	}
	if !IsNil(o.Notification) {
		toSerialize["notification"] = o.Notification
	}
	return toSerialize, nil
}

type NullableModDataUpdateEntry200ResponseFieldnotificationsInner struct {
	value *ModDataUpdateEntry200ResponseFieldnotificationsInner
	isSet bool
}

func (v NullableModDataUpdateEntry200ResponseFieldnotificationsInner) Get() *ModDataUpdateEntry200ResponseFieldnotificationsInner {
	return v.value
}

func (v *NullableModDataUpdateEntry200ResponseFieldnotificationsInner) Set(val *ModDataUpdateEntry200ResponseFieldnotificationsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableModDataUpdateEntry200ResponseFieldnotificationsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableModDataUpdateEntry200ResponseFieldnotificationsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModDataUpdateEntry200ResponseFieldnotificationsInner(val *ModDataUpdateEntry200ResponseFieldnotificationsInner) *NullableModDataUpdateEntry200ResponseFieldnotificationsInner {
	return &NullableModDataUpdateEntry200ResponseFieldnotificationsInner{value: val, isSet: true}
}

func (v NullableModDataUpdateEntry200ResponseFieldnotificationsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModDataUpdateEntry200ResponseFieldnotificationsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


