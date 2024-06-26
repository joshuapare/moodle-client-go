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

// checks if the ModQuizGetQuizFeedbackForGrade200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModQuizGetQuizFeedbackForGrade200Response{}

// ModQuizGetQuizFeedbackForGrade200Response struct for ModQuizGetQuizFeedbackForGrade200Response
type ModQuizGetQuizFeedbackForGrade200Response struct {
	Feedbackinlinefiles []CoreBlockGetDashboardBlocks200ResponseBlocksInnerContentsFilesInner `json:"feedbackinlinefiles,omitempty"`
	// the comment that corresponds to this grade (empty for none)
	Feedbacktext string `json:"feedbacktext"`
	// feedbacktext format (1 = HTML, 0 = MOODLE, 2 = PLAIN, or 4 = MARKDOWN)
	Feedbacktextformat *int32 `json:"feedbacktextformat,omitempty"`
	Warnings []AuthEmailSignupUser200ResponseWarningsInner `json:"warnings,omitempty"`
}

type _ModQuizGetQuizFeedbackForGrade200Response ModQuizGetQuizFeedbackForGrade200Response

// NewModQuizGetQuizFeedbackForGrade200Response instantiates a new ModQuizGetQuizFeedbackForGrade200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModQuizGetQuizFeedbackForGrade200Response(feedbacktext string) *ModQuizGetQuizFeedbackForGrade200Response {
	this := ModQuizGetQuizFeedbackForGrade200Response{}
	this.Feedbacktext = feedbacktext
	var feedbacktextformat int32 = null
	this.Feedbacktextformat = &feedbacktextformat
	return &this
}

// NewModQuizGetQuizFeedbackForGrade200ResponseWithDefaults instantiates a new ModQuizGetQuizFeedbackForGrade200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModQuizGetQuizFeedbackForGrade200ResponseWithDefaults() *ModQuizGetQuizFeedbackForGrade200Response {
	this := ModQuizGetQuizFeedbackForGrade200Response{}
	var feedbacktext string = "null"
	this.Feedbacktext = feedbacktext
	var feedbacktextformat int32 = null
	this.Feedbacktextformat = &feedbacktextformat
	return &this
}

// GetFeedbackinlinefiles returns the Feedbackinlinefiles field value if set, zero value otherwise.
func (o *ModQuizGetQuizFeedbackForGrade200Response) GetFeedbackinlinefiles() []CoreBlockGetDashboardBlocks200ResponseBlocksInnerContentsFilesInner {
	if o == nil || IsNil(o.Feedbackinlinefiles) {
		var ret []CoreBlockGetDashboardBlocks200ResponseBlocksInnerContentsFilesInner
		return ret
	}
	return o.Feedbackinlinefiles
}

// GetFeedbackinlinefilesOk returns a tuple with the Feedbackinlinefiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModQuizGetQuizFeedbackForGrade200Response) GetFeedbackinlinefilesOk() ([]CoreBlockGetDashboardBlocks200ResponseBlocksInnerContentsFilesInner, bool) {
	if o == nil || IsNil(o.Feedbackinlinefiles) {
		return nil, false
	}
	return o.Feedbackinlinefiles, true
}

// HasFeedbackinlinefiles returns a boolean if a field has been set.
func (o *ModQuizGetQuizFeedbackForGrade200Response) HasFeedbackinlinefiles() bool {
	if o != nil && !IsNil(o.Feedbackinlinefiles) {
		return true
	}

	return false
}

// SetFeedbackinlinefiles gets a reference to the given []CoreBlockGetDashboardBlocks200ResponseBlocksInnerContentsFilesInner and assigns it to the Feedbackinlinefiles field.
func (o *ModQuizGetQuizFeedbackForGrade200Response) SetFeedbackinlinefiles(v []CoreBlockGetDashboardBlocks200ResponseBlocksInnerContentsFilesInner) {
	o.Feedbackinlinefiles = v
}

// GetFeedbacktext returns the Feedbacktext field value
func (o *ModQuizGetQuizFeedbackForGrade200Response) GetFeedbacktext() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Feedbacktext
}

// GetFeedbacktextOk returns a tuple with the Feedbacktext field value
// and a boolean to check if the value has been set.
func (o *ModQuizGetQuizFeedbackForGrade200Response) GetFeedbacktextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Feedbacktext, true
}

// SetFeedbacktext sets field value
func (o *ModQuizGetQuizFeedbackForGrade200Response) SetFeedbacktext(v string) {
	o.Feedbacktext = v
}

// GetFeedbacktextformat returns the Feedbacktextformat field value if set, zero value otherwise.
func (o *ModQuizGetQuizFeedbackForGrade200Response) GetFeedbacktextformat() int32 {
	if o == nil || IsNil(o.Feedbacktextformat) {
		var ret int32
		return ret
	}
	return *o.Feedbacktextformat
}

// GetFeedbacktextformatOk returns a tuple with the Feedbacktextformat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModQuizGetQuizFeedbackForGrade200Response) GetFeedbacktextformatOk() (*int32, bool) {
	if o == nil || IsNil(o.Feedbacktextformat) {
		return nil, false
	}
	return o.Feedbacktextformat, true
}

// HasFeedbacktextformat returns a boolean if a field has been set.
func (o *ModQuizGetQuizFeedbackForGrade200Response) HasFeedbacktextformat() bool {
	if o != nil && !IsNil(o.Feedbacktextformat) {
		return true
	}

	return false
}

// SetFeedbacktextformat gets a reference to the given int32 and assigns it to the Feedbacktextformat field.
func (o *ModQuizGetQuizFeedbackForGrade200Response) SetFeedbacktextformat(v int32) {
	o.Feedbacktextformat = &v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *ModQuizGetQuizFeedbackForGrade200Response) GetWarnings() []AuthEmailSignupUser200ResponseWarningsInner {
	if o == nil || IsNil(o.Warnings) {
		var ret []AuthEmailSignupUser200ResponseWarningsInner
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModQuizGetQuizFeedbackForGrade200Response) GetWarningsOk() ([]AuthEmailSignupUser200ResponseWarningsInner, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *ModQuizGetQuizFeedbackForGrade200Response) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []AuthEmailSignupUser200ResponseWarningsInner and assigns it to the Warnings field.
func (o *ModQuizGetQuizFeedbackForGrade200Response) SetWarnings(v []AuthEmailSignupUser200ResponseWarningsInner) {
	o.Warnings = v
}

func (o ModQuizGetQuizFeedbackForGrade200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModQuizGetQuizFeedbackForGrade200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Feedbackinlinefiles) {
		toSerialize["feedbackinlinefiles"] = o.Feedbackinlinefiles
	}
	toSerialize["feedbacktext"] = o.Feedbacktext
	if !IsNil(o.Feedbacktextformat) {
		toSerialize["feedbacktextformat"] = o.Feedbacktextformat
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

func (o *ModQuizGetQuizFeedbackForGrade200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"feedbacktext",
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

	varModQuizGetQuizFeedbackForGrade200Response := _ModQuizGetQuizFeedbackForGrade200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModQuizGetQuizFeedbackForGrade200Response)

	if err != nil {
		return err
	}

	*o = ModQuizGetQuizFeedbackForGrade200Response(varModQuizGetQuizFeedbackForGrade200Response)

	return err
}

type NullableModQuizGetQuizFeedbackForGrade200Response struct {
	value *ModQuizGetQuizFeedbackForGrade200Response
	isSet bool
}

func (v NullableModQuizGetQuizFeedbackForGrade200Response) Get() *ModQuizGetQuizFeedbackForGrade200Response {
	return v.value
}

func (v *NullableModQuizGetQuizFeedbackForGrade200Response) Set(val *ModQuizGetQuizFeedbackForGrade200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableModQuizGetQuizFeedbackForGrade200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableModQuizGetQuizFeedbackForGrade200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModQuizGetQuizFeedbackForGrade200Response(val *ModQuizGetQuizFeedbackForGrade200Response) *NullableModQuizGetQuizFeedbackForGrade200Response {
	return &NullableModQuizGetQuizFeedbackForGrade200Response{value: val, isSet: true}
}

func (v NullableModQuizGetQuizFeedbackForGrade200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModQuizGetQuizFeedbackForGrade200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


