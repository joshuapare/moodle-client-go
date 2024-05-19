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

// checks if the BlockIomadCompanyAdminEnrolUsersRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BlockIomadCompanyAdminEnrolUsersRequest{}

// BlockIomadCompanyAdminEnrolUsersRequest struct for BlockIomadCompanyAdminEnrolUsersRequest
type BlockIomadCompanyAdminEnrolUsersRequest struct {
	Enrolments []BlockIomadCompanyAdminEnrolUsersRequestEnrolmentsInner `json:"enrolments"`
}

type _BlockIomadCompanyAdminEnrolUsersRequest BlockIomadCompanyAdminEnrolUsersRequest

// NewBlockIomadCompanyAdminEnrolUsersRequest instantiates a new BlockIomadCompanyAdminEnrolUsersRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlockIomadCompanyAdminEnrolUsersRequest(enrolments []BlockIomadCompanyAdminEnrolUsersRequestEnrolmentsInner) *BlockIomadCompanyAdminEnrolUsersRequest {
	this := BlockIomadCompanyAdminEnrolUsersRequest{}
	this.Enrolments = enrolments
	return &this
}

// NewBlockIomadCompanyAdminEnrolUsersRequestWithDefaults instantiates a new BlockIomadCompanyAdminEnrolUsersRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlockIomadCompanyAdminEnrolUsersRequestWithDefaults() *BlockIomadCompanyAdminEnrolUsersRequest {
	this := BlockIomadCompanyAdminEnrolUsersRequest{}
	return &this
}

// GetEnrolments returns the Enrolments field value
func (o *BlockIomadCompanyAdminEnrolUsersRequest) GetEnrolments() []BlockIomadCompanyAdminEnrolUsersRequestEnrolmentsInner {
	if o == nil {
		var ret []BlockIomadCompanyAdminEnrolUsersRequestEnrolmentsInner
		return ret
	}

	return o.Enrolments
}

// GetEnrolmentsOk returns a tuple with the Enrolments field value
// and a boolean to check if the value has been set.
func (o *BlockIomadCompanyAdminEnrolUsersRequest) GetEnrolmentsOk() ([]BlockIomadCompanyAdminEnrolUsersRequestEnrolmentsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Enrolments, true
}

// SetEnrolments sets field value
func (o *BlockIomadCompanyAdminEnrolUsersRequest) SetEnrolments(v []BlockIomadCompanyAdminEnrolUsersRequestEnrolmentsInner) {
	o.Enrolments = v
}

func (o BlockIomadCompanyAdminEnrolUsersRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BlockIomadCompanyAdminEnrolUsersRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["enrolments"] = o.Enrolments
	return toSerialize, nil
}

func (o *BlockIomadCompanyAdminEnrolUsersRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"enrolments",
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

	varBlockIomadCompanyAdminEnrolUsersRequest := _BlockIomadCompanyAdminEnrolUsersRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBlockIomadCompanyAdminEnrolUsersRequest)

	if err != nil {
		return err
	}

	*o = BlockIomadCompanyAdminEnrolUsersRequest(varBlockIomadCompanyAdminEnrolUsersRequest)

	return err
}

type NullableBlockIomadCompanyAdminEnrolUsersRequest struct {
	value *BlockIomadCompanyAdminEnrolUsersRequest
	isSet bool
}

func (v NullableBlockIomadCompanyAdminEnrolUsersRequest) Get() *BlockIomadCompanyAdminEnrolUsersRequest {
	return v.value
}

func (v *NullableBlockIomadCompanyAdminEnrolUsersRequest) Set(val *BlockIomadCompanyAdminEnrolUsersRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBlockIomadCompanyAdminEnrolUsersRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBlockIomadCompanyAdminEnrolUsersRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlockIomadCompanyAdminEnrolUsersRequest(val *BlockIomadCompanyAdminEnrolUsersRequest) *NullableBlockIomadCompanyAdminEnrolUsersRequest {
	return &NullableBlockIomadCompanyAdminEnrolUsersRequest{value: val, isSet: true}
}

func (v NullableBlockIomadCompanyAdminEnrolUsersRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlockIomadCompanyAdminEnrolUsersRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


