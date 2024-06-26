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

// checks if the ModFeedbackLaunchFeedback200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModFeedbackLaunchFeedback200Response{}

// ModFeedbackLaunchFeedback200Response struct for ModFeedbackLaunchFeedback200Response
type ModFeedbackLaunchFeedback200Response struct {
	// The next page to go (-1 if we were already in the last page). 0 for first page.
	Gopage int32 `json:"gopage"`
	Warnings []AuthEmailSignupUser200ResponseWarningsInner `json:"warnings,omitempty"`
}

type _ModFeedbackLaunchFeedback200Response ModFeedbackLaunchFeedback200Response

// NewModFeedbackLaunchFeedback200Response instantiates a new ModFeedbackLaunchFeedback200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModFeedbackLaunchFeedback200Response(gopage int32) *ModFeedbackLaunchFeedback200Response {
	this := ModFeedbackLaunchFeedback200Response{}
	this.Gopage = gopage
	return &this
}

// NewModFeedbackLaunchFeedback200ResponseWithDefaults instantiates a new ModFeedbackLaunchFeedback200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModFeedbackLaunchFeedback200ResponseWithDefaults() *ModFeedbackLaunchFeedback200Response {
	this := ModFeedbackLaunchFeedback200Response{}
	var gopage int32 = null
	this.Gopage = gopage
	return &this
}

// GetGopage returns the Gopage field value
func (o *ModFeedbackLaunchFeedback200Response) GetGopage() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Gopage
}

// GetGopageOk returns a tuple with the Gopage field value
// and a boolean to check if the value has been set.
func (o *ModFeedbackLaunchFeedback200Response) GetGopageOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Gopage, true
}

// SetGopage sets field value
func (o *ModFeedbackLaunchFeedback200Response) SetGopage(v int32) {
	o.Gopage = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *ModFeedbackLaunchFeedback200Response) GetWarnings() []AuthEmailSignupUser200ResponseWarningsInner {
	if o == nil || IsNil(o.Warnings) {
		var ret []AuthEmailSignupUser200ResponseWarningsInner
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModFeedbackLaunchFeedback200Response) GetWarningsOk() ([]AuthEmailSignupUser200ResponseWarningsInner, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *ModFeedbackLaunchFeedback200Response) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []AuthEmailSignupUser200ResponseWarningsInner and assigns it to the Warnings field.
func (o *ModFeedbackLaunchFeedback200Response) SetWarnings(v []AuthEmailSignupUser200ResponseWarningsInner) {
	o.Warnings = v
}

func (o ModFeedbackLaunchFeedback200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModFeedbackLaunchFeedback200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["gopage"] = o.Gopage
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

func (o *ModFeedbackLaunchFeedback200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"gopage",
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

	varModFeedbackLaunchFeedback200Response := _ModFeedbackLaunchFeedback200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModFeedbackLaunchFeedback200Response)

	if err != nil {
		return err
	}

	*o = ModFeedbackLaunchFeedback200Response(varModFeedbackLaunchFeedback200Response)

	return err
}

type NullableModFeedbackLaunchFeedback200Response struct {
	value *ModFeedbackLaunchFeedback200Response
	isSet bool
}

func (v NullableModFeedbackLaunchFeedback200Response) Get() *ModFeedbackLaunchFeedback200Response {
	return v.value
}

func (v *NullableModFeedbackLaunchFeedback200Response) Set(val *ModFeedbackLaunchFeedback200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableModFeedbackLaunchFeedback200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableModFeedbackLaunchFeedback200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModFeedbackLaunchFeedback200Response(val *ModFeedbackLaunchFeedback200Response) *NullableModFeedbackLaunchFeedback200Response {
	return &NullableModFeedbackLaunchFeedback200Response{value: val, isSet: true}
}

func (v NullableModFeedbackLaunchFeedback200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModFeedbackLaunchFeedback200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


