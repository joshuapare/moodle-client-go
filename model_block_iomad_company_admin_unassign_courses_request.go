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

// checks if the BlockIomadCompanyAdminUnassignCoursesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BlockIomadCompanyAdminUnassignCoursesRequest{}

// BlockIomadCompanyAdminUnassignCoursesRequest struct for BlockIomadCompanyAdminUnassignCoursesRequest
type BlockIomadCompanyAdminUnassignCoursesRequest struct {
	Courses []BlockIomadCompanyAdminUnassignCoursesRequestCoursesInner `json:"courses"`
}

type _BlockIomadCompanyAdminUnassignCoursesRequest BlockIomadCompanyAdminUnassignCoursesRequest

// NewBlockIomadCompanyAdminUnassignCoursesRequest instantiates a new BlockIomadCompanyAdminUnassignCoursesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlockIomadCompanyAdminUnassignCoursesRequest(courses []BlockIomadCompanyAdminUnassignCoursesRequestCoursesInner) *BlockIomadCompanyAdminUnassignCoursesRequest {
	this := BlockIomadCompanyAdminUnassignCoursesRequest{}
	this.Courses = courses
	return &this
}

// NewBlockIomadCompanyAdminUnassignCoursesRequestWithDefaults instantiates a new BlockIomadCompanyAdminUnassignCoursesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlockIomadCompanyAdminUnassignCoursesRequestWithDefaults() *BlockIomadCompanyAdminUnassignCoursesRequest {
	this := BlockIomadCompanyAdminUnassignCoursesRequest{}
	return &this
}

// GetCourses returns the Courses field value
func (o *BlockIomadCompanyAdminUnassignCoursesRequest) GetCourses() []BlockIomadCompanyAdminUnassignCoursesRequestCoursesInner {
	if o == nil {
		var ret []BlockIomadCompanyAdminUnassignCoursesRequestCoursesInner
		return ret
	}

	return o.Courses
}

// GetCoursesOk returns a tuple with the Courses field value
// and a boolean to check if the value has been set.
func (o *BlockIomadCompanyAdminUnassignCoursesRequest) GetCoursesOk() ([]BlockIomadCompanyAdminUnassignCoursesRequestCoursesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Courses, true
}

// SetCourses sets field value
func (o *BlockIomadCompanyAdminUnassignCoursesRequest) SetCourses(v []BlockIomadCompanyAdminUnassignCoursesRequestCoursesInner) {
	o.Courses = v
}

func (o BlockIomadCompanyAdminUnassignCoursesRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BlockIomadCompanyAdminUnassignCoursesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["courses"] = o.Courses
	return toSerialize, nil
}

func (o *BlockIomadCompanyAdminUnassignCoursesRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"courses",
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

	varBlockIomadCompanyAdminUnassignCoursesRequest := _BlockIomadCompanyAdminUnassignCoursesRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBlockIomadCompanyAdminUnassignCoursesRequest)

	if err != nil {
		return err
	}

	*o = BlockIomadCompanyAdminUnassignCoursesRequest(varBlockIomadCompanyAdminUnassignCoursesRequest)

	return err
}

type NullableBlockIomadCompanyAdminUnassignCoursesRequest struct {
	value *BlockIomadCompanyAdminUnassignCoursesRequest
	isSet bool
}

func (v NullableBlockIomadCompanyAdminUnassignCoursesRequest) Get() *BlockIomadCompanyAdminUnassignCoursesRequest {
	return v.value
}

func (v *NullableBlockIomadCompanyAdminUnassignCoursesRequest) Set(val *BlockIomadCompanyAdminUnassignCoursesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBlockIomadCompanyAdminUnassignCoursesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBlockIomadCompanyAdminUnassignCoursesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlockIomadCompanyAdminUnassignCoursesRequest(val *BlockIomadCompanyAdminUnassignCoursesRequest) *NullableBlockIomadCompanyAdminUnassignCoursesRequest {
	return &NullableBlockIomadCompanyAdminUnassignCoursesRequest{value: val, isSet: true}
}

func (v NullableBlockIomadCompanyAdminUnassignCoursesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlockIomadCompanyAdminUnassignCoursesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


