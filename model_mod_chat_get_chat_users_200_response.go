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

// checks if the ModChatGetChatUsers200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModChatGetChatUsers200Response{}

// ModChatGetChatUsers200Response struct for ModChatGetChatUsers200Response
type ModChatGetChatUsers200Response struct {
	Users []ModChatGetChatUsers200ResponseUsersInner `json:"users"`
	Warnings []AuthEmailSignupUser200ResponseWarningsInner `json:"warnings,omitempty"`
}

type _ModChatGetChatUsers200Response ModChatGetChatUsers200Response

// NewModChatGetChatUsers200Response instantiates a new ModChatGetChatUsers200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModChatGetChatUsers200Response(users []ModChatGetChatUsers200ResponseUsersInner) *ModChatGetChatUsers200Response {
	this := ModChatGetChatUsers200Response{}
	this.Users = users
	return &this
}

// NewModChatGetChatUsers200ResponseWithDefaults instantiates a new ModChatGetChatUsers200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModChatGetChatUsers200ResponseWithDefaults() *ModChatGetChatUsers200Response {
	this := ModChatGetChatUsers200Response{}
	return &this
}

// GetUsers returns the Users field value
func (o *ModChatGetChatUsers200Response) GetUsers() []ModChatGetChatUsers200ResponseUsersInner {
	if o == nil {
		var ret []ModChatGetChatUsers200ResponseUsersInner
		return ret
	}

	return o.Users
}

// GetUsersOk returns a tuple with the Users field value
// and a boolean to check if the value has been set.
func (o *ModChatGetChatUsers200Response) GetUsersOk() ([]ModChatGetChatUsers200ResponseUsersInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Users, true
}

// SetUsers sets field value
func (o *ModChatGetChatUsers200Response) SetUsers(v []ModChatGetChatUsers200ResponseUsersInner) {
	o.Users = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *ModChatGetChatUsers200Response) GetWarnings() []AuthEmailSignupUser200ResponseWarningsInner {
	if o == nil || IsNil(o.Warnings) {
		var ret []AuthEmailSignupUser200ResponseWarningsInner
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModChatGetChatUsers200Response) GetWarningsOk() ([]AuthEmailSignupUser200ResponseWarningsInner, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *ModChatGetChatUsers200Response) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []AuthEmailSignupUser200ResponseWarningsInner and assigns it to the Warnings field.
func (o *ModChatGetChatUsers200Response) SetWarnings(v []AuthEmailSignupUser200ResponseWarningsInner) {
	o.Warnings = v
}

func (o ModChatGetChatUsers200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModChatGetChatUsers200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["users"] = o.Users
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

func (o *ModChatGetChatUsers200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"users",
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

	varModChatGetChatUsers200Response := _ModChatGetChatUsers200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModChatGetChatUsers200Response)

	if err != nil {
		return err
	}

	*o = ModChatGetChatUsers200Response(varModChatGetChatUsers200Response)

	return err
}

type NullableModChatGetChatUsers200Response struct {
	value *ModChatGetChatUsers200Response
	isSet bool
}

func (v NullableModChatGetChatUsers200Response) Get() *ModChatGetChatUsers200Response {
	return v.value
}

func (v *NullableModChatGetChatUsers200Response) Set(val *ModChatGetChatUsers200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableModChatGetChatUsers200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableModChatGetChatUsers200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModChatGetChatUsers200Response(val *ModChatGetChatUsers200Response) *NullableModChatGetChatUsers200Response {
	return &NullableModChatGetChatUsers200Response{value: val, isSet: true}
}

func (v NullableModChatGetChatUsers200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModChatGetChatUsers200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


