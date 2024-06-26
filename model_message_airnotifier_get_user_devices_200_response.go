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

// checks if the MessageAirnotifierGetUserDevices200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MessageAirnotifierGetUserDevices200Response{}

// MessageAirnotifierGetUserDevices200Response struct for MessageAirnotifierGetUserDevices200Response
type MessageAirnotifierGetUserDevices200Response struct {
	Devices []MessageAirnotifierGetUserDevices200ResponseDevicesInner `json:"devices"`
	Warnings []AuthEmailSignupUser200ResponseWarningsInner `json:"warnings,omitempty"`
}

type _MessageAirnotifierGetUserDevices200Response MessageAirnotifierGetUserDevices200Response

// NewMessageAirnotifierGetUserDevices200Response instantiates a new MessageAirnotifierGetUserDevices200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessageAirnotifierGetUserDevices200Response(devices []MessageAirnotifierGetUserDevices200ResponseDevicesInner) *MessageAirnotifierGetUserDevices200Response {
	this := MessageAirnotifierGetUserDevices200Response{}
	this.Devices = devices
	return &this
}

// NewMessageAirnotifierGetUserDevices200ResponseWithDefaults instantiates a new MessageAirnotifierGetUserDevices200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessageAirnotifierGetUserDevices200ResponseWithDefaults() *MessageAirnotifierGetUserDevices200Response {
	this := MessageAirnotifierGetUserDevices200Response{}
	return &this
}

// GetDevices returns the Devices field value
func (o *MessageAirnotifierGetUserDevices200Response) GetDevices() []MessageAirnotifierGetUserDevices200ResponseDevicesInner {
	if o == nil {
		var ret []MessageAirnotifierGetUserDevices200ResponseDevicesInner
		return ret
	}

	return o.Devices
}

// GetDevicesOk returns a tuple with the Devices field value
// and a boolean to check if the value has been set.
func (o *MessageAirnotifierGetUserDevices200Response) GetDevicesOk() ([]MessageAirnotifierGetUserDevices200ResponseDevicesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Devices, true
}

// SetDevices sets field value
func (o *MessageAirnotifierGetUserDevices200Response) SetDevices(v []MessageAirnotifierGetUserDevices200ResponseDevicesInner) {
	o.Devices = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *MessageAirnotifierGetUserDevices200Response) GetWarnings() []AuthEmailSignupUser200ResponseWarningsInner {
	if o == nil || IsNil(o.Warnings) {
		var ret []AuthEmailSignupUser200ResponseWarningsInner
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageAirnotifierGetUserDevices200Response) GetWarningsOk() ([]AuthEmailSignupUser200ResponseWarningsInner, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *MessageAirnotifierGetUserDevices200Response) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []AuthEmailSignupUser200ResponseWarningsInner and assigns it to the Warnings field.
func (o *MessageAirnotifierGetUserDevices200Response) SetWarnings(v []AuthEmailSignupUser200ResponseWarningsInner) {
	o.Warnings = v
}

func (o MessageAirnotifierGetUserDevices200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MessageAirnotifierGetUserDevices200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["devices"] = o.Devices
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

func (o *MessageAirnotifierGetUserDevices200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"devices",
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

	varMessageAirnotifierGetUserDevices200Response := _MessageAirnotifierGetUserDevices200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMessageAirnotifierGetUserDevices200Response)

	if err != nil {
		return err
	}

	*o = MessageAirnotifierGetUserDevices200Response(varMessageAirnotifierGetUserDevices200Response)

	return err
}

type NullableMessageAirnotifierGetUserDevices200Response struct {
	value *MessageAirnotifierGetUserDevices200Response
	isSet bool
}

func (v NullableMessageAirnotifierGetUserDevices200Response) Get() *MessageAirnotifierGetUserDevices200Response {
	return v.value
}

func (v *NullableMessageAirnotifierGetUserDevices200Response) Set(val *MessageAirnotifierGetUserDevices200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageAirnotifierGetUserDevices200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageAirnotifierGetUserDevices200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageAirnotifierGetUserDevices200Response(val *MessageAirnotifierGetUserDevices200Response) *NullableMessageAirnotifierGetUserDevices200Response {
	return &NullableMessageAirnotifierGetUserDevices200Response{value: val, isSet: true}
}

func (v NullableMessageAirnotifierGetUserDevices200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageAirnotifierGetUserDevices200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


