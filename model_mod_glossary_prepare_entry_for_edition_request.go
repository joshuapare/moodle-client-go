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

// checks if the ModGlossaryPrepareEntryForEditionRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModGlossaryPrepareEntryForEditionRequest{}

// ModGlossaryPrepareEntryForEditionRequest struct for ModGlossaryPrepareEntryForEditionRequest
type ModGlossaryPrepareEntryForEditionRequest struct {
	// Glossary entry id to update
	Entryid int32 `json:"entryid"`
}

type _ModGlossaryPrepareEntryForEditionRequest ModGlossaryPrepareEntryForEditionRequest

// NewModGlossaryPrepareEntryForEditionRequest instantiates a new ModGlossaryPrepareEntryForEditionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModGlossaryPrepareEntryForEditionRequest(entryid int32) *ModGlossaryPrepareEntryForEditionRequest {
	this := ModGlossaryPrepareEntryForEditionRequest{}
	this.Entryid = entryid
	return &this
}

// NewModGlossaryPrepareEntryForEditionRequestWithDefaults instantiates a new ModGlossaryPrepareEntryForEditionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModGlossaryPrepareEntryForEditionRequestWithDefaults() *ModGlossaryPrepareEntryForEditionRequest {
	this := ModGlossaryPrepareEntryForEditionRequest{}
	var entryid int32 = null
	this.Entryid = entryid
	return &this
}

// GetEntryid returns the Entryid field value
func (o *ModGlossaryPrepareEntryForEditionRequest) GetEntryid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Entryid
}

// GetEntryidOk returns a tuple with the Entryid field value
// and a boolean to check if the value has been set.
func (o *ModGlossaryPrepareEntryForEditionRequest) GetEntryidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Entryid, true
}

// SetEntryid sets field value
func (o *ModGlossaryPrepareEntryForEditionRequest) SetEntryid(v int32) {
	o.Entryid = v
}

func (o ModGlossaryPrepareEntryForEditionRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModGlossaryPrepareEntryForEditionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["entryid"] = o.Entryid
	return toSerialize, nil
}

func (o *ModGlossaryPrepareEntryForEditionRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"entryid",
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

	varModGlossaryPrepareEntryForEditionRequest := _ModGlossaryPrepareEntryForEditionRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModGlossaryPrepareEntryForEditionRequest)

	if err != nil {
		return err
	}

	*o = ModGlossaryPrepareEntryForEditionRequest(varModGlossaryPrepareEntryForEditionRequest)

	return err
}

type NullableModGlossaryPrepareEntryForEditionRequest struct {
	value *ModGlossaryPrepareEntryForEditionRequest
	isSet bool
}

func (v NullableModGlossaryPrepareEntryForEditionRequest) Get() *ModGlossaryPrepareEntryForEditionRequest {
	return v.value
}

func (v *NullableModGlossaryPrepareEntryForEditionRequest) Set(val *ModGlossaryPrepareEntryForEditionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModGlossaryPrepareEntryForEditionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModGlossaryPrepareEntryForEditionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModGlossaryPrepareEntryForEditionRequest(val *ModGlossaryPrepareEntryForEditionRequest) *NullableModGlossaryPrepareEntryForEditionRequest {
	return &NullableModGlossaryPrepareEntryForEditionRequest{value: val, isSet: true}
}

func (v NullableModGlossaryPrepareEntryForEditionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModGlossaryPrepareEntryForEditionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


