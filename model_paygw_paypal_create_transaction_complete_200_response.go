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

// checks if the PaygwPaypalCreateTransactionComplete200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaygwPaypalCreateTransactionComplete200Response{}

// PaygwPaypalCreateTransactionComplete200Response struct for PaygwPaypalCreateTransactionComplete200Response
type PaygwPaypalCreateTransactionComplete200Response struct {
	// Message (usually the error message).
	Message string `json:"message"`
	// Whether everything was successful or not.
	Success bool `json:"success"`
}

type _PaygwPaypalCreateTransactionComplete200Response PaygwPaypalCreateTransactionComplete200Response

// NewPaygwPaypalCreateTransactionComplete200Response instantiates a new PaygwPaypalCreateTransactionComplete200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaygwPaypalCreateTransactionComplete200Response(message string, success bool) *PaygwPaypalCreateTransactionComplete200Response {
	this := PaygwPaypalCreateTransactionComplete200Response{}
	this.Message = message
	this.Success = success
	return &this
}

// NewPaygwPaypalCreateTransactionComplete200ResponseWithDefaults instantiates a new PaygwPaypalCreateTransactionComplete200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaygwPaypalCreateTransactionComplete200ResponseWithDefaults() *PaygwPaypalCreateTransactionComplete200Response {
	this := PaygwPaypalCreateTransactionComplete200Response{}
	var message string = "null"
	this.Message = message
	var success bool = null
	this.Success = success
	return &this
}

// GetMessage returns the Message field value
func (o *PaygwPaypalCreateTransactionComplete200Response) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *PaygwPaypalCreateTransactionComplete200Response) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *PaygwPaypalCreateTransactionComplete200Response) SetMessage(v string) {
	o.Message = v
}

// GetSuccess returns the Success field value
func (o *PaygwPaypalCreateTransactionComplete200Response) GetSuccess() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value
// and a boolean to check if the value has been set.
func (o *PaygwPaypalCreateTransactionComplete200Response) GetSuccessOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Success, true
}

// SetSuccess sets field value
func (o *PaygwPaypalCreateTransactionComplete200Response) SetSuccess(v bool) {
	o.Success = v
}

func (o PaygwPaypalCreateTransactionComplete200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaygwPaypalCreateTransactionComplete200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["message"] = o.Message
	toSerialize["success"] = o.Success
	return toSerialize, nil
}

func (o *PaygwPaypalCreateTransactionComplete200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"message",
		"success",
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

	varPaygwPaypalCreateTransactionComplete200Response := _PaygwPaypalCreateTransactionComplete200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPaygwPaypalCreateTransactionComplete200Response)

	if err != nil {
		return err
	}

	*o = PaygwPaypalCreateTransactionComplete200Response(varPaygwPaypalCreateTransactionComplete200Response)

	return err
}

type NullablePaygwPaypalCreateTransactionComplete200Response struct {
	value *PaygwPaypalCreateTransactionComplete200Response
	isSet bool
}

func (v NullablePaygwPaypalCreateTransactionComplete200Response) Get() *PaygwPaypalCreateTransactionComplete200Response {
	return v.value
}

func (v *NullablePaygwPaypalCreateTransactionComplete200Response) Set(val *PaygwPaypalCreateTransactionComplete200Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePaygwPaypalCreateTransactionComplete200Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePaygwPaypalCreateTransactionComplete200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaygwPaypalCreateTransactionComplete200Response(val *PaygwPaypalCreateTransactionComplete200Response) *NullablePaygwPaypalCreateTransactionComplete200Response {
	return &NullablePaygwPaypalCreateTransactionComplete200Response{value: val, isSet: true}
}

func (v NullablePaygwPaypalCreateTransactionComplete200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaygwPaypalCreateTransactionComplete200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


