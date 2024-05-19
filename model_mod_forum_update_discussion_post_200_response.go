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

// checks if the ModForumUpdateDiscussionPost200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModForumUpdateDiscussionPost200Response{}

// ModForumUpdateDiscussionPost200Response struct for ModForumUpdateDiscussionPost200Response
type ModForumUpdateDiscussionPost200Response struct {
	// True if the post/discussion was updated, false otherwise.
	Status bool `json:"status"`
	Warnings []AuthEmailSignupUser200ResponseWarningsInner `json:"warnings,omitempty"`
}

type _ModForumUpdateDiscussionPost200Response ModForumUpdateDiscussionPost200Response

// NewModForumUpdateDiscussionPost200Response instantiates a new ModForumUpdateDiscussionPost200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModForumUpdateDiscussionPost200Response(status bool) *ModForumUpdateDiscussionPost200Response {
	this := ModForumUpdateDiscussionPost200Response{}
	this.Status = status
	return &this
}

// NewModForumUpdateDiscussionPost200ResponseWithDefaults instantiates a new ModForumUpdateDiscussionPost200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModForumUpdateDiscussionPost200ResponseWithDefaults() *ModForumUpdateDiscussionPost200Response {
	this := ModForumUpdateDiscussionPost200Response{}
	var status bool = null
	this.Status = status
	return &this
}

// GetStatus returns the Status field value
func (o *ModForumUpdateDiscussionPost200Response) GetStatus() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ModForumUpdateDiscussionPost200Response) GetStatusOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ModForumUpdateDiscussionPost200Response) SetStatus(v bool) {
	o.Status = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *ModForumUpdateDiscussionPost200Response) GetWarnings() []AuthEmailSignupUser200ResponseWarningsInner {
	if o == nil || IsNil(o.Warnings) {
		var ret []AuthEmailSignupUser200ResponseWarningsInner
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModForumUpdateDiscussionPost200Response) GetWarningsOk() ([]AuthEmailSignupUser200ResponseWarningsInner, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *ModForumUpdateDiscussionPost200Response) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []AuthEmailSignupUser200ResponseWarningsInner and assigns it to the Warnings field.
func (o *ModForumUpdateDiscussionPost200Response) SetWarnings(v []AuthEmailSignupUser200ResponseWarningsInner) {
	o.Warnings = v
}

func (o ModForumUpdateDiscussionPost200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModForumUpdateDiscussionPost200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["status"] = o.Status
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

func (o *ModForumUpdateDiscussionPost200Response) UnmarshalJSON(data []byte) (err error) {
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

	varModForumUpdateDiscussionPost200Response := _ModForumUpdateDiscussionPost200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModForumUpdateDiscussionPost200Response)

	if err != nil {
		return err
	}

	*o = ModForumUpdateDiscussionPost200Response(varModForumUpdateDiscussionPost200Response)

	return err
}

type NullableModForumUpdateDiscussionPost200Response struct {
	value *ModForumUpdateDiscussionPost200Response
	isSet bool
}

func (v NullableModForumUpdateDiscussionPost200Response) Get() *ModForumUpdateDiscussionPost200Response {
	return v.value
}

func (v *NullableModForumUpdateDiscussionPost200Response) Set(val *ModForumUpdateDiscussionPost200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableModForumUpdateDiscussionPost200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableModForumUpdateDiscussionPost200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModForumUpdateDiscussionPost200Response(val *ModForumUpdateDiscussionPost200Response) *NullableModForumUpdateDiscussionPost200Response {
	return &NullableModForumUpdateDiscussionPost200Response{value: val, isSet: true}
}

func (v NullableModForumUpdateDiscussionPost200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModForumUpdateDiscussionPost200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

