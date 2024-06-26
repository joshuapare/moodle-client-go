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

// checks if the CoreAuthConfirmUser200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreAuthConfirmUser200Response{}

// CoreAuthConfirmUser200Response struct for CoreAuthConfirmUser200Response
type CoreAuthConfirmUser200Response struct {
	// True if the user was confirmed, false if he was already confirmed
	Success bool `json:"success"`
	Warnings []AuthEmailSignupUser200ResponseWarningsInner `json:"warnings,omitempty"`
}

type _CoreAuthConfirmUser200Response CoreAuthConfirmUser200Response

// NewCoreAuthConfirmUser200Response instantiates a new CoreAuthConfirmUser200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreAuthConfirmUser200Response(success bool) *CoreAuthConfirmUser200Response {
	this := CoreAuthConfirmUser200Response{}
	this.Success = success
	return &this
}

// NewCoreAuthConfirmUser200ResponseWithDefaults instantiates a new CoreAuthConfirmUser200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreAuthConfirmUser200ResponseWithDefaults() *CoreAuthConfirmUser200Response {
	this := CoreAuthConfirmUser200Response{}
	var success bool = null
	this.Success = success
	return &this
}

// GetSuccess returns the Success field value
func (o *CoreAuthConfirmUser200Response) GetSuccess() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value
// and a boolean to check if the value has been set.
func (o *CoreAuthConfirmUser200Response) GetSuccessOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Success, true
}

// SetSuccess sets field value
func (o *CoreAuthConfirmUser200Response) SetSuccess(v bool) {
	o.Success = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *CoreAuthConfirmUser200Response) GetWarnings() []AuthEmailSignupUser200ResponseWarningsInner {
	if o == nil || IsNil(o.Warnings) {
		var ret []AuthEmailSignupUser200ResponseWarningsInner
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreAuthConfirmUser200Response) GetWarningsOk() ([]AuthEmailSignupUser200ResponseWarningsInner, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *CoreAuthConfirmUser200Response) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []AuthEmailSignupUser200ResponseWarningsInner and assigns it to the Warnings field.
func (o *CoreAuthConfirmUser200Response) SetWarnings(v []AuthEmailSignupUser200ResponseWarningsInner) {
	o.Warnings = v
}

func (o CoreAuthConfirmUser200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreAuthConfirmUser200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["success"] = o.Success
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

func (o *CoreAuthConfirmUser200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"success",
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

	varCoreAuthConfirmUser200Response := _CoreAuthConfirmUser200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCoreAuthConfirmUser200Response)

	if err != nil {
		return err
	}

	*o = CoreAuthConfirmUser200Response(varCoreAuthConfirmUser200Response)

	return err
}

type NullableCoreAuthConfirmUser200Response struct {
	value *CoreAuthConfirmUser200Response
	isSet bool
}

func (v NullableCoreAuthConfirmUser200Response) Get() *CoreAuthConfirmUser200Response {
	return v.value
}

func (v *NullableCoreAuthConfirmUser200Response) Set(val *CoreAuthConfirmUser200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreAuthConfirmUser200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreAuthConfirmUser200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreAuthConfirmUser200Response(val *CoreAuthConfirmUser200Response) *NullableCoreAuthConfirmUser200Response {
	return &NullableCoreAuthConfirmUser200Response{value: val, isSet: true}
}

func (v NullableCoreAuthConfirmUser200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreAuthConfirmUser200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


