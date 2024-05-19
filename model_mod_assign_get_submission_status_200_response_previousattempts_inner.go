/*
Moodle Webservice API

Auto-generated OpenAPI spec for Moodle's Webservice API. Requires installation of the Catalyst RESTful API plugin: https://github.com/catalyst/moodle-webservice_restful.

API version: 4.3.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package moodleclient

import (
	"encoding/json"
)

// checks if the ModAssignGetSubmissionStatus200ResponsePreviousattemptsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModAssignGetSubmissionStatus200ResponsePreviousattemptsInner{}

// ModAssignGetSubmissionStatus200ResponsePreviousattemptsInner struct for ModAssignGetSubmissionStatus200ResponsePreviousattemptsInner
type ModAssignGetSubmissionStatus200ResponsePreviousattemptsInner struct {
	// Attempt number.
	Attemptnumber *int32 `json:"attemptnumber,omitempty"`
	Feedbackplugins []ModAssignGetSubmissionStatus200ResponseLastattemptSubmissionPluginsInner `json:"feedbackplugins,omitempty"`
	Grade *ModAssignGetSubmissionStatus200ResponseFeedbackGrade `json:"grade,omitempty"`
	Submission *ModAssignGetSubmissionStatus200ResponseLastattemptTeamsubmission `json:"submission,omitempty"`
}

// NewModAssignGetSubmissionStatus200ResponsePreviousattemptsInner instantiates a new ModAssignGetSubmissionStatus200ResponsePreviousattemptsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModAssignGetSubmissionStatus200ResponsePreviousattemptsInner() *ModAssignGetSubmissionStatus200ResponsePreviousattemptsInner {
	this := ModAssignGetSubmissionStatus200ResponsePreviousattemptsInner{}
	var attemptnumber int32 = null
	this.Attemptnumber = &attemptnumber
	return &this
}

// NewModAssignGetSubmissionStatus200ResponsePreviousattemptsInnerWithDefaults instantiates a new ModAssignGetSubmissionStatus200ResponsePreviousattemptsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModAssignGetSubmissionStatus200ResponsePreviousattemptsInnerWithDefaults() *ModAssignGetSubmissionStatus200ResponsePreviousattemptsInner {
	this := ModAssignGetSubmissionStatus200ResponsePreviousattemptsInner{}
	var attemptnumber int32 = null
	this.Attemptnumber = &attemptnumber
	return &this
}

// GetAttemptnumber returns the Attemptnumber field value if set, zero value otherwise.
func (o *ModAssignGetSubmissionStatus200ResponsePreviousattemptsInner) GetAttemptnumber() int32 {
	if o == nil || IsNil(o.Attemptnumber) {
		var ret int32
		return ret
	}
	return *o.Attemptnumber
}

// GetAttemptnumberOk returns a tuple with the Attemptnumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModAssignGetSubmissionStatus200ResponsePreviousattemptsInner) GetAttemptnumberOk() (*int32, bool) {
	if o == nil || IsNil(o.Attemptnumber) {
		return nil, false
	}
	return o.Attemptnumber, true
}

// HasAttemptnumber returns a boolean if a field has been set.
func (o *ModAssignGetSubmissionStatus200ResponsePreviousattemptsInner) HasAttemptnumber() bool {
	if o != nil && !IsNil(o.Attemptnumber) {
		return true
	}

	return false
}

// SetAttemptnumber gets a reference to the given int32 and assigns it to the Attemptnumber field.
func (o *ModAssignGetSubmissionStatus200ResponsePreviousattemptsInner) SetAttemptnumber(v int32) {
	o.Attemptnumber = &v
}

// GetFeedbackplugins returns the Feedbackplugins field value if set, zero value otherwise.
func (o *ModAssignGetSubmissionStatus200ResponsePreviousattemptsInner) GetFeedbackplugins() []ModAssignGetSubmissionStatus200ResponseLastattemptSubmissionPluginsInner {
	if o == nil || IsNil(o.Feedbackplugins) {
		var ret []ModAssignGetSubmissionStatus200ResponseLastattemptSubmissionPluginsInner
		return ret
	}
	return o.Feedbackplugins
}

// GetFeedbackpluginsOk returns a tuple with the Feedbackplugins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModAssignGetSubmissionStatus200ResponsePreviousattemptsInner) GetFeedbackpluginsOk() ([]ModAssignGetSubmissionStatus200ResponseLastattemptSubmissionPluginsInner, bool) {
	if o == nil || IsNil(o.Feedbackplugins) {
		return nil, false
	}
	return o.Feedbackplugins, true
}

// HasFeedbackplugins returns a boolean if a field has been set.
func (o *ModAssignGetSubmissionStatus200ResponsePreviousattemptsInner) HasFeedbackplugins() bool {
	if o != nil && !IsNil(o.Feedbackplugins) {
		return true
	}

	return false
}

// SetFeedbackplugins gets a reference to the given []ModAssignGetSubmissionStatus200ResponseLastattemptSubmissionPluginsInner and assigns it to the Feedbackplugins field.
func (o *ModAssignGetSubmissionStatus200ResponsePreviousattemptsInner) SetFeedbackplugins(v []ModAssignGetSubmissionStatus200ResponseLastattemptSubmissionPluginsInner) {
	o.Feedbackplugins = v
}

// GetGrade returns the Grade field value if set, zero value otherwise.
func (o *ModAssignGetSubmissionStatus200ResponsePreviousattemptsInner) GetGrade() ModAssignGetSubmissionStatus200ResponseFeedbackGrade {
	if o == nil || IsNil(o.Grade) {
		var ret ModAssignGetSubmissionStatus200ResponseFeedbackGrade
		return ret
	}
	return *o.Grade
}

// GetGradeOk returns a tuple with the Grade field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModAssignGetSubmissionStatus200ResponsePreviousattemptsInner) GetGradeOk() (*ModAssignGetSubmissionStatus200ResponseFeedbackGrade, bool) {
	if o == nil || IsNil(o.Grade) {
		return nil, false
	}
	return o.Grade, true
}

// HasGrade returns a boolean if a field has been set.
func (o *ModAssignGetSubmissionStatus200ResponsePreviousattemptsInner) HasGrade() bool {
	if o != nil && !IsNil(o.Grade) {
		return true
	}

	return false
}

// SetGrade gets a reference to the given ModAssignGetSubmissionStatus200ResponseFeedbackGrade and assigns it to the Grade field.
func (o *ModAssignGetSubmissionStatus200ResponsePreviousattemptsInner) SetGrade(v ModAssignGetSubmissionStatus200ResponseFeedbackGrade) {
	o.Grade = &v
}

// GetSubmission returns the Submission field value if set, zero value otherwise.
func (o *ModAssignGetSubmissionStatus200ResponsePreviousattemptsInner) GetSubmission() ModAssignGetSubmissionStatus200ResponseLastattemptTeamsubmission {
	if o == nil || IsNil(o.Submission) {
		var ret ModAssignGetSubmissionStatus200ResponseLastattemptTeamsubmission
		return ret
	}
	return *o.Submission
}

// GetSubmissionOk returns a tuple with the Submission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModAssignGetSubmissionStatus200ResponsePreviousattemptsInner) GetSubmissionOk() (*ModAssignGetSubmissionStatus200ResponseLastattemptTeamsubmission, bool) {
	if o == nil || IsNil(o.Submission) {
		return nil, false
	}
	return o.Submission, true
}

// HasSubmission returns a boolean if a field has been set.
func (o *ModAssignGetSubmissionStatus200ResponsePreviousattemptsInner) HasSubmission() bool {
	if o != nil && !IsNil(o.Submission) {
		return true
	}

	return false
}

// SetSubmission gets a reference to the given ModAssignGetSubmissionStatus200ResponseLastattemptTeamsubmission and assigns it to the Submission field.
func (o *ModAssignGetSubmissionStatus200ResponsePreviousattemptsInner) SetSubmission(v ModAssignGetSubmissionStatus200ResponseLastattemptTeamsubmission) {
	o.Submission = &v
}

func (o ModAssignGetSubmissionStatus200ResponsePreviousattemptsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModAssignGetSubmissionStatus200ResponsePreviousattemptsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Attemptnumber) {
		toSerialize["attemptnumber"] = o.Attemptnumber
	}
	if !IsNil(o.Feedbackplugins) {
		toSerialize["feedbackplugins"] = o.Feedbackplugins
	}
	if !IsNil(o.Grade) {
		toSerialize["grade"] = o.Grade
	}
	if !IsNil(o.Submission) {
		toSerialize["submission"] = o.Submission
	}
	return toSerialize, nil
}

type NullableModAssignGetSubmissionStatus200ResponsePreviousattemptsInner struct {
	value *ModAssignGetSubmissionStatus200ResponsePreviousattemptsInner
	isSet bool
}

func (v NullableModAssignGetSubmissionStatus200ResponsePreviousattemptsInner) Get() *ModAssignGetSubmissionStatus200ResponsePreviousattemptsInner {
	return v.value
}

func (v *NullableModAssignGetSubmissionStatus200ResponsePreviousattemptsInner) Set(val *ModAssignGetSubmissionStatus200ResponsePreviousattemptsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableModAssignGetSubmissionStatus200ResponsePreviousattemptsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableModAssignGetSubmissionStatus200ResponsePreviousattemptsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModAssignGetSubmissionStatus200ResponsePreviousattemptsInner(val *ModAssignGetSubmissionStatus200ResponsePreviousattemptsInner) *NullableModAssignGetSubmissionStatus200ResponsePreviousattemptsInner {
	return &NullableModAssignGetSubmissionStatus200ResponsePreviousattemptsInner{value: val, isSet: true}
}

func (v NullableModAssignGetSubmissionStatus200ResponsePreviousattemptsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModAssignGetSubmissionStatus200ResponsePreviousattemptsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


