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

// checks if the ModWikiEditPage200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModWikiEditPage200Response{}

// ModWikiEditPage200Response struct for ModWikiEditPage200Response
type ModWikiEditPage200Response struct {
	// Edited page id.
	Pageid int32 `json:"pageid"`
	Warnings []AuthEmailSignupUser200ResponseWarningsInner `json:"warnings,omitempty"`
}

type _ModWikiEditPage200Response ModWikiEditPage200Response

// NewModWikiEditPage200Response instantiates a new ModWikiEditPage200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModWikiEditPage200Response(pageid int32) *ModWikiEditPage200Response {
	this := ModWikiEditPage200Response{}
	this.Pageid = pageid
	return &this
}

// NewModWikiEditPage200ResponseWithDefaults instantiates a new ModWikiEditPage200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModWikiEditPage200ResponseWithDefaults() *ModWikiEditPage200Response {
	this := ModWikiEditPage200Response{}
	var pageid int32 = null
	this.Pageid = pageid
	return &this
}

// GetPageid returns the Pageid field value
func (o *ModWikiEditPage200Response) GetPageid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Pageid
}

// GetPageidOk returns a tuple with the Pageid field value
// and a boolean to check if the value has been set.
func (o *ModWikiEditPage200Response) GetPageidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pageid, true
}

// SetPageid sets field value
func (o *ModWikiEditPage200Response) SetPageid(v int32) {
	o.Pageid = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *ModWikiEditPage200Response) GetWarnings() []AuthEmailSignupUser200ResponseWarningsInner {
	if o == nil || IsNil(o.Warnings) {
		var ret []AuthEmailSignupUser200ResponseWarningsInner
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModWikiEditPage200Response) GetWarningsOk() ([]AuthEmailSignupUser200ResponseWarningsInner, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *ModWikiEditPage200Response) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []AuthEmailSignupUser200ResponseWarningsInner and assigns it to the Warnings field.
func (o *ModWikiEditPage200Response) SetWarnings(v []AuthEmailSignupUser200ResponseWarningsInner) {
	o.Warnings = v
}

func (o ModWikiEditPage200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModWikiEditPage200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pageid"] = o.Pageid
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

func (o *ModWikiEditPage200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"pageid",
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

	varModWikiEditPage200Response := _ModWikiEditPage200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModWikiEditPage200Response)

	if err != nil {
		return err
	}

	*o = ModWikiEditPage200Response(varModWikiEditPage200Response)

	return err
}

type NullableModWikiEditPage200Response struct {
	value *ModWikiEditPage200Response
	isSet bool
}

func (v NullableModWikiEditPage200Response) Get() *ModWikiEditPage200Response {
	return v.value
}

func (v *NullableModWikiEditPage200Response) Set(val *ModWikiEditPage200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableModWikiEditPage200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableModWikiEditPage200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModWikiEditPage200Response(val *ModWikiEditPage200Response) *NullableModWikiEditPage200Response {
	return &NullableModWikiEditPage200Response{value: val, isSet: true}
}

func (v NullableModWikiEditPage200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModWikiEditPage200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


