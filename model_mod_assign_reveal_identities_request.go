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

// checks if the ModAssignRevealIdentitiesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModAssignRevealIdentitiesRequest{}

// ModAssignRevealIdentitiesRequest struct for ModAssignRevealIdentitiesRequest
type ModAssignRevealIdentitiesRequest struct {
	// The assignment id to operate on
	Assignmentid int32 `json:"assignmentid"`
}

type _ModAssignRevealIdentitiesRequest ModAssignRevealIdentitiesRequest

// NewModAssignRevealIdentitiesRequest instantiates a new ModAssignRevealIdentitiesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModAssignRevealIdentitiesRequest(assignmentid int32) *ModAssignRevealIdentitiesRequest {
	this := ModAssignRevealIdentitiesRequest{}
	this.Assignmentid = assignmentid
	return &this
}

// NewModAssignRevealIdentitiesRequestWithDefaults instantiates a new ModAssignRevealIdentitiesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModAssignRevealIdentitiesRequestWithDefaults() *ModAssignRevealIdentitiesRequest {
	this := ModAssignRevealIdentitiesRequest{}
	return &this
}

// GetAssignmentid returns the Assignmentid field value
func (o *ModAssignRevealIdentitiesRequest) GetAssignmentid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Assignmentid
}

// GetAssignmentidOk returns a tuple with the Assignmentid field value
// and a boolean to check if the value has been set.
func (o *ModAssignRevealIdentitiesRequest) GetAssignmentidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Assignmentid, true
}

// SetAssignmentid sets field value
func (o *ModAssignRevealIdentitiesRequest) SetAssignmentid(v int32) {
	o.Assignmentid = v
}

func (o ModAssignRevealIdentitiesRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModAssignRevealIdentitiesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["assignmentid"] = o.Assignmentid
	return toSerialize, nil
}

func (o *ModAssignRevealIdentitiesRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"assignmentid",
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

	varModAssignRevealIdentitiesRequest := _ModAssignRevealIdentitiesRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModAssignRevealIdentitiesRequest)

	if err != nil {
		return err
	}

	*o = ModAssignRevealIdentitiesRequest(varModAssignRevealIdentitiesRequest)

	return err
}

type NullableModAssignRevealIdentitiesRequest struct {
	value *ModAssignRevealIdentitiesRequest
	isSet bool
}

func (v NullableModAssignRevealIdentitiesRequest) Get() *ModAssignRevealIdentitiesRequest {
	return v.value
}

func (v *NullableModAssignRevealIdentitiesRequest) Set(val *ModAssignRevealIdentitiesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModAssignRevealIdentitiesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModAssignRevealIdentitiesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModAssignRevealIdentitiesRequest(val *ModAssignRevealIdentitiesRequest) *NullableModAssignRevealIdentitiesRequest {
	return &NullableModAssignRevealIdentitiesRequest{value: val, isSet: true}
}

func (v NullableModAssignRevealIdentitiesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModAssignRevealIdentitiesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


