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

// checks if the ModLessonGetAttemptsOverview200ResponseDataStudentsInnerAttemptsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModLessonGetAttemptsOverview200ResponseDataStudentsInnerAttemptsInner{}

// ModLessonGetAttemptsOverview200ResponseDataStudentsInnerAttemptsInner struct for ModLessonGetAttemptsOverview200ResponseDataStudentsInnerAttemptsInner
type ModLessonGetAttemptsOverview200ResponseDataStudentsInnerAttemptsInner struct {
	// Attempt time ended.
	End *int32 `json:"end,omitempty"`
	// Attempt grade.
	Grade *float32 `json:"grade,omitempty"`
	// Attempt last time continued.
	Timeend *int32 `json:"timeend,omitempty"`
	// Attempt time started.
	Timestart *int32 `json:"timestart,omitempty"`
	// Attempt number.
	Try *int32 `json:"try,omitempty"`
}

// NewModLessonGetAttemptsOverview200ResponseDataStudentsInnerAttemptsInner instantiates a new ModLessonGetAttemptsOverview200ResponseDataStudentsInnerAttemptsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModLessonGetAttemptsOverview200ResponseDataStudentsInnerAttemptsInner() *ModLessonGetAttemptsOverview200ResponseDataStudentsInnerAttemptsInner {
	this := ModLessonGetAttemptsOverview200ResponseDataStudentsInnerAttemptsInner{}
	var end int32 = null
	this.End = &end
	var grade float32 = null
	this.Grade = &grade
	var timeend int32 = null
	this.Timeend = &timeend
	var timestart int32 = null
	this.Timestart = &timestart
	return &this
}

// NewModLessonGetAttemptsOverview200ResponseDataStudentsInnerAttemptsInnerWithDefaults instantiates a new ModLessonGetAttemptsOverview200ResponseDataStudentsInnerAttemptsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModLessonGetAttemptsOverview200ResponseDataStudentsInnerAttemptsInnerWithDefaults() *ModLessonGetAttemptsOverview200ResponseDataStudentsInnerAttemptsInner {
	this := ModLessonGetAttemptsOverview200ResponseDataStudentsInnerAttemptsInner{}
	var end int32 = null
	this.End = &end
	var grade float32 = null
	this.Grade = &grade
	var timeend int32 = null
	this.Timeend = &timeend
	var timestart int32 = null
	this.Timestart = &timestart
	return &this
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *ModLessonGetAttemptsOverview200ResponseDataStudentsInnerAttemptsInner) GetEnd() int32 {
	if o == nil || IsNil(o.End) {
		var ret int32
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModLessonGetAttemptsOverview200ResponseDataStudentsInnerAttemptsInner) GetEndOk() (*int32, bool) {
	if o == nil || IsNil(o.End) {
		return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *ModLessonGetAttemptsOverview200ResponseDataStudentsInnerAttemptsInner) HasEnd() bool {
	if o != nil && !IsNil(o.End) {
		return true
	}

	return false
}

// SetEnd gets a reference to the given int32 and assigns it to the End field.
func (o *ModLessonGetAttemptsOverview200ResponseDataStudentsInnerAttemptsInner) SetEnd(v int32) {
	o.End = &v
}

// GetGrade returns the Grade field value if set, zero value otherwise.
func (o *ModLessonGetAttemptsOverview200ResponseDataStudentsInnerAttemptsInner) GetGrade() float32 {
	if o == nil || IsNil(o.Grade) {
		var ret float32
		return ret
	}
	return *o.Grade
}

// GetGradeOk returns a tuple with the Grade field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModLessonGetAttemptsOverview200ResponseDataStudentsInnerAttemptsInner) GetGradeOk() (*float32, bool) {
	if o == nil || IsNil(o.Grade) {
		return nil, false
	}
	return o.Grade, true
}

// HasGrade returns a boolean if a field has been set.
func (o *ModLessonGetAttemptsOverview200ResponseDataStudentsInnerAttemptsInner) HasGrade() bool {
	if o != nil && !IsNil(o.Grade) {
		return true
	}

	return false
}

// SetGrade gets a reference to the given float32 and assigns it to the Grade field.
func (o *ModLessonGetAttemptsOverview200ResponseDataStudentsInnerAttemptsInner) SetGrade(v float32) {
	o.Grade = &v
}

// GetTimeend returns the Timeend field value if set, zero value otherwise.
func (o *ModLessonGetAttemptsOverview200ResponseDataStudentsInnerAttemptsInner) GetTimeend() int32 {
	if o == nil || IsNil(o.Timeend) {
		var ret int32
		return ret
	}
	return *o.Timeend
}

// GetTimeendOk returns a tuple with the Timeend field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModLessonGetAttemptsOverview200ResponseDataStudentsInnerAttemptsInner) GetTimeendOk() (*int32, bool) {
	if o == nil || IsNil(o.Timeend) {
		return nil, false
	}
	return o.Timeend, true
}

// HasTimeend returns a boolean if a field has been set.
func (o *ModLessonGetAttemptsOverview200ResponseDataStudentsInnerAttemptsInner) HasTimeend() bool {
	if o != nil && !IsNil(o.Timeend) {
		return true
	}

	return false
}

// SetTimeend gets a reference to the given int32 and assigns it to the Timeend field.
func (o *ModLessonGetAttemptsOverview200ResponseDataStudentsInnerAttemptsInner) SetTimeend(v int32) {
	o.Timeend = &v
}

// GetTimestart returns the Timestart field value if set, zero value otherwise.
func (o *ModLessonGetAttemptsOverview200ResponseDataStudentsInnerAttemptsInner) GetTimestart() int32 {
	if o == nil || IsNil(o.Timestart) {
		var ret int32
		return ret
	}
	return *o.Timestart
}

// GetTimestartOk returns a tuple with the Timestart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModLessonGetAttemptsOverview200ResponseDataStudentsInnerAttemptsInner) GetTimestartOk() (*int32, bool) {
	if o == nil || IsNil(o.Timestart) {
		return nil, false
	}
	return o.Timestart, true
}

// HasTimestart returns a boolean if a field has been set.
func (o *ModLessonGetAttemptsOverview200ResponseDataStudentsInnerAttemptsInner) HasTimestart() bool {
	if o != nil && !IsNil(o.Timestart) {
		return true
	}

	return false
}

// SetTimestart gets a reference to the given int32 and assigns it to the Timestart field.
func (o *ModLessonGetAttemptsOverview200ResponseDataStudentsInnerAttemptsInner) SetTimestart(v int32) {
	o.Timestart = &v
}

// GetTry returns the Try field value if set, zero value otherwise.
func (o *ModLessonGetAttemptsOverview200ResponseDataStudentsInnerAttemptsInner) GetTry() int32 {
	if o == nil || IsNil(o.Try) {
		var ret int32
		return ret
	}
	return *o.Try
}

// GetTryOk returns a tuple with the Try field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModLessonGetAttemptsOverview200ResponseDataStudentsInnerAttemptsInner) GetTryOk() (*int32, bool) {
	if o == nil || IsNil(o.Try) {
		return nil, false
	}
	return o.Try, true
}

// HasTry returns a boolean if a field has been set.
func (o *ModLessonGetAttemptsOverview200ResponseDataStudentsInnerAttemptsInner) HasTry() bool {
	if o != nil && !IsNil(o.Try) {
		return true
	}

	return false
}

// SetTry gets a reference to the given int32 and assigns it to the Try field.
func (o *ModLessonGetAttemptsOverview200ResponseDataStudentsInnerAttemptsInner) SetTry(v int32) {
	o.Try = &v
}

func (o ModLessonGetAttemptsOverview200ResponseDataStudentsInnerAttemptsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModLessonGetAttemptsOverview200ResponseDataStudentsInnerAttemptsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.End) {
		toSerialize["end"] = o.End
	}
	if !IsNil(o.Grade) {
		toSerialize["grade"] = o.Grade
	}
	if !IsNil(o.Timeend) {
		toSerialize["timeend"] = o.Timeend
	}
	if !IsNil(o.Timestart) {
		toSerialize["timestart"] = o.Timestart
	}
	if !IsNil(o.Try) {
		toSerialize["try"] = o.Try
	}
	return toSerialize, nil
}

type NullableModLessonGetAttemptsOverview200ResponseDataStudentsInnerAttemptsInner struct {
	value *ModLessonGetAttemptsOverview200ResponseDataStudentsInnerAttemptsInner
	isSet bool
}

func (v NullableModLessonGetAttemptsOverview200ResponseDataStudentsInnerAttemptsInner) Get() *ModLessonGetAttemptsOverview200ResponseDataStudentsInnerAttemptsInner {
	return v.value
}

func (v *NullableModLessonGetAttemptsOverview200ResponseDataStudentsInnerAttemptsInner) Set(val *ModLessonGetAttemptsOverview200ResponseDataStudentsInnerAttemptsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableModLessonGetAttemptsOverview200ResponseDataStudentsInnerAttemptsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableModLessonGetAttemptsOverview200ResponseDataStudentsInnerAttemptsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModLessonGetAttemptsOverview200ResponseDataStudentsInnerAttemptsInner(val *ModLessonGetAttemptsOverview200ResponseDataStudentsInnerAttemptsInner) *NullableModLessonGetAttemptsOverview200ResponseDataStudentsInnerAttemptsInner {
	return &NullableModLessonGetAttemptsOverview200ResponseDataStudentsInnerAttemptsInner{value: val, isSet: true}
}

func (v NullableModLessonGetAttemptsOverview200ResponseDataStudentsInnerAttemptsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModLessonGetAttemptsOverview200ResponseDataStudentsInnerAttemptsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


