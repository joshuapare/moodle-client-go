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

// checks if the CoreGroupGetCourseUserGroupsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreGroupGetCourseUserGroupsRequest{}

// CoreGroupGetCourseUserGroupsRequest struct for CoreGroupGetCourseUserGroupsRequest
type CoreGroupGetCourseUserGroupsRequest struct {
	// Id of course (empty or 0 for all the courses where the user is enrolled).
	Courseid *int32 `json:"courseid,omitempty"`
	// returns only groups in the specified grouping
	Groupingid *int32 `json:"groupingid,omitempty"`
	// Id of user (empty or 0 for current user).
	Userid *int32 `json:"userid,omitempty"`
}

// NewCoreGroupGetCourseUserGroupsRequest instantiates a new CoreGroupGetCourseUserGroupsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreGroupGetCourseUserGroupsRequest() *CoreGroupGetCourseUserGroupsRequest {
	this := CoreGroupGetCourseUserGroupsRequest{}
	var courseid int32 = 0
	this.Courseid = &courseid
	var groupingid int32 = 0
	this.Groupingid = &groupingid
	var userid int32 = 0
	this.Userid = &userid
	return &this
}

// NewCoreGroupGetCourseUserGroupsRequestWithDefaults instantiates a new CoreGroupGetCourseUserGroupsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreGroupGetCourseUserGroupsRequestWithDefaults() *CoreGroupGetCourseUserGroupsRequest {
	this := CoreGroupGetCourseUserGroupsRequest{}
	var courseid int32 = 0
	this.Courseid = &courseid
	var groupingid int32 = 0
	this.Groupingid = &groupingid
	var userid int32 = 0
	this.Userid = &userid
	return &this
}

// GetCourseid returns the Courseid field value if set, zero value otherwise.
func (o *CoreGroupGetCourseUserGroupsRequest) GetCourseid() int32 {
	if o == nil || IsNil(o.Courseid) {
		var ret int32
		return ret
	}
	return *o.Courseid
}

// GetCourseidOk returns a tuple with the Courseid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreGroupGetCourseUserGroupsRequest) GetCourseidOk() (*int32, bool) {
	if o == nil || IsNil(o.Courseid) {
		return nil, false
	}
	return o.Courseid, true
}

// HasCourseid returns a boolean if a field has been set.
func (o *CoreGroupGetCourseUserGroupsRequest) HasCourseid() bool {
	if o != nil && !IsNil(o.Courseid) {
		return true
	}

	return false
}

// SetCourseid gets a reference to the given int32 and assigns it to the Courseid field.
func (o *CoreGroupGetCourseUserGroupsRequest) SetCourseid(v int32) {
	o.Courseid = &v
}

// GetGroupingid returns the Groupingid field value if set, zero value otherwise.
func (o *CoreGroupGetCourseUserGroupsRequest) GetGroupingid() int32 {
	if o == nil || IsNil(o.Groupingid) {
		var ret int32
		return ret
	}
	return *o.Groupingid
}

// GetGroupingidOk returns a tuple with the Groupingid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreGroupGetCourseUserGroupsRequest) GetGroupingidOk() (*int32, bool) {
	if o == nil || IsNil(o.Groupingid) {
		return nil, false
	}
	return o.Groupingid, true
}

// HasGroupingid returns a boolean if a field has been set.
func (o *CoreGroupGetCourseUserGroupsRequest) HasGroupingid() bool {
	if o != nil && !IsNil(o.Groupingid) {
		return true
	}

	return false
}

// SetGroupingid gets a reference to the given int32 and assigns it to the Groupingid field.
func (o *CoreGroupGetCourseUserGroupsRequest) SetGroupingid(v int32) {
	o.Groupingid = &v
}

// GetUserid returns the Userid field value if set, zero value otherwise.
func (o *CoreGroupGetCourseUserGroupsRequest) GetUserid() int32 {
	if o == nil || IsNil(o.Userid) {
		var ret int32
		return ret
	}
	return *o.Userid
}

// GetUseridOk returns a tuple with the Userid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreGroupGetCourseUserGroupsRequest) GetUseridOk() (*int32, bool) {
	if o == nil || IsNil(o.Userid) {
		return nil, false
	}
	return o.Userid, true
}

// HasUserid returns a boolean if a field has been set.
func (o *CoreGroupGetCourseUserGroupsRequest) HasUserid() bool {
	if o != nil && !IsNil(o.Userid) {
		return true
	}

	return false
}

// SetUserid gets a reference to the given int32 and assigns it to the Userid field.
func (o *CoreGroupGetCourseUserGroupsRequest) SetUserid(v int32) {
	o.Userid = &v
}

func (o CoreGroupGetCourseUserGroupsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreGroupGetCourseUserGroupsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Courseid) {
		toSerialize["courseid"] = o.Courseid
	}
	if !IsNil(o.Groupingid) {
		toSerialize["groupingid"] = o.Groupingid
	}
	if !IsNil(o.Userid) {
		toSerialize["userid"] = o.Userid
	}
	return toSerialize, nil
}

type NullableCoreGroupGetCourseUserGroupsRequest struct {
	value *CoreGroupGetCourseUserGroupsRequest
	isSet bool
}

func (v NullableCoreGroupGetCourseUserGroupsRequest) Get() *CoreGroupGetCourseUserGroupsRequest {
	return v.value
}

func (v *NullableCoreGroupGetCourseUserGroupsRequest) Set(val *CoreGroupGetCourseUserGroupsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreGroupGetCourseUserGroupsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreGroupGetCourseUserGroupsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreGroupGetCourseUserGroupsRequest(val *CoreGroupGetCourseUserGroupsRequest) *NullableCoreGroupGetCourseUserGroupsRequest {
	return &NullableCoreGroupGetCourseUserGroupsRequest{value: val, isSet: true}
}

func (v NullableCoreGroupGetCourseUserGroupsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreGroupGetCourseUserGroupsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


