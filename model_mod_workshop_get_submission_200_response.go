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

// checks if the ModWorkshopGetSubmission200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModWorkshopGetSubmission200Response{}

// ModWorkshopGetSubmission200Response struct for ModWorkshopGetSubmission200Response
type ModWorkshopGetSubmission200Response struct {
	Submission ModWorkshopGetSubmission200ResponseSubmission `json:"submission"`
	Warnings []AuthEmailSignupUser200ResponseWarningsInner `json:"warnings,omitempty"`
}

type _ModWorkshopGetSubmission200Response ModWorkshopGetSubmission200Response

// NewModWorkshopGetSubmission200Response instantiates a new ModWorkshopGetSubmission200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModWorkshopGetSubmission200Response(submission ModWorkshopGetSubmission200ResponseSubmission) *ModWorkshopGetSubmission200Response {
	this := ModWorkshopGetSubmission200Response{}
	this.Submission = submission
	return &this
}

// NewModWorkshopGetSubmission200ResponseWithDefaults instantiates a new ModWorkshopGetSubmission200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModWorkshopGetSubmission200ResponseWithDefaults() *ModWorkshopGetSubmission200Response {
	this := ModWorkshopGetSubmission200Response{}
	return &this
}

// GetSubmission returns the Submission field value
func (o *ModWorkshopGetSubmission200Response) GetSubmission() ModWorkshopGetSubmission200ResponseSubmission {
	if o == nil {
		var ret ModWorkshopGetSubmission200ResponseSubmission
		return ret
	}

	return o.Submission
}

// GetSubmissionOk returns a tuple with the Submission field value
// and a boolean to check if the value has been set.
func (o *ModWorkshopGetSubmission200Response) GetSubmissionOk() (*ModWorkshopGetSubmission200ResponseSubmission, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Submission, true
}

// SetSubmission sets field value
func (o *ModWorkshopGetSubmission200Response) SetSubmission(v ModWorkshopGetSubmission200ResponseSubmission) {
	o.Submission = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *ModWorkshopGetSubmission200Response) GetWarnings() []AuthEmailSignupUser200ResponseWarningsInner {
	if o == nil || IsNil(o.Warnings) {
		var ret []AuthEmailSignupUser200ResponseWarningsInner
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModWorkshopGetSubmission200Response) GetWarningsOk() ([]AuthEmailSignupUser200ResponseWarningsInner, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *ModWorkshopGetSubmission200Response) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []AuthEmailSignupUser200ResponseWarningsInner and assigns it to the Warnings field.
func (o *ModWorkshopGetSubmission200Response) SetWarnings(v []AuthEmailSignupUser200ResponseWarningsInner) {
	o.Warnings = v
}

func (o ModWorkshopGetSubmission200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModWorkshopGetSubmission200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["submission"] = o.Submission
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

func (o *ModWorkshopGetSubmission200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"submission",
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

	varModWorkshopGetSubmission200Response := _ModWorkshopGetSubmission200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModWorkshopGetSubmission200Response)

	if err != nil {
		return err
	}

	*o = ModWorkshopGetSubmission200Response(varModWorkshopGetSubmission200Response)

	return err
}

type NullableModWorkshopGetSubmission200Response struct {
	value *ModWorkshopGetSubmission200Response
	isSet bool
}

func (v NullableModWorkshopGetSubmission200Response) Get() *ModWorkshopGetSubmission200Response {
	return v.value
}

func (v *NullableModWorkshopGetSubmission200Response) Set(val *ModWorkshopGetSubmission200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableModWorkshopGetSubmission200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableModWorkshopGetSubmission200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModWorkshopGetSubmission200Response(val *ModWorkshopGetSubmission200Response) *NullableModWorkshopGetSubmission200Response {
	return &NullableModWorkshopGetSubmission200Response{value: val, isSet: true}
}

func (v NullableModWorkshopGetSubmission200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModWorkshopGetSubmission200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


