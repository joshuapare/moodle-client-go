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

// checks if the EnrolManualUnenrolUsersRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EnrolManualUnenrolUsersRequest{}

// EnrolManualUnenrolUsersRequest struct for EnrolManualUnenrolUsersRequest
type EnrolManualUnenrolUsersRequest struct {
	Enrolments []EnrolManualUnenrolUsersRequestEnrolmentsInner `json:"enrolments"`
}

type _EnrolManualUnenrolUsersRequest EnrolManualUnenrolUsersRequest

// NewEnrolManualUnenrolUsersRequest instantiates a new EnrolManualUnenrolUsersRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnrolManualUnenrolUsersRequest(enrolments []EnrolManualUnenrolUsersRequestEnrolmentsInner) *EnrolManualUnenrolUsersRequest {
	this := EnrolManualUnenrolUsersRequest{}
	this.Enrolments = enrolments
	return &this
}

// NewEnrolManualUnenrolUsersRequestWithDefaults instantiates a new EnrolManualUnenrolUsersRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnrolManualUnenrolUsersRequestWithDefaults() *EnrolManualUnenrolUsersRequest {
	this := EnrolManualUnenrolUsersRequest{}
	return &this
}

// GetEnrolments returns the Enrolments field value
func (o *EnrolManualUnenrolUsersRequest) GetEnrolments() []EnrolManualUnenrolUsersRequestEnrolmentsInner {
	if o == nil {
		var ret []EnrolManualUnenrolUsersRequestEnrolmentsInner
		return ret
	}

	return o.Enrolments
}

// GetEnrolmentsOk returns a tuple with the Enrolments field value
// and a boolean to check if the value has been set.
func (o *EnrolManualUnenrolUsersRequest) GetEnrolmentsOk() ([]EnrolManualUnenrolUsersRequestEnrolmentsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Enrolments, true
}

// SetEnrolments sets field value
func (o *EnrolManualUnenrolUsersRequest) SetEnrolments(v []EnrolManualUnenrolUsersRequestEnrolmentsInner) {
	o.Enrolments = v
}

func (o EnrolManualUnenrolUsersRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnrolManualUnenrolUsersRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["enrolments"] = o.Enrolments
	return toSerialize, nil
}

func (o *EnrolManualUnenrolUsersRequest) UnmarshalJSON(data []byte) (err error) {
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

	varEnrolManualUnenrolUsersRequest := _EnrolManualUnenrolUsersRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEnrolManualUnenrolUsersRequest)

	if err != nil {
		return err
	}

	*o = EnrolManualUnenrolUsersRequest(varEnrolManualUnenrolUsersRequest)

	return err
}

type NullableEnrolManualUnenrolUsersRequest struct {
	value *EnrolManualUnenrolUsersRequest
	isSet bool
}

func (v NullableEnrolManualUnenrolUsersRequest) Get() *EnrolManualUnenrolUsersRequest {
	return v.value
}

func (v *NullableEnrolManualUnenrolUsersRequest) Set(val *EnrolManualUnenrolUsersRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEnrolManualUnenrolUsersRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEnrolManualUnenrolUsersRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnrolManualUnenrolUsersRequest(val *EnrolManualUnenrolUsersRequest) *NullableEnrolManualUnenrolUsersRequest {
	return &NullableEnrolManualUnenrolUsersRequest{value: val, isSet: true}
}

func (v NullableEnrolManualUnenrolUsersRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnrolManualUnenrolUsersRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


