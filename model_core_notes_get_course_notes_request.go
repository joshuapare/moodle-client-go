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

// checks if the CoreNotesGetCourseNotesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreNotesGetCourseNotesRequest{}

// CoreNotesGetCourseNotesRequest struct for CoreNotesGetCourseNotesRequest
type CoreNotesGetCourseNotesRequest struct {
	// course id, 0 for SITE
	Courseid int32 `json:"courseid"`
	// user id
	Userid *int32 `json:"userid,omitempty"`
}

type _CoreNotesGetCourseNotesRequest CoreNotesGetCourseNotesRequest

// NewCoreNotesGetCourseNotesRequest instantiates a new CoreNotesGetCourseNotesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreNotesGetCourseNotesRequest(courseid int32) *CoreNotesGetCourseNotesRequest {
	this := CoreNotesGetCourseNotesRequest{}
	this.Courseid = courseid
	var userid int32 = 0
	this.Userid = &userid
	return &this
}

// NewCoreNotesGetCourseNotesRequestWithDefaults instantiates a new CoreNotesGetCourseNotesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreNotesGetCourseNotesRequestWithDefaults() *CoreNotesGetCourseNotesRequest {
	this := CoreNotesGetCourseNotesRequest{}
	var courseid int32 = null
	this.Courseid = courseid
	var userid int32 = 0
	this.Userid = &userid
	return &this
}

// GetCourseid returns the Courseid field value
func (o *CoreNotesGetCourseNotesRequest) GetCourseid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Courseid
}

// GetCourseidOk returns a tuple with the Courseid field value
// and a boolean to check if the value has been set.
func (o *CoreNotesGetCourseNotesRequest) GetCourseidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Courseid, true
}

// SetCourseid sets field value
func (o *CoreNotesGetCourseNotesRequest) SetCourseid(v int32) {
	o.Courseid = v
}

// GetUserid returns the Userid field value if set, zero value otherwise.
func (o *CoreNotesGetCourseNotesRequest) GetUserid() int32 {
	if o == nil || IsNil(o.Userid) {
		var ret int32
		return ret
	}
	return *o.Userid
}

// GetUseridOk returns a tuple with the Userid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreNotesGetCourseNotesRequest) GetUseridOk() (*int32, bool) {
	if o == nil || IsNil(o.Userid) {
		return nil, false
	}
	return o.Userid, true
}

// HasUserid returns a boolean if a field has been set.
func (o *CoreNotesGetCourseNotesRequest) HasUserid() bool {
	if o != nil && !IsNil(o.Userid) {
		return true
	}

	return false
}

// SetUserid gets a reference to the given int32 and assigns it to the Userid field.
func (o *CoreNotesGetCourseNotesRequest) SetUserid(v int32) {
	o.Userid = &v
}

func (o CoreNotesGetCourseNotesRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreNotesGetCourseNotesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["courseid"] = o.Courseid
	if !IsNil(o.Userid) {
		toSerialize["userid"] = o.Userid
	}
	return toSerialize, nil
}

func (o *CoreNotesGetCourseNotesRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varCoreNotesGetCourseNotesRequest := _CoreNotesGetCourseNotesRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCoreNotesGetCourseNotesRequest)

	if err != nil {
		return err
	}

	*o = CoreNotesGetCourseNotesRequest(varCoreNotesGetCourseNotesRequest)

	return err
}

type NullableCoreNotesGetCourseNotesRequest struct {
	value *CoreNotesGetCourseNotesRequest
	isSet bool
}

func (v NullableCoreNotesGetCourseNotesRequest) Get() *CoreNotesGetCourseNotesRequest {
	return v.value
}

func (v *NullableCoreNotesGetCourseNotesRequest) Set(val *CoreNotesGetCourseNotesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreNotesGetCourseNotesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreNotesGetCourseNotesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreNotesGetCourseNotesRequest(val *CoreNotesGetCourseNotesRequest) *NullableCoreNotesGetCourseNotesRequest {
	return &NullableCoreNotesGetCourseNotesRequest{value: val, isSet: true}
}

func (v NullableCoreNotesGetCourseNotesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreNotesGetCourseNotesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


