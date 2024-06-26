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

// checks if the CoreMessageDataForMessageareaSearchMessages200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreMessageDataForMessageareaSearchMessages200Response{}

// CoreMessageDataForMessageareaSearchMessages200Response struct for CoreMessageDataForMessageareaSearchMessages200Response
type CoreMessageDataForMessageareaSearchMessages200Response struct {
	Contacts []CoreMessageDataForMessageareaSearchMessages200ResponseContactsInner `json:"contacts"`
}

type _CoreMessageDataForMessageareaSearchMessages200Response CoreMessageDataForMessageareaSearchMessages200Response

// NewCoreMessageDataForMessageareaSearchMessages200Response instantiates a new CoreMessageDataForMessageareaSearchMessages200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreMessageDataForMessageareaSearchMessages200Response(contacts []CoreMessageDataForMessageareaSearchMessages200ResponseContactsInner) *CoreMessageDataForMessageareaSearchMessages200Response {
	this := CoreMessageDataForMessageareaSearchMessages200Response{}
	this.Contacts = contacts
	return &this
}

// NewCoreMessageDataForMessageareaSearchMessages200ResponseWithDefaults instantiates a new CoreMessageDataForMessageareaSearchMessages200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreMessageDataForMessageareaSearchMessages200ResponseWithDefaults() *CoreMessageDataForMessageareaSearchMessages200Response {
	this := CoreMessageDataForMessageareaSearchMessages200Response{}
	return &this
}

// GetContacts returns the Contacts field value
func (o *CoreMessageDataForMessageareaSearchMessages200Response) GetContacts() []CoreMessageDataForMessageareaSearchMessages200ResponseContactsInner {
	if o == nil {
		var ret []CoreMessageDataForMessageareaSearchMessages200ResponseContactsInner
		return ret
	}

	return o.Contacts
}

// GetContactsOk returns a tuple with the Contacts field value
// and a boolean to check if the value has been set.
func (o *CoreMessageDataForMessageareaSearchMessages200Response) GetContactsOk() ([]CoreMessageDataForMessageareaSearchMessages200ResponseContactsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Contacts, true
}

// SetContacts sets field value
func (o *CoreMessageDataForMessageareaSearchMessages200Response) SetContacts(v []CoreMessageDataForMessageareaSearchMessages200ResponseContactsInner) {
	o.Contacts = v
}

func (o CoreMessageDataForMessageareaSearchMessages200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreMessageDataForMessageareaSearchMessages200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["contacts"] = o.Contacts
	return toSerialize, nil
}

func (o *CoreMessageDataForMessageareaSearchMessages200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"contacts",
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

	varCoreMessageDataForMessageareaSearchMessages200Response := _CoreMessageDataForMessageareaSearchMessages200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCoreMessageDataForMessageareaSearchMessages200Response)

	if err != nil {
		return err
	}

	*o = CoreMessageDataForMessageareaSearchMessages200Response(varCoreMessageDataForMessageareaSearchMessages200Response)

	return err
}

type NullableCoreMessageDataForMessageareaSearchMessages200Response struct {
	value *CoreMessageDataForMessageareaSearchMessages200Response
	isSet bool
}

func (v NullableCoreMessageDataForMessageareaSearchMessages200Response) Get() *CoreMessageDataForMessageareaSearchMessages200Response {
	return v.value
}

func (v *NullableCoreMessageDataForMessageareaSearchMessages200Response) Set(val *CoreMessageDataForMessageareaSearchMessages200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreMessageDataForMessageareaSearchMessages200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreMessageDataForMessageareaSearchMessages200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreMessageDataForMessageareaSearchMessages200Response(val *CoreMessageDataForMessageareaSearchMessages200Response) *NullableCoreMessageDataForMessageareaSearchMessages200Response {
	return &NullableCoreMessageDataForMessageareaSearchMessages200Response{value: val, isSet: true}
}

func (v NullableCoreMessageDataForMessageareaSearchMessages200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreMessageDataForMessageareaSearchMessages200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


