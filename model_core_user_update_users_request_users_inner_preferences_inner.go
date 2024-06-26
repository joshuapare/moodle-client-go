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

// checks if the CoreUserUpdateUsersRequestUsersInnerPreferencesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreUserUpdateUsersRequestUsersInnerPreferencesInner{}

// CoreUserUpdateUsersRequestUsersInnerPreferencesInner struct for CoreUserUpdateUsersRequestUsersInnerPreferencesInner
type CoreUserUpdateUsersRequestUsersInnerPreferencesInner struct {
	// The name of the preference
	Type *string `json:"type,omitempty"`
	// The value of the preference
	Value *string `json:"value,omitempty"`
}

// NewCoreUserUpdateUsersRequestUsersInnerPreferencesInner instantiates a new CoreUserUpdateUsersRequestUsersInnerPreferencesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreUserUpdateUsersRequestUsersInnerPreferencesInner() *CoreUserUpdateUsersRequestUsersInnerPreferencesInner {
	this := CoreUserUpdateUsersRequestUsersInnerPreferencesInner{}
	return &this
}

// NewCoreUserUpdateUsersRequestUsersInnerPreferencesInnerWithDefaults instantiates a new CoreUserUpdateUsersRequestUsersInnerPreferencesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreUserUpdateUsersRequestUsersInnerPreferencesInnerWithDefaults() *CoreUserUpdateUsersRequestUsersInnerPreferencesInner {
	this := CoreUserUpdateUsersRequestUsersInnerPreferencesInner{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CoreUserUpdateUsersRequestUsersInnerPreferencesInner) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreUserUpdateUsersRequestUsersInnerPreferencesInner) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CoreUserUpdateUsersRequestUsersInnerPreferencesInner) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *CoreUserUpdateUsersRequestUsersInnerPreferencesInner) SetType(v string) {
	o.Type = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CoreUserUpdateUsersRequestUsersInnerPreferencesInner) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreUserUpdateUsersRequestUsersInnerPreferencesInner) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CoreUserUpdateUsersRequestUsersInnerPreferencesInner) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *CoreUserUpdateUsersRequestUsersInnerPreferencesInner) SetValue(v string) {
	o.Value = &v
}

func (o CoreUserUpdateUsersRequestUsersInnerPreferencesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreUserUpdateUsersRequestUsersInnerPreferencesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableCoreUserUpdateUsersRequestUsersInnerPreferencesInner struct {
	value *CoreUserUpdateUsersRequestUsersInnerPreferencesInner
	isSet bool
}

func (v NullableCoreUserUpdateUsersRequestUsersInnerPreferencesInner) Get() *CoreUserUpdateUsersRequestUsersInnerPreferencesInner {
	return v.value
}

func (v *NullableCoreUserUpdateUsersRequestUsersInnerPreferencesInner) Set(val *CoreUserUpdateUsersRequestUsersInnerPreferencesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreUserUpdateUsersRequestUsersInnerPreferencesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreUserUpdateUsersRequestUsersInnerPreferencesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreUserUpdateUsersRequestUsersInnerPreferencesInner(val *CoreUserUpdateUsersRequestUsersInnerPreferencesInner) *NullableCoreUserUpdateUsersRequestUsersInnerPreferencesInner {
	return &NullableCoreUserUpdateUsersRequestUsersInnerPreferencesInner{value: val, isSet: true}
}

func (v NullableCoreUserUpdateUsersRequestUsersInnerPreferencesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreUserUpdateUsersRequestUsersInnerPreferencesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


