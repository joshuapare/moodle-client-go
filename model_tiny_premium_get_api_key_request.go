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

// checks if the TinyPremiumGetApiKeyRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TinyPremiumGetApiKeyRequest{}

// TinyPremiumGetApiKeyRequest struct for TinyPremiumGetApiKeyRequest
type TinyPremiumGetApiKeyRequest struct {
	// The current context ID.
	Contextid int32 `json:"contextid"`
}

type _TinyPremiumGetApiKeyRequest TinyPremiumGetApiKeyRequest

// NewTinyPremiumGetApiKeyRequest instantiates a new TinyPremiumGetApiKeyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTinyPremiumGetApiKeyRequest(contextid int32) *TinyPremiumGetApiKeyRequest {
	this := TinyPremiumGetApiKeyRequest{}
	this.Contextid = contextid
	return &this
}

// NewTinyPremiumGetApiKeyRequestWithDefaults instantiates a new TinyPremiumGetApiKeyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTinyPremiumGetApiKeyRequestWithDefaults() *TinyPremiumGetApiKeyRequest {
	this := TinyPremiumGetApiKeyRequest{}
	var contextid int32 = null
	this.Contextid = contextid
	return &this
}

// GetContextid returns the Contextid field value
func (o *TinyPremiumGetApiKeyRequest) GetContextid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Contextid
}

// GetContextidOk returns a tuple with the Contextid field value
// and a boolean to check if the value has been set.
func (o *TinyPremiumGetApiKeyRequest) GetContextidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Contextid, true
}

// SetContextid sets field value
func (o *TinyPremiumGetApiKeyRequest) SetContextid(v int32) {
	o.Contextid = v
}

func (o TinyPremiumGetApiKeyRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TinyPremiumGetApiKeyRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["contextid"] = o.Contextid
	return toSerialize, nil
}

func (o *TinyPremiumGetApiKeyRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"contextid",
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

	varTinyPremiumGetApiKeyRequest := _TinyPremiumGetApiKeyRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTinyPremiumGetApiKeyRequest)

	if err != nil {
		return err
	}

	*o = TinyPremiumGetApiKeyRequest(varTinyPremiumGetApiKeyRequest)

	return err
}

type NullableTinyPremiumGetApiKeyRequest struct {
	value *TinyPremiumGetApiKeyRequest
	isSet bool
}

func (v NullableTinyPremiumGetApiKeyRequest) Get() *TinyPremiumGetApiKeyRequest {
	return v.value
}

func (v *NullableTinyPremiumGetApiKeyRequest) Set(val *TinyPremiumGetApiKeyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTinyPremiumGetApiKeyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTinyPremiumGetApiKeyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTinyPremiumGetApiKeyRequest(val *TinyPremiumGetApiKeyRequest) *NullableTinyPremiumGetApiKeyRequest {
	return &NullableTinyPremiumGetApiKeyRequest{value: val, isSet: true}
}

func (v NullableTinyPremiumGetApiKeyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTinyPremiumGetApiKeyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


