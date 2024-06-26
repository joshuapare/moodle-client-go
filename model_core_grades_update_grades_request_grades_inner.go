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

// checks if the CoreGradesUpdateGradesRequestGradesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreGradesUpdateGradesRequestGradesInner{}

// CoreGradesUpdateGradesRequestGradesInner struct for CoreGradesUpdateGradesRequestGradesInner
type CoreGradesUpdateGradesRequestGradesInner struct {
	// Student grade
	Grade *float32 `json:"grade,omitempty"`
	// A string representation of the feedback from the grader
	StrFeedback *string `json:"str_feedback,omitempty"`
	// Student ID
	Studentid *int32 `json:"studentid,omitempty"`
}

// NewCoreGradesUpdateGradesRequestGradesInner instantiates a new CoreGradesUpdateGradesRequestGradesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreGradesUpdateGradesRequestGradesInner() *CoreGradesUpdateGradesRequestGradesInner {
	this := CoreGradesUpdateGradesRequestGradesInner{}
	var grade float32 = null
	this.Grade = &grade
	var strFeedback string = "null"
	this.StrFeedback = &strFeedback
	var studentid int32 = null
	this.Studentid = &studentid
	return &this
}

// NewCoreGradesUpdateGradesRequestGradesInnerWithDefaults instantiates a new CoreGradesUpdateGradesRequestGradesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreGradesUpdateGradesRequestGradesInnerWithDefaults() *CoreGradesUpdateGradesRequestGradesInner {
	this := CoreGradesUpdateGradesRequestGradesInner{}
	var grade float32 = null
	this.Grade = &grade
	var strFeedback string = "null"
	this.StrFeedback = &strFeedback
	var studentid int32 = null
	this.Studentid = &studentid
	return &this
}

// GetGrade returns the Grade field value if set, zero value otherwise.
func (o *CoreGradesUpdateGradesRequestGradesInner) GetGrade() float32 {
	if o == nil || IsNil(o.Grade) {
		var ret float32
		return ret
	}
	return *o.Grade
}

// GetGradeOk returns a tuple with the Grade field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreGradesUpdateGradesRequestGradesInner) GetGradeOk() (*float32, bool) {
	if o == nil || IsNil(o.Grade) {
		return nil, false
	}
	return o.Grade, true
}

// HasGrade returns a boolean if a field has been set.
func (o *CoreGradesUpdateGradesRequestGradesInner) HasGrade() bool {
	if o != nil && !IsNil(o.Grade) {
		return true
	}

	return false
}

// SetGrade gets a reference to the given float32 and assigns it to the Grade field.
func (o *CoreGradesUpdateGradesRequestGradesInner) SetGrade(v float32) {
	o.Grade = &v
}

// GetStrFeedback returns the StrFeedback field value if set, zero value otherwise.
func (o *CoreGradesUpdateGradesRequestGradesInner) GetStrFeedback() string {
	if o == nil || IsNil(o.StrFeedback) {
		var ret string
		return ret
	}
	return *o.StrFeedback
}

// GetStrFeedbackOk returns a tuple with the StrFeedback field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreGradesUpdateGradesRequestGradesInner) GetStrFeedbackOk() (*string, bool) {
	if o == nil || IsNil(o.StrFeedback) {
		return nil, false
	}
	return o.StrFeedback, true
}

// HasStrFeedback returns a boolean if a field has been set.
func (o *CoreGradesUpdateGradesRequestGradesInner) HasStrFeedback() bool {
	if o != nil && !IsNil(o.StrFeedback) {
		return true
	}

	return false
}

// SetStrFeedback gets a reference to the given string and assigns it to the StrFeedback field.
func (o *CoreGradesUpdateGradesRequestGradesInner) SetStrFeedback(v string) {
	o.StrFeedback = &v
}

// GetStudentid returns the Studentid field value if set, zero value otherwise.
func (o *CoreGradesUpdateGradesRequestGradesInner) GetStudentid() int32 {
	if o == nil || IsNil(o.Studentid) {
		var ret int32
		return ret
	}
	return *o.Studentid
}

// GetStudentidOk returns a tuple with the Studentid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreGradesUpdateGradesRequestGradesInner) GetStudentidOk() (*int32, bool) {
	if o == nil || IsNil(o.Studentid) {
		return nil, false
	}
	return o.Studentid, true
}

// HasStudentid returns a boolean if a field has been set.
func (o *CoreGradesUpdateGradesRequestGradesInner) HasStudentid() bool {
	if o != nil && !IsNil(o.Studentid) {
		return true
	}

	return false
}

// SetStudentid gets a reference to the given int32 and assigns it to the Studentid field.
func (o *CoreGradesUpdateGradesRequestGradesInner) SetStudentid(v int32) {
	o.Studentid = &v
}

func (o CoreGradesUpdateGradesRequestGradesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreGradesUpdateGradesRequestGradesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Grade) {
		toSerialize["grade"] = o.Grade
	}
	if !IsNil(o.StrFeedback) {
		toSerialize["str_feedback"] = o.StrFeedback
	}
	if !IsNil(o.Studentid) {
		toSerialize["studentid"] = o.Studentid
	}
	return toSerialize, nil
}

type NullableCoreGradesUpdateGradesRequestGradesInner struct {
	value *CoreGradesUpdateGradesRequestGradesInner
	isSet bool
}

func (v NullableCoreGradesUpdateGradesRequestGradesInner) Get() *CoreGradesUpdateGradesRequestGradesInner {
	return v.value
}

func (v *NullableCoreGradesUpdateGradesRequestGradesInner) Set(val *CoreGradesUpdateGradesRequestGradesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreGradesUpdateGradesRequestGradesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreGradesUpdateGradesRequestGradesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreGradesUpdateGradesRequestGradesInner(val *CoreGradesUpdateGradesRequestGradesInner) *NullableCoreGradesUpdateGradesRequestGradesInner {
	return &NullableCoreGradesUpdateGradesRequestGradesInner{value: val, isSet: true}
}

func (v NullableCoreGradesUpdateGradesRequestGradesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreGradesUpdateGradesRequestGradesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


