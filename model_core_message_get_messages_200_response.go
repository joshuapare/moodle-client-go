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

// checks if the CoreMessageGetMessages200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreMessageGetMessages200Response{}

// CoreMessageGetMessages200Response struct for CoreMessageGetMessages200Response
type CoreMessageGetMessages200Response struct {
	Messages []CoreMessageGetMessages200ResponseMessagesInner `json:"messages"`
	Warnings []AuthEmailSignupUser200ResponseWarningsInner `json:"warnings,omitempty"`
}

type _CoreMessageGetMessages200Response CoreMessageGetMessages200Response

// NewCoreMessageGetMessages200Response instantiates a new CoreMessageGetMessages200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreMessageGetMessages200Response(messages []CoreMessageGetMessages200ResponseMessagesInner) *CoreMessageGetMessages200Response {
	this := CoreMessageGetMessages200Response{}
	this.Messages = messages
	return &this
}

// NewCoreMessageGetMessages200ResponseWithDefaults instantiates a new CoreMessageGetMessages200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreMessageGetMessages200ResponseWithDefaults() *CoreMessageGetMessages200Response {
	this := CoreMessageGetMessages200Response{}
	return &this
}

// GetMessages returns the Messages field value
func (o *CoreMessageGetMessages200Response) GetMessages() []CoreMessageGetMessages200ResponseMessagesInner {
	if o == nil {
		var ret []CoreMessageGetMessages200ResponseMessagesInner
		return ret
	}

	return o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value
// and a boolean to check if the value has been set.
func (o *CoreMessageGetMessages200Response) GetMessagesOk() ([]CoreMessageGetMessages200ResponseMessagesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Messages, true
}

// SetMessages sets field value
func (o *CoreMessageGetMessages200Response) SetMessages(v []CoreMessageGetMessages200ResponseMessagesInner) {
	o.Messages = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *CoreMessageGetMessages200Response) GetWarnings() []AuthEmailSignupUser200ResponseWarningsInner {
	if o == nil || IsNil(o.Warnings) {
		var ret []AuthEmailSignupUser200ResponseWarningsInner
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreMessageGetMessages200Response) GetWarningsOk() ([]AuthEmailSignupUser200ResponseWarningsInner, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *CoreMessageGetMessages200Response) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []AuthEmailSignupUser200ResponseWarningsInner and assigns it to the Warnings field.
func (o *CoreMessageGetMessages200Response) SetWarnings(v []AuthEmailSignupUser200ResponseWarningsInner) {
	o.Warnings = v
}

func (o CoreMessageGetMessages200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreMessageGetMessages200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["messages"] = o.Messages
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

func (o *CoreMessageGetMessages200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"messages",
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

	varCoreMessageGetMessages200Response := _CoreMessageGetMessages200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCoreMessageGetMessages200Response)

	if err != nil {
		return err
	}

	*o = CoreMessageGetMessages200Response(varCoreMessageGetMessages200Response)

	return err
}

type NullableCoreMessageGetMessages200Response struct {
	value *CoreMessageGetMessages200Response
	isSet bool
}

func (v NullableCoreMessageGetMessages200Response) Get() *CoreMessageGetMessages200Response {
	return v.value
}

func (v *NullableCoreMessageGetMessages200Response) Set(val *CoreMessageGetMessages200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreMessageGetMessages200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreMessageGetMessages200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreMessageGetMessages200Response(val *CoreMessageGetMessages200Response) *NullableCoreMessageGetMessages200Response {
	return &NullableCoreMessageGetMessages200Response{value: val, isSet: true}
}

func (v NullableCoreMessageGetMessages200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreMessageGetMessages200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


