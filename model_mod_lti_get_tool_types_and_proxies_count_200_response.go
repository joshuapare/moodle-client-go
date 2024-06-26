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

// checks if the ModLtiGetToolTypesAndProxiesCount200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModLtiGetToolTypesAndProxiesCount200Response{}

// ModLtiGetToolTypesAndProxiesCount200Response struct for ModLtiGetToolTypesAndProxiesCount200Response
type ModLtiGetToolTypesAndProxiesCount200Response struct {
	// Total number of tool types and proxies
	Count int32 `json:"count"`
}

type _ModLtiGetToolTypesAndProxiesCount200Response ModLtiGetToolTypesAndProxiesCount200Response

// NewModLtiGetToolTypesAndProxiesCount200Response instantiates a new ModLtiGetToolTypesAndProxiesCount200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModLtiGetToolTypesAndProxiesCount200Response(count int32) *ModLtiGetToolTypesAndProxiesCount200Response {
	this := ModLtiGetToolTypesAndProxiesCount200Response{}
	this.Count = count
	return &this
}

// NewModLtiGetToolTypesAndProxiesCount200ResponseWithDefaults instantiates a new ModLtiGetToolTypesAndProxiesCount200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModLtiGetToolTypesAndProxiesCount200ResponseWithDefaults() *ModLtiGetToolTypesAndProxiesCount200Response {
	this := ModLtiGetToolTypesAndProxiesCount200Response{}
	var count int32 = null
	this.Count = count
	return &this
}

// GetCount returns the Count field value
func (o *ModLtiGetToolTypesAndProxiesCount200Response) GetCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *ModLtiGetToolTypesAndProxiesCount200Response) GetCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *ModLtiGetToolTypesAndProxiesCount200Response) SetCount(v int32) {
	o.Count = v
}

func (o ModLtiGetToolTypesAndProxiesCount200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModLtiGetToolTypesAndProxiesCount200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["count"] = o.Count
	return toSerialize, nil
}

func (o *ModLtiGetToolTypesAndProxiesCount200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"count",
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

	varModLtiGetToolTypesAndProxiesCount200Response := _ModLtiGetToolTypesAndProxiesCount200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModLtiGetToolTypesAndProxiesCount200Response)

	if err != nil {
		return err
	}

	*o = ModLtiGetToolTypesAndProxiesCount200Response(varModLtiGetToolTypesAndProxiesCount200Response)

	return err
}

type NullableModLtiGetToolTypesAndProxiesCount200Response struct {
	value *ModLtiGetToolTypesAndProxiesCount200Response
	isSet bool
}

func (v NullableModLtiGetToolTypesAndProxiesCount200Response) Get() *ModLtiGetToolTypesAndProxiesCount200Response {
	return v.value
}

func (v *NullableModLtiGetToolTypesAndProxiesCount200Response) Set(val *ModLtiGetToolTypesAndProxiesCount200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableModLtiGetToolTypesAndProxiesCount200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableModLtiGetToolTypesAndProxiesCount200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModLtiGetToolTypesAndProxiesCount200Response(val *ModLtiGetToolTypesAndProxiesCount200Response) *NullableModLtiGetToolTypesAndProxiesCount200Response {
	return &NullableModLtiGetToolTypesAndProxiesCount200Response{value: val, isSet: true}
}

func (v NullableModLtiGetToolTypesAndProxiesCount200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModLtiGetToolTypesAndProxiesCount200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


