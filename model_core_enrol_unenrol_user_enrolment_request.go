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

// checks if the CoreEnrolUnenrolUserEnrolmentRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreEnrolUnenrolUserEnrolmentRequest{}

// CoreEnrolUnenrolUserEnrolmentRequest struct for CoreEnrolUnenrolUserEnrolmentRequest
type CoreEnrolUnenrolUserEnrolmentRequest struct {
	// User enrolment ID
	Ueid int32 `json:"ueid"`
}

type _CoreEnrolUnenrolUserEnrolmentRequest CoreEnrolUnenrolUserEnrolmentRequest

// NewCoreEnrolUnenrolUserEnrolmentRequest instantiates a new CoreEnrolUnenrolUserEnrolmentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreEnrolUnenrolUserEnrolmentRequest(ueid int32) *CoreEnrolUnenrolUserEnrolmentRequest {
	this := CoreEnrolUnenrolUserEnrolmentRequest{}
	this.Ueid = ueid
	return &this
}

// NewCoreEnrolUnenrolUserEnrolmentRequestWithDefaults instantiates a new CoreEnrolUnenrolUserEnrolmentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreEnrolUnenrolUserEnrolmentRequestWithDefaults() *CoreEnrolUnenrolUserEnrolmentRequest {
	this := CoreEnrolUnenrolUserEnrolmentRequest{}
	var ueid int32 = null
	this.Ueid = ueid
	return &this
}

// GetUeid returns the Ueid field value
func (o *CoreEnrolUnenrolUserEnrolmentRequest) GetUeid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Ueid
}

// GetUeidOk returns a tuple with the Ueid field value
// and a boolean to check if the value has been set.
func (o *CoreEnrolUnenrolUserEnrolmentRequest) GetUeidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ueid, true
}

// SetUeid sets field value
func (o *CoreEnrolUnenrolUserEnrolmentRequest) SetUeid(v int32) {
	o.Ueid = v
}

func (o CoreEnrolUnenrolUserEnrolmentRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreEnrolUnenrolUserEnrolmentRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ueid"] = o.Ueid
	return toSerialize, nil
}

func (o *CoreEnrolUnenrolUserEnrolmentRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ueid",
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

	varCoreEnrolUnenrolUserEnrolmentRequest := _CoreEnrolUnenrolUserEnrolmentRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCoreEnrolUnenrolUserEnrolmentRequest)

	if err != nil {
		return err
	}

	*o = CoreEnrolUnenrolUserEnrolmentRequest(varCoreEnrolUnenrolUserEnrolmentRequest)

	return err
}

type NullableCoreEnrolUnenrolUserEnrolmentRequest struct {
	value *CoreEnrolUnenrolUserEnrolmentRequest
	isSet bool
}

func (v NullableCoreEnrolUnenrolUserEnrolmentRequest) Get() *CoreEnrolUnenrolUserEnrolmentRequest {
	return v.value
}

func (v *NullableCoreEnrolUnenrolUserEnrolmentRequest) Set(val *CoreEnrolUnenrolUserEnrolmentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreEnrolUnenrolUserEnrolmentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreEnrolUnenrolUserEnrolmentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreEnrolUnenrolUserEnrolmentRequest(val *CoreEnrolUnenrolUserEnrolmentRequest) *NullableCoreEnrolUnenrolUserEnrolmentRequest {
	return &NullableCoreEnrolUnenrolUserEnrolmentRequest{value: val, isSet: true}
}

func (v NullableCoreEnrolUnenrolUserEnrolmentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreEnrolUnenrolUserEnrolmentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


