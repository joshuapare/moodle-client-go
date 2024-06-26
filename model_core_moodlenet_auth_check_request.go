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

// checks if the CoreMoodlenetAuthCheckRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreMoodlenetAuthCheckRequest{}

// CoreMoodlenetAuthCheckRequest struct for CoreMoodlenetAuthCheckRequest
type CoreMoodlenetAuthCheckRequest struct {
	// Course ID
	Courseid int32 `json:"courseid"`
	// OAuth 2 issuer ID
	Issuerid int32 `json:"issuerid"`
}

type _CoreMoodlenetAuthCheckRequest CoreMoodlenetAuthCheckRequest

// NewCoreMoodlenetAuthCheckRequest instantiates a new CoreMoodlenetAuthCheckRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreMoodlenetAuthCheckRequest(courseid int32, issuerid int32) *CoreMoodlenetAuthCheckRequest {
	this := CoreMoodlenetAuthCheckRequest{}
	this.Courseid = courseid
	this.Issuerid = issuerid
	return &this
}

// NewCoreMoodlenetAuthCheckRequestWithDefaults instantiates a new CoreMoodlenetAuthCheckRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreMoodlenetAuthCheckRequestWithDefaults() *CoreMoodlenetAuthCheckRequest {
	this := CoreMoodlenetAuthCheckRequest{}
	var issuerid int32 = null
	this.Issuerid = issuerid
	return &this
}

// GetCourseid returns the Courseid field value
func (o *CoreMoodlenetAuthCheckRequest) GetCourseid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Courseid
}

// GetCourseidOk returns a tuple with the Courseid field value
// and a boolean to check if the value has been set.
func (o *CoreMoodlenetAuthCheckRequest) GetCourseidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Courseid, true
}

// SetCourseid sets field value
func (o *CoreMoodlenetAuthCheckRequest) SetCourseid(v int32) {
	o.Courseid = v
}

// GetIssuerid returns the Issuerid field value
func (o *CoreMoodlenetAuthCheckRequest) GetIssuerid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Issuerid
}

// GetIssueridOk returns a tuple with the Issuerid field value
// and a boolean to check if the value has been set.
func (o *CoreMoodlenetAuthCheckRequest) GetIssueridOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Issuerid, true
}

// SetIssuerid sets field value
func (o *CoreMoodlenetAuthCheckRequest) SetIssuerid(v int32) {
	o.Issuerid = v
}

func (o CoreMoodlenetAuthCheckRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreMoodlenetAuthCheckRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["courseid"] = o.Courseid
	toSerialize["issuerid"] = o.Issuerid
	return toSerialize, nil
}

func (o *CoreMoodlenetAuthCheckRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"courseid",
		"issuerid",
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

	varCoreMoodlenetAuthCheckRequest := _CoreMoodlenetAuthCheckRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCoreMoodlenetAuthCheckRequest)

	if err != nil {
		return err
	}

	*o = CoreMoodlenetAuthCheckRequest(varCoreMoodlenetAuthCheckRequest)

	return err
}

type NullableCoreMoodlenetAuthCheckRequest struct {
	value *CoreMoodlenetAuthCheckRequest
	isSet bool
}

func (v NullableCoreMoodlenetAuthCheckRequest) Get() *CoreMoodlenetAuthCheckRequest {
	return v.value
}

func (v *NullableCoreMoodlenetAuthCheckRequest) Set(val *CoreMoodlenetAuthCheckRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreMoodlenetAuthCheckRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreMoodlenetAuthCheckRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreMoodlenetAuthCheckRequest(val *CoreMoodlenetAuthCheckRequest) *NullableCoreMoodlenetAuthCheckRequest {
	return &NullableCoreMoodlenetAuthCheckRequest{value: val, isSet: true}
}

func (v NullableCoreMoodlenetAuthCheckRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreMoodlenetAuthCheckRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


