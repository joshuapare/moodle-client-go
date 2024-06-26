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

// checks if the ToolMobileGetTokensForQrLogin200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ToolMobileGetTokensForQrLogin200Response{}

// ToolMobileGetTokensForQrLogin200Response struct for ToolMobileGetTokensForQrLogin200Response
type ToolMobileGetTokensForQrLogin200Response struct {
	// Private token used for auto-login processes.
	Privatetoken string `json:"privatetoken"`
	// A valid WebService token for the official mobile app service.
	Token string `json:"token"`
	Warnings []AuthEmailSignupUser200ResponseWarningsInner `json:"warnings,omitempty"`
}

type _ToolMobileGetTokensForQrLogin200Response ToolMobileGetTokensForQrLogin200Response

// NewToolMobileGetTokensForQrLogin200Response instantiates a new ToolMobileGetTokensForQrLogin200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewToolMobileGetTokensForQrLogin200Response(privatetoken string, token string) *ToolMobileGetTokensForQrLogin200Response {
	this := ToolMobileGetTokensForQrLogin200Response{}
	this.Privatetoken = privatetoken
	this.Token = token
	return &this
}

// NewToolMobileGetTokensForQrLogin200ResponseWithDefaults instantiates a new ToolMobileGetTokensForQrLogin200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewToolMobileGetTokensForQrLogin200ResponseWithDefaults() *ToolMobileGetTokensForQrLogin200Response {
	this := ToolMobileGetTokensForQrLogin200Response{}
	var privatetoken string = "null"
	this.Privatetoken = privatetoken
	var token string = "null"
	this.Token = token
	return &this
}

// GetPrivatetoken returns the Privatetoken field value
func (o *ToolMobileGetTokensForQrLogin200Response) GetPrivatetoken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Privatetoken
}

// GetPrivatetokenOk returns a tuple with the Privatetoken field value
// and a boolean to check if the value has been set.
func (o *ToolMobileGetTokensForQrLogin200Response) GetPrivatetokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Privatetoken, true
}

// SetPrivatetoken sets field value
func (o *ToolMobileGetTokensForQrLogin200Response) SetPrivatetoken(v string) {
	o.Privatetoken = v
}

// GetToken returns the Token field value
func (o *ToolMobileGetTokensForQrLogin200Response) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *ToolMobileGetTokensForQrLogin200Response) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *ToolMobileGetTokensForQrLogin200Response) SetToken(v string) {
	o.Token = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *ToolMobileGetTokensForQrLogin200Response) GetWarnings() []AuthEmailSignupUser200ResponseWarningsInner {
	if o == nil || IsNil(o.Warnings) {
		var ret []AuthEmailSignupUser200ResponseWarningsInner
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ToolMobileGetTokensForQrLogin200Response) GetWarningsOk() ([]AuthEmailSignupUser200ResponseWarningsInner, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *ToolMobileGetTokensForQrLogin200Response) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []AuthEmailSignupUser200ResponseWarningsInner and assigns it to the Warnings field.
func (o *ToolMobileGetTokensForQrLogin200Response) SetWarnings(v []AuthEmailSignupUser200ResponseWarningsInner) {
	o.Warnings = v
}

func (o ToolMobileGetTokensForQrLogin200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ToolMobileGetTokensForQrLogin200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["privatetoken"] = o.Privatetoken
	toSerialize["token"] = o.Token
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

func (o *ToolMobileGetTokensForQrLogin200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"privatetoken",
		"token",
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

	varToolMobileGetTokensForQrLogin200Response := _ToolMobileGetTokensForQrLogin200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varToolMobileGetTokensForQrLogin200Response)

	if err != nil {
		return err
	}

	*o = ToolMobileGetTokensForQrLogin200Response(varToolMobileGetTokensForQrLogin200Response)

	return err
}

type NullableToolMobileGetTokensForQrLogin200Response struct {
	value *ToolMobileGetTokensForQrLogin200Response
	isSet bool
}

func (v NullableToolMobileGetTokensForQrLogin200Response) Get() *ToolMobileGetTokensForQrLogin200Response {
	return v.value
}

func (v *NullableToolMobileGetTokensForQrLogin200Response) Set(val *ToolMobileGetTokensForQrLogin200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableToolMobileGetTokensForQrLogin200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableToolMobileGetTokensForQrLogin200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableToolMobileGetTokensForQrLogin200Response(val *ToolMobileGetTokensForQrLogin200Response) *NullableToolMobileGetTokensForQrLogin200Response {
	return &NullableToolMobileGetTokensForQrLogin200Response{value: val, isSet: true}
}

func (v NullableToolMobileGetTokensForQrLogin200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableToolMobileGetTokensForQrLogin200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


