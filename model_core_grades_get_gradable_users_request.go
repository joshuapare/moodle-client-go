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

// checks if the CoreGradesGetGradableUsersRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreGradesGetGradableUsersRequest{}

// CoreGradesGetGradableUsersRequest struct for CoreGradesGetGradableUsersRequest
type CoreGradesGetGradableUsersRequest struct {
	// Course Id
	Courseid int32 `json:"courseid"`
	// Group Id
	Groupid *int32 `json:"groupid,omitempty"`
	// Only active enrolment
	Onlyactive *bool `json:"onlyactive,omitempty"`
}

type _CoreGradesGetGradableUsersRequest CoreGradesGetGradableUsersRequest

// NewCoreGradesGetGradableUsersRequest instantiates a new CoreGradesGetGradableUsersRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreGradesGetGradableUsersRequest(courseid int32) *CoreGradesGetGradableUsersRequest {
	this := CoreGradesGetGradableUsersRequest{}
	this.Courseid = courseid
	var groupid int32 = 0
	this.Groupid = &groupid
	var onlyactive bool = false
	this.Onlyactive = &onlyactive
	return &this
}

// NewCoreGradesGetGradableUsersRequestWithDefaults instantiates a new CoreGradesGetGradableUsersRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreGradesGetGradableUsersRequestWithDefaults() *CoreGradesGetGradableUsersRequest {
	this := CoreGradesGetGradableUsersRequest{}
	var groupid int32 = 0
	this.Groupid = &groupid
	var onlyactive bool = false
	this.Onlyactive = &onlyactive
	return &this
}

// GetCourseid returns the Courseid field value
func (o *CoreGradesGetGradableUsersRequest) GetCourseid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Courseid
}

// GetCourseidOk returns a tuple with the Courseid field value
// and a boolean to check if the value has been set.
func (o *CoreGradesGetGradableUsersRequest) GetCourseidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Courseid, true
}

// SetCourseid sets field value
func (o *CoreGradesGetGradableUsersRequest) SetCourseid(v int32) {
	o.Courseid = v
}

// GetGroupid returns the Groupid field value if set, zero value otherwise.
func (o *CoreGradesGetGradableUsersRequest) GetGroupid() int32 {
	if o == nil || IsNil(o.Groupid) {
		var ret int32
		return ret
	}
	return *o.Groupid
}

// GetGroupidOk returns a tuple with the Groupid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreGradesGetGradableUsersRequest) GetGroupidOk() (*int32, bool) {
	if o == nil || IsNil(o.Groupid) {
		return nil, false
	}
	return o.Groupid, true
}

// HasGroupid returns a boolean if a field has been set.
func (o *CoreGradesGetGradableUsersRequest) HasGroupid() bool {
	if o != nil && !IsNil(o.Groupid) {
		return true
	}

	return false
}

// SetGroupid gets a reference to the given int32 and assigns it to the Groupid field.
func (o *CoreGradesGetGradableUsersRequest) SetGroupid(v int32) {
	o.Groupid = &v
}

// GetOnlyactive returns the Onlyactive field value if set, zero value otherwise.
func (o *CoreGradesGetGradableUsersRequest) GetOnlyactive() bool {
	if o == nil || IsNil(o.Onlyactive) {
		var ret bool
		return ret
	}
	return *o.Onlyactive
}

// GetOnlyactiveOk returns a tuple with the Onlyactive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreGradesGetGradableUsersRequest) GetOnlyactiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Onlyactive) {
		return nil, false
	}
	return o.Onlyactive, true
}

// HasOnlyactive returns a boolean if a field has been set.
func (o *CoreGradesGetGradableUsersRequest) HasOnlyactive() bool {
	if o != nil && !IsNil(o.Onlyactive) {
		return true
	}

	return false
}

// SetOnlyactive gets a reference to the given bool and assigns it to the Onlyactive field.
func (o *CoreGradesGetGradableUsersRequest) SetOnlyactive(v bool) {
	o.Onlyactive = &v
}

func (o CoreGradesGetGradableUsersRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreGradesGetGradableUsersRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["courseid"] = o.Courseid
	if !IsNil(o.Groupid) {
		toSerialize["groupid"] = o.Groupid
	}
	if !IsNil(o.Onlyactive) {
		toSerialize["onlyactive"] = o.Onlyactive
	}
	return toSerialize, nil
}

func (o *CoreGradesGetGradableUsersRequest) UnmarshalJSON(data []byte) (err error) {
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

	varCoreGradesGetGradableUsersRequest := _CoreGradesGetGradableUsersRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCoreGradesGetGradableUsersRequest)

	if err != nil {
		return err
	}

	*o = CoreGradesGetGradableUsersRequest(varCoreGradesGetGradableUsersRequest)

	return err
}

type NullableCoreGradesGetGradableUsersRequest struct {
	value *CoreGradesGetGradableUsersRequest
	isSet bool
}

func (v NullableCoreGradesGetGradableUsersRequest) Get() *CoreGradesGetGradableUsersRequest {
	return v.value
}

func (v *NullableCoreGradesGetGradableUsersRequest) Set(val *CoreGradesGetGradableUsersRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreGradesGetGradableUsersRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreGradesGetGradableUsersRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreGradesGetGradableUsersRequest(val *CoreGradesGetGradableUsersRequest) *NullableCoreGradesGetGradableUsersRequest {
	return &NullableCoreGradesGetGradableUsersRequest{value: val, isSet: true}
}

func (v NullableCoreGradesGetGradableUsersRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreGradesGetGradableUsersRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


