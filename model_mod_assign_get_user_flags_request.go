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

// checks if the ModAssignGetUserFlagsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModAssignGetUserFlagsRequest{}

// ModAssignGetUserFlagsRequest struct for ModAssignGetUserFlagsRequest
type ModAssignGetUserFlagsRequest struct {
	Assignmentids []map[string]interface{} `json:"assignmentids"`
}

type _ModAssignGetUserFlagsRequest ModAssignGetUserFlagsRequest

// NewModAssignGetUserFlagsRequest instantiates a new ModAssignGetUserFlagsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModAssignGetUserFlagsRequest(assignmentids []map[string]interface{}) *ModAssignGetUserFlagsRequest {
	this := ModAssignGetUserFlagsRequest{}
	this.Assignmentids = assignmentids
	return &this
}

// NewModAssignGetUserFlagsRequestWithDefaults instantiates a new ModAssignGetUserFlagsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModAssignGetUserFlagsRequestWithDefaults() *ModAssignGetUserFlagsRequest {
	this := ModAssignGetUserFlagsRequest{}
	return &this
}

// GetAssignmentids returns the Assignmentids field value
func (o *ModAssignGetUserFlagsRequest) GetAssignmentids() []map[string]interface{} {
	if o == nil {
		var ret []map[string]interface{}
		return ret
	}

	return o.Assignmentids
}

// GetAssignmentidsOk returns a tuple with the Assignmentids field value
// and a boolean to check if the value has been set.
func (o *ModAssignGetUserFlagsRequest) GetAssignmentidsOk() ([]map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Assignmentids, true
}

// SetAssignmentids sets field value
func (o *ModAssignGetUserFlagsRequest) SetAssignmentids(v []map[string]interface{}) {
	o.Assignmentids = v
}

func (o ModAssignGetUserFlagsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModAssignGetUserFlagsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["assignmentids"] = o.Assignmentids
	return toSerialize, nil
}

func (o *ModAssignGetUserFlagsRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"assignmentids",
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

	varModAssignGetUserFlagsRequest := _ModAssignGetUserFlagsRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModAssignGetUserFlagsRequest)

	if err != nil {
		return err
	}

	*o = ModAssignGetUserFlagsRequest(varModAssignGetUserFlagsRequest)

	return err
}

type NullableModAssignGetUserFlagsRequest struct {
	value *ModAssignGetUserFlagsRequest
	isSet bool
}

func (v NullableModAssignGetUserFlagsRequest) Get() *ModAssignGetUserFlagsRequest {
	return v.value
}

func (v *NullableModAssignGetUserFlagsRequest) Set(val *ModAssignGetUserFlagsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModAssignGetUserFlagsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModAssignGetUserFlagsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModAssignGetUserFlagsRequest(val *ModAssignGetUserFlagsRequest) *NullableModAssignGetUserFlagsRequest {
	return &NullableModAssignGetUserFlagsRequest{value: val, isSet: true}
}

func (v NullableModAssignGetUserFlagsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModAssignGetUserFlagsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

