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

// checks if the CoreCohortAddCohortMembersRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreCohortAddCohortMembersRequest{}

// CoreCohortAddCohortMembersRequest struct for CoreCohortAddCohortMembersRequest
type CoreCohortAddCohortMembersRequest struct {
	Members []CoreCohortAddCohortMembersRequestMembersInner `json:"members"`
}

type _CoreCohortAddCohortMembersRequest CoreCohortAddCohortMembersRequest

// NewCoreCohortAddCohortMembersRequest instantiates a new CoreCohortAddCohortMembersRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreCohortAddCohortMembersRequest(members []CoreCohortAddCohortMembersRequestMembersInner) *CoreCohortAddCohortMembersRequest {
	this := CoreCohortAddCohortMembersRequest{}
	this.Members = members
	return &this
}

// NewCoreCohortAddCohortMembersRequestWithDefaults instantiates a new CoreCohortAddCohortMembersRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreCohortAddCohortMembersRequestWithDefaults() *CoreCohortAddCohortMembersRequest {
	this := CoreCohortAddCohortMembersRequest{}
	return &this
}

// GetMembers returns the Members field value
func (o *CoreCohortAddCohortMembersRequest) GetMembers() []CoreCohortAddCohortMembersRequestMembersInner {
	if o == nil {
		var ret []CoreCohortAddCohortMembersRequestMembersInner
		return ret
	}

	return o.Members
}

// GetMembersOk returns a tuple with the Members field value
// and a boolean to check if the value has been set.
func (o *CoreCohortAddCohortMembersRequest) GetMembersOk() ([]CoreCohortAddCohortMembersRequestMembersInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Members, true
}

// SetMembers sets field value
func (o *CoreCohortAddCohortMembersRequest) SetMembers(v []CoreCohortAddCohortMembersRequestMembersInner) {
	o.Members = v
}

func (o CoreCohortAddCohortMembersRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreCohortAddCohortMembersRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["members"] = o.Members
	return toSerialize, nil
}

func (o *CoreCohortAddCohortMembersRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"members",
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

	varCoreCohortAddCohortMembersRequest := _CoreCohortAddCohortMembersRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCoreCohortAddCohortMembersRequest)

	if err != nil {
		return err
	}

	*o = CoreCohortAddCohortMembersRequest(varCoreCohortAddCohortMembersRequest)

	return err
}

type NullableCoreCohortAddCohortMembersRequest struct {
	value *CoreCohortAddCohortMembersRequest
	isSet bool
}

func (v NullableCoreCohortAddCohortMembersRequest) Get() *CoreCohortAddCohortMembersRequest {
	return v.value
}

func (v *NullableCoreCohortAddCohortMembersRequest) Set(val *CoreCohortAddCohortMembersRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreCohortAddCohortMembersRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreCohortAddCohortMembersRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreCohortAddCohortMembersRequest(val *CoreCohortAddCohortMembersRequest) *NullableCoreCohortAddCohortMembersRequest {
	return &NullableCoreCohortAddCohortMembersRequest{value: val, isSet: true}
}

func (v NullableCoreCohortAddCohortMembersRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreCohortAddCohortMembersRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


