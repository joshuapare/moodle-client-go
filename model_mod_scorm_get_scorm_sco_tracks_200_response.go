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

// checks if the ModScormGetScormScoTracks200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModScormGetScormScoTracks200Response{}

// ModScormGetScormScoTracks200Response struct for ModScormGetScormScoTracks200Response
type ModScormGetScormScoTracks200Response struct {
	Data ModScormGetScormScoTracks200ResponseData `json:"data"`
	Warnings []AuthEmailSignupUser200ResponseWarningsInner `json:"warnings,omitempty"`
}

type _ModScormGetScormScoTracks200Response ModScormGetScormScoTracks200Response

// NewModScormGetScormScoTracks200Response instantiates a new ModScormGetScormScoTracks200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModScormGetScormScoTracks200Response(data ModScormGetScormScoTracks200ResponseData) *ModScormGetScormScoTracks200Response {
	this := ModScormGetScormScoTracks200Response{}
	this.Data = data
	return &this
}

// NewModScormGetScormScoTracks200ResponseWithDefaults instantiates a new ModScormGetScormScoTracks200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModScormGetScormScoTracks200ResponseWithDefaults() *ModScormGetScormScoTracks200Response {
	this := ModScormGetScormScoTracks200Response{}
	return &this
}

// GetData returns the Data field value
func (o *ModScormGetScormScoTracks200Response) GetData() ModScormGetScormScoTracks200ResponseData {
	if o == nil {
		var ret ModScormGetScormScoTracks200ResponseData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ModScormGetScormScoTracks200Response) GetDataOk() (*ModScormGetScormScoTracks200ResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ModScormGetScormScoTracks200Response) SetData(v ModScormGetScormScoTracks200ResponseData) {
	o.Data = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *ModScormGetScormScoTracks200Response) GetWarnings() []AuthEmailSignupUser200ResponseWarningsInner {
	if o == nil || IsNil(o.Warnings) {
		var ret []AuthEmailSignupUser200ResponseWarningsInner
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModScormGetScormScoTracks200Response) GetWarningsOk() ([]AuthEmailSignupUser200ResponseWarningsInner, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *ModScormGetScormScoTracks200Response) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []AuthEmailSignupUser200ResponseWarningsInner and assigns it to the Warnings field.
func (o *ModScormGetScormScoTracks200Response) SetWarnings(v []AuthEmailSignupUser200ResponseWarningsInner) {
	o.Warnings = v
}

func (o ModScormGetScormScoTracks200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModScormGetScormScoTracks200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

func (o *ModScormGetScormScoTracks200Response) UnmarshalJSON(data []byte) (err error) {
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

	varModScormGetScormScoTracks200Response := _ModScormGetScormScoTracks200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModScormGetScormScoTracks200Response)

	if err != nil {
		return err
	}

	*o = ModScormGetScormScoTracks200Response(varModScormGetScormScoTracks200Response)

	return err
}

type NullableModScormGetScormScoTracks200Response struct {
	value *ModScormGetScormScoTracks200Response
	isSet bool
}

func (v NullableModScormGetScormScoTracks200Response) Get() *ModScormGetScormScoTracks200Response {
	return v.value
}

func (v *NullableModScormGetScormScoTracks200Response) Set(val *ModScormGetScormScoTracks200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableModScormGetScormScoTracks200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableModScormGetScormScoTracks200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModScormGetScormScoTracks200Response(val *ModScormGetScormScoTracks200Response) *NullableModScormGetScormScoTracks200Response {
	return &NullableModScormGetScormScoTracks200Response{value: val, isSet: true}
}

func (v NullableModScormGetScormScoTracks200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModScormGetScormScoTracks200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


