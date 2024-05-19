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

// checks if the ModH5pactivityGetAttempts200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModH5pactivityGetAttempts200Response{}

// ModH5pactivityGetAttempts200Response struct for ModH5pactivityGetAttempts200Response
type ModH5pactivityGetAttempts200Response struct {
	// Activity course module ID
	Activityid int32 `json:"activityid"`
	Usersattempts []ModH5pactivityGetAttempts200ResponseUsersattemptsInner `json:"usersattempts"`
	Warnings []AuthEmailSignupUser200ResponseWarningsInner `json:"warnings,omitempty"`
}

type _ModH5pactivityGetAttempts200Response ModH5pactivityGetAttempts200Response

// NewModH5pactivityGetAttempts200Response instantiates a new ModH5pactivityGetAttempts200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModH5pactivityGetAttempts200Response(activityid int32, usersattempts []ModH5pactivityGetAttempts200ResponseUsersattemptsInner) *ModH5pactivityGetAttempts200Response {
	this := ModH5pactivityGetAttempts200Response{}
	this.Activityid = activityid
	this.Usersattempts = usersattempts
	return &this
}

// NewModH5pactivityGetAttempts200ResponseWithDefaults instantiates a new ModH5pactivityGetAttempts200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModH5pactivityGetAttempts200ResponseWithDefaults() *ModH5pactivityGetAttempts200Response {
	this := ModH5pactivityGetAttempts200Response{}
	var activityid int32 = null
	this.Activityid = activityid
	return &this
}

// GetActivityid returns the Activityid field value
func (o *ModH5pactivityGetAttempts200Response) GetActivityid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Activityid
}

// GetActivityidOk returns a tuple with the Activityid field value
// and a boolean to check if the value has been set.
func (o *ModH5pactivityGetAttempts200Response) GetActivityidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Activityid, true
}

// SetActivityid sets field value
func (o *ModH5pactivityGetAttempts200Response) SetActivityid(v int32) {
	o.Activityid = v
}

// GetUsersattempts returns the Usersattempts field value
func (o *ModH5pactivityGetAttempts200Response) GetUsersattempts() []ModH5pactivityGetAttempts200ResponseUsersattemptsInner {
	if o == nil {
		var ret []ModH5pactivityGetAttempts200ResponseUsersattemptsInner
		return ret
	}

	return o.Usersattempts
}

// GetUsersattemptsOk returns a tuple with the Usersattempts field value
// and a boolean to check if the value has been set.
func (o *ModH5pactivityGetAttempts200Response) GetUsersattemptsOk() ([]ModH5pactivityGetAttempts200ResponseUsersattemptsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Usersattempts, true
}

// SetUsersattempts sets field value
func (o *ModH5pactivityGetAttempts200Response) SetUsersattempts(v []ModH5pactivityGetAttempts200ResponseUsersattemptsInner) {
	o.Usersattempts = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *ModH5pactivityGetAttempts200Response) GetWarnings() []AuthEmailSignupUser200ResponseWarningsInner {
	if o == nil || IsNil(o.Warnings) {
		var ret []AuthEmailSignupUser200ResponseWarningsInner
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModH5pactivityGetAttempts200Response) GetWarningsOk() ([]AuthEmailSignupUser200ResponseWarningsInner, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *ModH5pactivityGetAttempts200Response) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []AuthEmailSignupUser200ResponseWarningsInner and assigns it to the Warnings field.
func (o *ModH5pactivityGetAttempts200Response) SetWarnings(v []AuthEmailSignupUser200ResponseWarningsInner) {
	o.Warnings = v
}

func (o ModH5pactivityGetAttempts200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModH5pactivityGetAttempts200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["activityid"] = o.Activityid
	toSerialize["usersattempts"] = o.Usersattempts
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

func (o *ModH5pactivityGetAttempts200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"activityid",
		"usersattempts",
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

	varModH5pactivityGetAttempts200Response := _ModH5pactivityGetAttempts200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModH5pactivityGetAttempts200Response)

	if err != nil {
		return err
	}

	*o = ModH5pactivityGetAttempts200Response(varModH5pactivityGetAttempts200Response)

	return err
}

type NullableModH5pactivityGetAttempts200Response struct {
	value *ModH5pactivityGetAttempts200Response
	isSet bool
}

func (v NullableModH5pactivityGetAttempts200Response) Get() *ModH5pactivityGetAttempts200Response {
	return v.value
}

func (v *NullableModH5pactivityGetAttempts200Response) Set(val *ModH5pactivityGetAttempts200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableModH5pactivityGetAttempts200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableModH5pactivityGetAttempts200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModH5pactivityGetAttempts200Response(val *ModH5pactivityGetAttempts200Response) *NullableModH5pactivityGetAttempts200Response {
	return &NullableModH5pactivityGetAttempts200Response{value: val, isSet: true}
}

func (v NullableModH5pactivityGetAttempts200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModH5pactivityGetAttempts200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

