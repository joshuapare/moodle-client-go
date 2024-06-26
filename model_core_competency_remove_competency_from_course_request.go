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

// checks if the CoreCompetencyRemoveCompetencyFromCourseRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreCompetencyRemoveCompetencyFromCourseRequest{}

// CoreCompetencyRemoveCompetencyFromCourseRequest struct for CoreCompetencyRemoveCompetencyFromCourseRequest
type CoreCompetencyRemoveCompetencyFromCourseRequest struct {
	// The competency id
	Competencyid int32 `json:"competencyid"`
	// The course id
	Courseid int32 `json:"courseid"`
}

type _CoreCompetencyRemoveCompetencyFromCourseRequest CoreCompetencyRemoveCompetencyFromCourseRequest

// NewCoreCompetencyRemoveCompetencyFromCourseRequest instantiates a new CoreCompetencyRemoveCompetencyFromCourseRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreCompetencyRemoveCompetencyFromCourseRequest(competencyid int32, courseid int32) *CoreCompetencyRemoveCompetencyFromCourseRequest {
	this := CoreCompetencyRemoveCompetencyFromCourseRequest{}
	this.Competencyid = competencyid
	this.Courseid = courseid
	return &this
}

// NewCoreCompetencyRemoveCompetencyFromCourseRequestWithDefaults instantiates a new CoreCompetencyRemoveCompetencyFromCourseRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreCompetencyRemoveCompetencyFromCourseRequestWithDefaults() *CoreCompetencyRemoveCompetencyFromCourseRequest {
	this := CoreCompetencyRemoveCompetencyFromCourseRequest{}
	return &this
}

// GetCompetencyid returns the Competencyid field value
func (o *CoreCompetencyRemoveCompetencyFromCourseRequest) GetCompetencyid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Competencyid
}

// GetCompetencyidOk returns a tuple with the Competencyid field value
// and a boolean to check if the value has been set.
func (o *CoreCompetencyRemoveCompetencyFromCourseRequest) GetCompetencyidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Competencyid, true
}

// SetCompetencyid sets field value
func (o *CoreCompetencyRemoveCompetencyFromCourseRequest) SetCompetencyid(v int32) {
	o.Competencyid = v
}

// GetCourseid returns the Courseid field value
func (o *CoreCompetencyRemoveCompetencyFromCourseRequest) GetCourseid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Courseid
}

// GetCourseidOk returns a tuple with the Courseid field value
// and a boolean to check if the value has been set.
func (o *CoreCompetencyRemoveCompetencyFromCourseRequest) GetCourseidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Courseid, true
}

// SetCourseid sets field value
func (o *CoreCompetencyRemoveCompetencyFromCourseRequest) SetCourseid(v int32) {
	o.Courseid = v
}

func (o CoreCompetencyRemoveCompetencyFromCourseRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreCompetencyRemoveCompetencyFromCourseRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["competencyid"] = o.Competencyid
	toSerialize["courseid"] = o.Courseid
	return toSerialize, nil
}

func (o *CoreCompetencyRemoveCompetencyFromCourseRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"competencyid",
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

	varCoreCompetencyRemoveCompetencyFromCourseRequest := _CoreCompetencyRemoveCompetencyFromCourseRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCoreCompetencyRemoveCompetencyFromCourseRequest)

	if err != nil {
		return err
	}

	*o = CoreCompetencyRemoveCompetencyFromCourseRequest(varCoreCompetencyRemoveCompetencyFromCourseRequest)

	return err
}

type NullableCoreCompetencyRemoveCompetencyFromCourseRequest struct {
	value *CoreCompetencyRemoveCompetencyFromCourseRequest
	isSet bool
}

func (v NullableCoreCompetencyRemoveCompetencyFromCourseRequest) Get() *CoreCompetencyRemoveCompetencyFromCourseRequest {
	return v.value
}

func (v *NullableCoreCompetencyRemoveCompetencyFromCourseRequest) Set(val *CoreCompetencyRemoveCompetencyFromCourseRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreCompetencyRemoveCompetencyFromCourseRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreCompetencyRemoveCompetencyFromCourseRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreCompetencyRemoveCompetencyFromCourseRequest(val *CoreCompetencyRemoveCompetencyFromCourseRequest) *NullableCoreCompetencyRemoveCompetencyFromCourseRequest {
	return &NullableCoreCompetencyRemoveCompetencyFromCourseRequest{value: val, isSet: true}
}

func (v NullableCoreCompetencyRemoveCompetencyFromCourseRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreCompetencyRemoveCompetencyFromCourseRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


