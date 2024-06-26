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

// checks if the ModFeedbackGetResponsesAnalysis200ResponseAttemptsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModFeedbackGetResponsesAnalysis200ResponseAttemptsInner{}

// ModFeedbackGetResponsesAnalysis200ResponseAttemptsInner struct for ModFeedbackGetResponsesAnalysis200ResponseAttemptsInner
type ModFeedbackGetResponsesAnalysis200ResponseAttemptsInner struct {
	// Course id
	Courseid *int32 `json:"courseid,omitempty"`
	// User full name
	Fullname *string `json:"fullname,omitempty"`
	// Completed id
	Id *int32 `json:"id,omitempty"`
	Responses []ModFeedbackGetResponsesAnalysis200ResponseAttemptsInnerResponsesInner `json:"responses,omitempty"`
	// Time modified for the response
	Timemodified *int32 `json:"timemodified,omitempty"`
	// User who responded
	Userid *int32 `json:"userid,omitempty"`
}

// NewModFeedbackGetResponsesAnalysis200ResponseAttemptsInner instantiates a new ModFeedbackGetResponsesAnalysis200ResponseAttemptsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModFeedbackGetResponsesAnalysis200ResponseAttemptsInner() *ModFeedbackGetResponsesAnalysis200ResponseAttemptsInner {
	this := ModFeedbackGetResponsesAnalysis200ResponseAttemptsInner{}
	var timemodified int32 = null
	this.Timemodified = &timemodified
	var userid int32 = null
	this.Userid = &userid
	return &this
}

// NewModFeedbackGetResponsesAnalysis200ResponseAttemptsInnerWithDefaults instantiates a new ModFeedbackGetResponsesAnalysis200ResponseAttemptsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModFeedbackGetResponsesAnalysis200ResponseAttemptsInnerWithDefaults() *ModFeedbackGetResponsesAnalysis200ResponseAttemptsInner {
	this := ModFeedbackGetResponsesAnalysis200ResponseAttemptsInner{}
	var timemodified int32 = null
	this.Timemodified = &timemodified
	var userid int32 = null
	this.Userid = &userid
	return &this
}

// GetCourseid returns the Courseid field value if set, zero value otherwise.
func (o *ModFeedbackGetResponsesAnalysis200ResponseAttemptsInner) GetCourseid() int32 {
	if o == nil || IsNil(o.Courseid) {
		var ret int32
		return ret
	}
	return *o.Courseid
}

// GetCourseidOk returns a tuple with the Courseid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModFeedbackGetResponsesAnalysis200ResponseAttemptsInner) GetCourseidOk() (*int32, bool) {
	if o == nil || IsNil(o.Courseid) {
		return nil, false
	}
	return o.Courseid, true
}

// HasCourseid returns a boolean if a field has been set.
func (o *ModFeedbackGetResponsesAnalysis200ResponseAttemptsInner) HasCourseid() bool {
	if o != nil && !IsNil(o.Courseid) {
		return true
	}

	return false
}

// SetCourseid gets a reference to the given int32 and assigns it to the Courseid field.
func (o *ModFeedbackGetResponsesAnalysis200ResponseAttemptsInner) SetCourseid(v int32) {
	o.Courseid = &v
}

// GetFullname returns the Fullname field value if set, zero value otherwise.
func (o *ModFeedbackGetResponsesAnalysis200ResponseAttemptsInner) GetFullname() string {
	if o == nil || IsNil(o.Fullname) {
		var ret string
		return ret
	}
	return *o.Fullname
}

// GetFullnameOk returns a tuple with the Fullname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModFeedbackGetResponsesAnalysis200ResponseAttemptsInner) GetFullnameOk() (*string, bool) {
	if o == nil || IsNil(o.Fullname) {
		return nil, false
	}
	return o.Fullname, true
}

// HasFullname returns a boolean if a field has been set.
func (o *ModFeedbackGetResponsesAnalysis200ResponseAttemptsInner) HasFullname() bool {
	if o != nil && !IsNil(o.Fullname) {
		return true
	}

	return false
}

// SetFullname gets a reference to the given string and assigns it to the Fullname field.
func (o *ModFeedbackGetResponsesAnalysis200ResponseAttemptsInner) SetFullname(v string) {
	o.Fullname = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ModFeedbackGetResponsesAnalysis200ResponseAttemptsInner) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModFeedbackGetResponsesAnalysis200ResponseAttemptsInner) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ModFeedbackGetResponsesAnalysis200ResponseAttemptsInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ModFeedbackGetResponsesAnalysis200ResponseAttemptsInner) SetId(v int32) {
	o.Id = &v
}

// GetResponses returns the Responses field value if set, zero value otherwise.
func (o *ModFeedbackGetResponsesAnalysis200ResponseAttemptsInner) GetResponses() []ModFeedbackGetResponsesAnalysis200ResponseAttemptsInnerResponsesInner {
	if o == nil || IsNil(o.Responses) {
		var ret []ModFeedbackGetResponsesAnalysis200ResponseAttemptsInnerResponsesInner
		return ret
	}
	return o.Responses
}

// GetResponsesOk returns a tuple with the Responses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModFeedbackGetResponsesAnalysis200ResponseAttemptsInner) GetResponsesOk() ([]ModFeedbackGetResponsesAnalysis200ResponseAttemptsInnerResponsesInner, bool) {
	if o == nil || IsNil(o.Responses) {
		return nil, false
	}
	return o.Responses, true
}

// HasResponses returns a boolean if a field has been set.
func (o *ModFeedbackGetResponsesAnalysis200ResponseAttemptsInner) HasResponses() bool {
	if o != nil && !IsNil(o.Responses) {
		return true
	}

	return false
}

// SetResponses gets a reference to the given []ModFeedbackGetResponsesAnalysis200ResponseAttemptsInnerResponsesInner and assigns it to the Responses field.
func (o *ModFeedbackGetResponsesAnalysis200ResponseAttemptsInner) SetResponses(v []ModFeedbackGetResponsesAnalysis200ResponseAttemptsInnerResponsesInner) {
	o.Responses = v
}

// GetTimemodified returns the Timemodified field value if set, zero value otherwise.
func (o *ModFeedbackGetResponsesAnalysis200ResponseAttemptsInner) GetTimemodified() int32 {
	if o == nil || IsNil(o.Timemodified) {
		var ret int32
		return ret
	}
	return *o.Timemodified
}

// GetTimemodifiedOk returns a tuple with the Timemodified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModFeedbackGetResponsesAnalysis200ResponseAttemptsInner) GetTimemodifiedOk() (*int32, bool) {
	if o == nil || IsNil(o.Timemodified) {
		return nil, false
	}
	return o.Timemodified, true
}

// HasTimemodified returns a boolean if a field has been set.
func (o *ModFeedbackGetResponsesAnalysis200ResponseAttemptsInner) HasTimemodified() bool {
	if o != nil && !IsNil(o.Timemodified) {
		return true
	}

	return false
}

// SetTimemodified gets a reference to the given int32 and assigns it to the Timemodified field.
func (o *ModFeedbackGetResponsesAnalysis200ResponseAttemptsInner) SetTimemodified(v int32) {
	o.Timemodified = &v
}

// GetUserid returns the Userid field value if set, zero value otherwise.
func (o *ModFeedbackGetResponsesAnalysis200ResponseAttemptsInner) GetUserid() int32 {
	if o == nil || IsNil(o.Userid) {
		var ret int32
		return ret
	}
	return *o.Userid
}

// GetUseridOk returns a tuple with the Userid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModFeedbackGetResponsesAnalysis200ResponseAttemptsInner) GetUseridOk() (*int32, bool) {
	if o == nil || IsNil(o.Userid) {
		return nil, false
	}
	return o.Userid, true
}

// HasUserid returns a boolean if a field has been set.
func (o *ModFeedbackGetResponsesAnalysis200ResponseAttemptsInner) HasUserid() bool {
	if o != nil && !IsNil(o.Userid) {
		return true
	}

	return false
}

// SetUserid gets a reference to the given int32 and assigns it to the Userid field.
func (o *ModFeedbackGetResponsesAnalysis200ResponseAttemptsInner) SetUserid(v int32) {
	o.Userid = &v
}

func (o ModFeedbackGetResponsesAnalysis200ResponseAttemptsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModFeedbackGetResponsesAnalysis200ResponseAttemptsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Courseid) {
		toSerialize["courseid"] = o.Courseid
	}
	if !IsNil(o.Fullname) {
		toSerialize["fullname"] = o.Fullname
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Responses) {
		toSerialize["responses"] = o.Responses
	}
	if !IsNil(o.Timemodified) {
		toSerialize["timemodified"] = o.Timemodified
	}
	if !IsNil(o.Userid) {
		toSerialize["userid"] = o.Userid
	}
	return toSerialize, nil
}

type NullableModFeedbackGetResponsesAnalysis200ResponseAttemptsInner struct {
	value *ModFeedbackGetResponsesAnalysis200ResponseAttemptsInner
	isSet bool
}

func (v NullableModFeedbackGetResponsesAnalysis200ResponseAttemptsInner) Get() *ModFeedbackGetResponsesAnalysis200ResponseAttemptsInner {
	return v.value
}

func (v *NullableModFeedbackGetResponsesAnalysis200ResponseAttemptsInner) Set(val *ModFeedbackGetResponsesAnalysis200ResponseAttemptsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableModFeedbackGetResponsesAnalysis200ResponseAttemptsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableModFeedbackGetResponsesAnalysis200ResponseAttemptsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModFeedbackGetResponsesAnalysis200ResponseAttemptsInner(val *ModFeedbackGetResponsesAnalysis200ResponseAttemptsInner) *NullableModFeedbackGetResponsesAnalysis200ResponseAttemptsInner {
	return &NullableModFeedbackGetResponsesAnalysis200ResponseAttemptsInner{value: val, isSet: true}
}

func (v NullableModFeedbackGetResponsesAnalysis200ResponseAttemptsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModFeedbackGetResponsesAnalysis200ResponseAttemptsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


