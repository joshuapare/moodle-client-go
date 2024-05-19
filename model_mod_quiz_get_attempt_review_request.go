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

// checks if the ModQuizGetAttemptReviewRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModQuizGetAttemptReviewRequest{}

// ModQuizGetAttemptReviewRequest struct for ModQuizGetAttemptReviewRequest
type ModQuizGetAttemptReviewRequest struct {
	// attempt id
	Attemptid int32 `json:"attemptid"`
	// page number, empty for all the questions in all the pages
	Page *int32 `json:"page,omitempty"`
}

type _ModQuizGetAttemptReviewRequest ModQuizGetAttemptReviewRequest

// NewModQuizGetAttemptReviewRequest instantiates a new ModQuizGetAttemptReviewRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModQuizGetAttemptReviewRequest(attemptid int32) *ModQuizGetAttemptReviewRequest {
	this := ModQuizGetAttemptReviewRequest{}
	this.Attemptid = attemptid
	var page int32 = -1
	this.Page = &page
	return &this
}

// NewModQuizGetAttemptReviewRequestWithDefaults instantiates a new ModQuizGetAttemptReviewRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModQuizGetAttemptReviewRequestWithDefaults() *ModQuizGetAttemptReviewRequest {
	this := ModQuizGetAttemptReviewRequest{}
	var page int32 = -1
	this.Page = &page
	return &this
}

// GetAttemptid returns the Attemptid field value
func (o *ModQuizGetAttemptReviewRequest) GetAttemptid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Attemptid
}

// GetAttemptidOk returns a tuple with the Attemptid field value
// and a boolean to check if the value has been set.
func (o *ModQuizGetAttemptReviewRequest) GetAttemptidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attemptid, true
}

// SetAttemptid sets field value
func (o *ModQuizGetAttemptReviewRequest) SetAttemptid(v int32) {
	o.Attemptid = v
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *ModQuizGetAttemptReviewRequest) GetPage() int32 {
	if o == nil || IsNil(o.Page) {
		var ret int32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModQuizGetAttemptReviewRequest) GetPageOk() (*int32, bool) {
	if o == nil || IsNil(o.Page) {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *ModQuizGetAttemptReviewRequest) HasPage() bool {
	if o != nil && !IsNil(o.Page) {
		return true
	}

	return false
}

// SetPage gets a reference to the given int32 and assigns it to the Page field.
func (o *ModQuizGetAttemptReviewRequest) SetPage(v int32) {
	o.Page = &v
}

func (o ModQuizGetAttemptReviewRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModQuizGetAttemptReviewRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["attemptid"] = o.Attemptid
	if !IsNil(o.Page) {
		toSerialize["page"] = o.Page
	}
	return toSerialize, nil
}

func (o *ModQuizGetAttemptReviewRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"attemptid",
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

	varModQuizGetAttemptReviewRequest := _ModQuizGetAttemptReviewRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModQuizGetAttemptReviewRequest)

	if err != nil {
		return err
	}

	*o = ModQuizGetAttemptReviewRequest(varModQuizGetAttemptReviewRequest)

	return err
}

type NullableModQuizGetAttemptReviewRequest struct {
	value *ModQuizGetAttemptReviewRequest
	isSet bool
}

func (v NullableModQuizGetAttemptReviewRequest) Get() *ModQuizGetAttemptReviewRequest {
	return v.value
}

func (v *NullableModQuizGetAttemptReviewRequest) Set(val *ModQuizGetAttemptReviewRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModQuizGetAttemptReviewRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModQuizGetAttemptReviewRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModQuizGetAttemptReviewRequest(val *ModQuizGetAttemptReviewRequest) *NullableModQuizGetAttemptReviewRequest {
	return &NullableModQuizGetAttemptReviewRequest{value: val, isSet: true}
}

func (v NullableModQuizGetAttemptReviewRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModQuizGetAttemptReviewRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


