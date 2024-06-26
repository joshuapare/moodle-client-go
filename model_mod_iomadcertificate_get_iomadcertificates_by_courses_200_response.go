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

// checks if the ModIomadcertificateGetIomadcertificatesByCourses200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModIomadcertificateGetIomadcertificatesByCourses200Response{}

// ModIomadcertificateGetIomadcertificatesByCourses200Response struct for ModIomadcertificateGetIomadcertificatesByCourses200Response
type ModIomadcertificateGetIomadcertificatesByCourses200Response struct {
	Iomadcertificates []ModIomadcertificateGetIomadcertificatesByCourses200ResponseIomadcertificatesInner `json:"iomadcertificates"`
	Warnings []AuthEmailSignupUser200ResponseWarningsInner `json:"warnings,omitempty"`
}

type _ModIomadcertificateGetIomadcertificatesByCourses200Response ModIomadcertificateGetIomadcertificatesByCourses200Response

// NewModIomadcertificateGetIomadcertificatesByCourses200Response instantiates a new ModIomadcertificateGetIomadcertificatesByCourses200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModIomadcertificateGetIomadcertificatesByCourses200Response(iomadcertificates []ModIomadcertificateGetIomadcertificatesByCourses200ResponseIomadcertificatesInner) *ModIomadcertificateGetIomadcertificatesByCourses200Response {
	this := ModIomadcertificateGetIomadcertificatesByCourses200Response{}
	this.Iomadcertificates = iomadcertificates
	return &this
}

// NewModIomadcertificateGetIomadcertificatesByCourses200ResponseWithDefaults instantiates a new ModIomadcertificateGetIomadcertificatesByCourses200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModIomadcertificateGetIomadcertificatesByCourses200ResponseWithDefaults() *ModIomadcertificateGetIomadcertificatesByCourses200Response {
	this := ModIomadcertificateGetIomadcertificatesByCourses200Response{}
	return &this
}

// GetIomadcertificates returns the Iomadcertificates field value
func (o *ModIomadcertificateGetIomadcertificatesByCourses200Response) GetIomadcertificates() []ModIomadcertificateGetIomadcertificatesByCourses200ResponseIomadcertificatesInner {
	if o == nil {
		var ret []ModIomadcertificateGetIomadcertificatesByCourses200ResponseIomadcertificatesInner
		return ret
	}

	return o.Iomadcertificates
}

// GetIomadcertificatesOk returns a tuple with the Iomadcertificates field value
// and a boolean to check if the value has been set.
func (o *ModIomadcertificateGetIomadcertificatesByCourses200Response) GetIomadcertificatesOk() ([]ModIomadcertificateGetIomadcertificatesByCourses200ResponseIomadcertificatesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Iomadcertificates, true
}

// SetIomadcertificates sets field value
func (o *ModIomadcertificateGetIomadcertificatesByCourses200Response) SetIomadcertificates(v []ModIomadcertificateGetIomadcertificatesByCourses200ResponseIomadcertificatesInner) {
	o.Iomadcertificates = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *ModIomadcertificateGetIomadcertificatesByCourses200Response) GetWarnings() []AuthEmailSignupUser200ResponseWarningsInner {
	if o == nil || IsNil(o.Warnings) {
		var ret []AuthEmailSignupUser200ResponseWarningsInner
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModIomadcertificateGetIomadcertificatesByCourses200Response) GetWarningsOk() ([]AuthEmailSignupUser200ResponseWarningsInner, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *ModIomadcertificateGetIomadcertificatesByCourses200Response) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []AuthEmailSignupUser200ResponseWarningsInner and assigns it to the Warnings field.
func (o *ModIomadcertificateGetIomadcertificatesByCourses200Response) SetWarnings(v []AuthEmailSignupUser200ResponseWarningsInner) {
	o.Warnings = v
}

func (o ModIomadcertificateGetIomadcertificatesByCourses200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModIomadcertificateGetIomadcertificatesByCourses200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["iomadcertificates"] = o.Iomadcertificates
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

func (o *ModIomadcertificateGetIomadcertificatesByCourses200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"iomadcertificates",
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

	varModIomadcertificateGetIomadcertificatesByCourses200Response := _ModIomadcertificateGetIomadcertificatesByCourses200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModIomadcertificateGetIomadcertificatesByCourses200Response)

	if err != nil {
		return err
	}

	*o = ModIomadcertificateGetIomadcertificatesByCourses200Response(varModIomadcertificateGetIomadcertificatesByCourses200Response)

	return err
}

type NullableModIomadcertificateGetIomadcertificatesByCourses200Response struct {
	value *ModIomadcertificateGetIomadcertificatesByCourses200Response
	isSet bool
}

func (v NullableModIomadcertificateGetIomadcertificatesByCourses200Response) Get() *ModIomadcertificateGetIomadcertificatesByCourses200Response {
	return v.value
}

func (v *NullableModIomadcertificateGetIomadcertificatesByCourses200Response) Set(val *ModIomadcertificateGetIomadcertificatesByCourses200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableModIomadcertificateGetIomadcertificatesByCourses200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableModIomadcertificateGetIomadcertificatesByCourses200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModIomadcertificateGetIomadcertificatesByCourses200Response(val *ModIomadcertificateGetIomadcertificatesByCourses200Response) *NullableModIomadcertificateGetIomadcertificatesByCourses200Response {
	return &NullableModIomadcertificateGetIomadcertificatesByCourses200Response{value: val, isSet: true}
}

func (v NullableModIomadcertificateGetIomadcertificatesByCourses200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModIomadcertificateGetIomadcertificatesByCourses200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


