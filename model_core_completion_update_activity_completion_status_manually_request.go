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

// checks if the CoreCompletionUpdateActivityCompletionStatusManuallyRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreCompletionUpdateActivityCompletionStatusManuallyRequest{}

// CoreCompletionUpdateActivityCompletionStatusManuallyRequest struct for CoreCompletionUpdateActivityCompletionStatusManuallyRequest
type CoreCompletionUpdateActivityCompletionStatusManuallyRequest struct {
	// course module id
	Cmid int32 `json:"cmid"`
	// activity completed or not
	Completed bool `json:"completed"`
}

type _CoreCompletionUpdateActivityCompletionStatusManuallyRequest CoreCompletionUpdateActivityCompletionStatusManuallyRequest

// NewCoreCompletionUpdateActivityCompletionStatusManuallyRequest instantiates a new CoreCompletionUpdateActivityCompletionStatusManuallyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreCompletionUpdateActivityCompletionStatusManuallyRequest(cmid int32, completed bool) *CoreCompletionUpdateActivityCompletionStatusManuallyRequest {
	this := CoreCompletionUpdateActivityCompletionStatusManuallyRequest{}
	this.Cmid = cmid
	this.Completed = completed
	return &this
}

// NewCoreCompletionUpdateActivityCompletionStatusManuallyRequestWithDefaults instantiates a new CoreCompletionUpdateActivityCompletionStatusManuallyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreCompletionUpdateActivityCompletionStatusManuallyRequestWithDefaults() *CoreCompletionUpdateActivityCompletionStatusManuallyRequest {
	this := CoreCompletionUpdateActivityCompletionStatusManuallyRequest{}
	var completed bool = null
	this.Completed = completed
	return &this
}

// GetCmid returns the Cmid field value
func (o *CoreCompletionUpdateActivityCompletionStatusManuallyRequest) GetCmid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Cmid
}

// GetCmidOk returns a tuple with the Cmid field value
// and a boolean to check if the value has been set.
func (o *CoreCompletionUpdateActivityCompletionStatusManuallyRequest) GetCmidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cmid, true
}

// SetCmid sets field value
func (o *CoreCompletionUpdateActivityCompletionStatusManuallyRequest) SetCmid(v int32) {
	o.Cmid = v
}

// GetCompleted returns the Completed field value
func (o *CoreCompletionUpdateActivityCompletionStatusManuallyRequest) GetCompleted() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Completed
}

// GetCompletedOk returns a tuple with the Completed field value
// and a boolean to check if the value has been set.
func (o *CoreCompletionUpdateActivityCompletionStatusManuallyRequest) GetCompletedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Completed, true
}

// SetCompleted sets field value
func (o *CoreCompletionUpdateActivityCompletionStatusManuallyRequest) SetCompleted(v bool) {
	o.Completed = v
}

func (o CoreCompletionUpdateActivityCompletionStatusManuallyRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreCompletionUpdateActivityCompletionStatusManuallyRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cmid"] = o.Cmid
	toSerialize["completed"] = o.Completed
	return toSerialize, nil
}

func (o *CoreCompletionUpdateActivityCompletionStatusManuallyRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"cmid",
		"completed",
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

	varCoreCompletionUpdateActivityCompletionStatusManuallyRequest := _CoreCompletionUpdateActivityCompletionStatusManuallyRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCoreCompletionUpdateActivityCompletionStatusManuallyRequest)

	if err != nil {
		return err
	}

	*o = CoreCompletionUpdateActivityCompletionStatusManuallyRequest(varCoreCompletionUpdateActivityCompletionStatusManuallyRequest)

	return err
}

type NullableCoreCompletionUpdateActivityCompletionStatusManuallyRequest struct {
	value *CoreCompletionUpdateActivityCompletionStatusManuallyRequest
	isSet bool
}

func (v NullableCoreCompletionUpdateActivityCompletionStatusManuallyRequest) Get() *CoreCompletionUpdateActivityCompletionStatusManuallyRequest {
	return v.value
}

func (v *NullableCoreCompletionUpdateActivityCompletionStatusManuallyRequest) Set(val *CoreCompletionUpdateActivityCompletionStatusManuallyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreCompletionUpdateActivityCompletionStatusManuallyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreCompletionUpdateActivityCompletionStatusManuallyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreCompletionUpdateActivityCompletionStatusManuallyRequest(val *CoreCompletionUpdateActivityCompletionStatusManuallyRequest) *NullableCoreCompletionUpdateActivityCompletionStatusManuallyRequest {
	return &NullableCoreCompletionUpdateActivityCompletionStatusManuallyRequest{value: val, isSet: true}
}

func (v NullableCoreCompletionUpdateActivityCompletionStatusManuallyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreCompletionUpdateActivityCompletionStatusManuallyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

