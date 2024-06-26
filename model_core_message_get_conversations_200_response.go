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

// checks if the CoreMessageGetConversations200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreMessageGetConversations200Response{}

// CoreMessageGetConversations200Response struct for CoreMessageGetConversations200Response
type CoreMessageGetConversations200Response struct {
	Conversations []CoreMessageGetConversations200ResponseConversationsInner `json:"conversations"`
}

type _CoreMessageGetConversations200Response CoreMessageGetConversations200Response

// NewCoreMessageGetConversations200Response instantiates a new CoreMessageGetConversations200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreMessageGetConversations200Response(conversations []CoreMessageGetConversations200ResponseConversationsInner) *CoreMessageGetConversations200Response {
	this := CoreMessageGetConversations200Response{}
	this.Conversations = conversations
	return &this
}

// NewCoreMessageGetConversations200ResponseWithDefaults instantiates a new CoreMessageGetConversations200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreMessageGetConversations200ResponseWithDefaults() *CoreMessageGetConversations200Response {
	this := CoreMessageGetConversations200Response{}
	return &this
}

// GetConversations returns the Conversations field value
func (o *CoreMessageGetConversations200Response) GetConversations() []CoreMessageGetConversations200ResponseConversationsInner {
	if o == nil {
		var ret []CoreMessageGetConversations200ResponseConversationsInner
		return ret
	}

	return o.Conversations
}

// GetConversationsOk returns a tuple with the Conversations field value
// and a boolean to check if the value has been set.
func (o *CoreMessageGetConversations200Response) GetConversationsOk() ([]CoreMessageGetConversations200ResponseConversationsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Conversations, true
}

// SetConversations sets field value
func (o *CoreMessageGetConversations200Response) SetConversations(v []CoreMessageGetConversations200ResponseConversationsInner) {
	o.Conversations = v
}

func (o CoreMessageGetConversations200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreMessageGetConversations200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["conversations"] = o.Conversations
	return toSerialize, nil
}

func (o *CoreMessageGetConversations200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"conversations",
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

	varCoreMessageGetConversations200Response := _CoreMessageGetConversations200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCoreMessageGetConversations200Response)

	if err != nil {
		return err
	}

	*o = CoreMessageGetConversations200Response(varCoreMessageGetConversations200Response)

	return err
}

type NullableCoreMessageGetConversations200Response struct {
	value *CoreMessageGetConversations200Response
	isSet bool
}

func (v NullableCoreMessageGetConversations200Response) Get() *CoreMessageGetConversations200Response {
	return v.value
}

func (v *NullableCoreMessageGetConversations200Response) Set(val *CoreMessageGetConversations200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreMessageGetConversations200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreMessageGetConversations200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreMessageGetConversations200Response(val *CoreMessageGetConversations200Response) *NullableCoreMessageGetConversations200Response {
	return &NullableCoreMessageGetConversations200Response{value: val, isSet: true}
}

func (v NullableCoreMessageGetConversations200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreMessageGetConversations200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


