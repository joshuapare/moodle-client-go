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

// checks if the EnrolManualEnrolUsersRequestEnrolmentsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EnrolManualEnrolUsersRequestEnrolmentsInner{}

// EnrolManualEnrolUsersRequestEnrolmentsInner struct for EnrolManualEnrolUsersRequestEnrolmentsInner
type EnrolManualEnrolUsersRequestEnrolmentsInner struct {
	// The course to enrol the user role in
	Courseid *int32 `json:"courseid,omitempty"`
	// Role to assign to the user
	Roleid *int32 `json:"roleid,omitempty"`
	// set to 1 to suspend the enrolment
	Suspend *int32 `json:"suspend,omitempty"`
	// Timestamp when the enrolment end
	Timeend *int32 `json:"timeend,omitempty"`
	// Timestamp when the enrolment start
	Timestart *int32 `json:"timestart,omitempty"`
	// The user that is going to be enrolled
	Userid *int32 `json:"userid,omitempty"`
}

// NewEnrolManualEnrolUsersRequestEnrolmentsInner instantiates a new EnrolManualEnrolUsersRequestEnrolmentsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnrolManualEnrolUsersRequestEnrolmentsInner() *EnrolManualEnrolUsersRequestEnrolmentsInner {
	this := EnrolManualEnrolUsersRequestEnrolmentsInner{}
	return &this
}

// NewEnrolManualEnrolUsersRequestEnrolmentsInnerWithDefaults instantiates a new EnrolManualEnrolUsersRequestEnrolmentsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnrolManualEnrolUsersRequestEnrolmentsInnerWithDefaults() *EnrolManualEnrolUsersRequestEnrolmentsInner {
	this := EnrolManualEnrolUsersRequestEnrolmentsInner{}
	return &this
}

// GetCourseid returns the Courseid field value if set, zero value otherwise.
func (o *EnrolManualEnrolUsersRequestEnrolmentsInner) GetCourseid() int32 {
	if o == nil || IsNil(o.Courseid) {
		var ret int32
		return ret
	}
	return *o.Courseid
}

// GetCourseidOk returns a tuple with the Courseid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnrolManualEnrolUsersRequestEnrolmentsInner) GetCourseidOk() (*int32, bool) {
	if o == nil || IsNil(o.Courseid) {
		return nil, false
	}
	return o.Courseid, true
}

// HasCourseid returns a boolean if a field has been set.
func (o *EnrolManualEnrolUsersRequestEnrolmentsInner) HasCourseid() bool {
	if o != nil && !IsNil(o.Courseid) {
		return true
	}

	return false
}

// SetCourseid gets a reference to the given int32 and assigns it to the Courseid field.
func (o *EnrolManualEnrolUsersRequestEnrolmentsInner) SetCourseid(v int32) {
	o.Courseid = &v
}

// GetRoleid returns the Roleid field value if set, zero value otherwise.
func (o *EnrolManualEnrolUsersRequestEnrolmentsInner) GetRoleid() int32 {
	if o == nil || IsNil(o.Roleid) {
		var ret int32
		return ret
	}
	return *o.Roleid
}

// GetRoleidOk returns a tuple with the Roleid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnrolManualEnrolUsersRequestEnrolmentsInner) GetRoleidOk() (*int32, bool) {
	if o == nil || IsNil(o.Roleid) {
		return nil, false
	}
	return o.Roleid, true
}

// HasRoleid returns a boolean if a field has been set.
func (o *EnrolManualEnrolUsersRequestEnrolmentsInner) HasRoleid() bool {
	if o != nil && !IsNil(o.Roleid) {
		return true
	}

	return false
}

// SetRoleid gets a reference to the given int32 and assigns it to the Roleid field.
func (o *EnrolManualEnrolUsersRequestEnrolmentsInner) SetRoleid(v int32) {
	o.Roleid = &v
}

// GetSuspend returns the Suspend field value if set, zero value otherwise.
func (o *EnrolManualEnrolUsersRequestEnrolmentsInner) GetSuspend() int32 {
	if o == nil || IsNil(o.Suspend) {
		var ret int32
		return ret
	}
	return *o.Suspend
}

// GetSuspendOk returns a tuple with the Suspend field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnrolManualEnrolUsersRequestEnrolmentsInner) GetSuspendOk() (*int32, bool) {
	if o == nil || IsNil(o.Suspend) {
		return nil, false
	}
	return o.Suspend, true
}

// HasSuspend returns a boolean if a field has been set.
func (o *EnrolManualEnrolUsersRequestEnrolmentsInner) HasSuspend() bool {
	if o != nil && !IsNil(o.Suspend) {
		return true
	}

	return false
}

// SetSuspend gets a reference to the given int32 and assigns it to the Suspend field.
func (o *EnrolManualEnrolUsersRequestEnrolmentsInner) SetSuspend(v int32) {
	o.Suspend = &v
}

// GetTimeend returns the Timeend field value if set, zero value otherwise.
func (o *EnrolManualEnrolUsersRequestEnrolmentsInner) GetTimeend() int32 {
	if o == nil || IsNil(o.Timeend) {
		var ret int32
		return ret
	}
	return *o.Timeend
}

// GetTimeendOk returns a tuple with the Timeend field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnrolManualEnrolUsersRequestEnrolmentsInner) GetTimeendOk() (*int32, bool) {
	if o == nil || IsNil(o.Timeend) {
		return nil, false
	}
	return o.Timeend, true
}

// HasTimeend returns a boolean if a field has been set.
func (o *EnrolManualEnrolUsersRequestEnrolmentsInner) HasTimeend() bool {
	if o != nil && !IsNil(o.Timeend) {
		return true
	}

	return false
}

// SetTimeend gets a reference to the given int32 and assigns it to the Timeend field.
func (o *EnrolManualEnrolUsersRequestEnrolmentsInner) SetTimeend(v int32) {
	o.Timeend = &v
}

// GetTimestart returns the Timestart field value if set, zero value otherwise.
func (o *EnrolManualEnrolUsersRequestEnrolmentsInner) GetTimestart() int32 {
	if o == nil || IsNil(o.Timestart) {
		var ret int32
		return ret
	}
	return *o.Timestart
}

// GetTimestartOk returns a tuple with the Timestart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnrolManualEnrolUsersRequestEnrolmentsInner) GetTimestartOk() (*int32, bool) {
	if o == nil || IsNil(o.Timestart) {
		return nil, false
	}
	return o.Timestart, true
}

// HasTimestart returns a boolean if a field has been set.
func (o *EnrolManualEnrolUsersRequestEnrolmentsInner) HasTimestart() bool {
	if o != nil && !IsNil(o.Timestart) {
		return true
	}

	return false
}

// SetTimestart gets a reference to the given int32 and assigns it to the Timestart field.
func (o *EnrolManualEnrolUsersRequestEnrolmentsInner) SetTimestart(v int32) {
	o.Timestart = &v
}

// GetUserid returns the Userid field value if set, zero value otherwise.
func (o *EnrolManualEnrolUsersRequestEnrolmentsInner) GetUserid() int32 {
	if o == nil || IsNil(o.Userid) {
		var ret int32
		return ret
	}
	return *o.Userid
}

// GetUseridOk returns a tuple with the Userid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnrolManualEnrolUsersRequestEnrolmentsInner) GetUseridOk() (*int32, bool) {
	if o == nil || IsNil(o.Userid) {
		return nil, false
	}
	return o.Userid, true
}

// HasUserid returns a boolean if a field has been set.
func (o *EnrolManualEnrolUsersRequestEnrolmentsInner) HasUserid() bool {
	if o != nil && !IsNil(o.Userid) {
		return true
	}

	return false
}

// SetUserid gets a reference to the given int32 and assigns it to the Userid field.
func (o *EnrolManualEnrolUsersRequestEnrolmentsInner) SetUserid(v int32) {
	o.Userid = &v
}

func (o EnrolManualEnrolUsersRequestEnrolmentsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnrolManualEnrolUsersRequestEnrolmentsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Courseid) {
		toSerialize["courseid"] = o.Courseid
	}
	if !IsNil(o.Roleid) {
		toSerialize["roleid"] = o.Roleid
	}
	if !IsNil(o.Suspend) {
		toSerialize["suspend"] = o.Suspend
	}
	if !IsNil(o.Timeend) {
		toSerialize["timeend"] = o.Timeend
	}
	if !IsNil(o.Timestart) {
		toSerialize["timestart"] = o.Timestart
	}
	if !IsNil(o.Userid) {
		toSerialize["userid"] = o.Userid
	}
	return toSerialize, nil
}

type NullableEnrolManualEnrolUsersRequestEnrolmentsInner struct {
	value *EnrolManualEnrolUsersRequestEnrolmentsInner
	isSet bool
}

func (v NullableEnrolManualEnrolUsersRequestEnrolmentsInner) Get() *EnrolManualEnrolUsersRequestEnrolmentsInner {
	return v.value
}

func (v *NullableEnrolManualEnrolUsersRequestEnrolmentsInner) Set(val *EnrolManualEnrolUsersRequestEnrolmentsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableEnrolManualEnrolUsersRequestEnrolmentsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableEnrolManualEnrolUsersRequestEnrolmentsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnrolManualEnrolUsersRequestEnrolmentsInner(val *EnrolManualEnrolUsersRequestEnrolmentsInner) *NullableEnrolManualEnrolUsersRequestEnrolmentsInner {
	return &NullableEnrolManualEnrolUsersRequestEnrolmentsInner{value: val, isSet: true}
}

func (v NullableEnrolManualEnrolUsersRequestEnrolmentsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnrolManualEnrolUsersRequestEnrolmentsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

