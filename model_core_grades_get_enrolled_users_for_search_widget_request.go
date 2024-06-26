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

// checks if the CoreGradesGetEnrolledUsersForSearchWidgetRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreGradesGetEnrolledUsersForSearchWidgetRequest{}

// CoreGradesGetEnrolledUsersForSearchWidgetRequest struct for CoreGradesGetEnrolledUsersForSearchWidgetRequest
type CoreGradesGetEnrolledUsersForSearchWidgetRequest struct {
	// The base URL for the user option
	Actionbaseurl string `json:"actionbaseurl"`
	// Course Id
	Courseid int32 `json:"courseid"`
	// Group Id
	Groupid *int32 `json:"groupid,omitempty"`
}

type _CoreGradesGetEnrolledUsersForSearchWidgetRequest CoreGradesGetEnrolledUsersForSearchWidgetRequest

// NewCoreGradesGetEnrolledUsersForSearchWidgetRequest instantiates a new CoreGradesGetEnrolledUsersForSearchWidgetRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreGradesGetEnrolledUsersForSearchWidgetRequest(actionbaseurl string, courseid int32) *CoreGradesGetEnrolledUsersForSearchWidgetRequest {
	this := CoreGradesGetEnrolledUsersForSearchWidgetRequest{}
	this.Actionbaseurl = actionbaseurl
	this.Courseid = courseid
	var groupid int32 = 0
	this.Groupid = &groupid
	return &this
}

// NewCoreGradesGetEnrolledUsersForSearchWidgetRequestWithDefaults instantiates a new CoreGradesGetEnrolledUsersForSearchWidgetRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreGradesGetEnrolledUsersForSearchWidgetRequestWithDefaults() *CoreGradesGetEnrolledUsersForSearchWidgetRequest {
	this := CoreGradesGetEnrolledUsersForSearchWidgetRequest{}
	var actionbaseurl string = "null"
	this.Actionbaseurl = actionbaseurl
	var courseid int32 = null
	this.Courseid = courseid
	var groupid int32 = 0
	this.Groupid = &groupid
	return &this
}

// GetActionbaseurl returns the Actionbaseurl field value
func (o *CoreGradesGetEnrolledUsersForSearchWidgetRequest) GetActionbaseurl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Actionbaseurl
}

// GetActionbaseurlOk returns a tuple with the Actionbaseurl field value
// and a boolean to check if the value has been set.
func (o *CoreGradesGetEnrolledUsersForSearchWidgetRequest) GetActionbaseurlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Actionbaseurl, true
}

// SetActionbaseurl sets field value
func (o *CoreGradesGetEnrolledUsersForSearchWidgetRequest) SetActionbaseurl(v string) {
	o.Actionbaseurl = v
}

// GetCourseid returns the Courseid field value
func (o *CoreGradesGetEnrolledUsersForSearchWidgetRequest) GetCourseid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Courseid
}

// GetCourseidOk returns a tuple with the Courseid field value
// and a boolean to check if the value has been set.
func (o *CoreGradesGetEnrolledUsersForSearchWidgetRequest) GetCourseidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Courseid, true
}

// SetCourseid sets field value
func (o *CoreGradesGetEnrolledUsersForSearchWidgetRequest) SetCourseid(v int32) {
	o.Courseid = v
}

// GetGroupid returns the Groupid field value if set, zero value otherwise.
func (o *CoreGradesGetEnrolledUsersForSearchWidgetRequest) GetGroupid() int32 {
	if o == nil || IsNil(o.Groupid) {
		var ret int32
		return ret
	}
	return *o.Groupid
}

// GetGroupidOk returns a tuple with the Groupid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreGradesGetEnrolledUsersForSearchWidgetRequest) GetGroupidOk() (*int32, bool) {
	if o == nil || IsNil(o.Groupid) {
		return nil, false
	}
	return o.Groupid, true
}

// HasGroupid returns a boolean if a field has been set.
func (o *CoreGradesGetEnrolledUsersForSearchWidgetRequest) HasGroupid() bool {
	if o != nil && !IsNil(o.Groupid) {
		return true
	}

	return false
}

// SetGroupid gets a reference to the given int32 and assigns it to the Groupid field.
func (o *CoreGradesGetEnrolledUsersForSearchWidgetRequest) SetGroupid(v int32) {
	o.Groupid = &v
}

func (o CoreGradesGetEnrolledUsersForSearchWidgetRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreGradesGetEnrolledUsersForSearchWidgetRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["actionbaseurl"] = o.Actionbaseurl
	toSerialize["courseid"] = o.Courseid
	if !IsNil(o.Groupid) {
		toSerialize["groupid"] = o.Groupid
	}
	return toSerialize, nil
}

func (o *CoreGradesGetEnrolledUsersForSearchWidgetRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"actionbaseurl",
		"courseid",
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

	varCoreGradesGetEnrolledUsersForSearchWidgetRequest := _CoreGradesGetEnrolledUsersForSearchWidgetRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCoreGradesGetEnrolledUsersForSearchWidgetRequest)

	if err != nil {
		return err
	}

	*o = CoreGradesGetEnrolledUsersForSearchWidgetRequest(varCoreGradesGetEnrolledUsersForSearchWidgetRequest)

	return err
}

type NullableCoreGradesGetEnrolledUsersForSearchWidgetRequest struct {
	value *CoreGradesGetEnrolledUsersForSearchWidgetRequest
	isSet bool
}

func (v NullableCoreGradesGetEnrolledUsersForSearchWidgetRequest) Get() *CoreGradesGetEnrolledUsersForSearchWidgetRequest {
	return v.value
}

func (v *NullableCoreGradesGetEnrolledUsersForSearchWidgetRequest) Set(val *CoreGradesGetEnrolledUsersForSearchWidgetRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreGradesGetEnrolledUsersForSearchWidgetRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreGradesGetEnrolledUsersForSearchWidgetRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreGradesGetEnrolledUsersForSearchWidgetRequest(val *CoreGradesGetEnrolledUsersForSearchWidgetRequest) *NullableCoreGradesGetEnrolledUsersForSearchWidgetRequest {
	return &NullableCoreGradesGetEnrolledUsersForSearchWidgetRequest{value: val, isSet: true}
}

func (v NullableCoreGradesGetEnrolledUsersForSearchWidgetRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreGradesGetEnrolledUsersForSearchWidgetRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


