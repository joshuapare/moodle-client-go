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

// checks if the CoreUserUpdateUsersRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreUserUpdateUsersRequest{}

// CoreUserUpdateUsersRequest struct for CoreUserUpdateUsersRequest
type CoreUserUpdateUsersRequest struct {
	Users []CoreUserUpdateUsersRequestUsersInner `json:"users"`
}

type _CoreUserUpdateUsersRequest CoreUserUpdateUsersRequest

// NewCoreUserUpdateUsersRequest instantiates a new CoreUserUpdateUsersRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreUserUpdateUsersRequest(users []CoreUserUpdateUsersRequestUsersInner) *CoreUserUpdateUsersRequest {
	this := CoreUserUpdateUsersRequest{}
	this.Users = users
	return &this
}

// NewCoreUserUpdateUsersRequestWithDefaults instantiates a new CoreUserUpdateUsersRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreUserUpdateUsersRequestWithDefaults() *CoreUserUpdateUsersRequest {
	this := CoreUserUpdateUsersRequest{}
	return &this
}

// GetUsers returns the Users field value
func (o *CoreUserUpdateUsersRequest) GetUsers() []CoreUserUpdateUsersRequestUsersInner {
	if o == nil {
		var ret []CoreUserUpdateUsersRequestUsersInner
		return ret
	}

	return o.Users
}

// GetUsersOk returns a tuple with the Users field value
// and a boolean to check if the value has been set.
func (o *CoreUserUpdateUsersRequest) GetUsersOk() ([]CoreUserUpdateUsersRequestUsersInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Users, true
}

// SetUsers sets field value
func (o *CoreUserUpdateUsersRequest) SetUsers(v []CoreUserUpdateUsersRequestUsersInner) {
	o.Users = v
}

func (o CoreUserUpdateUsersRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreUserUpdateUsersRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["users"] = o.Users
	return toSerialize, nil
}

func (o *CoreUserUpdateUsersRequest) UnmarshalJSON(data []byte) (err error) {
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

	varCoreUserUpdateUsersRequest := _CoreUserUpdateUsersRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCoreUserUpdateUsersRequest)

	if err != nil {
		return err
	}

	*o = CoreUserUpdateUsersRequest(varCoreUserUpdateUsersRequest)

	return err
}

type NullableCoreUserUpdateUsersRequest struct {
	value *CoreUserUpdateUsersRequest
	isSet bool
}

func (v NullableCoreUserUpdateUsersRequest) Get() *CoreUserUpdateUsersRequest {
	return v.value
}

func (v *NullableCoreUserUpdateUsersRequest) Set(val *CoreUserUpdateUsersRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreUserUpdateUsersRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreUserUpdateUsersRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreUserUpdateUsersRequest(val *CoreUserUpdateUsersRequest) *NullableCoreUserUpdateUsersRequest {
	return &NullableCoreUserUpdateUsersRequest{value: val, isSet: true}
}

func (v NullableCoreUserUpdateUsersRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreUserUpdateUsersRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


