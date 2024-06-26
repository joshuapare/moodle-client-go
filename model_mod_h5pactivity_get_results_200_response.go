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

// checks if the ModH5pactivityGetResults200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModH5pactivityGetResults200Response{}

// ModH5pactivityGetResults200Response struct for ModH5pactivityGetResults200Response
type ModH5pactivityGetResults200Response struct {
	// Activity course module ID
	Activityid int32 `json:"activityid"`
	Attempts []ModH5pactivityGetResults200ResponseAttemptsInner `json:"attempts"`
	Warnings []AuthEmailSignupUser200ResponseWarningsInner `json:"warnings,omitempty"`
}

type _ModH5pactivityGetResults200Response ModH5pactivityGetResults200Response

// NewModH5pactivityGetResults200Response instantiates a new ModH5pactivityGetResults200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModH5pactivityGetResults200Response(activityid int32, attempts []ModH5pactivityGetResults200ResponseAttemptsInner) *ModH5pactivityGetResults200Response {
	this := ModH5pactivityGetResults200Response{}
	this.Activityid = activityid
	this.Attempts = attempts
	return &this
}

// NewModH5pactivityGetResults200ResponseWithDefaults instantiates a new ModH5pactivityGetResults200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModH5pactivityGetResults200ResponseWithDefaults() *ModH5pactivityGetResults200Response {
	this := ModH5pactivityGetResults200Response{}
	return &this
}

// GetActivityid returns the Activityid field value
func (o *ModH5pactivityGetResults200Response) GetActivityid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Activityid
}

// GetActivityidOk returns a tuple with the Activityid field value
// and a boolean to check if the value has been set.
func (o *ModH5pactivityGetResults200Response) GetActivityidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Activityid, true
}

// SetActivityid sets field value
func (o *ModH5pactivityGetResults200Response) SetActivityid(v int32) {
	o.Activityid = v
}

// GetAttempts returns the Attempts field value
func (o *ModH5pactivityGetResults200Response) GetAttempts() []ModH5pactivityGetResults200ResponseAttemptsInner {
	if o == nil {
		var ret []ModH5pactivityGetResults200ResponseAttemptsInner
		return ret
	}

	return o.Attempts
}

// GetAttemptsOk returns a tuple with the Attempts field value
// and a boolean to check if the value has been set.
func (o *ModH5pactivityGetResults200Response) GetAttemptsOk() ([]ModH5pactivityGetResults200ResponseAttemptsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Attempts, true
}

// SetAttempts sets field value
func (o *ModH5pactivityGetResults200Response) SetAttempts(v []ModH5pactivityGetResults200ResponseAttemptsInner) {
	o.Attempts = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *ModH5pactivityGetResults200Response) GetWarnings() []AuthEmailSignupUser200ResponseWarningsInner {
	if o == nil || IsNil(o.Warnings) {
		var ret []AuthEmailSignupUser200ResponseWarningsInner
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModH5pactivityGetResults200Response) GetWarningsOk() ([]AuthEmailSignupUser200ResponseWarningsInner, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *ModH5pactivityGetResults200Response) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []AuthEmailSignupUser200ResponseWarningsInner and assigns it to the Warnings field.
func (o *ModH5pactivityGetResults200Response) SetWarnings(v []AuthEmailSignupUser200ResponseWarningsInner) {
	o.Warnings = v
}

func (o ModH5pactivityGetResults200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModH5pactivityGetResults200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["activityid"] = o.Activityid
	toSerialize["attempts"] = o.Attempts
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

func (o *ModH5pactivityGetResults200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"activityid",
		"attempts",
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

	varModH5pactivityGetResults200Response := _ModH5pactivityGetResults200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModH5pactivityGetResults200Response)

	if err != nil {
		return err
	}

	*o = ModH5pactivityGetResults200Response(varModH5pactivityGetResults200Response)

	return err
}

type NullableModH5pactivityGetResults200Response struct {
	value *ModH5pactivityGetResults200Response
	isSet bool
}

func (v NullableModH5pactivityGetResults200Response) Get() *ModH5pactivityGetResults200Response {
	return v.value
}

func (v *NullableModH5pactivityGetResults200Response) Set(val *ModH5pactivityGetResults200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableModH5pactivityGetResults200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableModH5pactivityGetResults200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModH5pactivityGetResults200Response(val *ModH5pactivityGetResults200Response) *NullableModH5pactivityGetResults200Response {
	return &NullableModH5pactivityGetResults200Response{value: val, isSet: true}
}

func (v NullableModH5pactivityGetResults200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModH5pactivityGetResults200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


