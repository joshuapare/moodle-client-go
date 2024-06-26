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

// checks if the CoreAuthIsMinorRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreAuthIsMinorRequest{}

// CoreAuthIsMinorRequest struct for CoreAuthIsMinorRequest
type CoreAuthIsMinorRequest struct {
	// Age
	Age int32 `json:"age"`
	// Country of residence
	Country string `json:"country"`
}

type _CoreAuthIsMinorRequest CoreAuthIsMinorRequest

// NewCoreAuthIsMinorRequest instantiates a new CoreAuthIsMinorRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreAuthIsMinorRequest(age int32, country string) *CoreAuthIsMinorRequest {
	this := CoreAuthIsMinorRequest{}
	this.Age = age
	this.Country = country
	return &this
}

// NewCoreAuthIsMinorRequestWithDefaults instantiates a new CoreAuthIsMinorRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreAuthIsMinorRequestWithDefaults() *CoreAuthIsMinorRequest {
	this := CoreAuthIsMinorRequest{}
	var age int32 = null
	this.Age = age
	var country string = "null"
	this.Country = country
	return &this
}

// GetAge returns the Age field value
func (o *CoreAuthIsMinorRequest) GetAge() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Age
}

// GetAgeOk returns a tuple with the Age field value
// and a boolean to check if the value has been set.
func (o *CoreAuthIsMinorRequest) GetAgeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Age, true
}

// SetAge sets field value
func (o *CoreAuthIsMinorRequest) SetAge(v int32) {
	o.Age = v
}

// GetCountry returns the Country field value
func (o *CoreAuthIsMinorRequest) GetCountry() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Country
}

// GetCountryOk returns a tuple with the Country field value
// and a boolean to check if the value has been set.
func (o *CoreAuthIsMinorRequest) GetCountryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Country, true
}

// SetCountry sets field value
func (o *CoreAuthIsMinorRequest) SetCountry(v string) {
	o.Country = v
}

func (o CoreAuthIsMinorRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreAuthIsMinorRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["age"] = o.Age
	toSerialize["country"] = o.Country
	return toSerialize, nil
}

func (o *CoreAuthIsMinorRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"age",
		"country",
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

	varCoreAuthIsMinorRequest := _CoreAuthIsMinorRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCoreAuthIsMinorRequest)

	if err != nil {
		return err
	}

	*o = CoreAuthIsMinorRequest(varCoreAuthIsMinorRequest)

	return err
}

type NullableCoreAuthIsMinorRequest struct {
	value *CoreAuthIsMinorRequest
	isSet bool
}

func (v NullableCoreAuthIsMinorRequest) Get() *CoreAuthIsMinorRequest {
	return v.value
}

func (v *NullableCoreAuthIsMinorRequest) Set(val *CoreAuthIsMinorRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreAuthIsMinorRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreAuthIsMinorRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreAuthIsMinorRequest(val *CoreAuthIsMinorRequest) *NullableCoreAuthIsMinorRequest {
	return &NullableCoreAuthIsMinorRequest{value: val, isSet: true}
}

func (v NullableCoreAuthIsMinorRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreAuthIsMinorRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


