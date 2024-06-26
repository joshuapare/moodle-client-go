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

// checks if the BlockIomadCompanyAdminMoveUsersRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BlockIomadCompanyAdminMoveUsersRequest{}

// BlockIomadCompanyAdminMoveUsersRequest struct for BlockIomadCompanyAdminMoveUsersRequest
type BlockIomadCompanyAdminMoveUsersRequest struct {
	Users []BlockIomadCompanyAdminMoveUsersRequestUsersInner `json:"users"`
}

type _BlockIomadCompanyAdminMoveUsersRequest BlockIomadCompanyAdminMoveUsersRequest

// NewBlockIomadCompanyAdminMoveUsersRequest instantiates a new BlockIomadCompanyAdminMoveUsersRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlockIomadCompanyAdminMoveUsersRequest(users []BlockIomadCompanyAdminMoveUsersRequestUsersInner) *BlockIomadCompanyAdminMoveUsersRequest {
	this := BlockIomadCompanyAdminMoveUsersRequest{}
	this.Users = users
	return &this
}

// NewBlockIomadCompanyAdminMoveUsersRequestWithDefaults instantiates a new BlockIomadCompanyAdminMoveUsersRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlockIomadCompanyAdminMoveUsersRequestWithDefaults() *BlockIomadCompanyAdminMoveUsersRequest {
	this := BlockIomadCompanyAdminMoveUsersRequest{}
	return &this
}

// GetUsers returns the Users field value
func (o *BlockIomadCompanyAdminMoveUsersRequest) GetUsers() []BlockIomadCompanyAdminMoveUsersRequestUsersInner {
	if o == nil {
		var ret []BlockIomadCompanyAdminMoveUsersRequestUsersInner
		return ret
	}

	return o.Users
}

// GetUsersOk returns a tuple with the Users field value
// and a boolean to check if the value has been set.
func (o *BlockIomadCompanyAdminMoveUsersRequest) GetUsersOk() ([]BlockIomadCompanyAdminMoveUsersRequestUsersInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Users, true
}

// SetUsers sets field value
func (o *BlockIomadCompanyAdminMoveUsersRequest) SetUsers(v []BlockIomadCompanyAdminMoveUsersRequestUsersInner) {
	o.Users = v
}

func (o BlockIomadCompanyAdminMoveUsersRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BlockIomadCompanyAdminMoveUsersRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["users"] = o.Users
	return toSerialize, nil
}

func (o *BlockIomadCompanyAdminMoveUsersRequest) UnmarshalJSON(data []byte) (err error) {
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

	varBlockIomadCompanyAdminMoveUsersRequest := _BlockIomadCompanyAdminMoveUsersRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBlockIomadCompanyAdminMoveUsersRequest)

	if err != nil {
		return err
	}

	*o = BlockIomadCompanyAdminMoveUsersRequest(varBlockIomadCompanyAdminMoveUsersRequest)

	return err
}

type NullableBlockIomadCompanyAdminMoveUsersRequest struct {
	value *BlockIomadCompanyAdminMoveUsersRequest
	isSet bool
}

func (v NullableBlockIomadCompanyAdminMoveUsersRequest) Get() *BlockIomadCompanyAdminMoveUsersRequest {
	return v.value
}

func (v *NullableBlockIomadCompanyAdminMoveUsersRequest) Set(val *BlockIomadCompanyAdminMoveUsersRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBlockIomadCompanyAdminMoveUsersRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBlockIomadCompanyAdminMoveUsersRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlockIomadCompanyAdminMoveUsersRequest(val *BlockIomadCompanyAdminMoveUsersRequest) *NullableBlockIomadCompanyAdminMoveUsersRequest {
	return &NullableBlockIomadCompanyAdminMoveUsersRequest{value: val, isSet: true}
}

func (v NullableBlockIomadCompanyAdminMoveUsersRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlockIomadCompanyAdminMoveUsersRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


