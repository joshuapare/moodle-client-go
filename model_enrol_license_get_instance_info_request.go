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

// checks if the EnrolLicenseGetInstanceInfoRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EnrolLicenseGetInstanceInfoRequest{}

// EnrolLicenseGetInstanceInfoRequest struct for EnrolLicenseGetInstanceInfoRequest
type EnrolLicenseGetInstanceInfoRequest struct {
	// instance id of license enrolment plugin.
	Instanceid int32 `json:"instanceid"`
}

type _EnrolLicenseGetInstanceInfoRequest EnrolLicenseGetInstanceInfoRequest

// NewEnrolLicenseGetInstanceInfoRequest instantiates a new EnrolLicenseGetInstanceInfoRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnrolLicenseGetInstanceInfoRequest(instanceid int32) *EnrolLicenseGetInstanceInfoRequest {
	this := EnrolLicenseGetInstanceInfoRequest{}
	this.Instanceid = instanceid
	return &this
}

// NewEnrolLicenseGetInstanceInfoRequestWithDefaults instantiates a new EnrolLicenseGetInstanceInfoRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnrolLicenseGetInstanceInfoRequestWithDefaults() *EnrolLicenseGetInstanceInfoRequest {
	this := EnrolLicenseGetInstanceInfoRequest{}
	var instanceid int32 = null
	this.Instanceid = instanceid
	return &this
}

// GetInstanceid returns the Instanceid field value
func (o *EnrolLicenseGetInstanceInfoRequest) GetInstanceid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Instanceid
}

// GetInstanceidOk returns a tuple with the Instanceid field value
// and a boolean to check if the value has been set.
func (o *EnrolLicenseGetInstanceInfoRequest) GetInstanceidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Instanceid, true
}

// SetInstanceid sets field value
func (o *EnrolLicenseGetInstanceInfoRequest) SetInstanceid(v int32) {
	o.Instanceid = v
}

func (o EnrolLicenseGetInstanceInfoRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnrolLicenseGetInstanceInfoRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["instanceid"] = o.Instanceid
	return toSerialize, nil
}

func (o *EnrolLicenseGetInstanceInfoRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"instanceid",
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

	varEnrolLicenseGetInstanceInfoRequest := _EnrolLicenseGetInstanceInfoRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEnrolLicenseGetInstanceInfoRequest)

	if err != nil {
		return err
	}

	*o = EnrolLicenseGetInstanceInfoRequest(varEnrolLicenseGetInstanceInfoRequest)

	return err
}

type NullableEnrolLicenseGetInstanceInfoRequest struct {
	value *EnrolLicenseGetInstanceInfoRequest
	isSet bool
}

func (v NullableEnrolLicenseGetInstanceInfoRequest) Get() *EnrolLicenseGetInstanceInfoRequest {
	return v.value
}

func (v *NullableEnrolLicenseGetInstanceInfoRequest) Set(val *EnrolLicenseGetInstanceInfoRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEnrolLicenseGetInstanceInfoRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEnrolLicenseGetInstanceInfoRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnrolLicenseGetInstanceInfoRequest(val *EnrolLicenseGetInstanceInfoRequest) *NullableEnrolLicenseGetInstanceInfoRequest {
	return &NullableEnrolLicenseGetInstanceInfoRequest{value: val, isSet: true}
}

func (v NullableEnrolLicenseGetInstanceInfoRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnrolLicenseGetInstanceInfoRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

