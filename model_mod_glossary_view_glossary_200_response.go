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

// checks if the ModGlossaryViewGlossary200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModGlossaryViewGlossary200Response{}

// ModGlossaryViewGlossary200Response struct for ModGlossaryViewGlossary200Response
type ModGlossaryViewGlossary200Response struct {
	// True on success
	Status bool `json:"status"`
	Warnings []AuthEmailSignupUser200ResponseWarningsInner `json:"warnings,omitempty"`
}

type _ModGlossaryViewGlossary200Response ModGlossaryViewGlossary200Response

// NewModGlossaryViewGlossary200Response instantiates a new ModGlossaryViewGlossary200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModGlossaryViewGlossary200Response(status bool) *ModGlossaryViewGlossary200Response {
	this := ModGlossaryViewGlossary200Response{}
	this.Status = status
	return &this
}

// NewModGlossaryViewGlossary200ResponseWithDefaults instantiates a new ModGlossaryViewGlossary200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModGlossaryViewGlossary200ResponseWithDefaults() *ModGlossaryViewGlossary200Response {
	this := ModGlossaryViewGlossary200Response{}
	return &this
}

// GetStatus returns the Status field value
func (o *ModGlossaryViewGlossary200Response) GetStatus() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ModGlossaryViewGlossary200Response) GetStatusOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ModGlossaryViewGlossary200Response) SetStatus(v bool) {
	o.Status = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *ModGlossaryViewGlossary200Response) GetWarnings() []AuthEmailSignupUser200ResponseWarningsInner {
	if o == nil || IsNil(o.Warnings) {
		var ret []AuthEmailSignupUser200ResponseWarningsInner
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModGlossaryViewGlossary200Response) GetWarningsOk() ([]AuthEmailSignupUser200ResponseWarningsInner, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *ModGlossaryViewGlossary200Response) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []AuthEmailSignupUser200ResponseWarningsInner and assigns it to the Warnings field.
func (o *ModGlossaryViewGlossary200Response) SetWarnings(v []AuthEmailSignupUser200ResponseWarningsInner) {
	o.Warnings = v
}

func (o ModGlossaryViewGlossary200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModGlossaryViewGlossary200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["status"] = o.Status
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

func (o *ModGlossaryViewGlossary200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"status",
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

	varModGlossaryViewGlossary200Response := _ModGlossaryViewGlossary200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModGlossaryViewGlossary200Response)

	if err != nil {
		return err
	}

	*o = ModGlossaryViewGlossary200Response(varModGlossaryViewGlossary200Response)

	return err
}

type NullableModGlossaryViewGlossary200Response struct {
	value *ModGlossaryViewGlossary200Response
	isSet bool
}

func (v NullableModGlossaryViewGlossary200Response) Get() *ModGlossaryViewGlossary200Response {
	return v.value
}

func (v *NullableModGlossaryViewGlossary200Response) Set(val *ModGlossaryViewGlossary200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableModGlossaryViewGlossary200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableModGlossaryViewGlossary200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModGlossaryViewGlossary200Response(val *ModGlossaryViewGlossary200Response) *NullableModGlossaryViewGlossary200Response {
	return &NullableModGlossaryViewGlossary200Response{value: val, isSet: true}
}

func (v NullableModGlossaryViewGlossary200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModGlossaryViewGlossary200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


