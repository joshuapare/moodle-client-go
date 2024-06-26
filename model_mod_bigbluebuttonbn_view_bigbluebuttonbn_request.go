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

// checks if the ModBigbluebuttonbnViewBigbluebuttonbnRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModBigbluebuttonbnViewBigbluebuttonbnRequest{}

// ModBigbluebuttonbnViewBigbluebuttonbnRequest struct for ModBigbluebuttonbnViewBigbluebuttonbnRequest
type ModBigbluebuttonbnViewBigbluebuttonbnRequest struct {
	// bigbluebuttonbn instance id
	Bigbluebuttonbnid int32 `json:"bigbluebuttonbnid"`
}

type _ModBigbluebuttonbnViewBigbluebuttonbnRequest ModBigbluebuttonbnViewBigbluebuttonbnRequest

// NewModBigbluebuttonbnViewBigbluebuttonbnRequest instantiates a new ModBigbluebuttonbnViewBigbluebuttonbnRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModBigbluebuttonbnViewBigbluebuttonbnRequest(bigbluebuttonbnid int32) *ModBigbluebuttonbnViewBigbluebuttonbnRequest {
	this := ModBigbluebuttonbnViewBigbluebuttonbnRequest{}
	this.Bigbluebuttonbnid = bigbluebuttonbnid
	return &this
}

// NewModBigbluebuttonbnViewBigbluebuttonbnRequestWithDefaults instantiates a new ModBigbluebuttonbnViewBigbluebuttonbnRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModBigbluebuttonbnViewBigbluebuttonbnRequestWithDefaults() *ModBigbluebuttonbnViewBigbluebuttonbnRequest {
	this := ModBigbluebuttonbnViewBigbluebuttonbnRequest{}
	return &this
}

// GetBigbluebuttonbnid returns the Bigbluebuttonbnid field value
func (o *ModBigbluebuttonbnViewBigbluebuttonbnRequest) GetBigbluebuttonbnid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Bigbluebuttonbnid
}

// GetBigbluebuttonbnidOk returns a tuple with the Bigbluebuttonbnid field value
// and a boolean to check if the value has been set.
func (o *ModBigbluebuttonbnViewBigbluebuttonbnRequest) GetBigbluebuttonbnidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Bigbluebuttonbnid, true
}

// SetBigbluebuttonbnid sets field value
func (o *ModBigbluebuttonbnViewBigbluebuttonbnRequest) SetBigbluebuttonbnid(v int32) {
	o.Bigbluebuttonbnid = v
}

func (o ModBigbluebuttonbnViewBigbluebuttonbnRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModBigbluebuttonbnViewBigbluebuttonbnRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["bigbluebuttonbnid"] = o.Bigbluebuttonbnid
	return toSerialize, nil
}

func (o *ModBigbluebuttonbnViewBigbluebuttonbnRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"bigbluebuttonbnid",
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

	varModBigbluebuttonbnViewBigbluebuttonbnRequest := _ModBigbluebuttonbnViewBigbluebuttonbnRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModBigbluebuttonbnViewBigbluebuttonbnRequest)

	if err != nil {
		return err
	}

	*o = ModBigbluebuttonbnViewBigbluebuttonbnRequest(varModBigbluebuttonbnViewBigbluebuttonbnRequest)

	return err
}

type NullableModBigbluebuttonbnViewBigbluebuttonbnRequest struct {
	value *ModBigbluebuttonbnViewBigbluebuttonbnRequest
	isSet bool
}

func (v NullableModBigbluebuttonbnViewBigbluebuttonbnRequest) Get() *ModBigbluebuttonbnViewBigbluebuttonbnRequest {
	return v.value
}

func (v *NullableModBigbluebuttonbnViewBigbluebuttonbnRequest) Set(val *ModBigbluebuttonbnViewBigbluebuttonbnRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModBigbluebuttonbnViewBigbluebuttonbnRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModBigbluebuttonbnViewBigbluebuttonbnRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModBigbluebuttonbnViewBigbluebuttonbnRequest(val *ModBigbluebuttonbnViewBigbluebuttonbnRequest) *NullableModBigbluebuttonbnViewBigbluebuttonbnRequest {
	return &NullableModBigbluebuttonbnViewBigbluebuttonbnRequest{value: val, isSet: true}
}

func (v NullableModBigbluebuttonbnViewBigbluebuttonbnRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModBigbluebuttonbnViewBigbluebuttonbnRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


