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

// checks if the ModDataDeleteEntryRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModDataDeleteEntryRequest{}

// ModDataDeleteEntryRequest struct for ModDataDeleteEntryRequest
type ModDataDeleteEntryRequest struct {
	// Record entry id.
	Entryid int32 `json:"entryid"`
}

type _ModDataDeleteEntryRequest ModDataDeleteEntryRequest

// NewModDataDeleteEntryRequest instantiates a new ModDataDeleteEntryRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModDataDeleteEntryRequest(entryid int32) *ModDataDeleteEntryRequest {
	this := ModDataDeleteEntryRequest{}
	this.Entryid = entryid
	return &this
}

// NewModDataDeleteEntryRequestWithDefaults instantiates a new ModDataDeleteEntryRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModDataDeleteEntryRequestWithDefaults() *ModDataDeleteEntryRequest {
	this := ModDataDeleteEntryRequest{}
	return &this
}

// GetEntryid returns the Entryid field value
func (o *ModDataDeleteEntryRequest) GetEntryid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Entryid
}

// GetEntryidOk returns a tuple with the Entryid field value
// and a boolean to check if the value has been set.
func (o *ModDataDeleteEntryRequest) GetEntryidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Entryid, true
}

// SetEntryid sets field value
func (o *ModDataDeleteEntryRequest) SetEntryid(v int32) {
	o.Entryid = v
}

func (o ModDataDeleteEntryRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModDataDeleteEntryRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["entryid"] = o.Entryid
	return toSerialize, nil
}

func (o *ModDataDeleteEntryRequest) UnmarshalJSON(data []byte) (err error) {
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

	varModDataDeleteEntryRequest := _ModDataDeleteEntryRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModDataDeleteEntryRequest)

	if err != nil {
		return err
	}

	*o = ModDataDeleteEntryRequest(varModDataDeleteEntryRequest)

	return err
}

type NullableModDataDeleteEntryRequest struct {
	value *ModDataDeleteEntryRequest
	isSet bool
}

func (v NullableModDataDeleteEntryRequest) Get() *ModDataDeleteEntryRequest {
	return v.value
}

func (v *NullableModDataDeleteEntryRequest) Set(val *ModDataDeleteEntryRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModDataDeleteEntryRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModDataDeleteEntryRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModDataDeleteEntryRequest(val *ModDataDeleteEntryRequest) *NullableModDataDeleteEntryRequest {
	return &NullableModDataDeleteEntryRequest{value: val, isSet: true}
}

func (v NullableModDataDeleteEntryRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModDataDeleteEntryRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

