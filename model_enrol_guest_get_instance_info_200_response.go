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

// checks if the EnrolGuestGetInstanceInfo200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EnrolGuestGetInstanceInfo200Response{}

// EnrolGuestGetInstanceInfo200Response struct for EnrolGuestGetInstanceInfo200Response
type EnrolGuestGetInstanceInfo200Response struct {
	Instanceinfo EnrolGuestGetInstanceInfo200ResponseInstanceinfo `json:"instanceinfo"`
	Warnings []AuthEmailSignupUser200ResponseWarningsInner `json:"warnings,omitempty"`
}

type _EnrolGuestGetInstanceInfo200Response EnrolGuestGetInstanceInfo200Response

// NewEnrolGuestGetInstanceInfo200Response instantiates a new EnrolGuestGetInstanceInfo200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnrolGuestGetInstanceInfo200Response(instanceinfo EnrolGuestGetInstanceInfo200ResponseInstanceinfo) *EnrolGuestGetInstanceInfo200Response {
	this := EnrolGuestGetInstanceInfo200Response{}
	this.Instanceinfo = instanceinfo
	return &this
}

// NewEnrolGuestGetInstanceInfo200ResponseWithDefaults instantiates a new EnrolGuestGetInstanceInfo200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnrolGuestGetInstanceInfo200ResponseWithDefaults() *EnrolGuestGetInstanceInfo200Response {
	this := EnrolGuestGetInstanceInfo200Response{}
	return &this
}

// GetInstanceinfo returns the Instanceinfo field value
func (o *EnrolGuestGetInstanceInfo200Response) GetInstanceinfo() EnrolGuestGetInstanceInfo200ResponseInstanceinfo {
	if o == nil {
		var ret EnrolGuestGetInstanceInfo200ResponseInstanceinfo
		return ret
	}

	return o.Instanceinfo
}

// GetInstanceinfoOk returns a tuple with the Instanceinfo field value
// and a boolean to check if the value has been set.
func (o *EnrolGuestGetInstanceInfo200Response) GetInstanceinfoOk() (*EnrolGuestGetInstanceInfo200ResponseInstanceinfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Instanceinfo, true
}

// SetInstanceinfo sets field value
func (o *EnrolGuestGetInstanceInfo200Response) SetInstanceinfo(v EnrolGuestGetInstanceInfo200ResponseInstanceinfo) {
	o.Instanceinfo = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *EnrolGuestGetInstanceInfo200Response) GetWarnings() []AuthEmailSignupUser200ResponseWarningsInner {
	if o == nil || IsNil(o.Warnings) {
		var ret []AuthEmailSignupUser200ResponseWarningsInner
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnrolGuestGetInstanceInfo200Response) GetWarningsOk() ([]AuthEmailSignupUser200ResponseWarningsInner, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *EnrolGuestGetInstanceInfo200Response) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []AuthEmailSignupUser200ResponseWarningsInner and assigns it to the Warnings field.
func (o *EnrolGuestGetInstanceInfo200Response) SetWarnings(v []AuthEmailSignupUser200ResponseWarningsInner) {
	o.Warnings = v
}

func (o EnrolGuestGetInstanceInfo200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnrolGuestGetInstanceInfo200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["instanceinfo"] = o.Instanceinfo
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

func (o *EnrolGuestGetInstanceInfo200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"instanceinfo",
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

	varEnrolGuestGetInstanceInfo200Response := _EnrolGuestGetInstanceInfo200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEnrolGuestGetInstanceInfo200Response)

	if err != nil {
		return err
	}

	*o = EnrolGuestGetInstanceInfo200Response(varEnrolGuestGetInstanceInfo200Response)

	return err
}

type NullableEnrolGuestGetInstanceInfo200Response struct {
	value *EnrolGuestGetInstanceInfo200Response
	isSet bool
}

func (v NullableEnrolGuestGetInstanceInfo200Response) Get() *EnrolGuestGetInstanceInfo200Response {
	return v.value
}

func (v *NullableEnrolGuestGetInstanceInfo200Response) Set(val *EnrolGuestGetInstanceInfo200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableEnrolGuestGetInstanceInfo200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableEnrolGuestGetInstanceInfo200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnrolGuestGetInstanceInfo200Response(val *EnrolGuestGetInstanceInfo200Response) *NullableEnrolGuestGetInstanceInfo200Response {
	return &NullableEnrolGuestGetInstanceInfo200Response{value: val, isSet: true}
}

func (v NullableEnrolGuestGetInstanceInfo200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnrolGuestGetInstanceInfo200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


