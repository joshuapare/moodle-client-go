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

// checks if the ModFeedbackGetUnfinishedResponses200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModFeedbackGetUnfinishedResponses200Response{}

// ModFeedbackGetUnfinishedResponses200Response struct for ModFeedbackGetUnfinishedResponses200Response
type ModFeedbackGetUnfinishedResponses200Response struct {
	Responses []ModFeedbackGetUnfinishedResponses200ResponseResponsesInner `json:"responses"`
	Warnings []AuthEmailSignupUser200ResponseWarningsInner `json:"warnings,omitempty"`
}

type _ModFeedbackGetUnfinishedResponses200Response ModFeedbackGetUnfinishedResponses200Response

// NewModFeedbackGetUnfinishedResponses200Response instantiates a new ModFeedbackGetUnfinishedResponses200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModFeedbackGetUnfinishedResponses200Response(responses []ModFeedbackGetUnfinishedResponses200ResponseResponsesInner) *ModFeedbackGetUnfinishedResponses200Response {
	this := ModFeedbackGetUnfinishedResponses200Response{}
	this.Responses = responses
	return &this
}

// NewModFeedbackGetUnfinishedResponses200ResponseWithDefaults instantiates a new ModFeedbackGetUnfinishedResponses200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModFeedbackGetUnfinishedResponses200ResponseWithDefaults() *ModFeedbackGetUnfinishedResponses200Response {
	this := ModFeedbackGetUnfinishedResponses200Response{}
	return &this
}

// GetResponses returns the Responses field value
func (o *ModFeedbackGetUnfinishedResponses200Response) GetResponses() []ModFeedbackGetUnfinishedResponses200ResponseResponsesInner {
	if o == nil {
		var ret []ModFeedbackGetUnfinishedResponses200ResponseResponsesInner
		return ret
	}

	return o.Responses
}

// GetResponsesOk returns a tuple with the Responses field value
// and a boolean to check if the value has been set.
func (o *ModFeedbackGetUnfinishedResponses200Response) GetResponsesOk() ([]ModFeedbackGetUnfinishedResponses200ResponseResponsesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Responses, true
}

// SetResponses sets field value
func (o *ModFeedbackGetUnfinishedResponses200Response) SetResponses(v []ModFeedbackGetUnfinishedResponses200ResponseResponsesInner) {
	o.Responses = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *ModFeedbackGetUnfinishedResponses200Response) GetWarnings() []AuthEmailSignupUser200ResponseWarningsInner {
	if o == nil || IsNil(o.Warnings) {
		var ret []AuthEmailSignupUser200ResponseWarningsInner
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModFeedbackGetUnfinishedResponses200Response) GetWarningsOk() ([]AuthEmailSignupUser200ResponseWarningsInner, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *ModFeedbackGetUnfinishedResponses200Response) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []AuthEmailSignupUser200ResponseWarningsInner and assigns it to the Warnings field.
func (o *ModFeedbackGetUnfinishedResponses200Response) SetWarnings(v []AuthEmailSignupUser200ResponseWarningsInner) {
	o.Warnings = v
}

func (o ModFeedbackGetUnfinishedResponses200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModFeedbackGetUnfinishedResponses200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["responses"] = o.Responses
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

func (o *ModFeedbackGetUnfinishedResponses200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"responses",
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

	varModFeedbackGetUnfinishedResponses200Response := _ModFeedbackGetUnfinishedResponses200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModFeedbackGetUnfinishedResponses200Response)

	if err != nil {
		return err
	}

	*o = ModFeedbackGetUnfinishedResponses200Response(varModFeedbackGetUnfinishedResponses200Response)

	return err
}

type NullableModFeedbackGetUnfinishedResponses200Response struct {
	value *ModFeedbackGetUnfinishedResponses200Response
	isSet bool
}

func (v NullableModFeedbackGetUnfinishedResponses200Response) Get() *ModFeedbackGetUnfinishedResponses200Response {
	return v.value
}

func (v *NullableModFeedbackGetUnfinishedResponses200Response) Set(val *ModFeedbackGetUnfinishedResponses200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableModFeedbackGetUnfinishedResponses200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableModFeedbackGetUnfinishedResponses200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModFeedbackGetUnfinishedResponses200Response(val *ModFeedbackGetUnfinishedResponses200Response) *NullableModFeedbackGetUnfinishedResponses200Response {
	return &NullableModFeedbackGetUnfinishedResponses200Response{value: val, isSet: true}
}

func (v NullableModFeedbackGetUnfinishedResponses200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModFeedbackGetUnfinishedResponses200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


