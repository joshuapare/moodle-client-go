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

// checks if the ModScormGetScormUserData200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModScormGetScormUserData200Response{}

// ModScormGetScormUserData200Response struct for ModScormGetScormUserData200Response
type ModScormGetScormUserData200Response struct {
	Data []ModScormGetScormUserData200ResponseDataInner `json:"data"`
	Warnings []AuthEmailSignupUser200ResponseWarningsInner `json:"warnings,omitempty"`
}

type _ModScormGetScormUserData200Response ModScormGetScormUserData200Response

// NewModScormGetScormUserData200Response instantiates a new ModScormGetScormUserData200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModScormGetScormUserData200Response(data []ModScormGetScormUserData200ResponseDataInner) *ModScormGetScormUserData200Response {
	this := ModScormGetScormUserData200Response{}
	this.Data = data
	return &this
}

// NewModScormGetScormUserData200ResponseWithDefaults instantiates a new ModScormGetScormUserData200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModScormGetScormUserData200ResponseWithDefaults() *ModScormGetScormUserData200Response {
	this := ModScormGetScormUserData200Response{}
	return &this
}

// GetData returns the Data field value
func (o *ModScormGetScormUserData200Response) GetData() []ModScormGetScormUserData200ResponseDataInner {
	if o == nil {
		var ret []ModScormGetScormUserData200ResponseDataInner
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ModScormGetScormUserData200Response) GetDataOk() ([]ModScormGetScormUserData200ResponseDataInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *ModScormGetScormUserData200Response) SetData(v []ModScormGetScormUserData200ResponseDataInner) {
	o.Data = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *ModScormGetScormUserData200Response) GetWarnings() []AuthEmailSignupUser200ResponseWarningsInner {
	if o == nil || IsNil(o.Warnings) {
		var ret []AuthEmailSignupUser200ResponseWarningsInner
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModScormGetScormUserData200Response) GetWarningsOk() ([]AuthEmailSignupUser200ResponseWarningsInner, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *ModScormGetScormUserData200Response) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []AuthEmailSignupUser200ResponseWarningsInner and assigns it to the Warnings field.
func (o *ModScormGetScormUserData200Response) SetWarnings(v []AuthEmailSignupUser200ResponseWarningsInner) {
	o.Warnings = v
}

func (o ModScormGetScormUserData200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModScormGetScormUserData200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

func (o *ModScormGetScormUserData200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
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

	varModScormGetScormUserData200Response := _ModScormGetScormUserData200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModScormGetScormUserData200Response)

	if err != nil {
		return err
	}

	*o = ModScormGetScormUserData200Response(varModScormGetScormUserData200Response)

	return err
}

type NullableModScormGetScormUserData200Response struct {
	value *ModScormGetScormUserData200Response
	isSet bool
}

func (v NullableModScormGetScormUserData200Response) Get() *ModScormGetScormUserData200Response {
	return v.value
}

func (v *NullableModScormGetScormUserData200Response) Set(val *ModScormGetScormUserData200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableModScormGetScormUserData200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableModScormGetScormUserData200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModScormGetScormUserData200Response(val *ModScormGetScormUserData200Response) *NullableModScormGetScormUserData200Response {
	return &NullableModScormGetScormUserData200Response{value: val, isSet: true}
}

func (v NullableModScormGetScormUserData200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModScormGetScormUserData200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

