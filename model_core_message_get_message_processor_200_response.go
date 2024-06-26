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

// checks if the CoreMessageGetMessageProcessor200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreMessageGetMessageProcessor200Response{}

// CoreMessageGetMessageProcessor200Response struct for CoreMessageGetMessageProcessor200Response
type CoreMessageGetMessageProcessor200Response struct {
	// Site configuration status
	Systemconfigured bool `json:"systemconfigured"`
	// The user configuration status
	Userconfigured bool `json:"userconfigured"`
}

type _CoreMessageGetMessageProcessor200Response CoreMessageGetMessageProcessor200Response

// NewCoreMessageGetMessageProcessor200Response instantiates a new CoreMessageGetMessageProcessor200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreMessageGetMessageProcessor200Response(systemconfigured bool, userconfigured bool) *CoreMessageGetMessageProcessor200Response {
	this := CoreMessageGetMessageProcessor200Response{}
	this.Systemconfigured = systemconfigured
	this.Userconfigured = userconfigured
	return &this
}

// NewCoreMessageGetMessageProcessor200ResponseWithDefaults instantiates a new CoreMessageGetMessageProcessor200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreMessageGetMessageProcessor200ResponseWithDefaults() *CoreMessageGetMessageProcessor200Response {
	this := CoreMessageGetMessageProcessor200Response{}
	var systemconfigured bool = null
	this.Systemconfigured = systemconfigured
	var userconfigured bool = null
	this.Userconfigured = userconfigured
	return &this
}

// GetSystemconfigured returns the Systemconfigured field value
func (o *CoreMessageGetMessageProcessor200Response) GetSystemconfigured() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Systemconfigured
}

// GetSystemconfiguredOk returns a tuple with the Systemconfigured field value
// and a boolean to check if the value has been set.
func (o *CoreMessageGetMessageProcessor200Response) GetSystemconfiguredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Systemconfigured, true
}

// SetSystemconfigured sets field value
func (o *CoreMessageGetMessageProcessor200Response) SetSystemconfigured(v bool) {
	o.Systemconfigured = v
}

// GetUserconfigured returns the Userconfigured field value
func (o *CoreMessageGetMessageProcessor200Response) GetUserconfigured() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Userconfigured
}

// GetUserconfiguredOk returns a tuple with the Userconfigured field value
// and a boolean to check if the value has been set.
func (o *CoreMessageGetMessageProcessor200Response) GetUserconfiguredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Userconfigured, true
}

// SetUserconfigured sets field value
func (o *CoreMessageGetMessageProcessor200Response) SetUserconfigured(v bool) {
	o.Userconfigured = v
}

func (o CoreMessageGetMessageProcessor200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreMessageGetMessageProcessor200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["systemconfigured"] = o.Systemconfigured
	toSerialize["userconfigured"] = o.Userconfigured
	return toSerialize, nil
}

func (o *CoreMessageGetMessageProcessor200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"systemconfigured",
		"userconfigured",
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

	varCoreMessageGetMessageProcessor200Response := _CoreMessageGetMessageProcessor200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCoreMessageGetMessageProcessor200Response)

	if err != nil {
		return err
	}

	*o = CoreMessageGetMessageProcessor200Response(varCoreMessageGetMessageProcessor200Response)

	return err
}

type NullableCoreMessageGetMessageProcessor200Response struct {
	value *CoreMessageGetMessageProcessor200Response
	isSet bool
}

func (v NullableCoreMessageGetMessageProcessor200Response) Get() *CoreMessageGetMessageProcessor200Response {
	return v.value
}

func (v *NullableCoreMessageGetMessageProcessor200Response) Set(val *CoreMessageGetMessageProcessor200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreMessageGetMessageProcessor200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreMessageGetMessageProcessor200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreMessageGetMessageProcessor200Response(val *CoreMessageGetMessageProcessor200Response) *NullableCoreMessageGetMessageProcessor200Response {
	return &NullableCoreMessageGetMessageProcessor200Response{value: val, isSet: true}
}

func (v NullableCoreMessageGetMessageProcessor200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreMessageGetMessageProcessor200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


