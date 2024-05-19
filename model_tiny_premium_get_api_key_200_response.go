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

// checks if the TinyPremiumGetApiKey200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TinyPremiumGetApiKey200Response{}

// TinyPremiumGetApiKey200Response struct for TinyPremiumGetApiKey200Response
type TinyPremiumGetApiKey200Response struct {
	// The API key for Tiny Premium
	Apikey string `json:"apikey"`
}

type _TinyPremiumGetApiKey200Response TinyPremiumGetApiKey200Response

// NewTinyPremiumGetApiKey200Response instantiates a new TinyPremiumGetApiKey200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTinyPremiumGetApiKey200Response(apikey string) *TinyPremiumGetApiKey200Response {
	this := TinyPremiumGetApiKey200Response{}
	this.Apikey = apikey
	return &this
}

// NewTinyPremiumGetApiKey200ResponseWithDefaults instantiates a new TinyPremiumGetApiKey200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTinyPremiumGetApiKey200ResponseWithDefaults() *TinyPremiumGetApiKey200Response {
	this := TinyPremiumGetApiKey200Response{}
	var apikey string = "null"
	this.Apikey = apikey
	return &this
}

// GetApikey returns the Apikey field value
func (o *TinyPremiumGetApiKey200Response) GetApikey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Apikey
}

// GetApikeyOk returns a tuple with the Apikey field value
// and a boolean to check if the value has been set.
func (o *TinyPremiumGetApiKey200Response) GetApikeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Apikey, true
}

// SetApikey sets field value
func (o *TinyPremiumGetApiKey200Response) SetApikey(v string) {
	o.Apikey = v
}

func (o TinyPremiumGetApiKey200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TinyPremiumGetApiKey200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["apikey"] = o.Apikey
	return toSerialize, nil
}

func (o *TinyPremiumGetApiKey200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"apikey",
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

	varTinyPremiumGetApiKey200Response := _TinyPremiumGetApiKey200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTinyPremiumGetApiKey200Response)

	if err != nil {
		return err
	}

	*o = TinyPremiumGetApiKey200Response(varTinyPremiumGetApiKey200Response)

	return err
}

type NullableTinyPremiumGetApiKey200Response struct {
	value *TinyPremiumGetApiKey200Response
	isSet bool
}

func (v NullableTinyPremiumGetApiKey200Response) Get() *TinyPremiumGetApiKey200Response {
	return v.value
}

func (v *NullableTinyPremiumGetApiKey200Response) Set(val *TinyPremiumGetApiKey200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableTinyPremiumGetApiKey200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableTinyPremiumGetApiKey200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTinyPremiumGetApiKey200Response(val *TinyPremiumGetApiKey200Response) *NullableTinyPremiumGetApiKey200Response {
	return &NullableTinyPremiumGetApiKey200Response{value: val, isSet: true}
}

func (v NullableTinyPremiumGetApiKey200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTinyPremiumGetApiKey200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


