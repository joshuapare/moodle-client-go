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

// checks if the CoreMessageMessageSearchUsers200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreMessageMessageSearchUsers200Response{}

// CoreMessageMessageSearchUsers200Response struct for CoreMessageMessageSearchUsers200Response
type CoreMessageMessageSearchUsers200Response struct {
	Contacts []CoreMessageGetConversationBetweenUsers200ResponseMembersInner `json:"contacts"`
	Noncontacts []CoreMessageGetConversationBetweenUsers200ResponseMembersInner `json:"noncontacts"`
}

type _CoreMessageMessageSearchUsers200Response CoreMessageMessageSearchUsers200Response

// NewCoreMessageMessageSearchUsers200Response instantiates a new CoreMessageMessageSearchUsers200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreMessageMessageSearchUsers200Response(contacts []CoreMessageGetConversationBetweenUsers200ResponseMembersInner, noncontacts []CoreMessageGetConversationBetweenUsers200ResponseMembersInner) *CoreMessageMessageSearchUsers200Response {
	this := CoreMessageMessageSearchUsers200Response{}
	this.Contacts = contacts
	this.Noncontacts = noncontacts
	return &this
}

// NewCoreMessageMessageSearchUsers200ResponseWithDefaults instantiates a new CoreMessageMessageSearchUsers200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreMessageMessageSearchUsers200ResponseWithDefaults() *CoreMessageMessageSearchUsers200Response {
	this := CoreMessageMessageSearchUsers200Response{}
	return &this
}

// GetContacts returns the Contacts field value
func (o *CoreMessageMessageSearchUsers200Response) GetContacts() []CoreMessageGetConversationBetweenUsers200ResponseMembersInner {
	if o == nil {
		var ret []CoreMessageGetConversationBetweenUsers200ResponseMembersInner
		return ret
	}

	return o.Contacts
}

// GetContactsOk returns a tuple with the Contacts field value
// and a boolean to check if the value has been set.
func (o *CoreMessageMessageSearchUsers200Response) GetContactsOk() ([]CoreMessageGetConversationBetweenUsers200ResponseMembersInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Contacts, true
}

// SetContacts sets field value
func (o *CoreMessageMessageSearchUsers200Response) SetContacts(v []CoreMessageGetConversationBetweenUsers200ResponseMembersInner) {
	o.Contacts = v
}

// GetNoncontacts returns the Noncontacts field value
func (o *CoreMessageMessageSearchUsers200Response) GetNoncontacts() []CoreMessageGetConversationBetweenUsers200ResponseMembersInner {
	if o == nil {
		var ret []CoreMessageGetConversationBetweenUsers200ResponseMembersInner
		return ret
	}

	return o.Noncontacts
}

// GetNoncontactsOk returns a tuple with the Noncontacts field value
// and a boolean to check if the value has been set.
func (o *CoreMessageMessageSearchUsers200Response) GetNoncontactsOk() ([]CoreMessageGetConversationBetweenUsers200ResponseMembersInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Noncontacts, true
}

// SetNoncontacts sets field value
func (o *CoreMessageMessageSearchUsers200Response) SetNoncontacts(v []CoreMessageGetConversationBetweenUsers200ResponseMembersInner) {
	o.Noncontacts = v
}

func (o CoreMessageMessageSearchUsers200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreMessageMessageSearchUsers200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["contacts"] = o.Contacts
	toSerialize["noncontacts"] = o.Noncontacts
	return toSerialize, nil
}

func (o *CoreMessageMessageSearchUsers200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"contacts",
		"noncontacts",
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

	varCoreMessageMessageSearchUsers200Response := _CoreMessageMessageSearchUsers200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCoreMessageMessageSearchUsers200Response)

	if err != nil {
		return err
	}

	*o = CoreMessageMessageSearchUsers200Response(varCoreMessageMessageSearchUsers200Response)

	return err
}

type NullableCoreMessageMessageSearchUsers200Response struct {
	value *CoreMessageMessageSearchUsers200Response
	isSet bool
}

func (v NullableCoreMessageMessageSearchUsers200Response) Get() *CoreMessageMessageSearchUsers200Response {
	return v.value
}

func (v *NullableCoreMessageMessageSearchUsers200Response) Set(val *CoreMessageMessageSearchUsers200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreMessageMessageSearchUsers200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreMessageMessageSearchUsers200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreMessageMessageSearchUsers200Response(val *CoreMessageMessageSearchUsers200Response) *NullableCoreMessageMessageSearchUsers200Response {
	return &NullableCoreMessageMessageSearchUsers200Response{value: val, isSet: true}
}

func (v NullableCoreMessageMessageSearchUsers200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreMessageMessageSearchUsers200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


