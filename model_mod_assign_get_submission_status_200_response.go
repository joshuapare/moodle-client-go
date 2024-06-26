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

// checks if the ModAssignGetSubmissionStatus200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModAssignGetSubmissionStatus200Response{}

// ModAssignGetSubmissionStatus200Response struct for ModAssignGetSubmissionStatus200Response
type ModAssignGetSubmissionStatus200Response struct {
	Assignmentdata *ModAssignGetSubmissionStatus200ResponseAssignmentdata `json:"assignmentdata,omitempty"`
	Feedback *ModAssignGetSubmissionStatus200ResponseFeedback `json:"feedback,omitempty"`
	Gradingsummary *ModAssignGetSubmissionStatus200ResponseGradingsummary `json:"gradingsummary,omitempty"`
	Lastattempt *ModAssignGetSubmissionStatus200ResponseLastattempt `json:"lastattempt,omitempty"`
	Previousattempts []ModAssignGetSubmissionStatus200ResponsePreviousattemptsInner `json:"previousattempts,omitempty"`
	Warnings []AuthEmailSignupUser200ResponseWarningsInner `json:"warnings,omitempty"`
}

// NewModAssignGetSubmissionStatus200Response instantiates a new ModAssignGetSubmissionStatus200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModAssignGetSubmissionStatus200Response() *ModAssignGetSubmissionStatus200Response {
	this := ModAssignGetSubmissionStatus200Response{}
	return &this
}

// NewModAssignGetSubmissionStatus200ResponseWithDefaults instantiates a new ModAssignGetSubmissionStatus200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModAssignGetSubmissionStatus200ResponseWithDefaults() *ModAssignGetSubmissionStatus200Response {
	this := ModAssignGetSubmissionStatus200Response{}
	return &this
}

// GetAssignmentdata returns the Assignmentdata field value if set, zero value otherwise.
func (o *ModAssignGetSubmissionStatus200Response) GetAssignmentdata() ModAssignGetSubmissionStatus200ResponseAssignmentdata {
	if o == nil || IsNil(o.Assignmentdata) {
		var ret ModAssignGetSubmissionStatus200ResponseAssignmentdata
		return ret
	}
	return *o.Assignmentdata
}

// GetAssignmentdataOk returns a tuple with the Assignmentdata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModAssignGetSubmissionStatus200Response) GetAssignmentdataOk() (*ModAssignGetSubmissionStatus200ResponseAssignmentdata, bool) {
	if o == nil || IsNil(o.Assignmentdata) {
		return nil, false
	}
	return o.Assignmentdata, true
}

// HasAssignmentdata returns a boolean if a field has been set.
func (o *ModAssignGetSubmissionStatus200Response) HasAssignmentdata() bool {
	if o != nil && !IsNil(o.Assignmentdata) {
		return true
	}

	return false
}

// SetAssignmentdata gets a reference to the given ModAssignGetSubmissionStatus200ResponseAssignmentdata and assigns it to the Assignmentdata field.
func (o *ModAssignGetSubmissionStatus200Response) SetAssignmentdata(v ModAssignGetSubmissionStatus200ResponseAssignmentdata) {
	o.Assignmentdata = &v
}

// GetFeedback returns the Feedback field value if set, zero value otherwise.
func (o *ModAssignGetSubmissionStatus200Response) GetFeedback() ModAssignGetSubmissionStatus200ResponseFeedback {
	if o == nil || IsNil(o.Feedback) {
		var ret ModAssignGetSubmissionStatus200ResponseFeedback
		return ret
	}
	return *o.Feedback
}

// GetFeedbackOk returns a tuple with the Feedback field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModAssignGetSubmissionStatus200Response) GetFeedbackOk() (*ModAssignGetSubmissionStatus200ResponseFeedback, bool) {
	if o == nil || IsNil(o.Feedback) {
		return nil, false
	}
	return o.Feedback, true
}

// HasFeedback returns a boolean if a field has been set.
func (o *ModAssignGetSubmissionStatus200Response) HasFeedback() bool {
	if o != nil && !IsNil(o.Feedback) {
		return true
	}

	return false
}

// SetFeedback gets a reference to the given ModAssignGetSubmissionStatus200ResponseFeedback and assigns it to the Feedback field.
func (o *ModAssignGetSubmissionStatus200Response) SetFeedback(v ModAssignGetSubmissionStatus200ResponseFeedback) {
	o.Feedback = &v
}

// GetGradingsummary returns the Gradingsummary field value if set, zero value otherwise.
func (o *ModAssignGetSubmissionStatus200Response) GetGradingsummary() ModAssignGetSubmissionStatus200ResponseGradingsummary {
	if o == nil || IsNil(o.Gradingsummary) {
		var ret ModAssignGetSubmissionStatus200ResponseGradingsummary
		return ret
	}
	return *o.Gradingsummary
}

// GetGradingsummaryOk returns a tuple with the Gradingsummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModAssignGetSubmissionStatus200Response) GetGradingsummaryOk() (*ModAssignGetSubmissionStatus200ResponseGradingsummary, bool) {
	if o == nil || IsNil(o.Gradingsummary) {
		return nil, false
	}
	return o.Gradingsummary, true
}

// HasGradingsummary returns a boolean if a field has been set.
func (o *ModAssignGetSubmissionStatus200Response) HasGradingsummary() bool {
	if o != nil && !IsNil(o.Gradingsummary) {
		return true
	}

	return false
}

// SetGradingsummary gets a reference to the given ModAssignGetSubmissionStatus200ResponseGradingsummary and assigns it to the Gradingsummary field.
func (o *ModAssignGetSubmissionStatus200Response) SetGradingsummary(v ModAssignGetSubmissionStatus200ResponseGradingsummary) {
	o.Gradingsummary = &v
}

// GetLastattempt returns the Lastattempt field value if set, zero value otherwise.
func (o *ModAssignGetSubmissionStatus200Response) GetLastattempt() ModAssignGetSubmissionStatus200ResponseLastattempt {
	if o == nil || IsNil(o.Lastattempt) {
		var ret ModAssignGetSubmissionStatus200ResponseLastattempt
		return ret
	}
	return *o.Lastattempt
}

// GetLastattemptOk returns a tuple with the Lastattempt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModAssignGetSubmissionStatus200Response) GetLastattemptOk() (*ModAssignGetSubmissionStatus200ResponseLastattempt, bool) {
	if o == nil || IsNil(o.Lastattempt) {
		return nil, false
	}
	return o.Lastattempt, true
}

// HasLastattempt returns a boolean if a field has been set.
func (o *ModAssignGetSubmissionStatus200Response) HasLastattempt() bool {
	if o != nil && !IsNil(o.Lastattempt) {
		return true
	}

	return false
}

// SetLastattempt gets a reference to the given ModAssignGetSubmissionStatus200ResponseLastattempt and assigns it to the Lastattempt field.
func (o *ModAssignGetSubmissionStatus200Response) SetLastattempt(v ModAssignGetSubmissionStatus200ResponseLastattempt) {
	o.Lastattempt = &v
}

// GetPreviousattempts returns the Previousattempts field value if set, zero value otherwise.
func (o *ModAssignGetSubmissionStatus200Response) GetPreviousattempts() []ModAssignGetSubmissionStatus200ResponsePreviousattemptsInner {
	if o == nil || IsNil(o.Previousattempts) {
		var ret []ModAssignGetSubmissionStatus200ResponsePreviousattemptsInner
		return ret
	}
	return o.Previousattempts
}

// GetPreviousattemptsOk returns a tuple with the Previousattempts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModAssignGetSubmissionStatus200Response) GetPreviousattemptsOk() ([]ModAssignGetSubmissionStatus200ResponsePreviousattemptsInner, bool) {
	if o == nil || IsNil(o.Previousattempts) {
		return nil, false
	}
	return o.Previousattempts, true
}

// HasPreviousattempts returns a boolean if a field has been set.
func (o *ModAssignGetSubmissionStatus200Response) HasPreviousattempts() bool {
	if o != nil && !IsNil(o.Previousattempts) {
		return true
	}

	return false
}

// SetPreviousattempts gets a reference to the given []ModAssignGetSubmissionStatus200ResponsePreviousattemptsInner and assigns it to the Previousattempts field.
func (o *ModAssignGetSubmissionStatus200Response) SetPreviousattempts(v []ModAssignGetSubmissionStatus200ResponsePreviousattemptsInner) {
	o.Previousattempts = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *ModAssignGetSubmissionStatus200Response) GetWarnings() []AuthEmailSignupUser200ResponseWarningsInner {
	if o == nil || IsNil(o.Warnings) {
		var ret []AuthEmailSignupUser200ResponseWarningsInner
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModAssignGetSubmissionStatus200Response) GetWarningsOk() ([]AuthEmailSignupUser200ResponseWarningsInner, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *ModAssignGetSubmissionStatus200Response) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []AuthEmailSignupUser200ResponseWarningsInner and assigns it to the Warnings field.
func (o *ModAssignGetSubmissionStatus200Response) SetWarnings(v []AuthEmailSignupUser200ResponseWarningsInner) {
	o.Warnings = v
}

func (o ModAssignGetSubmissionStatus200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModAssignGetSubmissionStatus200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Assignmentdata) {
		toSerialize["assignmentdata"] = o.Assignmentdata
	}
	if !IsNil(o.Feedback) {
		toSerialize["feedback"] = o.Feedback
	}
	if !IsNil(o.Gradingsummary) {
		toSerialize["gradingsummary"] = o.Gradingsummary
	}
	if !IsNil(o.Lastattempt) {
		toSerialize["lastattempt"] = o.Lastattempt
	}
	if !IsNil(o.Previousattempts) {
		toSerialize["previousattempts"] = o.Previousattempts
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullableModAssignGetSubmissionStatus200Response struct {
	value *ModAssignGetSubmissionStatus200Response
	isSet bool
}

func (v NullableModAssignGetSubmissionStatus200Response) Get() *ModAssignGetSubmissionStatus200Response {
	return v.value
}

func (v *NullableModAssignGetSubmissionStatus200Response) Set(val *ModAssignGetSubmissionStatus200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableModAssignGetSubmissionStatus200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableModAssignGetSubmissionStatus200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModAssignGetSubmissionStatus200Response(val *ModAssignGetSubmissionStatus200Response) *NullableModAssignGetSubmissionStatus200Response {
	return &NullableModAssignGetSubmissionStatus200Response{value: val, isSet: true}
}

func (v NullableModAssignGetSubmissionStatus200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModAssignGetSubmissionStatus200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


