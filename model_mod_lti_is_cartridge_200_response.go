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

// checks if the ModLtiIsCartridge200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModLtiIsCartridge200Response{}

// ModLtiIsCartridge200Response struct for ModLtiIsCartridge200Response
type ModLtiIsCartridge200Response struct {
	// True if the URL is a cartridge
	Iscartridge bool `json:"iscartridge"`
}

type _ModLtiIsCartridge200Response ModLtiIsCartridge200Response

// NewModLtiIsCartridge200Response instantiates a new ModLtiIsCartridge200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModLtiIsCartridge200Response(iscartridge bool) *ModLtiIsCartridge200Response {
	this := ModLtiIsCartridge200Response{}
	this.Iscartridge = iscartridge
	return &this
}

// NewModLtiIsCartridge200ResponseWithDefaults instantiates a new ModLtiIsCartridge200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModLtiIsCartridge200ResponseWithDefaults() *ModLtiIsCartridge200Response {
	this := ModLtiIsCartridge200Response{}
	var iscartridge bool = null
	this.Iscartridge = iscartridge
	return &this
}

// GetIscartridge returns the Iscartridge field value
func (o *ModLtiIsCartridge200Response) GetIscartridge() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Iscartridge
}

// GetIscartridgeOk returns a tuple with the Iscartridge field value
// and a boolean to check if the value has been set.
func (o *ModLtiIsCartridge200Response) GetIscartridgeOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Iscartridge, true
}

// SetIscartridge sets field value
func (o *ModLtiIsCartridge200Response) SetIscartridge(v bool) {
	o.Iscartridge = v
}

func (o ModLtiIsCartridge200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModLtiIsCartridge200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["iscartridge"] = o.Iscartridge
	return toSerialize, nil
}

func (o *ModLtiIsCartridge200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"iscartridge",
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

	varModLtiIsCartridge200Response := _ModLtiIsCartridge200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModLtiIsCartridge200Response)

	if err != nil {
		return err
	}

	*o = ModLtiIsCartridge200Response(varModLtiIsCartridge200Response)

	return err
}

type NullableModLtiIsCartridge200Response struct {
	value *ModLtiIsCartridge200Response
	isSet bool
}

func (v NullableModLtiIsCartridge200Response) Get() *ModLtiIsCartridge200Response {
	return v.value
}

func (v *NullableModLtiIsCartridge200Response) Set(val *ModLtiIsCartridge200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableModLtiIsCartridge200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableModLtiIsCartridge200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModLtiIsCartridge200Response(val *ModLtiIsCartridge200Response) *NullableModLtiIsCartridge200Response {
	return &NullableModLtiIsCartridge200Response{value: val, isSet: true}
}

func (v NullableModLtiIsCartridge200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModLtiIsCartridge200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

