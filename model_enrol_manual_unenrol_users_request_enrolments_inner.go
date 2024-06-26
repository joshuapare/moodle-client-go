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

// checks if the EnrolManualUnenrolUsersRequestEnrolmentsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EnrolManualUnenrolUsersRequestEnrolmentsInner{}

// EnrolManualUnenrolUsersRequestEnrolmentsInner struct for EnrolManualUnenrolUsersRequestEnrolmentsInner
type EnrolManualUnenrolUsersRequestEnrolmentsInner struct {
	// The course to unenrol the user from
	Courseid *int32 `json:"courseid,omitempty"`
	// The user role
	Roleid *int32 `json:"roleid,omitempty"`
	// The user that is going to be unenrolled
	Userid *int32 `json:"userid,omitempty"`
}

// NewEnrolManualUnenrolUsersRequestEnrolmentsInner instantiates a new EnrolManualUnenrolUsersRequestEnrolmentsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnrolManualUnenrolUsersRequestEnrolmentsInner() *EnrolManualUnenrolUsersRequestEnrolmentsInner {
	this := EnrolManualUnenrolUsersRequestEnrolmentsInner{}
	var courseid int32 = null
	this.Courseid = &courseid
	var roleid int32 = null
	this.Roleid = &roleid
	var userid int32 = null
	this.Userid = &userid
	return &this
}

// NewEnrolManualUnenrolUsersRequestEnrolmentsInnerWithDefaults instantiates a new EnrolManualUnenrolUsersRequestEnrolmentsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnrolManualUnenrolUsersRequestEnrolmentsInnerWithDefaults() *EnrolManualUnenrolUsersRequestEnrolmentsInner {
	this := EnrolManualUnenrolUsersRequestEnrolmentsInner{}
	var courseid int32 = null
	this.Courseid = &courseid
	var roleid int32 = null
	this.Roleid = &roleid
	var userid int32 = null
	this.Userid = &userid
	return &this
}

// GetCourseid returns the Courseid field value if set, zero value otherwise.
func (o *EnrolManualUnenrolUsersRequestEnrolmentsInner) GetCourseid() int32 {
	if o == nil || IsNil(o.Courseid) {
		var ret int32
		return ret
	}
	return *o.Courseid
}

// GetCourseidOk returns a tuple with the Courseid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnrolManualUnenrolUsersRequestEnrolmentsInner) GetCourseidOk() (*int32, bool) {
	if o == nil || IsNil(o.Courseid) {
		return nil, false
	}
	return o.Courseid, true
}

// HasCourseid returns a boolean if a field has been set.
func (o *EnrolManualUnenrolUsersRequestEnrolmentsInner) HasCourseid() bool {
	if o != nil && !IsNil(o.Courseid) {
		return true
	}

	return false
}

// SetCourseid gets a reference to the given int32 and assigns it to the Courseid field.
func (o *EnrolManualUnenrolUsersRequestEnrolmentsInner) SetCourseid(v int32) {
	o.Courseid = &v
}

// GetRoleid returns the Roleid field value if set, zero value otherwise.
func (o *EnrolManualUnenrolUsersRequestEnrolmentsInner) GetRoleid() int32 {
	if o == nil || IsNil(o.Roleid) {
		var ret int32
		return ret
	}
	return *o.Roleid
}

// GetRoleidOk returns a tuple with the Roleid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnrolManualUnenrolUsersRequestEnrolmentsInner) GetRoleidOk() (*int32, bool) {
	if o == nil || IsNil(o.Roleid) {
		return nil, false
	}
	return o.Roleid, true
}

// HasRoleid returns a boolean if a field has been set.
func (o *EnrolManualUnenrolUsersRequestEnrolmentsInner) HasRoleid() bool {
	if o != nil && !IsNil(o.Roleid) {
		return true
	}

	return false
}

// SetRoleid gets a reference to the given int32 and assigns it to the Roleid field.
func (o *EnrolManualUnenrolUsersRequestEnrolmentsInner) SetRoleid(v int32) {
	o.Roleid = &v
}

// GetUserid returns the Userid field value if set, zero value otherwise.
func (o *EnrolManualUnenrolUsersRequestEnrolmentsInner) GetUserid() int32 {
	if o == nil || IsNil(o.Userid) {
		var ret int32
		return ret
	}
	return *o.Userid
}

// GetUseridOk returns a tuple with the Userid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnrolManualUnenrolUsersRequestEnrolmentsInner) GetUseridOk() (*int32, bool) {
	if o == nil || IsNil(o.Userid) {
		return nil, false
	}
	return o.Userid, true
}

// HasUserid returns a boolean if a field has been set.
func (o *EnrolManualUnenrolUsersRequestEnrolmentsInner) HasUserid() bool {
	if o != nil && !IsNil(o.Userid) {
		return true
	}

	return false
}

// SetUserid gets a reference to the given int32 and assigns it to the Userid field.
func (o *EnrolManualUnenrolUsersRequestEnrolmentsInner) SetUserid(v int32) {
	o.Userid = &v
}

func (o EnrolManualUnenrolUsersRequestEnrolmentsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnrolManualUnenrolUsersRequestEnrolmentsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Courseid) {
		toSerialize["courseid"] = o.Courseid
	}
	if !IsNil(o.Roleid) {
		toSerialize["roleid"] = o.Roleid
	}
	if !IsNil(o.Userid) {
		toSerialize["userid"] = o.Userid
	}
	return toSerialize, nil
}

type NullableEnrolManualUnenrolUsersRequestEnrolmentsInner struct {
	value *EnrolManualUnenrolUsersRequestEnrolmentsInner
	isSet bool
}

func (v NullableEnrolManualUnenrolUsersRequestEnrolmentsInner) Get() *EnrolManualUnenrolUsersRequestEnrolmentsInner {
	return v.value
}

func (v *NullableEnrolManualUnenrolUsersRequestEnrolmentsInner) Set(val *EnrolManualUnenrolUsersRequestEnrolmentsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableEnrolManualUnenrolUsersRequestEnrolmentsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableEnrolManualUnenrolUsersRequestEnrolmentsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnrolManualUnenrolUsersRequestEnrolmentsInner(val *EnrolManualUnenrolUsersRequestEnrolmentsInner) *NullableEnrolManualUnenrolUsersRequestEnrolmentsInner {
	return &NullableEnrolManualUnenrolUsersRequestEnrolmentsInner{value: val, isSet: true}
}

func (v NullableEnrolManualUnenrolUsersRequestEnrolmentsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnrolManualUnenrolUsersRequestEnrolmentsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


