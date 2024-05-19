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

// checks if the ToolMobileGetAutologinKey200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ToolMobileGetAutologinKey200Response{}

// ToolMobileGetAutologinKey200Response struct for ToolMobileGetAutologinKey200Response
type ToolMobileGetAutologinKey200Response struct {
	// Auto-login URL.
	Autologinurl string `json:"autologinurl"`
	// Auto-login key for a single usage with time expiration.
	Key string `json:"key"`
	Warnings []AuthEmailSignupUser200ResponseWarningsInner `json:"warnings,omitempty"`
}

type _ToolMobileGetAutologinKey200Response ToolMobileGetAutologinKey200Response

// NewToolMobileGetAutologinKey200Response instantiates a new ToolMobileGetAutologinKey200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewToolMobileGetAutologinKey200Response(autologinurl string, key string) *ToolMobileGetAutologinKey200Response {
	this := ToolMobileGetAutologinKey200Response{}
	this.Autologinurl = autologinurl
	this.Key = key
	return &this
}

// NewToolMobileGetAutologinKey200ResponseWithDefaults instantiates a new ToolMobileGetAutologinKey200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewToolMobileGetAutologinKey200ResponseWithDefaults() *ToolMobileGetAutologinKey200Response {
	this := ToolMobileGetAutologinKey200Response{}
	var autologinurl string = "null"
	this.Autologinurl = autologinurl
	var key string = "null"
	this.Key = key
	return &this
}

// GetAutologinurl returns the Autologinurl field value
func (o *ToolMobileGetAutologinKey200Response) GetAutologinurl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Autologinurl
}

// GetAutologinurlOk returns a tuple with the Autologinurl field value
// and a boolean to check if the value has been set.
func (o *ToolMobileGetAutologinKey200Response) GetAutologinurlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Autologinurl, true
}

// SetAutologinurl sets field value
func (o *ToolMobileGetAutologinKey200Response) SetAutologinurl(v string) {
	o.Autologinurl = v
}

// GetKey returns the Key field value
func (o *ToolMobileGetAutologinKey200Response) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *ToolMobileGetAutologinKey200Response) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *ToolMobileGetAutologinKey200Response) SetKey(v string) {
	o.Key = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *ToolMobileGetAutologinKey200Response) GetWarnings() []AuthEmailSignupUser200ResponseWarningsInner {
	if o == nil || IsNil(o.Warnings) {
		var ret []AuthEmailSignupUser200ResponseWarningsInner
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ToolMobileGetAutologinKey200Response) GetWarningsOk() ([]AuthEmailSignupUser200ResponseWarningsInner, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *ToolMobileGetAutologinKey200Response) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []AuthEmailSignupUser200ResponseWarningsInner and assigns it to the Warnings field.
func (o *ToolMobileGetAutologinKey200Response) SetWarnings(v []AuthEmailSignupUser200ResponseWarningsInner) {
	o.Warnings = v
}

func (o ToolMobileGetAutologinKey200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ToolMobileGetAutologinKey200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["autologinurl"] = o.Autologinurl
	toSerialize["key"] = o.Key
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

func (o *ToolMobileGetAutologinKey200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"autologinurl",
		"key",
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

	varToolMobileGetAutologinKey200Response := _ToolMobileGetAutologinKey200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varToolMobileGetAutologinKey200Response)

	if err != nil {
		return err
	}

	*o = ToolMobileGetAutologinKey200Response(varToolMobileGetAutologinKey200Response)

	return err
}

type NullableToolMobileGetAutologinKey200Response struct {
	value *ToolMobileGetAutologinKey200Response
	isSet bool
}

func (v NullableToolMobileGetAutologinKey200Response) Get() *ToolMobileGetAutologinKey200Response {
	return v.value
}

func (v *NullableToolMobileGetAutologinKey200Response) Set(val *ToolMobileGetAutologinKey200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableToolMobileGetAutologinKey200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableToolMobileGetAutologinKey200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableToolMobileGetAutologinKey200Response(val *ToolMobileGetAutologinKey200Response) *NullableToolMobileGetAutologinKey200Response {
	return &NullableToolMobileGetAutologinKey200Response{value: val, isSet: true}
}

func (v NullableToolMobileGetAutologinKey200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableToolMobileGetAutologinKey200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

