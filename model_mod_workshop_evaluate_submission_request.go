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

// checks if the ModWorkshopEvaluateSubmissionRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModWorkshopEvaluateSubmissionRequest{}

// ModWorkshopEvaluateSubmissionRequest struct for ModWorkshopEvaluateSubmissionRequest
type ModWorkshopEvaluateSubmissionRequest struct {
	// The feedback format for text.
	Feedbackformat *int32 `json:"feedbackformat,omitempty"`
	// The feedback for the author.
	Feedbacktext *string `json:"feedbacktext,omitempty"`
	// The new submission grade.
	Gradeover *string `json:"gradeover,omitempty"`
	// Publish the submission for others?.
	Published *bool `json:"published,omitempty"`
	// submission id.
	Submissionid int32 `json:"submissionid"`
}

type _ModWorkshopEvaluateSubmissionRequest ModWorkshopEvaluateSubmissionRequest

// NewModWorkshopEvaluateSubmissionRequest instantiates a new ModWorkshopEvaluateSubmissionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModWorkshopEvaluateSubmissionRequest(submissionid int32) *ModWorkshopEvaluateSubmissionRequest {
	this := ModWorkshopEvaluateSubmissionRequest{}
	var feedbackformat int32 = 0
	this.Feedbackformat = &feedbackformat
	var feedbacktext string = ""
	this.Feedbacktext = &feedbacktext
	var gradeover string = ""
	this.Gradeover = &gradeover
	var published bool = false
	this.Published = &published
	this.Submissionid = submissionid
	return &this
}

// NewModWorkshopEvaluateSubmissionRequestWithDefaults instantiates a new ModWorkshopEvaluateSubmissionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModWorkshopEvaluateSubmissionRequestWithDefaults() *ModWorkshopEvaluateSubmissionRequest {
	this := ModWorkshopEvaluateSubmissionRequest{}
	var feedbackformat int32 = 0
	this.Feedbackformat = &feedbackformat
	var feedbacktext string = ""
	this.Feedbacktext = &feedbacktext
	var gradeover string = ""
	this.Gradeover = &gradeover
	var published bool = false
	this.Published = &published
	var submissionid int32 = null
	this.Submissionid = submissionid
	return &this
}

// GetFeedbackformat returns the Feedbackformat field value if set, zero value otherwise.
func (o *ModWorkshopEvaluateSubmissionRequest) GetFeedbackformat() int32 {
	if o == nil || IsNil(o.Feedbackformat) {
		var ret int32
		return ret
	}
	return *o.Feedbackformat
}

// GetFeedbackformatOk returns a tuple with the Feedbackformat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModWorkshopEvaluateSubmissionRequest) GetFeedbackformatOk() (*int32, bool) {
	if o == nil || IsNil(o.Feedbackformat) {
		return nil, false
	}
	return o.Feedbackformat, true
}

// HasFeedbackformat returns a boolean if a field has been set.
func (o *ModWorkshopEvaluateSubmissionRequest) HasFeedbackformat() bool {
	if o != nil && !IsNil(o.Feedbackformat) {
		return true
	}

	return false
}

// SetFeedbackformat gets a reference to the given int32 and assigns it to the Feedbackformat field.
func (o *ModWorkshopEvaluateSubmissionRequest) SetFeedbackformat(v int32) {
	o.Feedbackformat = &v
}

// GetFeedbacktext returns the Feedbacktext field value if set, zero value otherwise.
func (o *ModWorkshopEvaluateSubmissionRequest) GetFeedbacktext() string {
	if o == nil || IsNil(o.Feedbacktext) {
		var ret string
		return ret
	}
	return *o.Feedbacktext
}

// GetFeedbacktextOk returns a tuple with the Feedbacktext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModWorkshopEvaluateSubmissionRequest) GetFeedbacktextOk() (*string, bool) {
	if o == nil || IsNil(o.Feedbacktext) {
		return nil, false
	}
	return o.Feedbacktext, true
}

// HasFeedbacktext returns a boolean if a field has been set.
func (o *ModWorkshopEvaluateSubmissionRequest) HasFeedbacktext() bool {
	if o != nil && !IsNil(o.Feedbacktext) {
		return true
	}

	return false
}

// SetFeedbacktext gets a reference to the given string and assigns it to the Feedbacktext field.
func (o *ModWorkshopEvaluateSubmissionRequest) SetFeedbacktext(v string) {
	o.Feedbacktext = &v
}

// GetGradeover returns the Gradeover field value if set, zero value otherwise.
func (o *ModWorkshopEvaluateSubmissionRequest) GetGradeover() string {
	if o == nil || IsNil(o.Gradeover) {
		var ret string
		return ret
	}
	return *o.Gradeover
}

// GetGradeoverOk returns a tuple with the Gradeover field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModWorkshopEvaluateSubmissionRequest) GetGradeoverOk() (*string, bool) {
	if o == nil || IsNil(o.Gradeover) {
		return nil, false
	}
	return o.Gradeover, true
}

// HasGradeover returns a boolean if a field has been set.
func (o *ModWorkshopEvaluateSubmissionRequest) HasGradeover() bool {
	if o != nil && !IsNil(o.Gradeover) {
		return true
	}

	return false
}

// SetGradeover gets a reference to the given string and assigns it to the Gradeover field.
func (o *ModWorkshopEvaluateSubmissionRequest) SetGradeover(v string) {
	o.Gradeover = &v
}

// GetPublished returns the Published field value if set, zero value otherwise.
func (o *ModWorkshopEvaluateSubmissionRequest) GetPublished() bool {
	if o == nil || IsNil(o.Published) {
		var ret bool
		return ret
	}
	return *o.Published
}

// GetPublishedOk returns a tuple with the Published field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModWorkshopEvaluateSubmissionRequest) GetPublishedOk() (*bool, bool) {
	if o == nil || IsNil(o.Published) {
		return nil, false
	}
	return o.Published, true
}

// HasPublished returns a boolean if a field has been set.
func (o *ModWorkshopEvaluateSubmissionRequest) HasPublished() bool {
	if o != nil && !IsNil(o.Published) {
		return true
	}

	return false
}

// SetPublished gets a reference to the given bool and assigns it to the Published field.
func (o *ModWorkshopEvaluateSubmissionRequest) SetPublished(v bool) {
	o.Published = &v
}

// GetSubmissionid returns the Submissionid field value
func (o *ModWorkshopEvaluateSubmissionRequest) GetSubmissionid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Submissionid
}

// GetSubmissionidOk returns a tuple with the Submissionid field value
// and a boolean to check if the value has been set.
func (o *ModWorkshopEvaluateSubmissionRequest) GetSubmissionidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Submissionid, true
}

// SetSubmissionid sets field value
func (o *ModWorkshopEvaluateSubmissionRequest) SetSubmissionid(v int32) {
	o.Submissionid = v
}

func (o ModWorkshopEvaluateSubmissionRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModWorkshopEvaluateSubmissionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Feedbackformat) {
		toSerialize["feedbackformat"] = o.Feedbackformat
	}
	if !IsNil(o.Feedbacktext) {
		toSerialize["feedbacktext"] = o.Feedbacktext
	}
	if !IsNil(o.Gradeover) {
		toSerialize["gradeover"] = o.Gradeover
	}
	if !IsNil(o.Published) {
		toSerialize["published"] = o.Published
	}
	toSerialize["submissionid"] = o.Submissionid
	return toSerialize, nil
}

func (o *ModWorkshopEvaluateSubmissionRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"submissionid",
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

	varModWorkshopEvaluateSubmissionRequest := _ModWorkshopEvaluateSubmissionRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModWorkshopEvaluateSubmissionRequest)

	if err != nil {
		return err
	}

	*o = ModWorkshopEvaluateSubmissionRequest(varModWorkshopEvaluateSubmissionRequest)

	return err
}

type NullableModWorkshopEvaluateSubmissionRequest struct {
	value *ModWorkshopEvaluateSubmissionRequest
	isSet bool
}

func (v NullableModWorkshopEvaluateSubmissionRequest) Get() *ModWorkshopEvaluateSubmissionRequest {
	return v.value
}

func (v *NullableModWorkshopEvaluateSubmissionRequest) Set(val *ModWorkshopEvaluateSubmissionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModWorkshopEvaluateSubmissionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModWorkshopEvaluateSubmissionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModWorkshopEvaluateSubmissionRequest(val *ModWorkshopEvaluateSubmissionRequest) *NullableModWorkshopEvaluateSubmissionRequest {
	return &NullableModWorkshopEvaluateSubmissionRequest{value: val, isSet: true}
}

func (v NullableModWorkshopEvaluateSubmissionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModWorkshopEvaluateSubmissionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


