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

// checks if the ModQuizGetAttemptData200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModQuizGetAttemptData200Response{}

// ModQuizGetAttemptData200Response struct for ModQuizGetAttemptData200Response
type ModQuizGetAttemptData200Response struct {
	Attempt ModQuizGetAttemptData200ResponseAttempt `json:"attempt"`
	Messages []map[string]interface{} `json:"messages"`
	// next page number
	Nextpage int32 `json:"nextpage"`
	Questions []ModQuizGetAttemptData200ResponseQuestionsInner `json:"questions"`
	Warnings []AuthEmailSignupUser200ResponseWarningsInner `json:"warnings,omitempty"`
}

type _ModQuizGetAttemptData200Response ModQuizGetAttemptData200Response

// NewModQuizGetAttemptData200Response instantiates a new ModQuizGetAttemptData200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModQuizGetAttemptData200Response(attempt ModQuizGetAttemptData200ResponseAttempt, messages []map[string]interface{}, nextpage int32, questions []ModQuizGetAttemptData200ResponseQuestionsInner) *ModQuizGetAttemptData200Response {
	this := ModQuizGetAttemptData200Response{}
	this.Attempt = attempt
	this.Messages = messages
	this.Nextpage = nextpage
	this.Questions = questions
	return &this
}

// NewModQuizGetAttemptData200ResponseWithDefaults instantiates a new ModQuizGetAttemptData200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModQuizGetAttemptData200ResponseWithDefaults() *ModQuizGetAttemptData200Response {
	this := ModQuizGetAttemptData200Response{}
	var nextpage int32 = null
	this.Nextpage = nextpage
	return &this
}

// GetAttempt returns the Attempt field value
func (o *ModQuizGetAttemptData200Response) GetAttempt() ModQuizGetAttemptData200ResponseAttempt {
	if o == nil {
		var ret ModQuizGetAttemptData200ResponseAttempt
		return ret
	}

	return o.Attempt
}

// GetAttemptOk returns a tuple with the Attempt field value
// and a boolean to check if the value has been set.
func (o *ModQuizGetAttemptData200Response) GetAttemptOk() (*ModQuizGetAttemptData200ResponseAttempt, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attempt, true
}

// SetAttempt sets field value
func (o *ModQuizGetAttemptData200Response) SetAttempt(v ModQuizGetAttemptData200ResponseAttempt) {
	o.Attempt = v
}

// GetMessages returns the Messages field value
func (o *ModQuizGetAttemptData200Response) GetMessages() []map[string]interface{} {
	if o == nil {
		var ret []map[string]interface{}
		return ret
	}

	return o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value
// and a boolean to check if the value has been set.
func (o *ModQuizGetAttemptData200Response) GetMessagesOk() ([]map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Messages, true
}

// SetMessages sets field value
func (o *ModQuizGetAttemptData200Response) SetMessages(v []map[string]interface{}) {
	o.Messages = v
}

// GetNextpage returns the Nextpage field value
func (o *ModQuizGetAttemptData200Response) GetNextpage() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Nextpage
}

// GetNextpageOk returns a tuple with the Nextpage field value
// and a boolean to check if the value has been set.
func (o *ModQuizGetAttemptData200Response) GetNextpageOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Nextpage, true
}

// SetNextpage sets field value
func (o *ModQuizGetAttemptData200Response) SetNextpage(v int32) {
	o.Nextpage = v
}

// GetQuestions returns the Questions field value
func (o *ModQuizGetAttemptData200Response) GetQuestions() []ModQuizGetAttemptData200ResponseQuestionsInner {
	if o == nil {
		var ret []ModQuizGetAttemptData200ResponseQuestionsInner
		return ret
	}

	return o.Questions
}

// GetQuestionsOk returns a tuple with the Questions field value
// and a boolean to check if the value has been set.
func (o *ModQuizGetAttemptData200Response) GetQuestionsOk() ([]ModQuizGetAttemptData200ResponseQuestionsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Questions, true
}

// SetQuestions sets field value
func (o *ModQuizGetAttemptData200Response) SetQuestions(v []ModQuizGetAttemptData200ResponseQuestionsInner) {
	o.Questions = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *ModQuizGetAttemptData200Response) GetWarnings() []AuthEmailSignupUser200ResponseWarningsInner {
	if o == nil || IsNil(o.Warnings) {
		var ret []AuthEmailSignupUser200ResponseWarningsInner
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModQuizGetAttemptData200Response) GetWarningsOk() ([]AuthEmailSignupUser200ResponseWarningsInner, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *ModQuizGetAttemptData200Response) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []AuthEmailSignupUser200ResponseWarningsInner and assigns it to the Warnings field.
func (o *ModQuizGetAttemptData200Response) SetWarnings(v []AuthEmailSignupUser200ResponseWarningsInner) {
	o.Warnings = v
}

func (o ModQuizGetAttemptData200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModQuizGetAttemptData200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["attempt"] = o.Attempt
	toSerialize["messages"] = o.Messages
	toSerialize["nextpage"] = o.Nextpage
	toSerialize["questions"] = o.Questions
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

func (o *ModQuizGetAttemptData200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"attempt",
		"messages",
		"nextpage",
		"questions",
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

	varModQuizGetAttemptData200Response := _ModQuizGetAttemptData200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModQuizGetAttemptData200Response)

	if err != nil {
		return err
	}

	*o = ModQuizGetAttemptData200Response(varModQuizGetAttemptData200Response)

	return err
}

type NullableModQuizGetAttemptData200Response struct {
	value *ModQuizGetAttemptData200Response
	isSet bool
}

func (v NullableModQuizGetAttemptData200Response) Get() *ModQuizGetAttemptData200Response {
	return v.value
}

func (v *NullableModQuizGetAttemptData200Response) Set(val *ModQuizGetAttemptData200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableModQuizGetAttemptData200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableModQuizGetAttemptData200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModQuizGetAttemptData200Response(val *ModQuizGetAttemptData200Response) *NullableModQuizGetAttemptData200Response {
	return &NullableModQuizGetAttemptData200Response{value: val, isSet: true}
}

func (v NullableModQuizGetAttemptData200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModQuizGetAttemptData200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


